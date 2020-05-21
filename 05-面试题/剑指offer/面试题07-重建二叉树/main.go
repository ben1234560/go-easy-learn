package main

import (
"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		if inorder[k] == preorder[0] { // 遍历，利用中序和前序遍历的特性，中序[1]==前序[0]，也就是当前的根节点
			return &TreeNode{
				Val:   preorder[0],                              // 根为前序的第一个值
				Left:  buildTree(preorder[1:k+1], inorder[0:k]), // 左节点传入前序[9],中序[9]；第二次遍历得到9后则无遍历
				Right: buildTree(preorder[k+1:], inorder[k+1:]), // 右节点传入前序[20,15,7]，中序[15,20,7]；同样需要第三次遍历才能得到20=20
			}
		}
	}
	return nil
}

var res []int  // 用来存放结果值
// 前序遍历函数：先遍历根节点再遍历左子树，再遍历右子树
func preOrder(current *TreeNode) {
	if current == nil{
		return
	}
	res = append(res, current.Val)
	preOrder(current.Left)
	preOrder(current.Right)
}

// 中序遍历:先遍历左子树再遍历根，再遍历根节点再遍历右子树
func inOrder(current *TreeNode) {
	if current == nil{
		return
	}
	inOrder(current.Left)
	res = append(res, current.Val)
	inOrder(current.Right)
}

// 后序遍历:先遍历左子树再遍历根，再遍历右子树，再遍历根节点
func afterOrder(current *TreeNode) {
	if current == nil{
		return
	}
	afterOrder(current.Left)
	afterOrder(current.Right)
	res = append(res, current.Val)
}

func main() {
	// 制作数组
	var preorder = []int{3, 9, 20, 15, 7}
	var inorder = []int{9, 3, 15, 20, 7}
	result := buildTree(preorder, inorder)
	preOrder(result)
	fmt.Println(res)
}
