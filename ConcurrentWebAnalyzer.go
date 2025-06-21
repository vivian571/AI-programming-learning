package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// Task 定义了爬取任务的结构
type Task struct {
	URL   string
	Depth int
}

// SafeMap 是一个线程安全的map，用于存储已访问的URL和外部资源
type SafeMap struct {
	v   map[string]bool
	mux sync.Mutex
}

// Add 添加一个键，如果键已存在则返回false
func (s *SafeMap) Add(key string) bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	if _, ok := s.v[key]; ok {
		return false // 已存在
	}
	s.v[key] = true
	return true // 添加成功
}

// Value 返回map的快照
func (s *SafeMap) Value() map[string]bool {
	s.mux.Lock()
	defer s.mux.Unlock()
	// 返回一个副本以保证线程安全
	newMap := make(map[string]bool)
	for k, v := range s.v {
		newMap[k] = v
	}
	return newMap
}

// externalResources 用于存储所有外部资源的域名
var externalResources = &SafeMap{v: make(map[string]bool)}

// visitedUrls 存储所有已访问的URL
var visitedUrls = &SafeMap{v: make(map[string]bool)}

// main函数是程序的入口
func main() {
	startURL := "https://golang.org" // 你可以换成任何想分析的网站
	maxDepth := 2                    // 最大爬取深度
	workerCount := 10                // 并发worker数量

	tasks := make(chan Task, 100)
	var wg sync.WaitGroup

	// 启动worker池
	for i := 0; i < workerCount; i++ {
		go worker(i, tasks, &wg, maxDepth)
	}

	// 添加初始任务
	wg.Add(1)
	tasks <- Task{URL: startURL, Depth: 0}

	// 使用一个goroutine来等待所有任务完成，然后关闭channel
	// 这样可以防止主goroutine阻塞在wg.Wait()而无法关闭channel
	go func() {
		wg.Wait()
		close(tasks)
	}()

	// 结果分析
	fmt.Println("\n--- Analysis Complete ---")
	fmt.Printf("Crawled %d unique pages.\n", len(visitedUrls.Value()))
	fmt.Println("External domains this site depends on:")
	for domain := range externalResources.Value() {
		fmt.Printf("- %s\n", domain)
	}
}

// worker 是并发执行爬取任务的单元
func worker(id int, tasks chan Task, wg *sync.WaitGroup, maxDepth int) {
	for task := range tasks {
		if task.Depth > maxDepth || !visitedUrls.Add(task.URL) {
			wg.Done()
			continue
		}

		log.Printf("[Worker %d] Analyzing: %s (Depth: %d)", id, task.URL, task.Depth)
		crawl(task, tasks, wg)
	}
}

// crawl 执行单个URL的爬取和分析
func crawl(task Task, tasks chan<- Task, wg *sync.WaitGroup) {
	defer wg.Done()

	// 1. 获取页面
	resp, err := http.Get(task.URL)
	if err != nil {
		log.Printf("Failed to get %s: %v", task.URL, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Status error on %s: %s", task.URL, resp.Status)
		return
	}

	// 2. 解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("Failed to parse HTML from %s: %v", task.URL, err)
		return
	}

	// 3. 并发地寻找链接和外部资源
	var findWg sync.WaitGroup
	findWg.Add(2)

	// Goroutine 1: 寻找并添加新的爬取任务 (<a>标签)
	go func() {
		defer findWg.Done()
		findLinks(doc, task, tasks, wg)
	}()

	// Goroutine 2: 寻找并记录外部资源 (<link>, <script>)
	go func() {
		defer findWg.Done()
		findAssets(doc, task)
	}()

	findWg.Wait()
}

// findLinks 解析<a>标签，并将新链接作为任务添加到channel
func findLinks(n *html.Node, currentTask Task, tasks chan<- Task, wg *sync.WaitGroup) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link, err := url.Parse(a.Val)
				if err != nil {
					continue
				}
				base, _ := url.Parse(currentTask.URL)
				absURL := base.ResolveReference(link).String()

				// 仅爬取同域名下的链接
				if strings.HasPrefix(absURL, base.Scheme+"://"+base.Host) {
					wg.Add(1)
					// 使用select防止在wg.Wait()之后发送任务导致死锁
					select {
					case tasks <- Task{URL: absURL, Depth: currentTask.Depth + 1}:
					case <-time.After(1 * time.Second): // 如果channel阻塞，则放弃
						log.Println("Tasks channel full, dropping link.")
						wg.Done()
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinks(c, currentTask, tasks, wg)
	}
}

// findAssets 解析<link>和<script>标签，记录外部域名
func findAssets(n *html.Node, currentTask Task) {
	if n.Type == html.ElementNode {
		var srcAttr string
		if n.Data == "script" || n.Data == "img" {
			srcAttr = "src"
		} else if n.Data == "link" {
			// 只关心stylesheet类型的link
			isStylesheet := false
			for _, attr := range n.Attr {
				if attr.Key == "rel" && attr.Val == "stylesheet" {
					isStylesheet = true
					break
				}
			}
			if isStylesheet {
				srcAttr = "href"
			}
		}

		if srcAttr != "" {
			for _, a := range n.Attr {
				if a.Key == srcAttr {
					link, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					base, _ := url.Parse(currentTask.URL)
					assetURL := base.ResolveReference(link)

					// 如果资源域名与当前网站域名不同，则记录为外部依赖
					if assetURL.Host != "" && assetURL.Host != base.Host {
						externalResources.Add(assetURL.Host)
					}
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAssets(c, currentTask)
	}
}

// 如何运行这个程序:
// 确保你已经安装了 Go 语言环境。
// 将上面的代码保存到 ConcurrentWebAnalyzer.go。
// 在终端中，cd到文件所在的目录。
// 运行 go mod init webanalyzer 来初始化模块。
// 运行 go get golang.org/x/net/html 来下载依赖。
// 运行 go run ConcurrentWebAnalyzer.go 来启动程序。
// 我很乐意协助你完成后续步骤或解答任何关于代码的问题。
