package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/*

316. 去除重复字母
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。
需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

注意：该题与 1081 https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters 相同



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"


提示：

1 <= s.length <= 104
s 由小写英文字母组成

 */

func main() {
	str := "cbacdcbc"
	fmt.Println(removeDuplicateLetters2(str))
}

func removeDuplicateLetters(s string) string {
	return help(s)
}

func help(str string) string {
	counts := make([]int,26)
	for i := 0;i < len(str);i ++ {
		counts[str[i] - 'a'] ++
	}
	pos := 0
	for i := 0;i<len(str);i++ {
		// 找到字典序最小的
		if str[i] < str[pos] {
			pos = i
		}
		counts[str[i]-'a']--
		// caccbc
		if counts[str[i]-'a'] == 0 {
			break
		}
	}
	char := string(str[pos])
	str = strings.ReplaceAll(str[pos+1:],string(str[pos]),"")

	if len(str) == 0 {
		return char
	} else {
		return char + help(str)
	}
}

func removeDuplicateLetters2(s string) string {
	seen := map[byte]struct{}{}
	stack := []byte{}

	last_occurence := map[byte]int{}
	for i := 0;i < len(s);i ++ {
		last_occurence[s[i]] = i
	}

	for i := 0;i<len(s);i ++ {
		if _,ok := seen[s[i]];!ok {
			for len(stack) > 0 &&
				stack[len(stack)-1] > s[i] &&
				last_occurence[stack[len(stack)-1]] > i {
				delete(seen,stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			seen[s[i]] = struct{}{}
			stack = append(stack,s[i])
		}
	}
	return String(stack)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}