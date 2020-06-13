package main

import "fmt"

// 左右指针互换，
func exchange(nums []int) []int {
	l := 0             //左指针-队首
	r := len(nums) - 1 //右指针-队尾
	for l < r {
		// 如果左指针是偶数，且右指针是奇数，则两边调换
		if nums[l]%2 == 0 && nums[r]%2 != 0 {
			nums[l], nums[r] = nums[r], nums[l]
		}
		// 如果左指针（队首）是奇数，则l+1，注意把边界做好防止越界
		for l < len(nums) && nums[l]%2 != 0 {
			l++
		}
		// 如果右指针（队尾）是偶数，则r-1，注意把边界做好防止越界
		for r >= 0 && nums[r]%2 == 0 {
			r--
		}
	}
	return nums
}

func main() {
	// 测试
	var nums = []int{1, 3, 5}
	res := exchange(nums)
	fmt.Println(res)
}
