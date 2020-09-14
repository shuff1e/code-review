package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 3（一）：找出数组中重复的数字
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。

// A：可以先排序，然后从前往后看哪些元素是重复的，不过这样就不能利用所有数字都在0到n-1的范围内这个特性
// 类似计数排序
// 尝试将数字m放到index为m的位置上，如果该位置已有的元素等于m，那么就发现了一个重复的元素

func findDuplcation(arr []int) int {
	for index := 0;index < len(arr); {
		if arr[index] == index {
			index ++
			continue
		} else if arr[index] == arr[arr[index]]{
			return arr[index]
		} else {
			swap(arr,index,arr[index])
		}
	}
	return -1
}

func swap(arr []int,x,y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func GenerateSlice(size int) []int {
	slice := make([]int,size,size)
	rand.Seed(time.Now().UnixNano())
	for i := 0;i<size;i++ {
		slice[i] = rand.Intn(size)
	}
	return slice
}

func main() {
	for i := 0;i<10;i++ {
		arr := GenerateSlice(10)
		fmt.Printf("%#v\n",arr)
		dup := findDuplcation(arr)
		fmt.Println(dup)
	}
}