package main

import (
	"fmt"
)

// 思路：循环减去最右边得1
// 空间复杂度O(M)，M是二进制中1个个数，时间复杂度O(1)
func hammingWeight(num uint32) int {
	res := 0
	// 循环num
	for num > 0{
		// 如果进入说明是二进制num最右边的是1，统计变量+1
		res += 1
		// 并1变成0，
		num &= num - 1
	}
	return res
}

func main() {
	// 测试
	var n = uint32(00000000000000000000000000001011)
	result := hammingWeight(n)
	fmt.Println(result)
}
