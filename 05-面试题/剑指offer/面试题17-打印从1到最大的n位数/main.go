package main

import (
	"fmt"
)

// 思路：要打印的值是10的n次方-1以内的
func printNumbers(n int) []int {
	var res []int  // 接收结果
	num := 1
	for n > 0 {  // 完成10的n次方计算
		num *= 10
		n --
	}
	num = num - 1  // 完成10的n次方-1计算
	for i:= 1; i<=num; i++{  //遍历结果即可
		res = append(res, i)
	}
	return res
}

func main() {
	// 测试
	var n = int(1)
	result := printNumbers(n)
	fmt.Println(result)
}
