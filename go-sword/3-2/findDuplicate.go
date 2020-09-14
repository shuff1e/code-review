package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

// A：数组长度为n+1，但是所有数字范围是1到n，那么肯定有重复的数字
// 如果数组中，1到n/2的范围内的数字为n/2个，则可以判断出n/2到n的范围内肯定有重复的数字
// 这样用二分查找

func findDuplication(arr []int) int {
	for start,end := 1,len(arr)-1;start<=end; {
		mid := ((end-start) >> 1) + start
		count := getCountByRange(arr,start,mid)
		// 对于[1,1]这种情况，不在这里判断的话，会一直循环下去
		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		if count <= (mid-start+1) {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}

func getCountByRange(arr []int,left,right int) int {
	count := 0
	for i := 0;i<len(arr);i++ {
		if arr[i] >=left && arr[i] <= right {
			count ++
		}
	}
	return count
}

func main() {
	for i := 0;i<10;i++ {
		slice := GenerateSlice(10)
		fmt.Printf("%#v\n",slice)
		dup := findDuplication(slice)
		fmt.Println(dup)
	}
}

func GenerateSlice(size int) []int {
	slice := make([]int,size,size)
	rand.Seed(time.Now().UnixNano())
	for i := 0;i<size;i++ {
		slice[i] = rand.Intn(size-1)+1
	}
	return slice
}