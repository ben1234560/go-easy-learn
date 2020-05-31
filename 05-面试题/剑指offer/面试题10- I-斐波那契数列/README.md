# README

`动态规划` `迭代`

#### 面试题10- I. 斐波那契数列

> 难度：简单

写一个函数，输入 `n` ，求斐波那契（Fibonacci）数列的第 `n` 项。斐波那契数列的定义如下：

~~~
F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
~~~

斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

**示例1:**

~~~go
func fib(n int) int {
    
}

输入：n = 2
输出：1
~~~

**示例2：**

例如：给出

~~~
输入：n = 5
输出：5
~~~

**提示：**

```
0 <= n <= 100
```



可参考[解题思路](https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/solution/mian-shi-ti-10-i-fei-bo-na-qi-shu-lie-dong-tai-gui/)

