package main

/*
22. 括号生成
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例：

输入：n = 3
输出：[
"((()))",
"(()())",
"(())()",
"()(())",
"()()()"
]

 */

// A：左括号可以无限加，只要有
// 右括号的数量必须小于左括号

func generateParenthesis(n int) []string {
	result := []string{}
	help(0,0,n,"",&result)
	return result
}

func help(leftNumber ,rightNumber ,k int,temp string, result *[]string) {
	if leftNumber == k && rightNumber == k {
		if len(temp) > 0 {
			*result = append(*result,temp)
		}
		return
	}
	if leftNumber < k {
		help(leftNumber+1,rightNumber,k,temp+"(",result)
	}
	if rightNumber < leftNumber {
		help(leftNumber,rightNumber+1,k,temp+")",result)
	}
}
