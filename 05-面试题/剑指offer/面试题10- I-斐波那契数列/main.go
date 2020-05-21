package main

import "fmt"

// 动态规划：
// 原理： 以斐波那契数列性质 f(n + 1) = f(n) + f(n - 1) 为转移方程。百科的是：f(n) = f(n-1) + f(n-2)
// 时间复杂度O(n)，空间复杂度 O(1)
func fib(n int) int {
	a, b:=0, 1
	// 这里的a理解为f(n)，求现在；b理解为f(n+1)，求未来
	// a = 之前的b 这个不难理解，因为之前的b是f(n+1)，也就是未来的a值
	// b = 之前的 a + b，之前的a相对现在是f(n-1)，之前的b相对现在的f(n)
	for i:=1;i<=n;i++{  //总是从1开始，即f(1)，那么此时赋值后 a = 1（结果）
		a,b=b,(a+b)%1000000007  // 取模，如果不小于1000000007则无法取模，取模不生效
	}
	return a
}

func main() {
	res := fib(4)
	fmt.Println(res)
}
