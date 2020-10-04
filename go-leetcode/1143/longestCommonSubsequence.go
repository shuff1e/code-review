package main

/*
1143. 最长公共子序列
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

一个字符串的 子序列 是指这样一个新的字符串：
它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。

若这两个字符串没有公共子序列，则返回 0。

示例 1:

输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。
示例 2:

输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。
示例 3:

输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。


提示:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。

 */

// A：以str1[i],str2[j]结尾的公共子序列

// dp[i][j] = dp[i-1][j-1] + 1 if str1[i] == str2[j]
// dp[i][j] = max(dp[i-1][j],dp[i][j-1] if str1[i] != str2[j]

// dp[i][0] = 1 或者 0
//
// 	str1 := "1A2C3D4B56"
//	str2 := "B1D123CA45B6A"
//

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int,len(text1))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(text2))
	}

	if text1[0] == text2[0] {
		dp[0][0] = 1
	}

	for i := 1;i<len(text1);i++ {
		if text1[i] == text2[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1;j<len(text2);j++ {
		if text1[0] == text2[j] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1;i<len(text1);i++ {
		for j := 1;j<len(text2);j++ {
			dp[i][j] = Max(dp[i-1][j],dp[i][j-1])
			if text1[i] == text2[j] {
				dp[i][j] = Max(dp[i-1][j-1] + 1,dp[i][j])
			}
		}
	}

	return dp[len(text1)-1][len(text2)-1]
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}