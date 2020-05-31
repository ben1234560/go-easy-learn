package main

import "fmt"

// 二分法
// 时间复杂度 O(log_2 N)，空间复杂度 O(1)
func minArray(numbers []int) int {
	left,right :=0,len(numbers)-1  // 获取左右两个端点
	for left<right{  // 一直循环直到重合，即找到最小值
		mid :=(left+right)/2  //
		if numbers[mid]>numbers[right]{  // 表示最小值在右边，所以左端点往右移
			left = mid+1
		}else if numbers[mid]<numbers[right] {  //表示最小值肯定在m的位置，所以右端点移到m位
			right = mid
		}else{  // 表示最小值在左边，所以右端点往左移
			right--
		}
	}
	return numbers[left]  // 最终left和right会一致，取出即可
}

func main() {
	// 测试
	numbers := []int{1, 3, 5, 6}
	res := minArray(numbers)
	fmt.Println(res)
}
