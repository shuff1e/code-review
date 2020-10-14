package main

import "fmt"

/*
1277. 统计全为 1 的正方形子矩阵
给你一个 m * n 的矩阵，矩阵中的元素不是 0 就是 1，请你统计并返回其中完全由 1 组成的 正方形 子矩阵的个数。

示例 1：

输入：matrix =
[
[0,1,1,1],
[1,1,1,1],
[0,1,1,1]
]
输出：15
解释：
边长为 1 的正方形有 10 个。
边长为 2 的正方形有 4 个。
边长为 3 的正方形有 1 个。
正方形的总数 = 10 + 4 + 1 = 15.
示例 2：

输入：matrix =
[
[1,0,1],
[1,1,0],
[1,1,0]
]
输出：7
解释：
边长为 1 的正方形有 6 个。
边长为 2 的正方形有 1 个。
正方形的总数 = 6 + 1 = 7.


提示：

1 <= arr.length <= 300
1 <= arr[0].length <= 300
0 <= arr[i][j] <= 1
 */

/*

我们用 f[i][j] 表示以 (i, j) 为右下角的正方形的最大边长，
那么除此定义之外，f[i][j] = x 也表示以 (i, j) 为右下角的正方形的数目为 x（即边长为 1, 2, ..., x 的正方形各一个）。
在计算出所有的 f[i][j] 后，我们将它们进行累加，就可以得到矩阵中正方形的数目。

我们尝试挖掘 f[i][j] 与相邻位置的关系来计算出 f[i][j] 的值。

dp[i][j] = Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1

 */

func main() {
	matrix := [][]int{
		{1,0,1},
		{1,1,0},
		{1,1,0}}
	result := countSquares(matrix)
	fmt.Println(result)
}

func countSquares(matrix [][]int) int {
	result := 0

	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int,m)
	for i := 0;i<m;i++ {
		dp[i] = make([]int,n)
	}

	for i := 0;i<m;i++ {
		dp[i][0] = matrix[i][0]
		result += dp[i][0]
	}
	for j := 1;j<n;j++ {
		dp[0][j] = matrix[0][j]
		result += dp[0][j]
	}

	for i := 1;i<m;i++ {
		for j := 1;j<n;j++ {
			if matrix[i][j] != 0 {
				dp[i][j] = Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1
				result += dp[i][j]
			}
		}
	}
	return result
}

func Min(args ...int) int {
	result := args[0]
	for i:=1;i<len(args);i++ {
		if result > args[i] {
			result = args[i]
		}
	}
	return result
}