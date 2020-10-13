package main

import (
	"fmt"
	"unsafe"
)

/*
214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。

示例 1:

输入: "aacecaaa"
输出: "aaacecaaa"
示例 2:

输入: "abcd"
输出: "dcbabcd"

 */

func main() {
	//fmt.Println(shortestPalindrome("aacecaaa"))
	//fmt.Println(computeTempArray([]byte("aacecaaa")))
	text := "BBCABCDABABCDABCDABDE"
	pattern := "ABCDABDE"
	fmt.Println(Kmp([]byte(text),[]byte(pattern)))
}

func shortestPalindrome(s string) string {
	return help(s)
}

func help(s string) string {
	if len(s) <= 1 {
		return s
	}
	index := 0
	for i := len(s)-1;i>=0;i--{
		valid := getShortest(s,0,i)
		if valid {
			index = i
			break
		}
	}
	temp := s[index+1:]
	return reverse(temp) + s
}

// 先找到以s[0]开头的最短的回文串
func getShortest(s string,start,end int) bool {
	if start >= end {
		return true
	}
	if s[start] != s[end] {
		return false
	}
	return getShortest(s,start+1,end-1)
}

func reverse(s string) string {
	arr := make([]byte,len(s))
	for i := 0;i<len(s);i++ {
		arr[len(s)-1-i] = s[i]
	}
	return String(arr)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}
// 从暴力法可以看出，其实就是求 s 的「最长回文前缀」，然后在 rev_s 的后缀中砍掉这个回文，再加到 s 前面。
//
//这个最长前缀是回文的，它翻转之后等于它自己，出现在 rev_s 的后缀，
//这不就是公共前后缀吗？KMP 的 next 数组记录的就是一个字符串的每个位置上，最长公共前后缀的长度。公共前后缀指的是前后缀相同。
//
//因此，我们 “制造” 出公共前后缀，去套 KMP。
//
//s：abab，则 s + '#' + rev_s，得到 str ：abab#baba。
//
//求出 next 数组，最后一项就是 str 的最长公共前后缀的长度，即 s 的最长回文前缀的长度。
//

func shortestPalindrome2(s string) string {
	str := s + "#" + reverse(s)
	arr := computeTempArray([]byte(str))
	length := arr[len(arr)-1]
	return reverse(s[length:]) + s
}
// KMP 算法
func computeTempArray(arr []byte) []int {
	result := make([]int,len(arr))
	index := 0
	for i := 1;i<len(arr); {
		if arr[i] == arr[index] {
			result[i] = index+1
			i ++
			index ++
		} else {
			if index != 0 {
				index = result[index-1]
			} else {
				result[i] = 0
				i++
			}
		}
	}
	return result
}

func Kmp(text []byte,pattern []byte) bool {
	lps := computeTempArray(pattern)
	i,j := 0,0
	for i < len(text) && j < len(pattern) {
		if text[i] == pattern[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	if j == len(pattern) {
		return true
	}
	return false
}

func shortestPalindrome3(s string) string {
	base := 131
	mod := 1000000007
	left,right := 0,0
	best := -1
	mul := 1
	for i := 0;i<len(s);i++ {
		left = (left*base + int(s[i]))%mod
		right = (right + int(s[i])*mul)%mod
		if left == right {
			best = i
		}
		mul = mul*base%mod
	}
	if best == len(s) - 1 {
		return s
	}
	return reverse(s[best+1:]) + s
}