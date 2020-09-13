package main

import (
	"fmt"
	"unsafe"
)

// Q：n个数中，取k个，求这k个数的全组合

// A：每个数，都可以选择，也可以不选择

func combine(str string,k int) {
	if len(str) < k {
		return
	}
	if k == 0 {
		return
	}
	result := make([]byte,0)
	combineHelp(str,k,0,result)
}

func combineHelp(str string,k ,level int,result []byte) {
	if len(result) == k {
		fmt.Println(String(result))
		return
	}

	if len(str) - level < k - len(result) {
		return
	}

	if len(str) - level == k - len(result) {
		result = append(result,str[level:]...)
		fmt.Println(String(result))
		return
	}
	// 选择
	result = append(result,str[level])
	combineHelp(str,k,level+1,result)
	result = result[0:len(result)-1]

	// 不选择
	combineHelp(str,k,level+1,result)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}
func main() {
	Test("abcd",2)
	Test("abcdefg",3)
	Test("",0)
}

func Test(str string,k int) {
	combine(str,k)
}

