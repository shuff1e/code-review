package main

import "fmt"

/*
227. 基本计算器 II
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:

输入: "3+2*2"
输出: 7
示例 2:

输入: " 3/2 "
输出: 1
示例 3:

输入: " 3+5 / 2 "
输出: 5
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
 */

func main() {
	str := " 3+5 / 2 "
	str = "1+3*(3+2)*2"
	// 1 3 3 2
	// + * ( +

	// 1 3 5
	// + *

	// 1 15 2
	// + *
	fmt.Println(calculate(str))
}

// 我们其实可以直接利用两个栈，边遍历边进行的
//
//1.使用两个栈，stack0 用于存储操作数，stack1 用于存储操作符
//2.从左往右扫描，遇到操作数入栈 stack0
//3.遇到操作符时，如果当前优先级低于或等于栈顶操作符优先级，则从 stack0 弹出两个元素，从 stack1 弹出一个操作符，进行计算，将结果并压入stack0，继续与栈顶操作符的比较优先级。
//4.如果遇到操作符高于栈顶操作符优先级，则直接入栈 stack1
//5.遇到左括号，直接入栈 stack1。
//6.遇到右括号，则从 stack0 弹出两个元素，从 stack1 弹出一个操作符进行计算，并将结果加入到 stack0 中，重复这步直到遇到左括号
//
//
//第 3 条改成「遇到操作符时，则从 stack0 弹出两个元素进行计算，并压入stack0，直到栈空或者遇到左括号，最后将当前操作符压入 stack1 」

func calculate(s string) int {
	arr := []byte(s)
	numStack := []int{}
	operationStack := []byte{}
	temp := -1
	for i := 0;i<len(arr);i++ {
		if arr[i] == ' ' {
			continue
		}
		// 数字直接入栈
		if arr[i] >= '0' && arr[i] <= '9' {
			if temp == -1 {
				temp = int(arr[i] - '0')
			} else {
				temp = 10*temp + int(arr[i]- '0')
			}
		} else {
			if temp != -1 {
				numStack = append(numStack,temp)
				temp = -1
			}
			// 符号，看优先级
			if arr[i] == '-' || arr[i] == '+' || arr[i] == '*' || arr[i] == '/' {
				// 3+5/6-2
				for len(operationStack) > 0 &&
					!priority(arr[i],operationStack[len(operationStack)-1]){
					if operationStack[len(operationStack)-1] == '(' {
						break
					}
					num1 := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]

					num2 := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]

					operation := operationStack[len(operationStack)-1]
					operationStack = operationStack[:len(operationStack)-1]

					if operation == '+' {
						numStack = append(numStack,num2 + num1)
					} else if operation == '-' {
						numStack = append(numStack,num2 - num1)
					} else if operation == '*' {
						numStack = append(numStack,num2 * num1)
					} else if operation == '/' {
						numStack = append(numStack,num2 / num1)
					}
				}
				operationStack = append(operationStack,arr[i])
			}
			// 左括号直接入栈
			if arr[i] == '(' {
				operationStack = append(operationStack,arr[i])
			}
			// 右括号，直到计算到左括号
			if arr[i] == ')' {
				for operationStack[len(operationStack)-1] != '(' {
					num1 := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]

					num2 := numStack[len(numStack)-1]
					numStack = numStack[:len(numStack)-1]

					operation := operationStack[len(operationStack)-1]
					operationStack = operationStack[:len(operationStack)-1]

					if operation == '+' {
						numStack = append(numStack,num2 + num1)
					} else if operation == '-' {
						numStack = append(numStack,num2 - num1)
					} else if operation == '*' {
						numStack = append(numStack,num2 * num1)
					} else if operation == '/' {
						numStack = append(numStack,num2 / num1)
					}
				}
				operationStack = operationStack[:len(operationStack)-1]
			}
		}
	}

	if temp != -1 {
		numStack = append(numStack,temp)
	}
	for len(operationStack) > 0 {
		num1 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]

		num2 := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]

		operation := operationStack[len(operationStack)-1]
		operationStack = operationStack[:len(operationStack)-1]

		if operation == '+' {
			numStack = append(numStack,num2 + num1)
		} else if operation == '-' {
			numStack = append(numStack,num2 - num1)
		} else if operation == '*' {
			numStack = append(numStack,num2 * num1)
		} else if operation == '/' {
			numStack = append(numStack,num2 / num1)
		}
	}
	return numStack[len(numStack)-1]
}

// char1的优先级是否高于char2的优先级
func priority(char1,char2 byte) bool {
	if char1 == '*' || char1 == '/' {
		if char2 == '-' || char2 == '+' {
			return true
		}
		return false
	}
	return false
}