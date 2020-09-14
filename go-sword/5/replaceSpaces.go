package main

import (
	"fmt"
	"unsafe"
)

// 5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。
func replace(str string) string {
	spaceCount := 0
	for i := 0;i<len(str);i++ {
		if str[i] == ' ' {
			spaceCount ++
		}
	}

	result := make([]byte,len(str) + 2*spaceCount)
	index := 0
	for i := 0;i<len(str);i++ {
		if str[i] != ' ' {
			result[index] = str[i]
			index += 1
		} else {
			result[index] = '%'
			index += 1
			result[index] = '2'
			index += 1
			result[index] = '0'
			index += 1
		}
	}
	return String(result)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}
func main() {
	fmt.Println(replace("we are happy."))
	fmt.Println(replace("   "))
}
