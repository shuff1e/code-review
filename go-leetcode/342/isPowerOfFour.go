package main

import "fmt"

/*

342. 4的幂
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

输入: 16
输出: true
示例 2:

输入: 5
输出: false
进阶：
你能不使用循环或者递归来完成本题吗？

 */

func main() {
	x := -2147483648
	fmt.Println(isPowerOfFour2(x))
}

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		mod := n % 4
		if mod != 0 {
			return false
		}
		n /= 4
	}
	return true
}

// 首先 n > 0
// 其次2的幂 都满足 n & (n - 1) == 0
// 其次4的幂，
// 0000001
// 0000100
// 0010000
// 1000000
// 1 都是在 奇数 位
//
// 这种和 1 在 偶数 位的 与，就为0
// 101010

func isPowerOfFour2(n int) bool {
	return (n > 0) && (n & (n-1) == 0 ) && (n & 0xaaaaaaaa == 0)
}

