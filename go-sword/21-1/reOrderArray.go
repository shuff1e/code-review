package main

import "fmt"

// 21-1：调整数组顺序使奇数位于偶数前面
// 题目：输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有
// 奇数位于数组的前半部分，所有偶数位于数组的后半部分。

// A：类似快速排序的partition，或者荷兰国旗问题

func partition(arr []int) {
	mark := 0
	for i := 0;i<len(arr);i++ {
		if arr[i]%2 == 1 {
			swap(arr,mark,i)
			mark += 1
		}
	}
}

func swap(arr []int,mark,i int) {
	temp := arr[mark]
	arr[mark] = arr[i]
	arr[i] = temp
}

func main() {
	Test([]int{1, 2, 3, 4, 5, 6, 7})
	Test([]int{2, 4, 6, 1, 3, 5, 7})
	Test([]int{1, 3, 5, 7, 2, 4, 6})
	Test([]int{1})
	Test([]int{2})
	Test([]int{})
	Test(nil)
}

func Test(arr []int) {
	printArray(arr)
	partition(arr)
	printArray(arr)
	fmt.Println()
}

func printArray(arr []int) {
	for _,v := range arr {
		fmt.Print(v," ")
	}
	fmt.Println()
}