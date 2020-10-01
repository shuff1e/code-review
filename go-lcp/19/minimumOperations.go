package main

/*
LCP 19. 秋叶收藏集
小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves，
字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。
出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。每部分树叶数量可以不相等，但均需大于等于 1。
每次调整操作，小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。

示例 1：

输入：leaves = "rrryyyrryyyrr"

输出：2

解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"

示例 2：

输入：leaves = "ryr"

输出：0

解释：已符合要求，不需要额外操作

提示：

3 <= leaves.length <= 10^5
leaves 中只包含字符 'r' 和字符 'y'

 */

// 要排列成红黄红
// 设上述三种状态分别为0,1,2

// dp[i][j]表示将i位置设置为j状态，并且合法，需要的调整次数

// dp[i][0] = dp[i-1][0] + isYellow
//
// dp[i][1] = min(dp[i-1][1],dp[i-1][0]) + isRed
//
// dp[i][2] = min(dp[i-1][1],dp[i-1][2]) + isYellow
//
// dp[0][0] = isYellow
// dp[0][1] = dp[0][2] = dp[1][2] = INT_MAX

// return dp[len(arr)-1][2]

func minimumOperations(leaves string) int {
	const INT_MAX = 0x7fffffff
	dp := make([][]int,len(leaves))
	for i := 0;i<len(leaves);i++ {
		dp[i] = make([]int,3)
	}
	dp[0][1],dp[0][2],dp[1][2] = INT_MAX,INT_MAX,INT_MAX
	if leaves[0] == 'y' {
		dp[0][0] = 1
	}
	for i := 1;i<len(leaves);i++ {
		isRed,isYellow := 0,0
		if leaves[i] == 'r' {
			isRed = 1
		} else if leaves[i] == 'y' {
			isYellow = 1
		}
		dp[i][0] = dp[i-1][0] + isYellow
		dp[i][1] = min(dp[i-1][1],dp[i-1][0]) + isRed
		if i >= 2 {
			dp[i][2] = min(dp[i-1][1],dp[i-1][2]) + isYellow
		}
	}
	return dp[len(leaves)-1][2]

}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}