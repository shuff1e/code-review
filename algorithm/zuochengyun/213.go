package main

import "fmt"

// 最长公共子串
// 动态规划递推方程
// dp[i][j]表示以a[i]和b[j]结尾的最长公共子串的长度
// dp[i][j] = dp[i-1][j-1] + 1 (a[i]==b[j])
// dp[i][j] = 0 (a[i]!=b[j])

func maxCommonSequence(a,b string) [][]int {

	dp := make([][]int,len(a))
	for i := 0;i<len(a);i++ {
		dp[i] = make([]int,len(b))
	}

	for i :=0 ;i<len(a);i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	for j := 0;j<len(b);j++ {
		if b[j] == a[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}

	for i:=1;i<len(a);i++ {
		for j:=1;j<len(b);j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return dp
}

func lcst(a,b string) string {
	if a == "" || b == "" {
		return ""
	}
	dp := maxCommonSequence(a,b)
	max,end := dealDP(dp)
	return a[end-max+1:end+1]
}

func dealDP(dp [][]int) (int,int) {
	max := 0
	end := 0
	for i:=0;i<len(dp);i++ {
		for j:=0;j<len(dp[i]);j++ {
			if dp[i][j] > max {
				max = dp[i][j]
				end = i
			}
		}
	}
	return max,end
}

func main() {
	a := "abcde"
	b := "bebcd"
	result := lcst(a,b)
	fmt.Println(result)
}
