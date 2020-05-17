package main

import (
	"fmt"
)

// 从小到大往下排，也就是左下角是本列最大本行最小，而右上角是本行最大本列最小，利用这个特性，我们可以以行和列来剔除
// matrix定位标志位，从左下角开始，matrix > target，那么target一定在上边，matrix < target，那么target在右边
// 时间复杂度O(m+n)，空间复杂度O(1)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0 //i控制行,j控制列
	//从左下角开始
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {  // 判断如果一直，就直接返回
			return true
		} else if matrix[i][j] > target {  // 判断如果小于标志位，则表明在上边，所以i上移一位，即减1
			i--
		} else if matrix[i][j] < target {  // 判断如果大于标志位，则表明在右边，所以j右移一位，即加1
			j++
		} else {
			return true
		}
	}
	return false
}

func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	result := findNumberIn2DArray(arr, target)
	fmt.Println(result)
}
