package main

import "fmt"

/*
85. 最大矩形
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
["1","0","1","0","0"],
["1","0","1","1","1"],
["1","1","1","1","1"],
["1","0","0","1","0"]
]
输出: 6
 */

//

// 每一列都是一个矩形
// 随着row不断增大，列不断扩展

func main() {
	matrix := [][]byte{
		{'1','0','1','0','0'},
	{'1','0','1','1','1'},
	{'1','1','1','1','1'},
	{'1','0','0','1','0'},
	}
	fmt.Println(maximalRectangle(matrix))
	arr := []int{3, 1, 3, 2, 2}
	fmt.Println(getMaxArea(arr))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxArea := 0
	dp := make([]int,len(matrix[0]))
	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0]);j++ {
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		maxArea = Max(maxArea,getMaxArea(dp))
	}
	return maxArea
}

// 单调栈，得到最大的矩形
// 单调递增，出栈的时候结算
// for range arr
// 		for !stack.empty() && stack.peek() > arr[i]
// 		     stack.pop()
//			 left := stack.empty()? : -1 : stack.peek()
// 			 right := i - 1
// 	  stack.push(arr[i])

func getMaxArea(arr []int) int {
	maxArea := 0
	arr = append(arr,0)
	stack := []int{}
	for i := 0;i<len(arr);i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			temp := stack[len(stack)-1]
			stack = stack[0:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			maxArea = Max(maxArea,(i-1-left)*arr[temp])
		}
		stack = append(stack,i)
	}
	return maxArea
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}