package main

import "fmt"

/*

324. 摆动排序 II
给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

示例 1:

输入: nums = [1, 5, 1, 1, 6, 4]
输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
示例 2:

输入: nums = [1, 3, 2, 2, 3, 1]
输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
说明:
你可以假设所有输入都会得到有效的结果。

进阶:
你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？

 */
func main() {
	arr := []int{1,3,2,2,3,1}
	arr = []int{5,3,1,2,6,7,8,5,5,5}
	arr = []int{5,3,1,2,6,7,8,5,5}
	wiggleSort(arr)
	fmt.Printf("%#v\n",arr)
}

func wiggleSort(nums []int)  {
	findMid(nums,0,len(nums)-1)
	pivot := nums[len(nums)/2]
	partition(nums,0,len(nums)-1,pivot)

	// 9个数，左边取5个
	// 10个数，左边取5个
	p1,p2 := (len(nums)-1)/2,len(nums) - 1
	result := []int{}
	for p1 >= 0 && p2 > len(nums)/2 {
		result = append(result,nums[p1],nums[p2])
		p1 --
		p2 --
	}
	for p1 >= 0 {
		result = append(result,nums[p1])
		p1 --
	}
	for p2 > (len(nums)-1)/2 {
		result = append(result,nums[p2])
		p2 --
	}
	for i := 0;i<len(result);i++ {
		nums[i] = result[i]
	}
}

func partition(arr []int,start,end int,pivot int) (int,int) {
	less := start - 1
	more := end + 1

	left := start

	for left < more {
		if arr[left] < pivot {
			less ++
			arr[less],arr[left] = arr[left],arr[less]
			left ++
		} else if arr[left] > pivot {
			more --
			arr[more],arr[left] = arr[left],arr[more]
		} else {
			left ++
		}
	}
	return less + 1,more-1
}

func findMid(arr []int,start,end int) {
	pivot := arr[start]
	mark := start
	for i := start;i <= end;i++ {
		if arr[i] < pivot {
			mark ++
			arr[i],arr[mark] = arr[mark],arr[i]
		}
	}
	arr[start],arr[mark] = arr[mark],arr[start]
	if mark == len(arr)/2 {
		return
	} else if mark > len(arr)/2 {
		findMid(arr,start,mark-1)
	} else {
		findMid(arr,mark+1,end)
	}
}