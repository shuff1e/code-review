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

/*
package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "bbbab"
	s := "cbbd"
	length,dp := longestPalindromeSubseq(s)
	fmt.Println(length)
	fmt.Println(getStrFromDP(s,dp))
}

func longestPalindromeSubseq(s string) (int,[][]int) {
	//result := []string{}

	dp := make([][]int,len(s))

	for i := range dp {
		dp[i] = make([]int,len(s))
	}

	for i:= 0;i<len(s);i++ {
		//if len(result) < 1 {
		//	result = append(result,"")
		//}
		dp[i][i] = 1
		//result[0] = dp[i][i]
	}
	for i := 0;i<len(s)-1;i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
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

	for _,v := range dp {
		fmt.Println(v)
	}
	//fmt.Printf("%#v\n",dp)

	return dp[0][len(s)-1],dp
}

func Max(x,y int) int {
	if x < y {
		x = y
	}
	return x
}

// 生成dp的时候，怎么决策的，就怎么根据DP生成相应的字符串
func getStrFromDP(str string, dp [][]int) string {
	a,b := 0,len(str)-1

	result := []string{}
	for a < len(str) && b >= 0 && a<=b {
		//fmt.Println(a,b)
		// 1,2
		if a+1 < len(str) &&dp[a+1][b-1] == dp[a][b] -2 {
			result = append(result,string(str[b]))
			result = append([]string{string(str[a])},result...)
			a ++
			b --
		} else if b >=1 && dp[a][b-1] == dp[a][b] {
			b --
		} else if a + 1<len(str) && dp[a+1][b] == dp[a][b] {
			a ++
		}
	}

	return strings.Join(result,"")
}
*/
