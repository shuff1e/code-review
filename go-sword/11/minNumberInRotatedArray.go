package main

import "fmt"

// 11：旋转数组的最小数字
// 题目：把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如数组
// {3, 4, 5, 1, 2}为{1, 2, 3, 4, 5}的一个旋转，该数组的最小值为1。

// A：二分查找，需要每次能排除一边的元素
// 旋转后的数组，左边的元素递增，到中间突降，右边的元素仍然是递增的
// 中间元素mid
// 如果mid>最右边的元素，说明最小值肯定在mid右边
// 如果mid<最右边的元素，说明最小值在mid左边
// 如果mid=最右边的元素，最小值可能在mid右边，也可能在mid左边

func findMin(arr []int) int {
	left,right := 0,len(arr)-1
	for left < right {
		mid := (left + right) >> 1
		if arr[mid] < arr[right] {
			right = mid
		} else if arr[mid] > arr[right] {
			left = mid + 1
		} else {
			return searchMin(arr[left:right])
		}
	}
	return arr[left]
}

func searchMin(arr []int) int {
	min := arr[0]
	for i := 1;i<len(arr);i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	arr := []int{ 3, 4, 5, 1, 2 }
	fmt.Println(findMin(arr))
	arr = []int{ 3, 4, 5, 1, 1, 2 }
	fmt.Println(findMin(arr))
	arr = []int{ 3, 4, 5, 1, 2, 2 }
	fmt.Println(findMin(arr))
	arr = []int{ 1, 0, 1, 1, 1 }
	fmt.Println(findMin(arr))
	arr = []int{ 1, 2, 3, 4, 5 }
	fmt.Println(findMin(arr))
	arr = []int{ 2 }
	fmt.Println(findMin(arr))
}
