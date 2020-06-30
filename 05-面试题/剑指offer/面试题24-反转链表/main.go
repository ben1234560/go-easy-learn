package main

import "fmt"

// 定义结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针方法
//func reverseList(head *ListNode) *ListNode {
//	if head == nil{
//		return nil
//	}
//	var cur *ListNode  // 定义一个正常的
//	pre := head  // 定义一个先往前走的
//	for pre != nil{
//		next := pre.Next  // 把pre下一个提出来，后面赋予表示往前走；next=2，3，4，5，nil
//		pre.Next = cur  // 把pre的下一个置空；pre.Next=nil，1，2，3，4
//		cur = pre  // 把pre现在赋予cur；cur=1，2，3，4，5
//		pre = next  // 把提出来的下一个赋予给pre；pre=2，3，4，5，nil
//	}
//	return cur
//}
// 妖魔化的双指针
func reverseList(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	cur := head  // 定义一个目前的
	for head.Next != nil{
		nNext := head.Next.Next  // 下下个提取出来，nNext=3，4，5，nil
		head.Next.Next = cur  // 目前的复制给下下个，head.Next.Next=1，2，3，4，5
		cur = head.Next  // 下一个复制给cur，表示往前走，cur=2，3，4，5，nil
		head.Next = nNext  // 下下个复制给下一个，head.Next=3，4，5，nil
	}
	return cur
}

func main() {
	// 测试
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
	result := reverseList(head1)
	var res []int // 用一个数组来接受结果
	// 从头遍历并添加到数组中，直到为空
	for result != nil {
		res = append(res, result.Val)
		result = result.Next
	}
	fmt.Println(res)
}
