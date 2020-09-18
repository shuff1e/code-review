package main

import (
	"fmt"
)

// 选出基准，小于基准的放左边，大于基准的放右边

func partition(arr []int,start,end int) int {
	pivot := arr[start]
	less := start - 1
	more := end + 1

	for start < more {
		if arr[start] < pivot {
			less ++
			swap(arr,start,less)
			start ++
		} else if arr[start] > pivot {
			more--
			swap(arr,start,more)
		} else {
			start ++
		}
	}
	return less + 1
	//right := end
	//for i := start;i<right;i++ {
	//	for arr[i] > pivot {
	//		swap(arr,i,right)
	//		right --
	//	}
	//}
	//swap(arr,start,right)
	//return right
	//left := start
	//right := end
	//for left < right {
	//	for left < right && arr[right] > pivot {
	//		right --
	//	}
	//	for left < right && arr[left] <= pivot {
	//		left ++
	//	}
	//	if left < right {
	//		swap(arr,left,right)
	//	}
	//}
	//swap(arr,start,left)
	//return left
	//mark := start
	// 遇到小的，就替换位置
	//for i := start;i<=end;i++ {
	//	if pivot > arr[i] {
	//		mark ++
	//		swap(arr,mark,i)
	//	}
	//}
	//swap(arr,start,mark)
	//return mark
}

func quickSort(a []int,start,end int) {
	if start >= end {
		return
	}
	pivot := partition(a,start,end)
	quickSort(a,start,pivot-1)
	quickSort(a,pivot+1,end)
}

func QuickSort(a []int) {
	quickSort(a,0,len(a)-1)
}

func swap(a []int,left,right int) {
	temp := a[left]
	a[left] = a[right]
	a[right] = temp
}

func main() {
	//slice := help.GenerateSlice(20)
	slice := []int{3,7,5,4,2,1,9}
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	QuickSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
