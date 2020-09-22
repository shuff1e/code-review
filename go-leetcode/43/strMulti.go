package main

import "fmt"

//给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
//
//示例 1:
//
//输入: num1 = "2", num2 = "3"
//输出: "6"
//示例 2:
//
//输入: num1 = "123", num2 = "456"
//输出: "56088"
//说明：
//
//num1 和 num2 的长度小于110。
//num1 和 num2 只包含数字 0-9。
//num1 和 num2 均不以零开头，除非是数字 0 本身。
//不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
//

func multiply(num1 string, num2 string) string {
	if num2 == "0" || num1 == "0" {
		return "0"
	}
	result := "0"
	base := 0
	for i := len(num2) - 1;i>=0;i -- {
		temp := multi(num1,num2[i],base)
		result = add(temp,result)
		base ++
	}
	return result
}

// num1和num2的每一位相乘
func multi(num1 string,num2 uint8,base int) string {
	if num2 == '0' {
		return "0"
	}

	result := ""
	for i := 0;i<base;i++ {
		result = "0" + result
	}

	var carry uint8 = 0
	for i := len(num1)-1;i>=0;i-- {
		sum := (num1[i] - '0') * (num2 - '0') + carry
		result = string(sum%10 + '0') + result
		carry = sum/10
	}
	if carry > 0 {
		result = string('0' + carry) + result
	}
	return result
}

// 字符串加法
func add(num1 string,num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}

	if len(num1) < len(num2) {
		temp := num1
		num1 = num2
		num2 = temp
	}

	result := ""
	index1,index2 := len(num1) - 1,len(num2) - 1

	var carry uint8 = 0

	for index2 >= 0 {
		sum := num1[index1] - '0' + num2[index2] - '0' + carry
		result = string('0' + sum%10) + result
		carry = sum/10
		index2 --
		index1--
	}

	for index1 >= 0 {
		sum := num1[index1] - '0' + carry
		result = string('0' + sum%10) + result
		carry = sum/10
		index1--
	}
	if carry > 0 {
		result = string('0' + carry) + result
	}
	return result
}

func main() {
	fmt.Println(add("123","877"))
	fmt.Println(multiply("123","456"))
}

