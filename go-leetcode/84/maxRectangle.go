package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
//图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
//
//示例:
//
//输入: [2,1,5,6,2,3]
//输出: 10

// A：单调栈，单调递增，arr[i]小于栈中的数字，数字出栈时，计算面积

/*
模板

for i:= 0;i<len(arr);i++ {
	for !stack.empty() && stack.peek() > arr[i] {
		height := arr[stack.pop()]
		left := stack.empty() ? -1 : stack.peek()
		right := i - 1
	}
	stack.push(i)
}
 */

// https://zhuanlan.zhihu.com/p/26465701
// https://oi-wiki.org/ds/monotonous-stack/
// https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/dan-tiao-zhan

func largestRectangleArea(heights []int) int {
	return getMaxRectan(heights)
}

func getMaxRectan(arr []int) int {
	// 如果是单调递增的，没有出栈的机会
	// 不会计算面积，因此加上一个最小值0
	res := 0
	arr = append(arr,0)
	stack := linkedliststack.New()
	for i := 0;i<len(arr);i++ {
		for {
			if stack.Empty() {
				break
			}
			temp,_ := stack.Peek()
			if arr[temp.(int)] <= arr[i] {
				break
			}
			stack.Pop()
			left := -1
			if !stack.Empty() {
				temp2,_ := stack.Peek()
				left = temp2.(int)
			}
			right := i - 1
			res = Max(res,(right-left)*arr[temp.(int)])
		}
		stack.Push(i)
	}
	return res
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{2,1,5,6,2,3}
	arr = []int{1,2,3,4,5,6}
	arr = []int{1}
	fmt.Println(getMaxRectan(arr))
}