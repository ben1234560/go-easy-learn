package main

import (
	"fmt"
)

// 遍历字符串，将空格替换成%20。时间复杂度O(n)，空间复杂度O(1)
func replaceSpace(s string) string {
	// 由于golang中string底层通过byte实现(ASCII码)，这时候就需要rune类型来处理Unicode字符
	var result []rune
	for _, v := range s{  // 此时遍历的v是ASCII字符
		if v == 32 {  // 32在ASCII码代表空格，如果是空格则替换
			result = append(result, 37, 50, 48)
		}else{
			// 这里需要用rune来处理Unicode字符，如果result是byte类型则无法处理
			result = append(result, v)
		}
	}
	// 最后再将rune转成string类型
	return string(result)
}

func main() {
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println(result)
}
