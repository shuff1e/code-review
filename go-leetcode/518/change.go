package main

/*

518. 零钱兑换 II
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
示例 2:

输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。
示例 3:

输入: amount = 10, coins = [10]
输出: 1


注意:

你可以假设：

0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数

 */

// dp[i][j] = dp[i-1][j] + dp[i][j-arr[i]]

func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	dp := make([][]int,len(coins))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,amount+1)
	}

	for i := 0;i<len(dp);i++ {
		dp[i][0] = 1
	}

	for j:=1;j<len(dp[0]);j++ {
		if j >= coins[0] {
			dp[0][j] = dp[0][j-coins[0]]
		}
	}

	for i := 1;i<len(dp);i++ {
		for j := 1;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coins[i] {
				dp[i][j] += dp[i][j-coins[i]]
			}
		}
	}
	return dp[len(dp)-1][amount]
}