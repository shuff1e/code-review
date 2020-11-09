package main

import "fmt"

/*

5. 最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

 */

func main() {
	str := "babad"
	str = "cbbd"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	start := 0
	end := 0
	for i := 0;i<len(s);i++ {
		s1,e1 := help5(s,i,i)
		s2,e2 := help5(s,i,i+1)

		if (e2 - s2) > (e1-s1) {
			e1 = e2
			s1 = s2
		}

		if (e1 - s1) > (end - start) {
			end = e1
			start = s1
		}
	}
	return s[start:end+1]
}

func help5(str string, left, right int) (int,int) {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left --
		right ++
	}
	return left + 1,right - 1
}