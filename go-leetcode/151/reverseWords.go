package main

import (
	"fmt"
	"strings"
	"unsafe"
)

/*
151. 翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。

说明：

无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

示例 1：

输入："the sky is blue"
输出："blue is sky the"
示例 2：

输入："  hello world!  "
输出："world! hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入："a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
示例 4：

输入：s = "  Bob    Loves  Alice   "
输出："Alice Loves Bob"
示例 5：

输入：s = "Alice does not even like bob"
输出："bob like even not does Alice"

提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词
 */

func main() {
	fmt.Println(reverseWords("123 456"))
	fmt.Println(reverseWords("  Bob    Loves  Alice   "))
	fmt.Println(reverseWords("   123"))
}

func reverseWords(s string) string {
	// 转换成数组
	arr := make([]byte,len(s))
	for i := 0;i<len(arr);i++ {
		arr[i] = s[i]
	}

	// 翻转整体
	reverse(arr,0,len(arr)-1)

	result := []string{}
	start := 0
	for start < len(s) {
		for start < len(s) && arr[start] == ' ' {
			start ++
		}
		end := start
		for end < len(s) && arr[end] != ' ' {
			end ++
		}
		// 翻转每个单词
		reverse(arr,start,end-1)
		// 将单词记录下来
		if end > start {
			result = append(result,String(arr[start:end]))
		}
		start = end
	}
	return strings.Join(result," ")
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func reverse(arr []byte,left,right int) {
	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left ++
		right --
	}
}