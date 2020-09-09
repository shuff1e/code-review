package main

import "fmt"

// 16：数值的整数次方
// 题目：实现函数double Power(double base, int exponent)，求base的exponent
// 次方。不得使用库函数，同时不需要考虑大数问题。

// A：10^13 = 10^8 * 10^4 * 10
// 13的二进制为 1101
// 同时要考虑到指数为负数的情况

func Power(base float64,exponent int) (float64,bool) {
	if base ==0 && exponent < 0 {
		return 0,false
	}
	flag := 1
	if exponent < 0 {
		flag = -1
		exponent = -exponent
	}

	result := 1.0
	for exponent > 0 {
		if exponent & 1 > 0 {
			result = result * base
		}
		base = base * base
		exponent >>= 1
	}
	if flag == 1 {
		return result,true
	} else {
		return 1/result,true
	}
}

func main() {
	Test(2, 3, 8,true)
	Test(-2, 3, -8, true);
	Test(2, -3, 0.125,true)
	Test(2, 0, 1,true)
	Test( 0, 0, 1,true)
	Test(0, 4, 0,true)

	Test(0,-4,0,false)
}

func Test(base float64,exponent int,expected float64,isValid bool) {
	result,valid := Power(base,exponent)
	fmt.Println(result == expected)
	fmt.Println(valid == isValid)
}