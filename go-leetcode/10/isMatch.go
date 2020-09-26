package main

import "fmt"

/*
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

 */

//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
// if j+1 == '*' ,匹配0个，dp[i][j+2]
//   if p[j] == '.'，匹配一个或者多个,dp[i+1][j]
//   if p[j] == s[i]，匹配一个或者多个

// if p[j] != s[i]，匹配0个，dp[i][j] = false

// 初始条件
// if p[len(p)-1] == '*'
//    if p[len(p)-2] == s[len(s)-1] 或者 '.' dp[len(s)-1][len(p)-1] = true,dp[len(s)-1][len(p)-2] = true
//    dp[i][len(p)-1] = (s[i] == p[len(p)-2] 或者 p[len(p)-2] == '.') && dp[i+1][len(p)-1]

//    dp[len(s)-1][j] =

func main() {
	s := "mississippi"
	p := "mis*is*p*."
	s = ""
	p = ".*"
	s = ""
	p = "c*c*"
	fmt.Println(isMatch(s,p))
}

func isMatch(s string, p string) bool {
	memo := make([][]int,len(s))
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(p))
	}
	result := help(s,p,0,0,memo)
	return result == 1
}

func help(s,p string,start1,start2 int,memo [][]int) int {
	if start1 >= len(s) && start2 >= len(p) {
		return 1
	}
	if start1 == len(s) {
		if start2 + 1 < len(p) && p[start2+1] == '*' {
			return help(s,p,start1,start2+2,memo)
		}
		return -1
	}
	if start1 < len(s) && start2 == len(p) {
		return -1
	}

	if memo[start1][start2] != 0 {
		return memo[start1][start2]
	}

	if start2 + 1 < len(p) && p[start2+1] == '*' {
		// if j+1 == '*' ,匹配0个，dp[i][j+2]
		result := help(s,p,start1,start2+2,memo)
		if result == 1 {
			memo[start1][start2] = 1
			return 1
		}
		//   if p[j] == '.'，匹配一个或者多个,dp[i+1][j]
		//   if p[j] == s[i]，匹配一个或者多个
		if p[start2] == '.' || p[start2] == s[start1] {
			result = help(s,p,start1+1,start2,memo)
			if result == 1 {
				memo[start1][start2] = 1
				return 1
			}
		}
	}
	if p[start2] == '.' || p[start2] == s[start1] {
		result := help(s,p,start1+1,start2+1,memo)
		if result == 1 {
			memo[start1][start2] = 1
			return 1
		}
	}
	memo[start1][start2] = -1
	return -1
}