package main

import "fmt"

/*

279. 完全平方数
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

 */

func main() {
	n := 0
	fmt.Println(numSquares(n))
}

func numSquares(n int) int {
	index := 1
	arr := []int{}
	for index * index <= n {
		arr = append(arr,index*index)
		index ++
	}
	return help(arr,n)
}

// dp[i][j] = Min(dp[i-1][j],dp[i][j-arr[i]]+1)

func help(arr []int,amount int) int {
	const INTEGER_MAX = 0x7fffffff

	dp := make([]int,amount+1)
	dp[0] = 0

	for j := 1;j<len(dp);j++ {
		dp[j] = INTEGER_MAX
		if dp[j-arr[0]] != INTEGER_MAX {
			dp[j] = dp[j-arr[0]] + 1
		}
	}

	for i := 1;i<len(arr);i++ {
		for j:=arr[i];j<len(dp);j++ {
			if dp[j-arr[i]] != INTEGER_MAX {
				dp[j] = Min(dp[j],dp[j-arr[i]]+1)
			}
		}
	}
	return dp[amount]
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}