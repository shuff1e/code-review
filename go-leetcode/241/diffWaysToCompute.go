package main

import (
	"fmt"
	"strconv"
)

/*
241. 为运算表达式设计优先级
给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。
你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

示例 1:

输入: "2-1-1"
输出: [0, 2]
解释:
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2:

输入: "2*3-4*5"
输出: [-34, -14, -10, -10, 10]
解释:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
 */

// 括号唯一的作用就是改变了运算的优先级

// 对于一个形如 x op y（op 为运算符，x 和 y 为数） 的算式而言，
//它的结果组合取决于 x 和 y 的结果组合数，而 x 和 y 又可以写成形如 x op y 的算式。
//
//因此，该问题的子问题就是 x op y 中的 x 和 y：以运算符分隔的左右两侧算式解。
//
//然后我们来进行 分治算法三步走：
//
//分解：按运算符分成左右两部分，分别求解
//解决：实现一个递归函数，输入算式，返回算式解
//合并：根据运算符合并左右两部分的解，得出最终解

func main() {
	str := "2*3-4*5"
	fmt.Println(diffWaysToCompute(str))
}

func diffWaysToCompute(input string) []int {
	if v,err := strconv.Atoi(input);err == nil {
		return []int{v}
	}
	result := []int{}
	for i := 0;i<len(input);i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			leftParts := diffWaysToCompute(input[:i])
			rightParts := diffWaysToCompute(input[i+1:])
			for j := 0;j<len(leftParts);j++ {
				for k := 0;k<len(rightParts);k++ {
					if input[i] == '+' {
						result = append(result,leftParts[j] + rightParts[k])
					} else if input[i] == '-' {
						result = append(result,leftParts[j] - rightParts[k])
					} else if input[i] == '*' {
						result = append(result,leftParts[j] * rightParts[k])
					}
				}
			}
		}
	}
	return result
}