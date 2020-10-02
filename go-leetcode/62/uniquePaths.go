package main

/*
62. 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？



例如，上图是一个7 x 3 的网格。有多少可能的路径？



示例 1:

1 1 1
1 1 1

输入: m = 3, n = 2
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向右 -> 向下
2. 向右 -> 向下 -> 向右
3. 向下 -> 向右 -> 向右
示例 2:

输入: m = 7, n = 3
输出: 28


提示：

1 <= m, n <= 100
题目数据保证答案小于等于 2 * 10 ^ 9

 */

// 动态方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
//
// 注意，对于第一行 dp[0][j]，或者第一列 dp[i][0]，由于都是在边界，所以只能为 1
//

func uniquePaths(m int, n int) int {
	visited := make([][]bool,m)
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,n)
	}
	memo := make([][]int,m)
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,n)
		for j := 0;j<len(memo[i]);j++ {
			memo[i][j] = -1
		}
	}

	return help(m,n,0,0,visited,memo)
}

func help(m,n,row,col int,visited [][]bool,memo [][]int) int {
	if row >= m || col >= n {
		return 0
	}
	if visited[row][col] {
		return 0
	}
	if row == m-1&& col == n-1 {
		return 1
	}
	if memo[row][col] != -1 {
		return memo[row][col]
	}

	visited[row][col] = true
	// 向下
	result1 := help(m,n,row+1,col,visited,memo)
	// 向右
	result2 := help(m,n,row,col+1,visited,memo)
	visited[row][col] = false
	memo[row][col] = result1 + result2

	return result1 + result2
}