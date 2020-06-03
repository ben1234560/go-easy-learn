package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 思路：循环遍历，直至找到val
func deleteNode(head *ListNode, val int) *ListNode {
	//要删除的节点为头节点，直接返回链表的head.Next
	if head.Val==val{
		return head.Next
	}
	pre:=head   // pre保存头节点
	for head.Next.Val!=val{  // 遍历链表，如果下一个的值是要找的则跳出循环，否则继续赋值往下
		head=head.Next
	}
	head.Next=head.Next.Next  // 跳出循环后表示下一个就是要删除的值，所以把下下一个赋值给下一个
	return pre
}

func main() {
	// 测试
	var head9 = &ListNode{
		Val: 9,
		Next: nil,
	}
	var head1 = &ListNode{
		Val:  1,
		Next: head9,
	}
	var head5 = &ListNode{
		Val: 5,
		Next: head1,
	}
	var head4 = &ListNode{
		Val: 4,
		Next: head5,
	}
	result := deleteNode(head4, 5)
	// 以数组的形式打印结果
	var res []int  // 用一个数组来接受结果
	// 从头遍历并添加到数组中，直到为空
	for result != nil {
		res = append(res, result.Val)
		result = result.Next
	}
	fmt.Println(res)
}
