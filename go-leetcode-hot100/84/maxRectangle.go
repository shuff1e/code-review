package main

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

// 单调栈，出栈的时候，结算
// 单调递增，可以相等，遇到第一个小的元素的时候，就知道左右边界
// 然后计算面积
func getMaxRectan(arr []int) int {
	res := 0
	stack := NewStack()
	// 如果单调递增的，没法出栈
	// 因此结尾搞个0
	arr = append(arr,0)
	for i:=0;i<len(arr);i++ {
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
	if x > y{
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
