package main

import (
	"fmt"
	"math"
)

// 推导解题：① 当所有绳段长度相等时，乘积最大。② 最优的绳段长度为 3。
// 设将绳子按照 x 长度等分为 a 段，即 n = ax ，则乘积为 x^a
// 空间复杂度O(1)，时间复杂度O(1)
func cuttingRope(n int) int {
	// 处理特殊情况
	if n <= 3 {
		return n - 1
	}
	// a表示能整除的部分，b表示整除完的余数
	a, b := n/3, n%3
	// 最终会有三种情况：全为3、有余数为1、有余数为2
	// 当没有余数时，直接得出返回x^a
	if b == 0 {
		return int(math.Pow(3, float64(a)))
	}
	// 当余数为1时，2*2 > 3*1，所以我们把a拆一个3出来，形成2*2
	if b == 1 {
		return int(math.Pow(3, float64(a-1)) * 4)
	}
	// 剩余的情况就是余数为2，那么则保留不再拆分
	return int(math.Pow(3, float64(a)) * 2)
}

func main() {
	// 测试
	n := 6
	result := cuttingRope(n)
	fmt.Println(result)
}
