package main

import "fmt"

func main() {
	arr := []int{1,1,2,2,3,5,6}
	k := 4
	result := help(arr,k)
	// 2
	fmt.Println(result)
}

func help(arr []int,k int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := (l + r)/2
		if arr[mid] < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
