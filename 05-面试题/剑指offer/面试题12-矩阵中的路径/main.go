package main

import "fmt"

// 解题思路，把所有的可能性进行判断，false和true
// 时间复杂度0(n2)
func dfs(board [][]byte,word string,i,j,k int)bool{
	// i表示行长度、j表示列长度、k表示字符索引
	// 首先判断如果i小于0 或 i大于等于实际行长度 或 j大于实际列长度 或 board字符串不等于word里面的字符串
	if i<0||i>=len(board)||j<0||j>=len(board[0])||board[i][j]!=word[k]{
		return false
	}
	// 因为k是从0开始，如果k的长度等同于word-1的长度，则表示已全部找到
	if k==len(word)-1{
		return true
	}
	// 把遍历的字符串存起来，然后用/代表已经遍历过，无法再遍历
	tmp:=board[i][j]
	board[i][j] = '/'
	// 如果满足一下任意一个则返回true，右移动、左移动、下移动、上移动
	// 第一个if已经用 || 限制了很多且是返回false，所以如果下面的任意位置移动不能返回true，则表示没有对应路径的结果
	res := dfs(board,word,i,j+1,k+1)||dfs(board,word,i,j-1,k+1)||dfs(board,word,i+1,j,k+1)||dfs(board,word,i-1,j,k+1)
	// 最后再把原来存起来的再放回去
	board[i][j] = tmp
	return res
}

func exist(board [][]byte, word string) bool {
	for i:=0;i<len(board);i++{  // 获取board的行的长度并循环，例子的是3，首次i=0
		for j:=0;j<len(board[0]);j++{  // 获取board的列的长度并循环，例子的是4，首次j=0
			if dfs(board,word,i,j,0){  // 如果dfs正常返回，则成功
				return true
			}
		}
	}
	return false
}

func main() {
	// 测试
	sli := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'},{'A','D','C','E'}}
	word := "ABCCEEE"
	res := exist(sli, word)
	fmt.Println(res)
}
