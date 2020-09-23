package main

import (
	"fmt"
)

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
两个字符串完全匹配才算匹配成功。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。
示例 3:

输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
示例 4:

输入:
s = "adceb"
p = "*a*b"
输出: true

// *可以匹配，也可以不匹配
// 可以匹配一个，也可以匹配多个
// s和p都走完，才是true

解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
示例 5:

输入:
s = "acdcb"
p = "a*c?b"
输出: false

 */

func main() {
	s := "acdcb"
	p := "a*c?b"
	//fmt.Println(isMatch(s,p))
	s = "aa"
	p = "*"
	fmt.Println(isMatch(s,p))
	fmt.Println(isMatchBetter(s,p))
}

func isMatch(s string, p string) bool {
	memo := make([][]int,len(s)+1)
	for i := 0;i<len(s);i++ {
		memo[i] = make([]int,len(p)+1)
	}
	return helper(s,p,0,0,memo)
}

func checkAllStar(str string,start int) bool {
	for i := start;i<len(str);i++ {
		if str[i] != '*' {
			return false
		}
	}
	return true
}

func helper(s string,p string,sIndex,pIndex int,memo [][]int) bool {
	if sIndex == len(s) && pIndex == len(p) {
		return true
	}
	if sIndex < len(s) && pIndex == len(p) {
		return false
	}
	if sIndex == len(s) && pIndex < len(p) {
		result := checkAllStar(p,pIndex)
		return result
	}

	if memo[sIndex][pIndex] == 1 {
		return true
	} else if memo[sIndex][pIndex] == -1 {
		return false
	}

	if s[sIndex] == p[pIndex] {
		result :=  helper(s,p,sIndex+1,pIndex+1,memo)
		if result {
			memo[sIndex][pIndex] = 1
		} else {
			memo[sIndex][pIndex] = -1
		}
		return result
	}
	if p[pIndex] == '?' {
		result := helper(s,p,sIndex+1,pIndex+1,memo)
		if result {
			memo[sIndex][pIndex] = 1
		} else {
			memo[sIndex][pIndex] = -1
		}
		return result
	}
	if p[pIndex] == '*' {
		// 匹配0个
		last := findLastStar(p,pIndex)
		result := helper(s,p,sIndex,last+1,memo)
		for i := sIndex + 1;i<=len(s);i++ {
			result = result || helper(s,p,i,last+1,memo)
		}
		if result {
			memo[sIndex][pIndex] = 1
		} else {
			memo[sIndex][pIndex] = -1
		}
		return result
	}
	return false
}

func findLastStar(str string,start int) int {
	for i := start;i<len(str)-1;i++ {
		if str[i] == '*' && str[i+1] != '*' {
			return i
		}
	}
	return len(str) - 1
}

// 递归+memo可以改成dp
// dp[i][j] = dp[i+1][j+1]  s[i]=p[j],或者p[j]=?
// dp[i][j] = dp[i+x][j+1]  p[j]=='*' (*匹配0个，或者多个)(x从0到len(s)-i)
// dp[i][j]依赖右下角的，以及右下角那一列的

func isMatchBetter(s string, p string) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	}
	dp := make([][]bool,len(s)+1)
	for i := 0;i<len(s)+1;i++ {
		dp[i] = make([]bool,len(p)+1)
	}

	dp[len(s)][len(p)] = true

	for i := 0;i<len(s);i++ {
		dp[i][len(p)] = false
	}

	if p[len(p)-1] == '*' {
		dp[len(s)][len(p)-1] = true
	} else {
		dp[len(s)][len(p)-1] = false
	}

	for j :=len(p)-2;j>=0;j-- {
		dp[len(s)][j] = dp[len(s)][j+1] && (p[j] == '*')
	}

	for j := len(p) - 1;j>=0;j-- {
		for i := len(s)-1;i>=0;i-- {
			if s[i] == p[j] || p[j] == '?' {
				dp[i][j] = dp[i+1][j+1]
			} else if p[j] == '*' {
				dp[i][j] = dp[i][j+1]
				for start := i+1;start<=len(s);start++ {
					dp[i][j] = dp[i][j] || dp[start][j+1]
				}
			}
		}
	}
	//for i :=0;i<=len(s);i++ {
	//	fmt.Printf("%#v\n",dp[i])
	//}
	return dp[0][0]
}