package main

import "fmt"

/*

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

示例 1:

输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:

输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
说明: 你可以假设 n 不小于 2 且不大于 58。

 */

func main() {
	x := 10
	fmt.Println(getMax(x))
}

func getMax(x int) int {
	if x < 2 {
		return 0
	}

	memo := make([][]int,x)
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,x + 1)
		for j := 0;j<len(memo[i]);j++ {
			memo[i][j] = -1
		}
	}

	result := 0
	for i := 1;i<x;i++ {
		result = Max2(result,help3(0,i,memo)*help3(i,x,memo))
	}
	return result
}

func help3(x int,memo [][]int) int {
	if start >= end {
		return 1
	}
	if end - start == 1 {
		return start
	}

	if memo[start][end] != -1 {
		return memo[start][end]
	}

	result := end
	for i := start;i<end;i++ {
		result = Max2(result,help3(start,i,memo) * help3(i,end,memo))
	}
	memo[start][end] = result
	return result
}

func Max2(x,y int) int {
	if x > y {
		return x
	}
	return y
}