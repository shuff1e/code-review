package main

import "fmt"

/*
96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

1         3     3      2      1
\       /     /      / \      \
3     2     1      1   3      2
/     /       \                 \
2     1         2                 3
 */

// f(0) = 1   f(1)=1 f(2)=2 f(3)=5
// A：f(3) = f(0)*f(2) + f(2)*f(0) + f(1)*f(1) = 5
// f(4) = f(0)*f(3) + f(1)*f(2) + f(2)*f(1) + f(3)*f(0)
//      = 1*5 + 2*2 + 5*1

func main() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	if n == 0 || n== 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int,n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i :=3;i<=n;i++ {
		for j :=1;j<=i;j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}