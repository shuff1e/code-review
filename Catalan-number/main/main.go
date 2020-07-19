package main

import (
	"fmt"
	Catalan_number "github.com/shuff1e/code-review/Catalan-number"
)

func main() {
	//fmt.Println(Catalan_number.BSTNumber(7))
	// catalan-4
	str := "(())())(())())()"
	fmt.Println(Catalan_number.GetMaxLength(str))
}
