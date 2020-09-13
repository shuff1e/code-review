package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 31：栈的压入、弹出序列
// 题目：输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是
// 否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1、2、3、4、
// 5是某栈的压栈序列，序列4、5、3、2、1是该压栈序列对应的一个弹出序列，但
// 4、3、5、1、2就不可能是该压栈序列的弹出序列。

func isValid(arr1,arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	index2 := 0
	stack := linkedliststack.New()
	for i := 0;i<len(arr1);i++ {
		stack.Push(arr1[i])
		for {
			temp,ok := stack.Peek()
			if !ok {
				break
			}
			if temp.(int) == arr2[index2] {
				index2 ++
				stack.Pop()
			} else {
				break
			}
		}
	}
	return stack.Empty()
}

func main() {
	Test([]int{1, 2, 3, 4, 5},[]int{4, 5, 3, 2, 1},true)
	Test([]int{1, 2, 3, 4, 5},[]int{3, 5, 4, 2, 1},true)
	Test([]int{1, 2, 3, 4, 5},[]int{4, 3, 5, 1, 2},false)
	Test([]int{1, 2, 3, 4, 5},[]int{3, 5, 4, 1, 2},false)
	Test([]int{1},[]int{2},false)
	Test([]int{1},[]int{1},true)
	Test([]int{},[]int{},true)
	Test(nil,nil,true)
}

func Test(arr1,arr2 []int,valid bool) {
	fmt.Println(isValid(arr1,arr2) == valid)
}