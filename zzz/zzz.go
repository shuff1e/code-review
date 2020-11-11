package main

import "fmt"

func main() {
	arr := []int{1,1,2,2,9,13,15}
	fmt.Println(lowBound(arr,5))
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

