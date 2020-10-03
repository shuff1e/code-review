package main

/*
72. 编辑距离
给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
 */

// 输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
//

// A：每次都有4种选择
// 例如对于
// A = horse，B = ros
//
// A插入字符: dp[i][j] =  dp[i-1][j] + 1
// (horse 已经匹配了 ro，现在要求 horse匹配ros，则只能在horse后面加上s)
//
// B插入字符: d[[i][j] =  dp[i][j-1] + 1
// (hors 已经匹配了 ros，现在要求 horse 匹配ros，则只能在ros后面加上e，
// 这样相当于从A中删除了e)
//
// A修改字符: d[[i][j] =  dp[i-1][j-1] + 1 (hors 已经匹配了 ro ,现在要求horse匹配ros)
//
// A从""变为"ros"，代价为3
// B从""变为"hors"，代价为4
//
// dp[i][j] 表示 A 的前 i 个字母和 B 的前 j 个字母之间的编辑距离
// if a[i] == b[j]  dp[i][j] = Min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1])
// if a[i] != b[j]  dp[i][j] = Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1

func minDistance(word1 string, word2 string) int {
	length1 := len(word1)
	length2 := len(word2)
	if length2 * length1 == 0 {
		return length2 + length1
	}

	dp := make([][]int,length1+1)
	for i := 0;i<length1+1;i ++ {
		dp[i] = make([]int,length2+1)
	}

	for i := 0;i<length1+1;i++ {
		dp[i][0] = i
	}
	for j := 0;j<length2+1;j++ {
		dp[0][j] = j
	}
	for i := 1;i<length1+1;i++ {
		for j := 1;j<length2+1;j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = Min(dp[i-1][j]+1,dp[i][j-1]+1,dp[i-1][j-1])
			} else {
				dp[i][j] = Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[length1][length2]
}

func Min(x,y,z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}