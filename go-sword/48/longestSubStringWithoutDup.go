package main

import (
	"fmt"
)

// 48：最长不含重复字符的子字符串
// 题目：请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子
// 字符串的长度。假设字符串中只包含从'a'到'z'的字符。
// 例如，在字符串"arabcacfr"中，最长的不含重复字符的子字符串是"acfr"，长度为4

// A：记录每个字符最后一次出现的位置，
// 遍历到一个字符时，计算以该字符结尾的最长子串的长度
// 检查该字符上次出现的位置

func getMaxLength(str string) int {
	if len(str) == 0 {
		return 0
	}
	memo :=[26]int{}
	for i := 0;i<len(memo);i++ {
		memo[i] = -1
	}

	start := 0
	result := 0

	for i := 0;i<len(str);i++ {
		if memo[str[i]-'a'] >= start {
			start = memo[str[i]-'a'] + 1
		}
		memo[str[i]-'a'] = i
		result = Max(result,i-start+1)
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	Test("Test1","abcacfrar",4)
	Test("Test2","acfrarabc",4)
	Test("Test3","arabcacfr",4)
	Test("Test4","aaaa",1)
	Test("Test5","abcdefg",7)
	Test("Test6","aaabbbccc",2)
	Test("Test7","abcdcba",4)
	Test("Test8","abcdaef",6)
	Test("Test9","a",1)
	Test("Test10","",0)
}

func Test(name string,str string,expected int) {
	fmt.Println(name)
	fmt.Println(getMaxLength(str),expected)
	if getMaxLength(str) != expected {
		panic("fuck")
	}
}


















