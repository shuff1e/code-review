package main

import "fmt"

func main() {
	arr := []int{1,2,5,15,16,21}
	fmt.Println(lowBound(arr,22))
}

func lowBound(arr []int,x int ) int {
	left := 0
	right := len(arr)-1
	for left <= right {
		mid := (left + right)/2
		if x < arr[mid] {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return right
}

