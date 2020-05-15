## 剑指offer

> [出处：LeetCode](https://leetcode-cn.com/)

##### 面试题03. 数组中重复的数字

找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

**示例 1：**

~~~
func findRepeatNumber(nums []int) int {
    
}

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3 
~~~

**限制：**

```
2 <= n <= 100000
```

**解法：**

~~~go
// method_1：利用sort排序，将同一值排到一起，然后循环拿当前的对比下一个。空间复杂度O(nlogn)、时间复杂度O(1)
func findRepeatNumber(nums []int) int {
	sort.Ints(nums)  //排序，相同数字会靠近 [0 1 2 2 3 3 5]
	for idx := range nums{  // 循环nums的下标
		if idx > 0{  // 确保下标是大于0的
			if nums[idx] == nums[idx-1]{  // 判断当前下标值和前一下标值是否相同，相同则返回
				return nums[idx]
			}
		}
	}
	return 0   // 如果没有则返回0，必须是int型和首行对应
}

// method_2：利用map的key作为判断，循环如果有该key则返回。空间复杂度O(n)、时间复杂度O(n)
func findRepeatNumber(nums []int) int {
	maps := make(map[int]bool)  // 建map，以值为key，bool为value
	for _,val:=range nums{
		if maps[val]{	// 利用map的key，判断是否已存在了，如果存在就返回
			return val
		}else{
			maps[val]=true  // 否则添加到maps里
		}
	}
	return 0 // 如果没有则返回0，必须是int型和首行对应
}
~~~

