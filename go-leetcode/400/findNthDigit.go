package main

import "fmt"

/*

400. 第N个数字
在无限的整数序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...中找到第 n 个数字。

注意:
n 是正数且在32位整数范围内 ( n < 231)。

示例 1:

输入:
3

输出:
3
示例 2:

输入:
11

输出:
0

说明:
第11个数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是0，它是10的一部分。

 */

// 1-9
// 10-99
// 100-999

// 9+ 2*90 + 3*900 +


func main() {
	fmt.Println(Pow(2,5))
	fmt.Println(findNthDigit(28))
}

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}

	base := 9
	index := 1
	for n > base + (index+1)*9*Pow(10,index) {
		base += (index+1)*9*Pow(10,index)
		index ++
	}

	// 每个位置是index+1个数字
	n -= base
	// 第几个数字
	count := n/(index+1)
	rest := n%(index+1)
	if rest == 0 {
		realNum := base + count
		return realNum%10
	}

	realNum := Pow(10,index) + count
	for i:=0;i<index + 1 -rest;i++ {
		realNum /= 10
	}
	return realNum%10

}

func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 >0 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}