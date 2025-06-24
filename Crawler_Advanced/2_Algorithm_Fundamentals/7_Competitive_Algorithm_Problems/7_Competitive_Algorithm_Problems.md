# 经典算法问题与解题策略

---

### 1. 核心内容 (Core Content)

在掌握了基础数据结构和排序算法之后，本章将引导你进入更广阔的算法世界——解决经典的算法问题。这些问题不仅是面试的重头戏，更是锻炼逻辑思维、应用数据结构和优化代码能力的绝佳途径。我们将介绍几种常见的问题类型和通用的解题策略。

**核心知识点：**
- **经典问题类型**:
    - **双指针 (Two Pointers)**: 如“三数之和”、“盛最多水的容器”。
    - **滑动窗口 (Sliding Window)**: 如“无重复字符的最长子串”、“最小覆盖子串”。
    - **递归与回溯 (Recursion & Backtracking)**: 如“全排列”、“组合总和”、“N皇后问题”。
    - **动态规划 (Dynamic Programming, DP)**: 如“爬楼梯”、“最长递增子序列”、“背包问题”。
    - **贪心算法 (Greedy Algorithm)**: 如“买卖股票的最佳时机”、“分发饼干”。
    - **二分查找 (Binary Search)**: 如“搜索旋转排序数组”、“寻找峰值”。
- **通用解题策略 (五步法)**:
    1.  **Clarification (问题澄清)**: 仔细阅读题目，向面试官确认所有不明确的细节、边界条件和输入输出格式。
    2.  **Possible Solutions (寻找所有可能解)**: 思考解决该问题的不同方法，包括暴力解法。分析它们的时间和空间复杂度。
    3.  **Optimization (优化)**: 在暴力解法的基础上，思考如何使用更优的数据结构或算法来降低复杂度。
    4.  **Coding (编码)**: 将优化后的思路清晰、准确地用代码实现。注意代码风格、变量命名和边界处理。
    5.  **Testing (测试)**: 用几个典型的测试用例（包括正常、边界、异常情况）来验证代码的正确性。

---

### 2. 学习目标 (Learning Objectives)

- **熟悉双指针、滑动窗口、回溯、动态规划等经典算法问题的模式。**
- **掌握“解题五步法”，并将其应用到每一次的算法解题实践中。**
- **能够识别一个问题背后所考察的核心数据结构和算法思想。**
- **提升将抽象问题转化为具体代码实现的编码能力。**
- **培养分析和优化算法复杂度的能力。**

---

### 3. 价值与目的 (Purpose and Value)

- **决胜技术面试**: 算法题是衡量候选人逻辑思维和编程功底的黄金标准，是进入大厂的必经关卡。
- **提升解决问题的能力**: 算法训练能系统性地提升分析问题、拆解问题和解决问题的能力，这种能力在任何技术领域都是通用的。
- **编写高质量代码**: 算法题的训练过程促使你思考代码的效率、健壮性和可读性，从而养成良好的编程习惯。
- **享受思维的乐趣**: 解决一个复杂的算法问题能带来巨大的成就感和智力上的愉悦。

---

### 4. 高级技巧与注意事项 (Advanced Techniques & Pitfalls)

- **从暴力解开始**: 不要害怕暴力解法，它通常是思考的起点，也是检验你对问题理解是否正确的基准。优化往往是在暴力解的基础上进行的。
- **空间换时间**: 很多时候，使用额外的数据结构（如哈希表）可以大幅降低时间复杂度。这是算法优化中最常见的思想。
- **画图辅助思考**: 对于链表、树、图以及一些抽象的问题，画图是帮助你理清逻辑、发现规律的绝佳方法。
- **模板化记忆**: 对于回溯、动态规划等有固定套路的算法，可以总结和记忆解题模板，在遇到同类问题时能快速上手。
- **刻意练习**: 算法能力的提升没有捷径，唯一的途径就是持续、大量的刻意练习。定期刷题，保持手感。
- **动态规划的难点**: DP 是算法中的难点。关键在于定义好 `dp` 数组的含义、找出正确的状态转移方程和确定好初始化条件。

---

### 5. 蓝图计划 (Project Blueprint)

#### LeetCode 刷题进阶之路

1.  **第一阶段：数组与字符串 (T0-T1)**
    -   **重点**: 双指针、滑动窗口。
    -   **题目**: [15. 三数之和](https://leetcode-cn.com/problems/3sum/), [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/), [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)。

2.  **第二阶段：链表与树 (T2-T3)**
    -   **重点**: 递归思想，树的遍历 (前序、中序、后序、层序)。
    -   **题目**: [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/), [104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/), [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)。

3.  **第三阶段：搜索与回溯 (T4-T5)**
    -   **重点**: 深度优先搜索 (DFS), 广度优先搜索 (BFS), 回溯法解题模板。
    -   **题目**: [46. 全排列](https://leetcode-cn.com/problems/permutations/), [78. 子集](https://leetcode-cn.com/problems/subsets/), [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)。

4.  **第四阶段：动态规划入门 (T6-T7)**
    -   **重点**: 理解状态和状态转移。
    -   **题目**: [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/), [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/), [300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)。

---

### 6. 推荐资源 (Recommended Resources)

-   **刷题平台**:
    -   [LeetCode](https://leetcode-cn.com/) (首选)
    -   [牛客网](https://www.nowcoder.com/) (国内面试题库)
-   **题解与教程**:
    -   **labuladong 的算法小抄**: [fucking-algorithm](https://github.com/labuladong/fucking-algorithm) - 提供了大量算法题的解题模板和思路。
    -   **花花酱 LeetCode**: [huahua.tech](https://zxi.mytechroad.com/blog/) - C++ 题解，思路清晰。
-   **书籍**: 《剑指 Offer》, 《程序员代码面试指南》
