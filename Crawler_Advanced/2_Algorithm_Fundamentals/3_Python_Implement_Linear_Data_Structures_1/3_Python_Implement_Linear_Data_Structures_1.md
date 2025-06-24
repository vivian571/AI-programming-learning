# Python 实现线性数据结构（一）：链表与栈

---

### 1. 核心内容 (Core Content)

理论知识需要通过编码实践来巩固。本章将重点介绍如何使用 Python 从零开始实现两种核心的线性数据结构：链表和栈。这不仅是面试的高频考点，也是深入理解数据结构内部工作原理的最佳途径。

**核心知识点：**
- **用类 (Class) 模拟数据结构**:
    - Python 的 `class` 是封装数据和行为的完美工具，非常适合用来定义数据结构。
    - `__init__` 方法用于初始化数据结构。
    - `__str__` 或 `__repr__` 方法用于友好地打印数据结构的内容，方便调试。
- **实现单向链表 (Singly Linked List)**:
    - **`Node` 类**: 包含 `data` (数据域) 和 `next` (指针域) 两个属性。
    - **`LinkedList` 类**: 包含一个 `head` 属性，指向链表的第一个节点。
    - **核心方法实现**: `append` (尾部添加), `prepend` (头部添加), `insert` (中间插入), `remove` (删除节点), `find` (查找节点)。
- **实现栈 (Stack)**:
    - **基于 Python `list` 实现**: 利用 `list` 的 `append()` 方法模拟 `push`，`pop()` 方法模拟 `pop`。这是最简单直接的方式。
    - **基于自定义链表实现**: 使用自己实现的 `LinkedList`，通过在头部添加和删除节点来模拟栈的 LIFO 行为，这能更好地理解栈的底层逻辑。

---

### 2. 学习目标 (Learning Objectives)

- **能够使用 Python 的 `class` 来定义一个节点 (`Node`) 和一个链表 (`LinkedList`)。**
- **掌握链表各项核心操作（增、删、改、查）的 Python 代码实现。**
- **能够独立处理链表操作中的各种边界情况（如空链表、单节点链表、操作头尾节点）。**
- **能够使用两种不同方式（基于 `list` 和基于链表）实现一个功能完整的栈。**
- **通过编码实践，加深对链表和栈工作原理的理解。**

---

### 3. 价值与目的 (Purpose and Value)

- **面试必备技能**: 手写链表和栈是国内外大厂技术面试中的“敲门砖”。
- **打下坚实基础**: 链表是树、图等更复杂数据结构的基础，掌握链表是后续学习的关键。
- **理解抽象与封装**: 通过 `class` 实现数据结构，能深刻体会面向对象编程中封装和抽象的思想。
- **提升代码能力**: 实现数据结构的过程涉及大量的细节和边界处理，是提升代码严谨性和健壮性的绝佳练习。

---

### 4. 高级技巧与注意事项 (Advanced Techniques & Pitfalls)

- **虚拟头节点 (Dummy Head)**: 在处理链表时，创建一个虚拟的头节点（dummy node），其 `next` 指向真正的头节点。这可以极大地简化边界条件的处理，例如在头部插入或删除节点时，无需编写特殊的 `if/else` 判断。
- **双指针技术**: 在链表中，使用两个指针（快慢指针、前后指针）同步移动，可以巧妙地解决很多问题，如查找中间节点、判断链表是否有环、删除倒数第 N 个节点等。
- **Python `list` 作为栈的效率**: Python 的 `list` 是动态数组，虽然 `append` 和 `pop` (尾部操作) 的平均时间复杂度是 O(1)，但在列表扩容时会产生额外的开销。对于超大规模或性能敏感的场景，`collections.deque` 是更好的选择。
- **注意 `None` 的处理**: 在链表操作中，要时刻注意当前节点的 `next` 是否为 `None`，防止出现 `AttributeError: 'NoneType' object has no attribute 'next'` 的错误。

---

### 5. 蓝图计划与代码示例 (Blueprint & Code Examples)

#### 阶段一：实现单向链表

```python
# 1. 定义节点类
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 2. 定义链表类
class LinkedList:
    def __init__(self):
        self.head = None

    # 3. 实现 append 方法
    def append(self, data):
        new_node = Node(data)
        if self.head is None:
            self.head = new_node
            return
        last_node = self.head
        while last_node.next:
            last_node = last_node.next
        last_node.next = new_node

    # 4. 实现打印方法，方便调试
    def __str__(self):
        nodes = []
        current = self.head
        while current:
            nodes.append(str(current.data))
            current = current.next
        return " -> ".join(nodes)

# 5. 测试
ll = LinkedList()
ll.append(1)
ll.append(2)
ll.append(3)
print(ll)  # 输出: 1 -> 2 -> 3
```

#### 阶段二：实现栈

```python
# 方式一：基于 Python list 实现
class ListStack:
    def __init__(self):
        self._items = []

    def push(self, item):
        self._items.append(item)

    def pop(self):
        if not self.is_empty():
            return self._items.pop()
        return None # 或者抛出异常

    def peek(self):
        if not self.is_empty():
            return self._items[-1]
        return None

    def is_empty(self):
        return len(self._items) == 0

    def size(self):
        return len(self._items)

# 方式二：基于 collections.deque (推荐)
from collections import deque

class DequeStack:
    def __init__(self):
        self._items = deque()
    # push, pop, peek 等实现与 ListStack 类似，但效率更高
    # ...
```

---

### 6. 推荐资源 (Recommended Resources)

-   **书籍**: 《流畅的 Python》 - 深入理解 Python 的数据模型。
-   **LeetCode**:
    -   [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)
    -   [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)
    -   [设计链表](https://leetcode-cn.com/problems/design-linked-list/) (一个很好的综合练习题)
-   **Real Python**: [Python Stacks, Queues, and Priority Queues](https://realpython.com/python-stack-queue-priority-queue/)
