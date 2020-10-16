package main

import (
	"fmt"
	"strconv"
)

/*
224. 基本计算器
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。
 */

func main() {
	str := "(1+(4+5+2)-3)+(6+8)"
	arr := getPolish(str)
	fmt.Println(evalRPN(arr))
	//str = "2-1+2"
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
				// 3/5+6
				for len(operationStack) > 0 {
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

func calculate2(s string) int {
	polish := getPolish(s)
	return evalRPN(polish)
}

//中缀表达式转后缀表达式
// 有了上边的代码，我们只需要把题目给的中缀表达式转成后缀表达式，直接调用上边计算逆波兰式就可以了。
//
//中缀表达式转后缀表达式也有一个通用的方法，我直接复制 这里 的规则过来。
//
//1）如果遇到操作数，我们就直接将其加入到后缀表达式。
//
//2）如果遇到左括号，则我们将其放入到栈中。
//
//3）如果遇到一个右括号，则将栈元素弹出，将弹出的操作符加入到后缀表达式直到遇到左括号为止，
// 接着将左括号弹出，但不加入到结果中。
//
//4）如果遇到其他的操作符，如（“+”， “-”）等，从栈中弹出元素将其加入到后缀表达式，
// 直到栈顶的元素优先级比当前的优先级低（或者遇到左括号或者栈为空）为止。
// 弹出完这些元素后，最后将当前遇到的操作符压入到栈中。
//
//5）如果我们读到了输入的末尾，则将栈中所有元素依次弹出。
//
//这里的话注意一下第四条规则，因为题目中只有加法和减法，加法和减法是同优先级的，
// 所以一定不会遇到更低优先级的元素，所以「直到栈顶的元素优先级比当前的优先级低（或者遇到左括号或者栈为空）为止。」
// 这句话可以改成「直到遇到左括号或者栈为空为止」。
//
//然后就是对数字的处理，因为数字可能并不只有一位，所以遇到数字的时候要不停的累加。
//
//当遇到运算符或者括号的时候就将累加的数字加到后缀表达式中。

func getPolish(str string) []string {
	res := []string{}
	stack := []string{}
	arr := []byte(str)
	temp := -1
	for i := 0;i<len(arr);i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
			if temp == -1 {
				temp = int(arr[i] - '0')
			} else {
				temp = temp*10 + int(arr[i] - '0')
			}
		} else {
			if temp != -1 {
				res = append(res,strconv.Itoa(temp))
				temp = -1
			}
			if arr[i] == '+' || arr[i] == '-' || arr[i] == '*' || arr[i] == '/' {
				for len(stack) > 0 {
					if stack[len(stack)-1] == "(" {
						break
					}
					res = append(res,stack[len(stack)-1])
					stack = stack[:len(stack)-1]
				}
				stack = append(stack,string(arr[i]))
			} else {
				if arr[i] == '(' {
					stack = append(stack,string(arr[i]))
				}
				if arr[i] == ')' {
					for len(stack) > 0 && stack[len(stack)-1] != "(" {
						res = append(res,stack[len(stack)-1])
						stack = stack[:len(stack)-1]
					}
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	if temp != -1 {
		res = append(res,strconv.Itoa(temp))
	}
	for len(stack) > 0 {
		res = append(res,stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return res
}

// lc 150
func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0;i<len(tokens);i++ {
		if v,err := strconv.Atoi(tokens[i]);err == nil {
			// push
			stack = append(stack,v)
		} else {
			// pop
			op1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// pop
			op2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result := 0
			switch tokens[i] {
			case "*":
				result = op2 * op1
			case "/":
				result = op2 / op1
			case "+":
				result = op2 + op1
			case "-":
				result = op2 - op1
			}
			stack = append(stack,result)
		}
	}
	return stack[0]
}