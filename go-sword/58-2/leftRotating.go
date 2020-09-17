package main

import (
	"fmt"
	"unsafe"
)

// 58（二）：左旋转字符串
// 题目：字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。比如输入字符串"abcdefg"和数
// 字2，该函数将返回左旋转2位得到的结果"cdefgab"。

// A：类似58-1
// 先整体翻转字符串
// 再翻转部分

func reverse(arr []byte,k int) []byte {
	reverseHelp(arr,0,len(arr)-1)
	reverseHelp(arr,len(arr)-k,len(arr)-1)
	reverseHelp(arr,0,len(arr)-k-1)
	return arr
}

func reverseHelp(arr []byte,start,end int) {
	if start < 0 {
		return
	}

	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start ++
		end --
	}
}

func main() {
	Test("Test1","abcdefg",2,"cdefgab")
	Test("Test2","abcdefg",1,"bcdefga")
	Test("Test3","abcdefg",6,"gabcdef")
	Test("Test4","",6,"")
	Test("Test5","abcdefg",0,"abcdefg")
	Test("test6","abcdefg",7,"abcdefg")
}

func Test(name string,input string,k int,expected string) {
	fmt.Println(name)
	input1 := []byte(input)
	expected1 := []byte(expected)
	fmt.Println(String(reverse(input1,k)) == String(expected1))
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

