package main

import "fmt"

//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
//
//示例 1:
//
//输入: [1,2,1]
//输出: [2,-1,2]
//解释: 第一个 1 的下一个更大的数是 2；
//数字 2 找不到下一个更大的数；
//第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
//注意: 输入数组的长度不会超过 10000。

// 示例 2：
// 输入：[2,1,2,4,3]
// 输出：[4,2,4,-1,4]

// https://labuladong.gitbook.io/algo/gao-pin-mian-shi-xi-lie/pan-duan-hui-wen-lian-biao
// https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/dan-tiao-zhan

// 单调栈，单调递减

// 类似山峰的问题
// 将原始数组“翻倍”，就是在后面再接一个原始数组，
// 这样的话，按照之前“比身高”的流程，每个元素不仅可以比较自己右边的元素，而且也可以和左边的元素比较了。

// for range
//		 !empty && peek < arr[i]
//	  push

func main() {
	arr := []int{2,1,2,4,3}
	arr = []int{1,2,1}
	result := nextGreaterElements(arr)
	fmt.Printf("%#v\n",result)
}

// 单调栈，入栈之前结算
func nextGreaterElements(nums []int) []int {
	result := make([]int,len(nums))
	stack := NewStack()
	for i := 2*len(nums)-1;i>=0;i-- {
		for {
			if stack.Empty() {
				break
			}
			temp,_ := stack.Peek()
			if nums[temp.(int)%len(nums)] > nums[i%len(nums)] {
				break
			}
			stack.Pop()
		}
		if stack.Empty() {
			result[i%len(nums)] = -1
		} else {
			temp2,_ := stack.Peek()
			result[i%len(nums)] = nums[temp2.(int)%len(nums)]
		}
		stack.Push(i)
	}
	return result
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

