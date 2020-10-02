package main

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
[1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
 */

// dp[i][j] = min(dp[i+1][j],dp[i][j+1])+matrix[i][j]

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make([][]int,len(grid))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(grid[0]))
	}
	dp[len(grid)-1][len(grid[0])-1] = grid[len(grid)-1][len(grid[0])-1]
	for i := len(grid)-2;i>=0;i-- {
		dp[i][len(grid[0])-1] = dp[i+1][len(grid[0])-1] + grid[i][len(grid[0])-1]
	}
	for j := len(grid[0])-2;j>=0;j-- {
		dp[len(grid)-1][j] = dp[len(grid)-1][j+1] + grid[len(grid)-1][j]
	}

	for i := len(grid)-2;i>=0;i-- {
		for j := len(grid[0])-2;j>=0;j-- {
			dp[i][j] = min(dp[i+1][j],dp[i][j+1])+grid[i][j]
		}
	}
	return dp[0][0]
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}