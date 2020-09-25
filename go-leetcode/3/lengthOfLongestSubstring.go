package main

import "fmt"

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 */

// 一次遍历，记录当前有效字符串的开始位置start
// 遍历到某个字符时，查看map中是否有这个字符，如果没有该字符，继续
// 如果有，index < start，不用管
// 如果 index >= start ，start = index + 1
// 并将该字符更新到map

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	start := 0
	mmp := make(map[byte]int,0)

	for i := 0;i<len(s);i++ {
		if index,ok := mmp[s[i]];!ok {
			result = Max(result,i-start+1)
		} else if index >= start {
			start = index + 1
		} else {
			result = Max(result,i-start+1)
		}
		mmp[s[i]] = i
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}