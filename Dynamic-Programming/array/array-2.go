package array

import (
	"github.com/shuff1e/code-review/util/math"
	"unsafe"
)

// 给定2个字符串，返回两个字符串的最长公共子序列
// 例如给定 ABCBDAB 和 BDCABA
// 返回BCBA 和 BDAB 和 BCAB都可以

// dp[i][j]表示以str1[i]，str2[j]为结尾的最长字串的长度

//              dp[i-1][j-1] + 1 ,str1[i] == str2[j]
// dp[i][j] =   max(dp[i-1][j],dp[i][j-1])

func getDP(str1 ,str2 string) [][]int {
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
			dp[i][j] = math.Max(dp[i-1][j],dp[i][j-1])
			if str1[i] == str2[j] {
				dp[i][j] = math.Max(dp[i-1][j-1] + 1,dp[i][j])
			}
		}
	}

	return dp
}

func Lcse(str1,str2 string) string {
	dp := getDP(str1,str2)

	result := make([]byte,dp[len(str1)-1][len(str2)-1])
	index := len(result) - 1
	i := len(str1) - 1
	j := len(str2) - 1
	for index >= 0 {
		if i > 0 && dp[i][j] == dp[i-1][j] {
			i--
		} else if j > 0 && dp[i][j] == dp[i][j-1] {
			j--
		} else {
			result[index] = str1[i]
			index --
			i --
			j --
		}
	}
	return *(*string)(unsafe.Pointer(&result))
}