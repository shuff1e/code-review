package main

import "fmt"

/*
97. 交错字符串
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出：true
示例 2：

输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出：false
 */

// A：f(i,j)=[f(i−1,j) and s1(i)=s3(p)] or [f(i,j−1) and s2(j)=s3(p)]
//

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	//fmt.Println(isInterleave2(s1,s2,s3))
	//fmt.Println(isInterleave(s1,s2,s3))
	s1 = "aabcc"
	s2 = "dbbca"
	s3 = "aadbbcbcac"
	fmt.Println(isInterleave2(s1,s2,s3))
	fmt.Println(isInterleave(s1,s2,s3))
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	m,n := len(s1),len(s2)
	if m + n != len(s3) {
		return false
	}

	dp := make([][]bool,m+1)
	for i := 0;i<=m;i++ {
		dp[i] = make([]bool,n+1)
	}

	dp[0][0] = true
	for i := 0;i<=m;i++ {
		for j := 0;j<=n;j++ {
			if i > 0 {
				dp[i][j] = s1[i-1] == s3[i+j-1] && dp[i-1][j]
			}
			if dp[i][j] {
				continue
			}
			if j > 0 {
				dp[i][j] = s2[j-1] == s3[i+j-1] && dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	dict := map[byte]int{}

	for i := 0;i<len(s1);i++ {
		dict[s1[i]] = dict[s1[i]] + 1
	}
	for i := 0;i<len(s2);i++ {
		dict[s2[i]] = dict[s2[i]] + 1
	}
	for i := 0;i<len(s3);i++ {
		dict[s3[i]] = dict[s3[i]] - 1
	}
	for _,v := range dict {
		if v != 0 {
			return false
		}
	}

	memo := make([][][]int,len(s1)+1)
	for i := 0;i<len(s1)+1;i++ {
		memo[i] = make([][]int,len(s2)+1)
		for j :=0;j<len(s2)+1;j++ {
			memo[i][j] = make([]int,len(s3)+1)
		}
	}
	result := help(s1,s2,s3,0,0,0,memo)
	return result
}


func help(s1,s2,s3 string,index1,index2,index3 int,memo [][][]int) bool {
	if index3 == len(s3) {
		return true
	}
	//if index1 == len(s1) && index2 == len(s2) {
	//	return false
	//}
	if index1 < len(s1)+1 && index2 < len(s2)+1 && index3 < len(s3)+1 && memo[index1][index2][index3] != 0 {
		if memo[index1][index2][index3] == 1 {
			return true
		}
		return false
	}
	if index1 < len(s1) && index3 < len(s3) && s3[index3] == s1[index1] {
		result := help(s1,s2,s3,index1+1,index2,index3+1,memo)
		if result {
			memo[index1][index2][index3] = 1
			return true
		}
	}

	if index2 < len(s2) && index3 < len(s3) && s3[index3] == s2[index2] {
		result := help(s1,s2,s3,index1,index2+1,index3+1,memo)
		if result {
			memo[index1][index2][index3] = 1
			return true
		}
	}
	memo[index1][index2][index3] = 0
	return false
}