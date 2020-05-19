package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

// 解题思路：循环进行两端置换，两端同时往中间靠拢，继续循环
func reversePrint(head *ListNode) []int {
	if head == nil {  // 如果没有则直接返回
		return nil
	}

	var res []int  // 用一个数组来接受结果
	// 先把所有值按选后顺序放进来，直到为空
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
		// 循环完后res的内容是[1 3 2]
	}

	// 两端置换，并同时往中间靠拢，循环；直到两端都靠拢到最中间
	for i, j := 0, len(res)-1; i < j; {  // i=0，j=2，开始循环
		res[i], res[j] = res[j], res[i]  // 两端互换，即1，2 = 2， 1
		// 两端同时往中间靠拢
		i++
		j--
	}
	return res
}



func main() {
	// 制作链表
	var head1 = &ListNode{
		Val:2,
		Next:nil,
	}
	var head2 = &ListNode{
		Val:  3,
		Next: head1,
	}
	var head3 = &ListNode{
		Val:  1,
		Next: head2,
	}
	result := reversePrint(head3)
	fmt.Println(result)
}
