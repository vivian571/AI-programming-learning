# Python 实现线性数据结构（二）：队列与双端队列

---

### 1. 核心内容 (Core Content)

本章继续我们的 Python 实现之旅，重点是队列 (Queue) 和其更通用的形式——双端队列 (Deque)。我们将探讨如何用不同的方式实现它们，并理解在 Python 中使用内置模块 `collections.deque` 的巨大优势。

**核心知识点：**
- **实现队列 (Queue)**:
    - **基于 `list` 的低效实现**: 直接使用 `list` 的 `append()` 入队和 `pop(0)` 出队。这种方式非常直观，但效率低下，因为 `pop(0)` 的时间复杂度是 O(n)。
    - **基于自定义链表实现**: 维护 `head` 和 `tail` 两个指针，`enqueue` 在尾部添加节点，`dequeue` 在头部移除节点，两个操作的时间复杂度都是 O(1)，效率高。
- **循环队列 (Circular Queue)**:
    - **概念**: 基于定长数组实现队列时，为了避免空间的浪费和数据的频繁移动，我们把数组看作一个环形结构。
    - **实现关键**: 使用两个指针 `front` 和 `rear`，并通过模运算 `(index + 1) % capacity` 来实现环形移动，从而高效地复用数组空间。
- **Python 的 `collections.deque`**:
    - **定义**: `deque` (double-ended queue) 是 Python 标准库中一个高度优化的双端队列。
    - **底层实现**: 它基于双向链表实现，因此在两端进行添加 (`append`, `appendleft`) 和删除 (`pop`, `popleft`) 操作的时间复杂度都是 O(1)。
    - **应用**: 是在 Python 中实现队列和栈的最推荐、最高效的方式。

---

### 2. 学习目标 (Learning Objectives)

- **能够分析并指出直接使用 Python `list` 实现队列的性能瓶颈。**
- **能够使用自定义链表，通过维护头尾指针，高效地实现一个队列。**
- **理解循环队列的设计思想，并能用 Python 数组和模运算实现一个循环队列。**
- **熟练掌握 `collections.deque` 的使用方法，并将其作为实现队列和栈的首选工具。**
- **能够在不同场景下，选择最合适的队列实现方式。**

---

### 3. 价值与目的 (Purpose and Value)

- **深入理解队列原理**: 通过手写不同版本的队列，深刻理解其 FIFO（先进先出）特性和底层实现机制。
- **掌握性能优化技巧**: 学习循环队列是理解如何用巧妙设计来优化数据结构性能的经典案例。
- **熟悉 Python 标准库**: 了解并善用 `collections.deque` 这样的高效内置模块，是 Pythonic 编程风格的体现，能极大提升代码质量和运行效率。
- **解决实际问题**: 队列是广度优先搜索 (BFS)、操作系统任务调度、消息中间件等众多应用场景的核心数据结构。

---

### 4. 高级技巧与注意事项 (Advanced Techniques & Pitfalls)

- **循环队列的判满与判空**:
    - **牺牲一个空间**: 保持数组中至少有一个空位。当 `(rear + 1) % capacity == front` 时，队列为满。当 `rear == front` 时，队列为空。这是最常见的实现方式。
    - **使用计数器**: 额外维护一个 `size` 变量来记录元素数量。`size == capacity` 则满，`size == 0` 则空。
- **线程安全**: Python 的 `collections.deque` 和 `list` 都不是线程安全的。在多线程环境下，需要使用 `queue.Queue` 模块，它内部实现了线程同步机制。
- **`deque` 的其他妙用**: `deque` 还提供了 `rotate(n)` 方法，可以方便地将队列中的元素进行循环移动，这在某些算法问题中非常有用。

---

### 5. 蓝图计划与代码示例 (Blueprint & Code Examples)

#### 阶段一：基于链表实现队列

```python
# (需先实现上一章的 Node 和 LinkedList)
class LinkedListQueue:
    def __init__(self):
        self._items = LinkedList()
        # 为了 O(1) 的入队效率，需要维护一个 tail 指针
        self.tail = None

    def enqueue(self, item):
        new_node = Node(item)
        if self.is_empty():
            self._items.head = new_node
            self.tail = new_node
        else:
            self.tail.next = new_node
            self.tail = new_node

    def dequeue(self):
        if self.is_empty():
            return None
        item_node = self._items.head
        self._items.head = item_node.next
        if self._items.head is None: # 如果队列变空
            self.tail = None
        return item_node.data

    def is_empty(self):
        return self._items.head is None
```

#### 阶段二：使用 `collections.deque` (推荐方式)

```python
from collections import deque

# Deque 本身就是功能完善的队列
queue = deque()

# 入队
queue.append("task1")
queue.append("task2")
queue.append("task3")
print("Queue:", queue) # Queue: deque(['task1', 'task2', 'task3'])

# 出队
item = queue.popleft()
print("Processing:", item) # Processing: task1
print("Queue after dequeue:", queue) # Queue after dequeue: deque(['task2', 'task3'])

# 查看队头
print("Next task:", queue[0]) # Next task: task2

# 作为栈使用
stack = deque()
stack.append("A") # push
stack.append("B") # push
print(stack.pop()) # pop -> 'B'
```

---

### 6. 推荐资源 (Recommended Resources)

-   **Python 官方文档**: [collections.deque](https://docs.python.org/3/library/collections.html#collections.deque)
-   **LeetCode**:
    -   [用队列实现栈 (225)](https://leetcode-cn.com/problems/implement-stack-using-queues/)
    -   [设计循环队列 (622)](https://leetcode-cn.com/problems/design-circular-queue/)
-   **Real Python**: [How to Implement a Python Queue (FIFO)](https://realpython.com/how-to-implement-a-python-queue/)
