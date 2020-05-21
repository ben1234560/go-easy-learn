# README

#### 面试题07. 重建二叉树

> 难度：简单

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

**示例:**

~~~go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
}
~~~

例如：给出

~~~
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
~~~

返回如下的二叉树：

~~~
    3
   / \
  9  20
    /  \
   15   7
~~~

**限制：**

```
0 <= 节点个数 <= 5000
```



**相关知识点：**

如果你不理解二叉树，你需要[深入学习二叉树](https://www.baidu.com/link?url=zQMCO6fYcRpt14jTaFvwfEfFtu-CBiuHzuXfURKnYnYsVEUmdKWruZovtptRl9fa&wd=&eqid=8ef4bae60038f923000000045ec486c0)；

总得来说：二叉树的访问次序可以分为四种：

~~~
前序遍历：父节点->左子树->右子树
中序遍历：左子树->父节点->右子树
后序遍历：左子树 -> 右子树 ->父节点
层序遍历：从左至右遍历
~~~

> 都是从上至下