package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 两个已经有序的数组，合并成一个有序的

func merge(input []int,left ,middle,right int, temp []int) {
	start := left
	start1 := left
	end1 := middle
	start2 := middle + 1
	end2 := right
	for start1 <= end1 && start2 <= end2 {
		if input[start1] < input[start2] {
			temp[start] = input[start1]
			start1 ++
		} else {
			temp[start] = input[start2]
			start2 ++
		}
		start ++
	}
	for start1 <= end1 {
		temp[start] = input[start1]
		start1 ++
		start ++
	}
	for start2 <= end2 {
		temp[start] = input[start2]
		start2 ++
		start ++
	}
	for left <= right {
		input[left] = temp[left]
		left ++
	}
}

func mergeSort(input []int,left,right int,temp []int) {
	if left < right {
		middle := (left+right)/2
		mergeSort(input,left,middle,temp)
		mergeSort(input,middle+1,right,temp)
		merge(input,left,middle,right,temp)
	}
}
func MergeSort(input []int) {
	temp := make([]int,len(input))
	mergeSort(input,0,len(input)-1,temp)
}

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func main() {
	slice := GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	MergeSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}