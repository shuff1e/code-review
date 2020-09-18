package main

import (
	"algorithm/common/help"
	"fmt"
)

func selectSort(a []int) {
	length := len(a)
	// i对应每个位置，每个位置找到合适的元素
	for i:=0;i<length;i++ {
		min := i
		for j := i + 1;j<length;j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		help.Swap(a,min,i)
	}
}

func main() {
	slice := help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}