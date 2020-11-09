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

func main() {
	str := "pwwkew"
	str = "a"
	str = "abcdefghif"
	fmt.Println(lengthOfLongestSubstring(str))
}

// 记录字符最后一次出现的位置

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	result := 1
	cnt := 1
	dict := map[byte]int{}
	dict[s[0]] = 0

	for i := 1;i<len(s);i++ {
		if _,ok := dict[s[i]];!ok {
			dict[s[i]] = i
			cnt ++
			result = Max3(result,cnt)
		} else {
			if i - dict[s[i]] <= cnt {
				cnt = i - dict[s[i]]
			} else {
				cnt ++
				result = Max3(result,cnt)
			}
			dict[s[i]] = i
		}
	}
	result = Max3(result,cnt)
	return result
}

func Max3(x,y int) int {
	if x > y {
		return x
	}
	return y
}
