package main

import "fmt"

// 53（三）：数组中数值和下标相等的元素
// 题目：假设一个单调递增的数组里的每个元素都是整数并且是唯一的。请编程实
// 现一个函数找出数组中任意一个数值等于其下标的元素。例如，在数组{-3, -1,
// 1, 3, 5}中，数字3和它的下标相等。

// A：单调递增
// index的增长速度一定
// arr[index]的增长速度更高

// index > arr[index]，往右边找
// index < arr[index]，往左边找

// -3, -1, 1, 3, 5
func getIdenticalNumber(arr []int,start,end int) int {
	if start > end {
		return -1
	}
	mid := (start+end)/2
	if mid > arr[mid] {
		return getIdenticalNumber(arr,mid+1,end)
	} else if mid < arr[mid] {
		return getIdenticalNumber(arr,start,mid-1)
	} else {
		return mid
	}
}

func getIdenticcal(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	return getIdenticalNumber(arr,0,len(arr)-1)
}

func main() {
	Test([]int{ -3, -1, 1, 3, 5 },3)
	Test([]int{ 0, 1, 3, 5, 6 },0)
	Test([]int{ -1, 0, 1, 2, 4 },4)
	Test([]int{ -1, 0, 1, 2, 5 },-1)
	Test([]int{ 0 },0)
	Test([]int{ 10 },-1)
	Test([]int{},-1)
	Test(nil,-1)
}

func Test(arr []int,expected int) {
	fmt.Println(arr,getIdenticcal(arr))
	if getIdenticcal(arr) != expected {
		panic("fcuk")
	}
}