 
# Day 3: 计算机网络 - HTTP/2 与 HTTP/3 的演进

## 学习目标

- **宏观理解**：清晰地阐述 HTTP 协议从 HTTP/1.1 到 HTTP/2，再到 HTTP/3 的主要演进动机和解决的核心问题。
- **核心区别**：能够准确说出 HTTP/2 与 HTTP/3 在传输层协议、多路复用、队头阻塞（HOL Blocking）解决方案上的本质区别。
- **实践能力**：掌握使用 Wireshark 工具抓取并分析 HTTPS 流量的基本方法，能从抓包结果中识别出 TLS 握手过程的关键步骤。

## 核心知识：从 HTTP/2 到 HTTP/3 的飞跃

### 1. HTTP/2 的革新与局限

HTTP/2 相比于 HTTP/1.1 是一次巨大的进步，它通过以下技术显著提升了性能：

- **二进制分帧 (Binary Framing)**：将所有传输的信息分割为更小的消息和帧，并对它们采用二进制格式编码。
- **多路复用 (Multiplexing)**：在单个 TCP 连接上，客户端和服务器可以同时、并行地发送和接收多个请求和响应，彻底解决了 HTTP/1.1 的队头阻塞问题。
- **头部压缩 (Header Compression)**：使用 HPACK 算法压缩请求和响应的头部，减少了传输开销。
- **服务器推送 (Server Push)**：服务器可以主动向客户端推送资源，而无需客户端明确请求。

**HTTP/2 的核心局限：TCP 队头阻塞 (HOL Blocking)**

尽管 HTTP/2 解决了应用层的队头阻塞，但它构建于 TCP 协议之上。TCP 是一个可靠的、按序传输的协议。如果在一个 TCP 连接中，某个数据包（Packet）丢失了，那么所有后续的数据包（即使已经到达）都必须等待这个丢失的包被重传并成功接收后，才能被 TCP 协议栈向上层（HTTP/2）交付。

这就导致了一个问题：在一个繁忙的 HTTP/2 连接上，哪怕只有一个小小的 JS 文件的数据包丢失，也会阻塞住后面所有正在传输的资源（比如图片、CSS等），因为它们共享同一个 TCP 连接。这就是**TCP 层的队头阻塞**。

### 2. HTTP/3 的革命：拥抱 QUIC

为了彻底解决 TCP 的队头阻塞问题，HTTP/3 做出了一个革命性的改变：**将传输层从 TCP 更换为 UDP，并在 UDP 之上构建了一个全新的协议——QUIC**。

**QUIC (Quick UDP Internet Connections) 的核心特性：**

- **基于 UDP，实现可靠传输**：QUIC 在 UDP 这个“不可靠”的协议之上，自己实现了一套可靠传输机制，包括包的确认（ACK）、重传（Retransmission）和流量控制（Flow Control）。
- **真正的多路复用**：QUIC 的多路复用是建立在“连接ID”（Connection ID）和“流”（Stream）的概念上的。每个流都是独立的。如果一个流中的某个 UDP 包丢失了，它只会阻塞住当前这个流，而不会影响其他并行的流。这从根本上解决了队头阻塞问题。
- **集成的 TLS 加密**：QUIC 将 TLS 1.3 的加密和握手过程深度集成。相比于 TCP + TLS 的模型（TCP 握手 -> TLS 握手），QUIC 握手和加密协商通常只需要 1-RTT（往返时间），对于已有连接甚至可以实现 0-RTT，大大加快了连接建立的速度。
- **连接迁移 (Connection Migration)**：当你的网络环境发生变化时（例如从 Wi-Fi 切换到 4G），TCP 连接会中断。而 QUIC 使用连接 ID 来标识一次会话，而不是依赖 IP 地址和端口四元组。因此，即使网络切换导致 IP 地址变化，QUIC 连接也能无缝地迁移，保持会话不中断。

### 核心区别总结

| 特性 | HTTP/2 | HTTP/3 |
| :--- | :--- | :--- |
| **底层协议** | TCP | UDP (之上构建 QUIC) |
| **多路复用** | 在单个 TCP 连接上复用 | 在 QUIC 连接上，通过独立的“流”实现 |
| **队头阻塞** | 解决了应用层HOL，但存在TCP层HOL | 从根本上解决了HOL |
| **连接建立** | TCP 握手(1-RTT) + TLS 握手(1-2 RTT) | QUIC 握手 (集成TLS1.3, 0-1 RTT) |
| **连接迁移** | 不支持，IP改变则连接中断 | 支持，通过连接ID实现无缝迁移 |

