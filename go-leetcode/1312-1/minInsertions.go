package main

import "fmt"

/*
1312-1. 让字符串成为回文串的最少插入次数
给你一个字符串 s ，每一次操作你都可以在字符串的任意位置插入任意字符。

请你返回让 s 成为回文串的 最少操作次数 。

「回文串」是正读和反读都相同的字符串。



示例 1：

输入：s = "zzazz"
输出：0
解释：字符串 "zzazz" 已经是回文串了，所以不需要做任何插入操作。
示例 2：

输入：s = "mbadm"
输出：2
解释：字符串可变为 "mbdadbm" 或者 "mdbabdm" 。
示例 3：

输入：s = "leetcode"
输出：5
解释：插入 5 个字符后字符串变为 "leetcodocteel" 。
示例 4：

输入：s = "g"
输出：0
示例 5：

输入：s = "no"
输出：1


提示：

1 <= s.length <= 500
s 中所有字符都是小写字母。
 */

/*
A：对于leetcode

leetcodocteel

如果str[i] = str[j]
dp[i][j] = dp[i+1][j-1]

如果str[i] != str[j]
dp[i][j] = dp[i+1][j] + arr[i], leetcodel
dp[i][j] = dp[i][j-1] + arr[j], eleetcode

dp[i][i]=0
 */

func main() {
	str := "mbadm"
	fmt.Println(minInsertions(str))
}

func minInsertions(s string) int {
	dp := make([][]int,len(s))
	for i := 0;i<len(s);i++ {
		dp[i] = make([]int,len(s))
	}
	for i := 0;i<len(s);i++ {
		dp[i][i] = 0
	}
	for length := 1;length<len(s);length++ {
		for i := 0;i<len(s);i++ {
			if i + length < len(s) {
				if s[i] == s[i+length] {
					dp[i][i+length] = dp[i+1][i+length-1]
				} else {
					dp[i][i+length] = Min(dp[i+1][i+length],dp[i][i+length-1])+1
				}
			}
		}
	}

	return dp[0][len(s)-1]

}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}