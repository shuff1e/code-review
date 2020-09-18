package main

import (
	"algorithm/common/help"
	"fmt"
)

func heapSort(a []int) []int{
	heapSize := len(a)
	heap := make([]int,heapSize)
	for i,v := range a {
		heap[i] = v
		insert(heap,i)
	}

	for heapSize > 0 {
		help.Swap(heap,0,heapSize-1)
		heapSize--
		buildHeap(heap,0,heapSize)
	}
	return heap
}

func buildHeap(a []int,index,heapSize int) {
	largest := index
	left := 2*index +1
	right := 2*index +2
	for left < heapSize {
		if a[left] > a[index] {
			largest = left
		}
		if right < heapSize && a[right] > a[largest] {
			largest = right
		}
		if largest != index {
			help.Swap(a,index,largest)
		} else {
			break
		}
		index = largest
		left = 2*index+1
		right = 2 *index +2
	}
}

func insert(a []int,index int) {
	for index > 0 {
		parent := (index-1)/2
		if a[parent] < a[index] {
			help.Swap(a,parent,index)
		}
		index = parent
	}
}

func main() {
	slice := help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	slice = heapSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}