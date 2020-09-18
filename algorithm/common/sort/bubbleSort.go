package main

import (
	"algorithm/common/help"
	"fmt"
)

func bubbleSort(a []int) {
	length := len(a)
	// i表明对每个元素
	// j表明对每个元素的冒泡
	for i:=0;i<length-1;i++ {
		for j:=0;j<length-1-i;j++ {
			if a[j+1] < a[j] {
				help.Swap(a,j,j+1)
			}
		}
	}
}

func main() {
	slice := help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}