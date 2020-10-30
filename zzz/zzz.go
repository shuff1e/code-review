package main

import "fmt"

func main() {
	str1 := "A1234123"
	str2 := "ACD4123ABC"
	fmt.Println(LCS(str1,str2))
}


/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS( str1 string ,  str2 string ) string {
	// write code here
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}

	dp := getDP(str1,str2)
	count,end := dealDP(dp)

	if count == 0 {
		return ""
	}

	return str1[end-count+1:end+1]
}

// str1[6] str2[4]
// index = 6

// count = 4
// 3:7

// dp[i][j] =0  str1[i] != str2[j]
// dp[i][j] = dp[i-1][j-1] + 1 if str1[i] == str2[j]

// count, index

func getDP(str1,str2 string) [][]int {
	dp := make([][]int,len(str1))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(str2))
	}

	for i := 0;i<len(str1);i++ {
		if str1[i] == str2[0] {
			dp[i][0] = 1

		} else {
			dp[i][0] = 0
		}
	}

	for j := 0;j<len(str2);j++ {
		if str2[j] == str1[0] {
			dp[0][j] = 1

		} else {
			dp[0][j] = 0
		}
	}

	for i := 1;i<len(str1);i++ {
		for j := 1;j<len(str2);j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp
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

