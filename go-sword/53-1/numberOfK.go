package main

// 53-1（一）：数字在排序数组中出现的次数
// 题目：统计一个数字在排序数组中出现的次数。例如输入排序数组{1, 2, 3, 3,
// 3, 3, 4, 5}和数字3，由于3在这个数组中出现了4次，因此输出4。

// A：二分法找到第一个3和最后一个3
// 如果mid左边不是3，mid是3，则找到了第一个3
// 如果mid左边是3，则在start,mid-1里面找

func getFirstOccur(arr []int,start,end,k int) int {
	if start > end {
		return -1
	}
	if start == end {
		if arr[start] != k {
			return -1
		} else {
			return start
		}
	}

	mid := (start+end)/2

	if arr[mid] < k {
		return getFirstOccur(arr,mid+1,end,k)
	} else if arr[mid] > k {
		return getFirstOccur(arr,0,mid-1,k)
	} else {
		if mid > 0 && arr[mid-1] == k {
			return getFirstOccur(arr,0,mid-1,k)
		} else {
			return mid
		}
	}
}

func getLastOccur(arr []int,start,end,k int) int {
	if start > end {
		return -1
	}
	if start == end {
		if arr[start] != k {
			return -1
		} else {
			return start
		}
	}

	mid := (start+end)/2
	if arr[mid] < k {
		return getLastOccur(arr,mid+1,end,k)
	} else if arr[mid] > k {
		return getLastOccur(arr,0,mid-1,k)
	} else {
		if mid < end && arr[mid+1] == k {
			return getLastOccur(arr,mid+1,end,k)
		} else {
			return mid
		}
	}
}

func getLength(arr []int,k int) int {
	left := getFirstOccur(arr,0,len(arr)-1,k)
	if left == -1 {
		return 0
	}
	right := getLastOccur(arr,0,len(arr)-1,k)
	return right-left+1
}
func main() {
	Test([]int{1, 2, 3, 3, 3, 3, 4, 5},3,4)
	Test([]int{3, 3, 3, 3, 4, 5},3,4)
	Test([]int{1, 2, 3, 3, 3, 3},3,4)
	Test([]int{1, 3, 3, 3, 3, 4, 5},2,0)
	Test([]int{1, 3, 3, 3, 3, 4, 5},0,0)
	Test([]int{1, 3, 3, 3, 3, 4, 5},6,0)
	Test([]int{3, 3, 3, 3},3,4)
	Test([]int{3, 3, 3, 3},4,0)
	Test([]int{3},3,1)
	Test([]int{3},4,0)
	Test(nil,0,0)
}

func Test(arr []int,k ,expected int) {
	if getLength(arr,k) != expected {
		panic("fuck")
	}
}