package main

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
 */

func searchRange(nums []int, target int) []int {
	left,right := 0,len(nums)-1
	index := -1
	for left <= right {
		mid := (left + right)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if index == -1 {
		return []int{-1,-1}
	}
	first := findFirst(nums,target,0,index)
	last := findLast(nums,target,index,len(nums)-1)
	return []int{first,last}
}

// 找到第一个
func findFirst(arr []int,target ,start ,end int) int {
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			if mid - 1 >= start && arr[mid] != arr[mid-1] {
				return mid
			} else if mid == start{
				return mid
			} else {
				end = mid - 1
			}
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// 找到最后一个
func findLast(arr []int,target ,start ,end int) int {
	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			if mid == end {
				return mid
			} else if arr[mid] != arr[mid+1] {
				return mid
			} else {
				start = mid + 1
			}
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

