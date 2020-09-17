package main

import (
	"fmt"
	"unsafe"
)

// 58（一）：翻转单词顺序
// 题目：输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，
// 则输出"student. a am I"。

// A：先翻转整个句子，再反转其中的每个单词

func reverse(arr []byte,start,end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}

func reverseAll(arr []byte) []byte {
	reverse(arr,0,len(arr)-1)

	first := 0

	for {
		first = findFirstNotEmpty(arr,first)
		if first == -1 {
			return arr
		}
		last := findLastNotEmpty(arr,first)
		reverse(arr,first,last)
		first = last + 1
	}
}

func findFirstNotEmpty(arr []byte,start int) int {
	for i := start;i<len(arr);i++ {
		if arr[i] != ' ' {
			return i
		}
	}
	return -1
}

func findLastNotEmpty(arr []byte,start int) int {
	for start < len(arr) && arr[start] != ' ' {
		start ++
	}
	return start - 1
}


func main() {
	Test("Test1","I am a student.","student. a am I")
	Test("Test2","Wonderful","Wonderful")
	Test("Test3","","")
	Test("Test4","   ","   ")
}

func Test(name string,input,expected string) {
	fmt.Println(name)
	input1 := []byte(input)
	expected1 := []byte(expected)
	fmt.Println(String(reverseAll(input1)) == String(expected1))
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}