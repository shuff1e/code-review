package main

import "fmt"

/*
172. 阶乘后的零
给定一个整数 n，返回 n! 结果尾数中零的数量。

示例 1:

输入: 3
输出: 0
解释: 3! = 6, 尾数中没有零。
示例 2:

输入: 5
输出: 1
解释: 5! = 120, 尾数中有 1 个零.
说明: 你算法的时间复杂度应为 O(log n) 。
 */

func main() {
	fmt.Println(trailingZeroes(625))
}

// 其实就是求2和5的个数
// 5的个数更少，因此就是求2的个数

/*
只关注5的个数：例如求125!中有多少个5相乘

125! 转化为
125*120*115*...*25*20..*5 =
(25*5) * (24*5) * (23*5) * (22*5) * ...* (1*5)=25!*(25个5相乘)
125！转化为了再求25！中有多少个5  +  25个5

 */

// f(n) = f(n/5) + n/5
//
// f(5) = 1
// f(n) = 0 (n<5)

func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n/5
		n = n/5
	}
	return count
}

func trailingZeroes2(n int) int {
	result := 0
	for i :=1;i<=n;i++ {
		temp := getFive(i)
		result += temp
	}
	return result
}

func getFive(n int) int {
	count := 0
	for n % 5 == 0 {
		count ++
		n = n/5
	}
	return count
}
