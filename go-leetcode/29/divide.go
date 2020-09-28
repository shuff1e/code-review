package main

/*
29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2


提示：

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
 */

const INT_MIN = -0x80000000
const INT_MAX = 0x7fffffff

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > INT_MIN { // 不是最小的整数，返回相反的值
			return -dividend
		}
		return INT_MAX // 是最小的整数，返回最大的值
	}

	symbol := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		symbol = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0{
		divisor = -divisor
	}
	res := div(dividend,divisor)

	if symbol == 1 {
		return res
	}
	return -res
}

// 10 /3 =3
// 7/-3 = -2

// 类似pow
// 11/3
// 3翻倍=6，6翻倍是12>11，不行
// 所以计算11-6=5，5/3

func div(a,b int) int {
	if a < b {
		return 0
	}
	count := 1
	temp := b
	for (temp + temp) <= a {
		count = count + count
		temp = temp + temp
	}
	return count + div(a-temp,b)
}

