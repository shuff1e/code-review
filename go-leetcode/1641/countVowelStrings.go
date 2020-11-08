package main

import "fmt"

/*

1641. 统计字典序元音字符串的数目
给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。

字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。

示例 1：

输入：n = 1
输出：5
解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
示例 2：

输入：n = 2
输出：15
解释：仅由元音组成的 15 个字典序字符串为
["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
示例 3：

输入：n = 33
输出：66045


提示：

1 <= n <= 50

*/

// 类似换钱的方法数

// arr = [a,e,i,o,u]

// dp[i][j]表示,使用到 arr[i],并且 使用了j个位置的方法数

// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[0][j] = 1
// dp[i][0] = 0

func main() {
	fmt.Println(countVowelStrings(33))
}

func countVowelStrings(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([][]int,5)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,n+1)
	}

	for i := 0;i<len(dp);i++ {
		dp[i][0] = 0
		dp[i][1] = i + 1
	}

	for j := 0;j<len(dp[0]);j++ {
		dp[0][j] = 1
	}

	for i := 1;i<len(dp);i++ {
		for j := 2;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[len(dp)-1][n]
	//result := 0
	//for i := 0;i<len(dp);i++ {
	//	result += dp[i][n]
	//}
	//return result
}