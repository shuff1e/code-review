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
	//str1 := "ABCBDAB"
	//str2 := "BDCABA"
	//fmt.Println(array.Lcse(str1,str2))
	// array-03
	//str1 := "1234567ABCDEF123"
	//str2 := "23457ABCDEF456"
	//fmt.Println(array.Lcss(str1,str2))
	// array-05
	//arr := []int{5,10,25,1}
	//arr = []int{3,5}
	//aim := 2
	//fmt.Println(array.GetChange(arr,aim))
	// arr-06
	arr := []int{1,0,1,7,2,4}
	fmt.Println(array.GetDiff(arr))
	// memo-1
	//n := 1
	//fmt.Println(memo.Fibonacci(n))
	// memo-2
	//n := 4
	//fmt.Println(memo.Fact(n))
}

