package main

import "fmt"

// 50（一）：字符串中第一个只出现一次的字符
// 题目：在字符串中找出第一个只出现一次的字符。如输入"abaccdeff"，则输出
// 'b'。

func findFirst(str string) byte {
	if len(str) == 0 {
		return '0'
	}
	table := make([]int,256)
	for i := 0;i<len(str);i++ {
		table[str[i]] ++
	}
	for i := 0;i<len(str);i++ {
		if table[str[i]] == 1 {
			return str[i]
		}
	}
	return '0'
}

func Test(str string,expected byte) {
	fmt.Println(findFirst(str))
	if findFirst(str) != expected {
		panic("fuck")
	}
}

func main() {
	// 常规输入测试，存在只出现一次的字符
	Test("google", 'l');

	// 常规输入测试，不存在只出现一次的字符
	Test("aabccdbd", '0');

	// 常规输入测试，所有字符都只出现一次
	Test("abcdefg", 'a');

	// 鲁棒性测试，输入nullptr
	Test("", '0');
}


























