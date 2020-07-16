package Sort

import "github.com/shuff1e/code-review/util"

func partition1(arr []int,start,end int) int {
	pivot := arr[start]

	left := start
	right := end
	for left < right {
		for left < right && pivot < arr[right] {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}
		if left < right {
			util.Swap(arr,left,right)
		}
	}
	util.Swap(arr,start,left)
	return left
}

func partition2(arr []int,start,end int) int {
	pivot := arr[start]
	mark := start
	// 将小于pivot的数放到前面
	for i := start;i<=end;i++ {
		if arr[i] < pivot {
			mark ++
			util.Swap(arr,mark,i)
		}
	}
	util.Swap(arr,mark,start)
	return mark
}

// 荷兰国旗问题
// 问题：现有红白蓝三个不同颜色的小球，乱序排列在一起，
//请重新排列这些小球，使得红白蓝三色的同颜色的球在一起。
//这个问题之所以叫荷兰国旗问题，是因为我们可以将红白蓝三色小球想象成条状物，
//有序排列后正好组成荷兰国旗。

// 将小于pivot的放到前面
// 等于pivot的放到中间
// 大于pivot的放到后面
func partition3(arr []int,start,end int,pivot int) (int,int) {
	less := start - 1
	more := end + 1
	left := start

	for left < more {
		// 将小的放到前面
		if arr[left] < pivot {
			less ++
			util.Swap(arr,less,left)
			left ++
			// 大的放到后面
		} else if arr[left] > pivot {
			more --
			util.Swap(arr,left,more)
		} else {
			left ++
		}
	}
	return less + 1,more-1
}

func QuickSort(arr []int) {
	quickSort(arr,0,len(arr)-1)
}

func quickSort(arr []int,start,end int) {
	if start >= end {
		return
	}
	partition := partition2(arr,start,end)
	quickSort(arr,start,partition-1)
	quickSort(arr,partition+1,end)
}

func QuickSortNetherlandsFlag(arr []int) {
	quickSortNetherlandsFlag(arr,0,len(arr)-1)
}

func quickSortNetherlandsFlag(arr []int,start,end int) {
	if start >= end {
		return
	}
	less,more := partition3(arr,start,end,arr[start])
	quickSortNetherlandsFlag(arr,start,less-1)
	quickSortNetherlandsFlag(arr,more+1,end)
}