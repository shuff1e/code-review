package main

import "fmt"

/*

343. 整数拆分
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
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak2(10))
}

func integerBreak(n int) int {
	result := -1
	memo := map[int]int{}

	for i := 1;i<n;i++ {
		temp := i*help(n-i,memo)
		result = Max(result,temp)
	}
	return result
}

func help(n int,memo map[int]int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if v,ok := memo[n];ok {
		return v
	}

	result := n

	for i := 1;i<n;i++ {
		temp := i*help(n-i,memo)
		result = Max(result,temp)
	}
	memo[n] = result
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func integerBreak2(n int) int {
	if n <= 3 {
		return n - 1
	}
	a := n/3
	b := n % 3
	if b ==0 {
		return Pow(3,a)
	}
	if b == 1 {
		return Pow(3,a-1)*4
	}
	return Pow(3,a)*2
}

func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 != 0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}