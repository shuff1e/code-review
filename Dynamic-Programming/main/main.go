package main

import (
	"fmt"
	"github.com/shuff1e/code-review/Dynamic-Programming/array"
)

func main() {
	// array-01
	//arr := []int{2,1,5,3,6,4,8,9,7}
	//result := array.SolveArray01(arr)
	//fmt.Printf("%#v\n",result)
	// array-02
	str1 := "ABCBDAB"
	str2 := "BDCABA"
	fmt.Println(array.Lcse(str1,str2))
	// memo-1
	//n := 1
	//fmt.Println(memo.Fibonacci(n))
	// memo-2
	//n := 4
	//fmt.Println(memo.Fact(n))
}

