package main

import "fmt"

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回 s 所有可能的分割方案。

示例:

输入: "aab"
输出:
[
["aa","b"],
["a","a","b"]
]

 */

/*
aab
dp[i][i] = 1
dp[0][1] = 1 str[0] == str[j]
dp[1][2] = 0

// 右上角
dp[i-1][j+1] = dp[i][j] && (str[i] == str[j])

生成矩阵之后
for 以i开头的dp为1的
	dp[i][j] = 1 ，start从j
	dp[i][j+1] = 1,start从j+1往后递归
 */

func main() {
	str := "aab"
	result := partition(str)
	fmt.Printf("%#v\n",result)
}

func printMatrix(dp [][]int) {
	for _,v := range dp {
		fmt.Printf("%#v\n",v)
	}
}

func partition(s string) [][]string {
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
	result := [][]string{}
	temp := []string{}
	help(s,0,&result,dp,temp)
	return result
}

func help(s string,start int,result *[][]string,dp [][]int,temp []string) {
	if start == len(s) {
		temp2 := make([]string,len(temp))
		copy(temp2,temp)
		*result = append(*result,temp2)
	}

	for j := start;j<len(s);j++ {
		if dp[start][j] == 1 {
			temp = append(temp,s[start:j+1])
			help(s,j+1,result,dp,temp)
			temp = temp[0:len(temp)-1]
		}
	}
}