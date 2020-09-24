package main

import "fmt"

/*
132. 分割回文串 II
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的最少分割次数。

示例:

输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

 */

func main() {
	str := "aab"
	result := minCut(str)
	fmt.Printf("%#v\n",result)
}

func printMatrix(dp [][]int) {
	for _,v := range dp {
		fmt.Printf("%#v\n",v)
	}
}

func minCut(s string) int {
	dp := make([][]int,len(s))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(s))
	}
	// init
	for i := 0;i<len(s);i++ {
		dp[i][i] = 1
	}
	for i := 0;i<len(s)-1;i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
		} else {
			dp[i][i+1] = 0
		}
	}

	// dp[i][j] = dp[i+1][j-1] && (str[i] == str[j])
	for length := 2;length<len(s);length ++ {
		for i := 0;i<len(s);i++ {
			// dp[i][i+length]
			if i + 1 <len(s) && i + length < len(s) {
				if s[i] == s[i+length] {
					dp[i][i+length] = dp[i+1][i+length-1]
				} else {
					dp[i][i+length] = 0
				}
			}
		}
	}
	//printMatrix(dp)

	memo := make([]int,len(s))
	for i := 0;i<len(s);i++ {
		memo[i] = -1
	}

	return help(s,0,dp,memo) - 1
}

// 131中这里是回溯，需要到达每一个可能的结果
// 这里只求个数，可以使用memo加速
func help(s string,start int,dp [][]int,memo []int) int {
	if start == len(s) {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}

	curSum := len(s)

	for j := start;j<len(s);j++ {
		if dp[start][j] == 1 {
			result := help(s, j+1, dp, memo)
			curSum = Min(result + 1, curSum)
		}
	}
	memo[start] = curSum
	return curSum
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}