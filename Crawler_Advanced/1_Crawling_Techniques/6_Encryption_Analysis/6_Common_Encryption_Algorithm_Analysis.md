# 常见加密算法分析与爬虫破解思路

---

### 1. 核心内容 (Core Content)

在现代网络爬虫实战中，常常会遇到目标网站对数据接口、参数或内容进行加密处理。理解常见加密算法的原理和破解思路，是突破反爬机制、实现数据采集的关键能力。

**核心知识点：**
- **对称加密算法**：如 AES、DES，特点是加密解密密钥相同，常用于数据传输加密。
- **非对称加密算法**：如 RSA，公钥加密私钥解密，常用于登录、签名等场景。
- **哈希算法**：如 MD5、SHA1、SHA256，常用于数据完整性校验、密码存储。
- **Base64 编码**：常见的内容混淆方式，非加密但常与加密算法组合使用。
- **加密算法在前端的实现**：JavaScript 加密代码的逆向与还原。

---

### 2. 学习目标 (Learning Objectives)

- 理解常见加密算法的基本原理和应用场景。
- 掌握加密参数的识别与逆向分析方法。
- 能够利用 Python、Node.js 等工具还原前端加密逻辑。
- 学会结合抓包、调试工具定位加密流程。
- 掌握常见加密算法的破解与绕过技巧。

---

### 3. 价值与目的 (Purpose and Value)

- **突破反爬壁垒**：破解加密参数是高阶爬虫工程师的必备技能。
- **提升逆向能力**：加深对前端与后端数据交互安全机制的理解。
- **工程实用性强**：广泛应用于数据采集、接口测试、自动化运维等场景。
- **安全意识提升**：理解加密算法有助于提升自身开发的安全性。

---

### 4. 高级技巧与注意事项 (Advanced Techniques & Pitfalls)

- **前端加密逆向**：利用浏览器开发者工具、Fiddler、Charles 等抓包工具，结合 JS 断点调试，定位加密函数。
- **Python 调用 JS**：通过 PyExecJS、PyV8、Node.js 等方式在 Python 中执行前端加密 JS 代码。
- **密钥获取陷阱**：部分密钥为动态生成或后端下发，需结合请求链路分析。
- **哈希不可逆陷阱**：如 MD5、SHA256 等哈希算法不可逆，只能通过彩虹表、爆破等方式尝试还原。
- **加密算法组合**：实际场景中常见多种加密算法嵌套组合，需逐层还原。
- **法律与道德风险**：破解加密仅限于学习和授权测试，切勿用于非法用途。

---

### 5. 蓝图计划 (Project Blueprint)

#### 项目：模拟登录加密接口数据采集

1. **目标分析 (T0)**
    - 通过浏览器抓包工具分析登录请求，识别加密参数及算法类型。

2. **前端加密逆向 (T1)**
    - 利用 Chrome DevTools 设置断点，定位加密 JS 函数。
    - 提取加密核心代码，分析加密流程。

3. **Python 集成加密算法 (T2)**
    - 使用 PyExecJS 或 Node.js 在 Python 中调用 JS 加密函数。
    - 或用 Python 实现等价加密算法（如 AES、RSA、MD5 等）。

4. **模拟请求与数据采集 (T3)**
    - 构造加密参数，模拟登录或数据请求。
    - 采集目标数据，处理解密与还原。

5. **自动化与优化 (T4)**
    - 封装加密与请求流程，实现批量自动化采集。
    - 优化性能与异常处理。

---

### 6. 代码示例 (Code Example)

```python
# Python 调用 JS 实现加密（以 MD5 为例）
import execjs

js_code = """
function md5Encrypt(text) {
    return hex_md5(text); // 需引入 md5.js
}
"""

ctx = execjs.compile(js_code + open('md5.js', encoding='utf-8').read())
encrypted = ctx.call('md5Encrypt', '123456')
print(encrypted)
```

```python
# Python 实现 AES 加密
from Crypto.Cipher import AES
import base64

def pad(text):
    return text + (16 - len(text) % 16) * chr(16 - len(text) % 16)

key = b'1234567890abcdef'
iv = b'abcdef1234567890'
cipher = AES.new(key, AES.MODE_CBC, iv)
encrypted = cipher.encrypt(pad('hello world').encode())
print(base64.b64encode(encrypted).decode())


