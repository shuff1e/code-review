package main

import "fmt"

//给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
//
//示例 1:
//
//输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
//输出: [-1,3,-1]
//解释:
//对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
//对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
//对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
//示例 2:
//
//输入: nums1 = [2,4], nums2 = [1,2,3,4].
//输出: [3,-1]
//解释:
//    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
//对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。
//

// A：单调栈，单调递减，从右向左

func main() {
	arr1 := []int{2,4}
	arr2 := []int{1,2,3,4}
	result := nextGreaterElement(arr1,arr2)
	fmt.Printf("%#v\n",result)
	arr1 = []int{4,1,2}
	arr2 = []int{1,3,4,2}
	result = nextGreaterElement(arr1,arr2)
	fmt.Printf("%#v\n",result)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := NewStack()
	for i := len(nums2) - 1;i>=0;i-- {
		for {
			if stack.Empty() {
				break
			}
			temp,_ := stack.Peek()
			if nums2[temp.(int)] > nums2[i] {
				break
			}
			stack.Pop()
		}
		if stack.Empty() {
			m[nums2[i]] = -1
		} else {
			temp,_ := stack.Peek()
			m[nums2[i]] = nums2[temp.(int)]
		}
		stack.Push(i)
	}

	result := make([]int,len(nums1))
	for i := 0;i<len(nums1);i++ {
		result[i] = m[nums1[i]]
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

