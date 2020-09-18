package main

import "fmt"

// 最长公共子序列
// 动态规划递推关系
// dp[i][j]是以a[i],b[j]为结尾的最长公共子序列的长度
// dp[i][j] = max(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]+1)

func getDP(a,b string) [][]int {

	dp := make([][]int,len(a))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(b))
	}

	dp[0][0] = 0
	if a[0] == b[0] {
		dp[0][0] = 1
	}

	for i:=1;i<len(a);i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j:=1;j<len(b);j++ {
		if b[j] == a[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1;i<len(a);i++ {
		for j:=1;j<len(b);j++ {
			dp[i][j] = Max(dp[i-1][j],dp[i][j-1])
			if a[i] == b[j] {
				dp[i][j] = Max(dp[i][j],dp[i-1][j-1]+1)
			}
		}
	}

	return dp
}

func getStr(a,b string,dp [][]int) string {
	m := len(a)-1
	n := len(b)-1
	max := dp[m][n]
	result := make([]byte,max)
	index := max -1
	for index >= 0 {
		if m >0 && dp[m][n] == dp[m-1][n] {
			m = m-1
		} else if n >0 && dp[m][n] == dp[m][n-1] {
			n = n-1
		} else {
			result[index] = a[m]
			index = index -1
			m = m-1
			n = n-1
		}
	}
	return string(result)
}

func Max(x,y int) int{
	if x>y {
		return x
	}
	return y
}

func main() {
	str1 := "1A2C3D4B56"
	str2 := "B1D123CA45B6A"
	dp := getDP(str1,str2)
	fmt.Println(dp)
	fmt.Println(getStr(str1,str2,dp))
}