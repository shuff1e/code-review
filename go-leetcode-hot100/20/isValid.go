package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

 */


// A：栈，左括号入栈，
// 右括号pop
// 题32
func isValid(s string) bool {
	dict := map[rune]rune{
		'}':'{',
		']':'[',
		')':'(',
	}

	stack := []rune{}

	for _,v := range s {
		// 右括号
		if temp,ok := dict[v];ok {
			if len(stack) == 0 {
				return false
			}
			// peek
			if rune(stack[len(stack)-1]) != temp {
				return false
			}
			// pop
			stack = stack[0:len(stack)-1]
		} else {
			stack = append(stack,v)
		}
	}
	return len(stack) == 0
}
