package main

import (
	"algorithm/common/help"
	"fmt"
)

func valid(in []int,out []int) bool {
	if len(in) == 0 {
		return false
	}

	if len(in) != len(out) {
		return false
	}
	inIndex := 0
	outIndex := 0
	stack := *help.NewStack()
	for inIndex < len(in) {
		stack.Push(in[inIndex])
		inIndex ++
		for stack.Length() > 0{
			temp,_ := stack.Peek();
			if temp == out[outIndex] {
				stack.Pop()
				outIndex ++
			} else {
				break
			}
		}
	}
	fmt.Println(stack.Length())
	fmt.Println(outIndex)
	if stack.Length() != 0 || outIndex != len(out) {
		return false
	}
	return true
}

func main() {
	in := []int{1,2,3,4,5,6,7,8}
	out := []int{3,2,1,5,4,8,7,6}
	flag := valid(in,out)
	fmt.Println(flag)
}