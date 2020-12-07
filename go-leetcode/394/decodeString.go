package main

import (
	"fmt"
	"unsafe"
)

/*

394. 字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"

 */

// 本题中可能出现括号嵌套的情况，比如 2[a2[bc]]，这种情况下我们可以先转化成 2[abcbc]，
// 在转化成 abcbcabcbc。
// 我们可以把字母、数字和括号看成是独立的 TOKEN，并用栈来维护这些 TOKEN。

func main() {
	str := "2[abc]3[cd]ef"
	fmt.Println(decodeString2(str))
}

func decodeString(s string) string {
	stack := []byte{}

	numStack := []int{}
	ptr := 0

	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			// 如果是数字，
			times := getDigits(s,&ptr)
			numStack = append(numStack,times)
		} else if cur >= 'a' && cur <= 'z' || cur == '[' {
			// 如果是字符或者是[，直接入栈
			stack = append(stack,cur)
			ptr ++
		} else {
			// 右]，直接过滤掉
			ptr ++
			sub := []byte{}
			for stack[len(stack)-1] >= 'a' && stack[len(stack)-1] <= 'z' {
				sub = append(sub,stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			sub = reverse(sub)
			// 有一个右括号，必有一个times
			times := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			t := []byte{}
			for times > 0 {
				t = append(t,sub...)
				times --
			}
			stack = append(stack,t...)
		}
	}
	return String(stack)
}

func getDigits(str string,ptr *int) int {
	cur := 0
	for str[*ptr] >= '0' && str[*ptr] <= '9' {
		cur = cur * 10 + int(str[*ptr] - '0')
		*ptr ++
	}
	return cur
}

func reverse(arr []byte) []byte {
	left,right := 0,len(arr)-1
	for left < right {
		arr[left],arr[right] = arr[right],arr[left]
		left ++
		right --
	}
	return arr
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func decodeString2(s string) string {
	ptr := 0
	result := getString(s,&ptr)
	return result
}

func getString(str string, ptr *int) string {
	if *ptr == len(str) || str[*ptr] == ']' {
		return ""
	}
	cur := str[*ptr]
	times := 1
	result := ""

	if cur >= '0' && cur <= '9' {
		times = getDigits(str,ptr)
		// 过滤掉 左括号
		*ptr ++
		x := getString(str,ptr)
		// 过滤掉右括号
		*ptr ++
		for times > 0 {
			times --
			result += x
		}
	} else if cur >= 'a' && cur <= 'z' {
		result = string(cur)
		*ptr ++
	}
	result += getString(str,ptr)
	return result
}