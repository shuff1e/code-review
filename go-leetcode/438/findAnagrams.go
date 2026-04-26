package main

import "fmt"

/*
438. 找到字符串中所有字母异位词

给定两个字符串 s 和 p，找出 s 中所有是 p 的字母异位词的子串，返回这些子串的起始下标。

字母异位词指由相同字母重新排列得到的字符串。答案可以按任意顺序返回。

示例 1：
输入：s = "cbaebabacd", p = "abc"
输出：[0,6]
解释：起始下标 0 的子串 "cba" 和起始下标 6 的子串 "bac" 都是 "abc" 的字母异位词。

示例 2：
输入：s = "abab", p = "ab"
输出：[0,1,2]

提示：
1 <= s.length, p.length <= 3 * 10^4
s 和 p 仅包含小写英文字母。
*/

func main() {
	s := "cbaebabacd"
	p := "abc"
	result := findAnagrams(s, p)
	fmt.Println(result)

	s = "abab"
	p = "ab"
	result = findAnagrams(s, p)
	fmt.Println(result)
}

func findAnagrams(s string, p string) []int {
	dictS := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		dictS[s[i]]++
	}

	dictP := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		dictP[p[i]]++
	}

	result := []int{}

	if checkValid(dictS, dictP) {
		result = append(result, 0)
	}

	for i := 0; i < len(s)-len(p); i++ {
		dictS[s[i]]--
		dictS[s[i+len(p)]]++
		if checkValid(dictS, dictP) {
			result = append(result, i+1)
		}
	}

	return result
}

func checkValid(dictS map[byte]int, dictP map[byte]int) bool {
	for k, v := range dictP {
		if dictS[k] < v {
			return false
		}
	}

	return true
}
