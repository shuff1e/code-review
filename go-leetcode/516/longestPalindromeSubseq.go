package main

import "fmt"

func longestPalindromeSubseq(s string) int {

	dp := make([][]int,len(s))

	for i := range dp {
		dp[i] = make([]int,len(s))
	}

	for i:= 0;i<len(s);i++ {
		dp[i][i] = 1
	}
	for i := 0;i<len(s)-1;i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 0
		}
	}

	for length := 2 ;length < len(s);length++ {
		for i := 0;i+length < len(s);i++ {
			if s[i] == s[i+length] {
				dp[i][i+length] = dp[i+1][i+length-1] + 2
			} else {
				dp[i][i+length] = Max(dp[i][i+length-1],dp[i+1][i+length])
			}
		}
	}
	//fmt.Printf("%#v\n",dp)

	return dp[0][len(s)-1]
}

func Max(x,y int) int {
	if x < y {
		x = y
	}
	return x
}

// dp[i][j] = max(dp[x][y]) if s[i] == s[j]

func main() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}
