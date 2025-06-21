# 第1天: 编程语言精通 (并发模型) - 脚本实践版

## 学习目标

1.  **理论理解**: 深入理解Go的`goroutine`、`channel`和Java的JUC核心组件。
2.  **实践能力**: 将理论知识应用于实际的、有用的脚本和程序中，并最终完成一个并发Web爬虫。

---

## Go 并发模型 (CSP - Communicating Sequential Processes)

Go的哲学是“**通过通信共享内存**”，而不是“通过共享内存通信”。这从根本上鼓励了更安全的并发模式。

### 1. Goroutine & Channel

*   **Goroutine**: Go语言自己管理的“轻量级线程”，创建成本极低，可以轻松创建成千上万个。
*   **Channel**: Goroutine之间通信的“管道”，用于安全地传递数据。

#### 有趣且实用的脚本案例: **并发日志处理器**

想象一下，你的服务器上有一个目录，多个服务在不断地向这个目录里写入日志文件（如 `service-A.log`, `service-B.log`）。你需要写一个脚本来实时监控并处理这些日志，统计每个文件中有多少个 "ERROR" 级别的日志。

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// worker函数负责处理单个文件
func processLogFile(filepath string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Failed to open %s: %v", filepath, err)
		return
	}
	defer file.Close()

	errorCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "ERROR") {
			errorCount++
		}
	}

	results <- fmt.Sprintf("File: %s, Error Count: %d", filepath, errorCount)
}