## 实践操作：使用 Wireshark 分析 HTTPS 握手

**目标**：抓取并分析一次你访问 `https://www.baidu.com` 的 TLS 握手流量。

### 第一步：准备工作

1.  **安装 Wireshark**：从 [Wireshark 官网](https://www.wireshark.org/) 下载并安装最新版本。
2.  **关闭无关应用**：关闭浏览器、聊天工具等所有可能产生网络流量的应用，以减少干扰。

### 第二步：配置 Wireshark (仅首次需要)

由于 HTTPS 流量是加密的，直接抓包看到的是乱码。为了让 Wireshark 能解密，我们需要配置它使用浏览器的 TLS 密钥。

1.  **设置系统环境变量**：
    - 新建一个名为 `SSLKEYLOGFILE` 的环境变量。
    - 变量值为一个文件的绝对路径，例如 `C:\Users\YourUser\Desktop\sslkeylog.log`。
    - **验证**：重启你的命令行工具（如 PowerShell 或 CMD），输入 `echo $env:SSLKEYLOGFILE` (PowerShell) 或 `echo %SSLKEYLOGFILE%` (CMD)，确保能看到你设置的路径。
2.  **配置 Wireshark**：
    - 打开 Wireshark，进入 `编辑 (Edit)` -> `首选项 (Preferences)`。
    - 导航到 `Protocols` -> `TLS`。
    - 在 `(Pre)-Master-Secret log filename` 字段，填入或选择你刚刚设置的 `sslkeylog.log` 文件路径。
    - 点击“确定”保存。

### 第三步：开始抓包

1.  在 Wireshark 主界面，双击选择你正在用于上网的那个网络接口（例如 `WLAN` 或 `以太网`）。
2.  在顶部的筛选器输入框中，输入 `tls.handshake`，这样可以只显示 TLS 握手相关的包，方便观察。
3.  打开你的浏览器（如 Chrome 或 Edge，它们支持导出密钥），访问 `https://www.baidu.com`。
4.  观察 Wireshark，你会看到一系列的 TLS 握手报文被捕获。
5.  页面加载完成后，点击 Wireshark 左上角的红色方块按钮停止抓包。

### 第四步：分析握手过程

在 Wireshark 的包列表窗口，你应该能清晰地看到以下几个关键步骤：

1.  **Client Hello**:
    - 这是握手的开始，由你的浏览器（客户端）发送。
    - **看点**：点击这个包，在下方的“协议详情”窗口展开 `Transport Layer Security` -> `Handshake Protocol: Client Hello`。
    - 你可以看到客户端支持的 TLS 版本 (`Version`)、支持的加密套件 (`Cipher Suites`)、以及一个关键的 `Random` 值。

2.  **Server Hello**:
    - 这是百度服务器（服务端）的响应。
    - **看点**：服务器会从客户端支持的列表中，选择一个它也支持的加密套件，并确认使用的 TLS 版本。同时，它也会生成自己的一个 `Random` 值。

3.  **Certificate, Server Key Exchange, Server Hello Done**:
    - **Certificate**: 服务器将其数字证书发送给客户端，客户端用以验证服务器的身份。你可以看到证书的颁发机构、有效期等信息。
    - **Server Key Exchange**: 服务器发送密钥交换所需的参数。
    - **Server Hello Done**: 标志着服务器的初始协商消息发送完毕。

4.  **Client Key Exchange, Change Cipher Spec, Encrypted Handshake Message**:
    - **Client Key Exchange**: 客户端生成一个用于对称加密的“预主密钥”(Pre-Master Secret)，并用服务器证书中的公钥加密后发送给服务器。
    - **Change Cipher Spec**: 客户端通知服务器，从现在开始，后续的通信都将使用刚刚协商好的对称密钥进行加密。
    - **Encrypted Handshake Message**: 这是客户端发出的第一条加密消息，内容是之前所有握手消息的摘要，用于校验。

5.  **New Session Ticket, Change Cipher Spec, Encrypted Handshake Message (from Server)**:
    - 服务器用自己的私钥解密得到“预主密钥”，并同样计算出对称密钥。
    - **Change Cipher Spec**: 服务器同样通知客户端，后续通信将采用加密方式。
    - **Encrypted Handshake Message**: 服务器也发送一条加密的握手摘要给客户端进行校验。

至此，TLS 握手完成。双方都拥有了相同的对称密钥，后续的 HTTP 应用数据将通过这个安全的加密通道进行传输。你可以清除 `tls.handshake` 筛选器，就能看到后续被解密出来的 HTTP/2 或 HTTP/3 流量了。
