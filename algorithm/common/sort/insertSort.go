package main

import (
	"algorithm/common/help"
	"fmt"
)

func insertSort(a []int) {
	for  i:=0;i<len(a);i++ {
		v := a[i]
		j := i-1
		for ;j>=0;j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		// 找位子，插入数据
		a[j+1] = v
	}
}

func main() {
	slice := help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}
