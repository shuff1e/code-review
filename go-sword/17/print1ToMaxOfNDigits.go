package main

import (
	"fmt"
	"unsafe"
)

// 17：打印1到最大的n位数
// 题目：输入数字n，按顺序打印出从1最大的n位十进制数。比如输入3，则
// 打印出1、2、3一直到最大的3位数即999。

// A：考虑到可能溢出的问题，使用字符串加法

// 将数字加1
func increment(arr []byte) bool {
	over := byte(0)
	for i := len(arr) - 1;i>=0;i-- {
		sum := arr[i] - '0' + over
		if i == len(arr) - 1 {
			sum ++
		}
		if sum >=10 {
			if i ==0 {
				return true
			}
			arr[i] = sum - 10 + '0'
			over = 1
		} else {
			arr[i] = sum + '0'
			break
		}
	}
	return false
}

func myPrintHelper(arr []byte) {
	for i := 0;i<len(arr);i++ {
		if arr[i] != '0' {
			fmt.Println(String(arr[i:]))
			break
		}
	}
}

func myPrint(n int) {
	arr := make([]byte,n)
	for i := 0;i<len(arr);i++ {
		arr[i] = '0'
	}
	for !increment(arr) {
		myPrintHelper(arr)
	}
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

func main() {
	//myPrint(3)
	doBetter(1)
}

// A：转换为0~9一共9个数组的全排列的问题

func doBetter(n int) {
	arr := make([]byte,n)
	better(0,arr)
}
func better(index int,arr []byte) {
	if index == len(arr) {
		myPrintHelper(arr)
		return
	}

	for i := byte(0);i<=9;i++ {
		arr[index] = '0' + i
		better(index+1,arr)
	}
}