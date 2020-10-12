package main

import "fmt"

/*
201. 数字范围按位与
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1:

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0
 */

func main() {
	fmt.Println(rangeBitwiseAnd(5,7))
}

/*
鉴于上述问题的陈述，我们的目的是求出两个给定数字的二进制字符串的公共前缀，这里给出的第一个方法是采用位移操作。

我们的想法是将两个数字不断向右移动，直到数字相等，即数字被缩减为它们的公共前缀。
然后，通过将公共前缀向左移动，将零添加到公共前缀的右边以获得最终结果。

 */

func rangeBitwiseAnd(m int, n int) int {
	shift := 0
	for m < n {
		m = m >> 1
		n = n >> 1
		shift ++
	}
	return m << shift
}

func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n = n & (n-1)
	}
	return n
}
