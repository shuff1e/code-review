package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
166. 分数到小数
给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

示例 1:

输入: numerator = 1, denominator = 2
输出: "0.5"
示例 2:

输入: numerator = 2, denominator = 1
输出: "2"
示例 3:

输入: numerator = 2, denominator = 3
输出: "0.(6)"
 */

// 1/4=0，1%4=1
// 10/4=2， 10%4=2
// 20/4=5， 20%4=0
//
// 2/3=0，2%3=2
// put(2,1)
// 20/3=6，20%3=2

func main()  {
	fmt.Println(fractionToDecimal(-50,8))
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	result := []byte{}
	// boolean的异或
	if ((numerator < 0) != (denominator < 0)) {
		result = append(result,'-')
	}

	result = append(result,[]byte(strconv.Itoa(Abs(numerator/denominator)))...)
	remainder := numerator%denominator
	if remainder == 0 {
		return String(result)
	}
	result = append(result,'.')
	// 2/3
	// 0.
	// remainder = 2
	// 20/3=6,remainder = 2
	dict := map[int]int{}
	for remainder != 0 {
		if index,ok := dict[remainder];ok {
			rear := append([]byte{},result[index:]...)
			temp := result[0:index]
			result = append(append(temp,'('),rear...)
			result = append(result,')')
			break
		}
		dict[remainder] = len(result)
		remainder *= 10
		result = append(result,[]byte(strconv.Itoa(Abs(remainder/denominator)))...)
		remainder = remainder%denominator
	}
	return String(result)
}


func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}