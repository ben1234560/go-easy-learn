package main

import "fmt"

// 定义结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组索引
//func getKthFromEnd(head *ListNode, k int) *ListNode {
//	var res []*ListNode // 用来存放结果
//	for head != nil {
//		res = append(res, head)
//		head = head.Next
//	}
//	l := len(res) // 获取长度
//	if l >= k && k > 0 { // 如果有值，则把多余部分去掉即可，并处理k为0的情况
//		return res[l-k]
//	}
//	return nil
//}

// 快慢双指针, 快指针是全部,k为快慢指针的距离,慢指针为结果
func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast = head // 先走指针,以K为距离,先往尾部走
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	var slow = head //后走指针,以fast长度走,其中fast已经走了k距离
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	// 测试
	var k = 2
	var head5 = &ListNode{
		Val:  5,
		Next: nil,
	}
	var head4 = &ListNode{
		Val:  4,
		Next: head5,
	}
	var head3 = &ListNode{
		Val:  3,
		Next: head4,
	}
	var head2 = &ListNode{
		Val:  2,
		Next: head3,
	}
	var head1 = &ListNode{
		Val:  1,
		Next: head2,
	}
	result := getKthFromEnd(head1, k)
	var res []int // 用一个数组来接受结果
	// 从头遍历并添加到数组中，直到为空
	for result != nil {
		res = append(res, result.Val)
		result = result.Next
	}
	fmt.Println(res)
}
