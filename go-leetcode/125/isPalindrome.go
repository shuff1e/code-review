package main

import (
	"fmt"
	"strings"
)

/*
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
 */

func main() {
	str := "A man, a plan, a canal: Panama"
	str = "race a car"
	str = "0P"
	fmt.Println(isPalindrome(str))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left <len(s) && !isCharOrNumber(s[left]) {
			left ++
		}
		for right >= 0 && !isCharOrNumber(s[right]) {
			right --
		}
		if right <= left {
			break
		}
		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left ++
		right --
	}
	return true
}

func isCharOrNumber(x byte) bool {
	result := (x >= 'A' && x <= 'Z') || (x >= 'a' && x <= 'z') || (x >= '0' && x<= '9')
	return result
}