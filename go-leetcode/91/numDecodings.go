package main

import "fmt"

/*
91. 解码方法
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

题目数据保证答案肯定是一个 32 位的整数。

示例 1：

输入："12"
输出：2
解释：它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2：

输入："226"
输出：3
解释：它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。
示例 3：

输入：s = "0"
输出：0
示例 4：

输入：s = "1"
输出：1
示例 5：

输入：s = "2"
输出：1

提示：

1 <= s.length <= 100
s 只包含数字，并且可以包含前导零。
 */

// dp[i] = dp[i+1] '0' < str[i] < '9'
// continue			str[i] == '0'

// dp[i] = dp[i+2]  "10" <= str[i:i+2] <= "26"
// else if str[i+1] == '0' && (str[i] > '2' || str[i] =='0') return 0

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("1"))
	fmt.Println(numDecodings("2"))
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}

	dp := make([]int,len(s)+1)
	dp[len(s)] = 1

	for i := len(s)-1;i>=0;i-- {
		if i+1 < len(s) {
			if s[i:i+2] <="26" && s[i:i+2] >= "10" {
				dp[i] = dp[i+2]
			} else if s[i+1] == '0' && (s[i] > '2' || s[i] =='0') {
				return 0
			}
		}

		if s[i] > '0' {
			dp[i] += dp[i+1]
		} else if s[i] == '0' {
			continue
		}
	}
	return dp[0]
}