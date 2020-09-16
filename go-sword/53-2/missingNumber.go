package main

import "fmt"

// 53（二）：0到n-1中缺失的数字
// 题目：一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字
// 都在范围0到n-1之内。在范围0到n-1的n个数字中有且只有一个数字不在该数组
// 中，请找出这个数字。

// A：mid=arr[mid]，缺失的数字在右边
// mid < arr[mid]，在左边
// mid > arr[mid]，错误

// start==end && arr[start] == start,return start+1
// start==end && arr[start] != start,return start-1

// 0 1 2 3 4 6 7

// 0 1 2 3 4 5 6

// 0 2 4 5 6 7 9

// 0 1 1 1 1 1 1

func getMissing(arr []int) int {
	// 注意边界条件
	if len(arr) == 0 {
		return -1
	}
	return getMissingNumber(arr,0,len(arr)-1)
}

func getMissingNumber(arr []int,start,end int) int {
	if start > end {
		return arr[start] - 1
	}
	if start == end {
		if arr[start] == start {
			return arr[start] + 1
		} else if arr[start] > start {
			return arr[start] - 1
		} else {
			return -1
		}
	}
	mid := (start + end)/2
	if arr[mid] == mid {
		return getMissingNumber(arr,mid+1,end)
	} else if arr[mid] > mid {
		return getMissingNumber(arr,start,mid-1)
	} else {
		return -1
	}
}

func main() {
	Test([]int{ 1, 2, 3, 4, 5 },0)
	Test([]int{ 0, 1, 2, 3, 4 },5)
	Test([]int{ 0, 1, 2, 4, 5 },3)
	Test([]int{ 0, 1, 3, 4, 5},2)
	Test([]int{ 1 },0)
	Test([]int{ 0 },1)
	Test(nil,-1)
	Test([]int{1,1,1,1},-1)
	Test([]int{},-1)
	Test([]int{1,1,2,2,3,3,4},-1)
	Test([]int{1,2},0)
}

func Test(arr []int,expected int) {
	fmt.Println(arr,getMissing(arr))
	if getMissing(arr) != expected {
		panic("fuck")
	}
}

