package main

import "fmt"

func movingCount(m int, n int, k int) int {
	// 获取机器人最大的行动范围
	flag := make([][]int, m+1)
	for i := range flag {
		flag[i] = make([]int, n+1)
	}
	return dfs(m, n, 0, 0, k, flag)
}

func dfs(m, n, i, j, k int, flag [][]int) int {
	// 判断如果条件不符合，则退出
	if i < 0 || j < 0 || i >= m || j >= n || flag[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	// 用来标识机器人已经走过的地方
	flag[i][j] = 1
	// 递归自己的深度，遍历所有可能性，每次成功一个位置就+1，最终得出能够达到多少个格子
	return 1 + dfs(m, n, i, j+1, k, flag) + dfs(m, n, i, j-1, k, flag) + dfs(m, n, i+1, j, k, flag) + dfs(m, n, i-1, j, k, flag)
}

// 求所有位之和
func sumPos(n int) int {
	var sum int
	for n != 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

func main() {
	// 测试
	m, n, k := 1, 2, 1
	count := movingCount(m, n, k)
	fmt.Println(count)
}
