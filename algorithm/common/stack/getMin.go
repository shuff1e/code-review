package main

import (
	"algorithm/common/help"
	"errors"
	"fmt"
	"sync"
)

// 设计一个stack，具有getMin的功能

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []int
}

func NewStack() *stack {
	return &stack {sync.Mutex{}, make([]int,0), }
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()


	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) Peek() (int,error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	return s.s[l-1],nil
}

func (s *stack) length() int {
	return len(s.s)
}

type minStack struct {
	data stack
	min stack
}

func (ms *minStack) push(v int) {
	if ms.min.length() == 0 {
		ms.min.Push(v)
	} else if temp,_ := ms.min.Peek(); temp >= v {
		ms.min.Push(v)
	}

	ms.data.Push(v)
}

func (ms *minStack) pop() int {
	result,_ := ms.data.Pop()
	if temp,_ := ms.min.Peek();temp == result {
		ms.min.Pop()
	}
	return result
}

func (ms *minStack) getMin() int {
	result,_ := ms.min.Peek()
	return result
}

func main() {
	ms := &minStack{
		data: stack{s: []int{}},
		min:  stack{s: []int{}},
	}
	slice := help.GenerateSlice(20)


	for _,v := range slice {
		ms.push(v)
	}

	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")

	for i := 0;i<10;i++ {
		fmt.Println(ms.getMin())
		ms.pop()
	}
}

func bubbleSort(a []int) {
	length := len(a)
	// i表明对每个元素
	// j表明对每个元素的冒泡
	for i:=0;i<length-1;i++ {
		for j:=0;j<length-1-i;j++ {
			if a[j+1] < a[j] {
				help.Swap(a,j,j+1)
			}
		}
	}
}