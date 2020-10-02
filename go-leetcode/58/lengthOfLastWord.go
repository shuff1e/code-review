package main

import "fmt"

/*
58. 最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。
如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。

如果不存在最后一个单词，请返回 0 。

说明：一个单词是指仅由字母组成、不包含任何空格字符的最大子字符串。



示例:

输入: "Hello World"
输出: 5
 */

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	index := -1
	index2 := -1
	for i := len(s)-1;i>=0;i-- {
		if s[i] != ' ' {
			if index == -1 {
				index = i
			}
		} else if index != -1 {
			index2 = i
			break
		}
	}
	// _abc
	// 3-0
	return index - index2
}
