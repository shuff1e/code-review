package main

import "fmt"

/*
5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

 */

// A：dp[i][j] = dp[i+1][j-1] + 2 if s[i] == s[j] && dp[i+1][j-1] > 0
//  else 0

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
}
func longestPalindrome(s string) string {
	result := ""
	dp := make([][]int,len(s))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(s))
	}

	for i := 0;i<len(dp);i++ {
		dp[i][i] = 1
		result = s[i:i+1]
	}
	for i :=0;i+1<len(dp);i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
			result = s[i:i+2]
		} else {
			dp[i][i+1] = 0
		}
	}

	for length := 2;length < len(s);length ++ {
		for i := 0;i+length<len(s);i++ {
			if s[i] == s[i+length] && dp[i+1][i+length-1] > 0 {
				dp[i][i+length] = dp[i+1][i+length-1] + 2
				result = s[i:i+length+1]
			} else {
				dp[i][i+length] = 0
			}
		}
	}
	return result
}
