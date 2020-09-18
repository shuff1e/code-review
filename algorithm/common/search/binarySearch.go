package main

import (
	"algorithm/common/help"
	"fmt"
)

func binarySearch(array []int,value int) int{
	length := len(array)
	if length == 0 {
		return -1
	}
	left := 0
	right := length -1
	mid := (left + right)/2

	if array[mid] == value {
		return mid
	} else if array[mid] > value {
		right = mid - 1
	} else {
		left = mid + 1
	}
	return binarySearch(array[left:right+1],value)
}

func main() {
	source := []int{1,2,3,4,12,23,34,44,47}
	value := 12
	index := bSearch(source,value)
	fmt.Println(index)

	value = 13
	index = bSearch(source,value)
	fmt.Println(index)
}

func bSearch(array []int,value int) int {
	stack := &help.Stack{}
	left := 0
	right := len(array) -1
	for (right - left + 1) > 0 || stack.Length() > 0 {
		mid := (left+right)/2
		if array[mid] == value {
			return mid
		}
		if array[mid] > value {
			right = mid-1
		} else if array[mid] < value {
			left = mid + 1
		}
		fmt.Println(left,right)
	}
	return -1
}

