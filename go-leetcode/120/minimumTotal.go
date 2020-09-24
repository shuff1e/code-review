package main

/*
120. 三角形最小路径和
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。

 */

// A：因为是求最值，因此可以不用回溯到最后，在某些位置可以缓存当前的最值，就是递归+memo
// 可以改成dp，然后dp的空间应该是n*n
// 不过有些情况下，dp可以进一步有华为空间为n的

// dp[i][j] = min(dp[i+1][j],dp[i+1][j+1]) + matrix[i][j]
// dp[len(matrix)-1][j] = matrix[len(matrix)-1][j]

// 只与右边的有关系，因此可以从左到右的计算

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([]int,len(triangle))
	for i := 0;i<len(dp);i++ {
		dp[i] = triangle[len(triangle)-1][i]
	}

	for i := len(triangle) - 2;i>=0;i-- {
		for j := 0;j<i+1;j++ {
			dp[j] = Min(dp[j],dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func minimumTotalBad(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	dp := make([][]int,len(triangle))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,i+1)
	}
	for j := 0;j<len(triangle);j++ {
		dp[len(triangle)-1][j] = triangle[len(triangle)-1][j]
	}

	for i := len(triangle) - 2;i>=0;i-- {
		for j := 0;j<len(dp[i]);j++ {
			dp[i][j] = Min(dp[i+1][j],dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}