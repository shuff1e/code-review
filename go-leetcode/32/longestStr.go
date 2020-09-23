package main

import "fmt"

//给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
//
//示例 1:

//输入: "(()"
//输出: 2
//解释: 最长有效括号子串为 "()"
//示例 2:
//
//输入: ")()())"
//输出: 4
//解释: 最长有效括号子串为 "()()"
//

func main() {
	str := "(()(((()"
	str = ")()())"
	str = "()(())"
	fmt.Println(longestValidParenthesesBetter(str))
}

// A：用栈，当有)不匹配时，就隔断开了
// 如果我们能记录最后一个不匹配的位置
// 然后(和)匹配的时候，就能得到长度
func longestValidParentheses(s string) int {
	res := 0
	stack := NewStack()
	// 开始的时候，第一个不匹配的）的位置为-1
	stack.Push(-1)
	for i := 0;i<len(s);i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			// 右括号的话出栈
			stack.Pop()

			// 如果stack不是空的话，
			// 说明还有不匹配的右括号垫底
			if !stack.Empty() {
				temp,_ := stack.Peek()
				res = Max(res,i-temp.(int))
			} else {
				// 如果stack时空的话，说明当前右括号时不匹配的右括号
				stack.Push(i)
			}
		}
	}
	return res
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

type Stack struct {
	s []interface{}
}

func NewStack() *Stack {
	return &Stack{
		make([]interface{},0),
	}
}

func (s *Stack) Push(v interface{}) {
	s.s = append(s.s,v)
}

func (s *Stack) Pop() (interface{},bool) {

	l := len(s.s)
	if l == 0 {
		return 0, false
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, true
}

func (s *Stack) Peek() (interface{},bool) {

	l := len(s.s)
	if l == 0 {
		return 0, false
	}

	return s.s[l-1],true
}

func (s *Stack) Length() int {
	return len(s.s)
}

func (s *Stack) Empty() bool {
	return s.Length() == 0
}

// 用递归的思路解决问题
// dp
// 如果当前i是),i-1是(,dp[i]=dp[i-1]+2
// 如果当前i是),i-1是),dp[i]要看i-dp[i-1]-1,如果这个位置是(，则dp[i]=dp[i-1]+2
// 如果当前i是(,i-1是(,dp[i]=0
// 如果当前i是(,i-1是),dp[i]=dp[i-1]

func longestValidParenthesesBetter(s string) int {
	if len(s) <= 1 {
		return 0
	}
	max := 0

	dp := make([]int,len(s))
	dp[0] = 0
	for i := 1;i<len(s);i++ {
		if s[i] == '(' {
			if s[i-1] == '(' {
				dp[i] = 0
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			if s[i-1] == '(' {
				dp[i] = dp[i-1] + 2
				max = Max(max,dp[i])
			} else {
				if i-dp[i-1]-1 >=0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-1]
					max = Max(max,dp[i])
				}
			}
		}
	}
	return max
}
