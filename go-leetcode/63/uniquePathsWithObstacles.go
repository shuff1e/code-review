package main

/*
63. 不同路径 II
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

网格中的障碍物和空位置分别用 1 和 0 来表示。

说明：m 和 n 的值均不超过 100。

示例 1:

输入:
[
[0,0,0],
[0,1,0],
[0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
 */

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	visited := make([][]bool,len(obstacleGrid))
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,len(obstacleGrid[0]))
	}
	memo := make([][]int,len(obstacleGrid))
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(obstacleGrid[0]))
		for j := 0;j<len(memo[i]);j++ {
			memo[i][j] = -1
		}
	}

	return help(len(obstacleGrid),len(obstacleGrid[0]),0,0,obstacleGrid,visited,memo)
}

func help(m,n,row,col int,obstacleGrid [][]int,visited [][]bool,memo [][]int) int {
	if row >= m || col >= n {
		return 0
	}
	if visited[row][col] {
		return 0
	}
	if obstacleGrid[row][col] == 1 {
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
	result1 := help(m,n,row+1,col,obstacleGrid,visited,memo)
	// 向右
	result2 := help(m,n,row,col+1,obstacleGrid,visited,memo)
	visited[row][col] = false
	memo[row][col] = result1 + result2

	return result1 + result2
}
