package main

import "fmt"

/*
115. 不同的子序列
给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。

一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

题目数据保证答案符合 32 位带符号整数范围。



示例 1：

输入：S = "rabbbit", T = "rabbit"
输出：3
解释：

如下图所示, 有 3 种可以从 S 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)

rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
示例 2：

输入：S = "babgbag", T = "bag"
输出：5
解释：

如下图所示, 有 5 种可以从 S 中得到 "bag" 的方案。
(上箭头符号 ^ 表示选取的字母)

babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
^  ^^
babgbag
^^^

 */

/*
rabbbit
rabbit

str1[i] == str2[j]
dp[i][j] = dp[i+1][j+1] + dp[i+1][j]

str1[i] != str2[j]
dp[i][j] = dp[i+1][j]

 */

func main() {
	S := "rabbbit"; T := "rabbit"
	S = "babgbag";T = "bag"
	S = "rabbbit";T = "rabbit"
	fmt.Println(numDistinct(S,T))
}

func numDistinct(s string, t string) int {
	memo := make([][]int,len(s)+1)
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(t)+1)
		for j := 0;j<len(t)+1;j++ {
			memo[i][j] = -1
		}
	}
	return help(s,t,0,0,memo)
}

// memo!!!

func help(s string,t string,sIndex ,tIndex int,memo [][]int) int {
	// t没消耗完但是s消耗完了
	if sIndex == len(s) {
		if tIndex < len(t) {
			return 0
		} else {
			return 1
		}
	}
	// t消耗完了，s消耗完了 或者没消耗完
	if tIndex == len(t) {
		if sIndex <= len(s) {
			return 1
		} else {
			return 0
		}
	}

	if memo[sIndex][tIndex] != -1 {
		return memo[sIndex][tIndex]
	}

	if s[sIndex] == t[tIndex] {
		result :=  help(s,t,sIndex + 1,tIndex + 1,memo) +
			help(s,t,sIndex + 1,tIndex,memo)
		memo[sIndex][tIndex] = result
		return result
	} else {
		result := help(s,t,sIndex+1,tIndex,memo)
		memo[sIndex][tIndex] = result
		return result
	}
}