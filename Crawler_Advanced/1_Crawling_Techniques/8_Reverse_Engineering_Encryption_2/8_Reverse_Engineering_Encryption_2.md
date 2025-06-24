# 爬虫逆向与 JS 破解实战 (二)

---

### 1. 核心内容 (Core Content)

在掌握了基本的 JS 逆向调试技巧后，本章将深入探讨更复杂的逆向场景和高级技术。这包括应对代码混淆、理解 WebAssembly (WASM)、使用 Hook 技术以及自动化反混淆工具的应用。

**核心知识点：**
- **高级混淆技术与反混淆**:
    - **obfuscator.io**: 一种流行的、高强度的 JS 混淆工具，其特征包括字符串阵列、控制流平坦化、死代码注入等。
    - **AST (Abstract Syntax Tree)**: 将 JS 代码解析成树状结构，通过编写脚本对 AST 进行遍历和修改，实现自动化反混淆，如还原变量名、合并字符串、解平坦化控制流。
    - **常用工具**: Babel、esprima、estraverse、escodegen。
- **WebAssembly (WASM) 逆向**:
    - **WASM 简介**: 一种可移植、体积小、加载快的二进制格式，常用于在浏览器中运行高性能代码（如 C/C++/Rust 编译而来）。
    - **逆向思路**: 通过分析 JS 调用 WASM 的接口，理解其输入输出；使用 WASM 反编译工具 (如 `wasm2wat`, `Wabt`) 将其转换为可读的文本格式 (`.wat`) 进行分析。
- **Hook 技术实战**:
    - **目的**: 拦截和修改原生函数或对象方法的行为，用于监控加密过程、绕过反调试或直接获取加密结果。
    - **实现**: 通过 JavaScript 的 `Proxy` 对象或直接重写函数 (`window.btoa = function(...)`) 来实现。
- **RPC (Remote Procedure Call)**:
    - **概念**: 当逆向成本过高时，可以考虑搭建一个 Node.js 服务，直接运行前端的加密 JS，Python 爬虫通过 RPC 调用该服务来获取加密结果。这是一种工程化的妥协方案。

---

### 2. 学习目标 (Learning Objectives)

- **能够识别并分析由 `obfuscator.io` 等工具生成的复杂混淆代码。**
- **理解 AST 的基本概念，并能使用相关工具进行简单的自动化反混淆操作。**
- **了解 WebAssembly 技术及其在前端加密中的应用，掌握初步的 WASM 逆向方法。**
- **熟练运用 Hook 技术来辅助逆向分析，如 Hook `cookie`、`eval` 等关键函数。**
- **掌握 RPC 的思想，并能够在必要时采用此方案解决复杂的 JS 逆向问题。**

---

### 3. 价值与目的 (Purpose and Value)

- **攻克高难度反爬**: 掌握高级逆向技术，能够应对绝大部分前端加密防护，是顶级爬虫工程师的核心竞争力。
- **提升自动化水平**: 从手动调试转向利用 AST 等工具进行半自动化、自动化的逆向分析，大幅提升效率。
- **拓展技术视野**: 接触并理解 WebAssembly 等前沿技术，保持技术栈的先进性。
- **培养工程化思维**: 学会评估逆向成本，并在“完全破解”和“RPC 调用”等不同方案间做出权衡，找到最优解。

---

### 4. 高级技巧与注意事项 (Advanced Techniques & Pitfalls)

- **AST 反混淆的复杂性**: 编写通用的反混淆脚本非常困难，通常需要针对特定网站的混淆特征进行定制。
- **WASM 的黑盒特性**: WASM 内部逻辑通常是编译后的底层代码，可读性差，逆向分析难度远大于 JS，通常优先分析调用它的 JS 代码。
- **Hook 被检测**: 一些网站会通过 `Function.prototype.toString.call(func)` 等方式检测原生函数是否被修改，需要更隐蔽的 Hook 手段。
- **环境差异**: 前端的 JS 代码可能依赖特定的浏览器环境 (如 `window`, `document` 对象)，在 Node.js 中执行时需要使用 `JSDOM` 等库来模拟浏览器环境。
- **代码的持续迭代**: 网站前端代码会不断更新，逆向脚本可能随时失效，需要考虑脚本的健壮性和可维护性。

---

### 5. 蓝图计划 (Project Blueprint)

#### 项目：逆向一个使用 `obfuscator.io` 混淆并结合 WASM 加密的接口

1.  **初步侦查 (T0)**
    -   抓包分析，确定目标接口和加密参数。
    -   查看 JS 源码，识别出 `obfuscator.io` 的典型特征 (如 `_0x...` 变量名，字符串阵列)。

2.  **AST 反混淆 (T1)**
    -   将混淆的 JS 代码保存到本地。
    -   编写一个基于 Babel 的脚本，尝试进行自动化反混淆：
        -   解密字符串阵列。
        -   简化常量计算 (如 `_0x123 + _0x456` 直接替换为结果)。
        -   尝试还原部分控制流。
    -   输出初步反混淆后的代码，提升可读性。

3.  **定位与调试 (T2)**
    -   在反混淆后的代码中，通过动态调试，定位到关键的加密逻辑。
    -   分析发现加密过程调用了 WASM 模块。

4.  **WASM 分析与 Hook (T3)**
    -   分析 JS 如何加载和调用 WASM 函数，理解其传入的参数和返回的数据。
    -   使用 Hook 技术拦截 JS 对 WASM 函数的调用，打印输入和输出，确认其功能。
    -   (可选) 使用 `wasm2wat` 等工具反编译 WASM，尝试理解其内部算法。

5.  **构建 RPC 服务 (T4)**
    -   由于 WASM 逆向成本高，决定采用 RPC 方案。
    -   创建一个 Node.js 服务 (使用 Express 或 Koa)。
    -   将调用 WASM 加密的相关 JS 逻辑整体搬入 Node.js 服务，并封装成一个 API 接口。

6.  **Python 集成 (T5)**
    -   Python 爬虫在需要加密时，通过 HTTP 请求调用本地的 Node.js RPC 服务。
    -   获取加密结果，并用于后续的爬取请求。

---

### 6. 推荐资源 (Recommended Resources)

-   **Obfuscator.io**: [https://obfuscator.io/](https://obfuscator.io/) - 在线的 JS 混淆工具，可以用来学习混淆特征。
-   **Babel REPL**: [https://babeljs.io/repl](https://babeljs.io/repl) - 在线查看 JS 代码转换成 AST 的结果。
-   **《JavaScript 语言精粹》**: 深入理解 JS 语言特性是逆向的基础。
-   **MDN WebAssembly**: [WebAssembly Docs](https://developer.mozilla.org/en-US/docs/WebAssembly)
-   **Node.js**: [https://nodejs.org/](https://nodejs.org/)
-   **PyExecJS / PyMiniRacer**: Python 中执行 JS 的库。
