package main

import (
	"fmt"
	"strconv"
)

// 46：把数字翻译成字符串
// 题目：给定一个数字，我们按照如下规则把它翻译为字符串：0翻译成"a"，1翻
// 译成"b"，……，11翻译成"l"，……，25翻译成"z"。一个数字可能有多个翻译。例
// 如12258有5种不同的翻译，它们分别是"bccfi"、"bwfi"、"bczi"、"mcfi"和
// "mzi"。请编程实现一个函数用来计算一个数字有多少种不同的翻译方法。

// A：一次取2个，或者一次取1个翻译

func getCount(str string,index int) int {
	if index == len(str) {
		return 1
	}
	count := 0
	if checkValid1(str[index:index+1]) {
		count += getCount(str,index+1)
	}
	if index + 2 <= len(str) && checkValid2(str[index:index+2]) {
		count += getCount(str,index+2)
	}
	return count
}

func checkValid1(str string) bool {
	result,err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return result <= 9 && result >= 0
}

func checkValid2(str string) bool {
	 result := 10*(str[0] - '0') + str[1] - '0'
	 return result >=10 && result <=25
}

func main() {
	Test("Test1","0",1)
	Test("Test2","10",2)
	Test("Test3","125",3)
	Test("Test4","126",2)
	Test("Test5","426",1)
	Test("Test6","100",2)
	Test("Test7","101",2)
	Test("Test8","12258",5)
	Test("Test9","-100",0)
}

func Test(name string,str string,expected int) {
	fmt.Println(name)
	if getCount(str,0) != expected {
		panic("fuck")
	}
}