func main() {
	logFiles := []string{"service-A.log", "service-B.log", "service-C.log"} // 假设的日志文件列表
	// 提前创建一些假的日志文件来测试
	for _, f := range logFiles {
		os.WriteFile(f, []byte("INFO: starting...\nERROR: failed to connect\nINFO: running..."), 0644)
	}


	var wg sync.WaitGroup
	results := make(chan string, len(logFiles))

	fmt.Println("Starting log processing...")

	// 为每个日志文件启动一个goroutine
	for _, file := range logFiles {
		wg.Add(1)
		go processLogFile(file, &wg, results)
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(results) // 关闭channel，否则下面的range会一直阻塞

	// 收集并打印所有结果
	fmt.Println("\nProcessing finished. Results:")
	for result := range results {
		fmt.Println(result)
	}
	
	// 清理测试文件
	for _, f := range logFiles {
		os.Remove(f)
	}
}
```
**这个脚本如何工作?**
1. `main`函数为每个日志文件启动一个`processLogFile` goroutine。
2. 每个`worker`并发地读取和分析自己的文件，互不干扰。
3. 分析完成后，每个`worker`将结果（一个格式化的字符串）发送到`results` channel中。
4. `main`函数等待所有`worker`都完成后（`wg.Wait()`），安全地关闭`results` channel，并遍历打印所有结果。

### 2. Select - 多路复用

*   `select` 语句可以让你同时等待多个channel操作，它会选择第一个就绪的case执行。

#### 有趣且实用的脚本案例: **最快API路由**

假设你需要从3个不同的天气服务API获取数据，你只关心哪个API最先返回结果，就用它的数据，以保证你的应用响应速度最快。

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getWeatherFromAPI(name string) chan string {
	result := make(chan string)
	go func() {
		// 模拟不同的网络延迟
		latency := time.Duration(rand.Intn(120)) * time.Millisecond
		time.Sleep(latency)
		result <- fmt.Sprintf("Weather data from %s (took %v)", name, latency)
	}()
	return result
}

func main() {
	// 同时请求3个API
	api1 := getWeatherFromAPI("FastAPI")
	api2 := getWeatherFromAPI("StableAPI")
	api3 := getWeatherFromAPI("SlowAPI")

	select {
	case result := <-api1:
		fmt.Println(result)
	case result := <-api2:
		fmt.Println(result)
	case result := <-api3:
		fmt.Println(result)
	case <-time.After(100 * time.Millisecond): // 设置一个总超时
		fmt.Println("Timeout: No API responded in time.")
	}
}
```
**这个脚本如何工作?**
- `select`会同时监听3个API的channel以及一个超时channel。
- 哪个channel最先有数据返回，对应的`case`就会被执行，程序结束。
- 如果超过100毫秒没有任何API返回，`time.After`的case就会被触发，实现超时控制。

---

## Java 并发模型 (JUC - java.util.concurrent)

Java的并发基于共享内存和显式锁，JUC包提供了高级工具来管理线程和同步。

### 1. ExecutorService - 线程池

*   **核心优势**: 避免线程的频繁创建与销毁，复用线程资源，并提供统一的线程管理。

#### 有趣且实用的脚本案例: **批量图片加水印**

你需要为一个目录下的所有图片批量添加水印。这是一个典型的CPU密集型任务，非常适合使用固定大小的线程池来处理，以充分利用多核CPU，同时又不会因为创建过多线程而耗尽系统资源。

```java
import java.io.File;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class BatchImageWatermarker {

    // 模拟加水印的过程
    public static void addWatermark(File image) {
        System.out.println("Processing " + image.getName() + " on thread: " + Thread.currentThread().getName());
        try {
            // 模拟IO和CPU密集型操作
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        System.out.println("Finished " + image.getName());
    }

    public static void main(String[] args) throws InterruptedException {
        // 获取CPU核心数，创建最优大小的线程池
        int coreCount = Runtime.getRuntime().availableProcessors();
        ExecutorService executor = Executors.newFixedThreadPool(coreCount);

        File imageDir = new File("./images"); // 假设图片都在这个目录下
        imageDir.mkdir();
        // 创建一些假的图片文件
        // for(int i=0; i<10; i++) new File("./images/pic"+i+".jpg").createNewFile();


        File[] images = imageDir.listFiles();
        if (images != null) {
            for (File image : images) {
                executor.submit(() -> addWatermark(image));
            }
        }

        // 关闭线程池
        executor.shutdown(); // 不再接受新任务
        System.out.println("All tasks submitted. Waiting for completion...");
        // 等待所有已提交任务完成，最多等待1小时
        executor.awaitTermination(1, TimeUnit.HOURS);
        System.out.println("All images processed!");
    }
}
```
**这个脚本如何工作?**
1. 创建一个大小等于CPU核心数的`FixedThreadPool`，这是处理CPU密集型任务的理想配置。
2. 遍历图片目录，每张图片都作为一个任务`submit`给线程池。
3. 提交完所有任务后，调用`shutdown()`和`awaitTermination()`来优雅地关闭线程池，确保所有图片都被处理完毕。

### 2. Callable & Future - 获取异步结果

*   `Callable<V>` 是一个可以返回结果的任务。
*   `Future<V>` 代表异步计算的结果，你可以用`future.get()`来阻塞等待并获取结果。

#### 有趣且实用的脚本案例: **多源数据聚合**

假设你需要为一个报表聚合数据，数据源来自3个部分：一个来自数据库，一个来自外部HTTP API，一个来自本地文件。这三个数据源可以并行获取。

```java
import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.*;

public class DataAggregator {

    // 模拟从数据库获取数据
    static Callable<String> fetchFromDB = () -> {
        TimeUnit.SECONDS.sleep(2);
        return "Data from DB";
    };

    // 模拟从API获取数据
    static Callable<String> fetchFromAPI = () -> {
        TimeUnit.SECONDS.sleep(3);
        return "Data from API";
    };

    // 模拟从文件读取数据
    static Callable<String> fetchFromFile = () -> {
        TimeUnit.SECONDS.sleep(1);
        return "Data from File";
    };

    public static void main(String[] args) throws Exception {
        ExecutorService executor = Executors.newCachedThreadPool();
        List<Callable<String>> tasks = List.of(fetchFromDB, fetchFromAPI, fetchFromFile);

        // 提交所有任务，并获取Future列表
        List<Future<String>> futures = executor.invokeAll(tasks);

        System.out.println("Tasks submitted. Aggregating results...");
        
        List<String> results = new ArrayList<>();
        for (Future<String> future : futures) {
            // future.get()会阻塞，直到该任务完成
            results.add(future.get());
        }

        System.out.println("\nFinal Report Data:");
        results.forEach(System.out::println);

        executor.shutdown();
    }
}
```
**这个脚本如何工作?**
1. 三个耗时的数据获取操作被分别定义为`Callable`任务。
2. `executor.invokeAll(tasks)` 会一次性提交所有任务，并返回一个`Future`列表。这个方法会阻塞，直到所有任务都完成。
3. 遍历`Future`列表，调用`get()`方法（此时不会阻塞，因为任务已完成）来安全地获取每个任务的结果。
4. 整个过程的耗时取决于最慢的那个任务（3秒），而不是所有任务耗时之和（2+3+1=6秒）。

---
## 最终实践: 并发Web爬虫

现在，你可以运用上面学到的知识，挑战实现一个并发Web爬虫。参考之前提供的详细步骤，并尝试将这些有趣的脚本思想融入其中。

**祝你学习愉快，编码顺利！**