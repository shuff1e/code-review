package main

import (
	"fmt"
	"github.com/shuff1e/code-review/Dynamic-Programming/array"
)

func main() {
	// array-01
	arr := []int{2,1,5,3,6,4,8,9,7}
	result := array.SolveArray01(arr)
	fmt.Printf("%#v\n",result)
}
