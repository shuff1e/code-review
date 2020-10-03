package main

import "fmt"

/*
81. 搜索旋转排序数组 II
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:

输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
进阶:

这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
 */

// mid > left (left，mid有序）
// mid < right (mid,right有序）
// mid == left

func main() {
	arr := []int{2,5,6,0,0,1,2}
	target := 0
	fmt.Println(search(arr,target))
	target = 4
	fmt.Println(search(arr,target))
}

func search(nums []int, target int) bool {
	l,r := 0,len(nums) - 1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return true
		} else if nums[l] < nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if nums[mid] < nums[r] {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid -1
			}
		} else {
			// 10111 和 1110111101 这种。此种情况下 nums[start] == nums[mid]，
			// 分不清到底是前面有序还是后面有序，此时 start++ 即可。相当于去掉一个重复的干扰项。
			//
			//l ++
			return help(nums[l:r+1],target)
		}
	}
	return false
}

func help(arr []int,target int) bool {
	for i := 0;i<len(arr);i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func search2(nums []int, target int) bool {
	l,r := 0,len(nums) - 1
	for l <= r {
		mid := (l+r)/2
		if nums[mid] == target {
			return true
		} else if nums[l] == nums[mid] {
			// 10111 和 1110111101 这种。此种情况下 nums[start] == nums[mid]，
			// 分不清到底是前面有序还是后面有序，此时 start++ 即可。相当于去掉一个重复的干扰项。
			//
			l ++

		} else if nums[l] < nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid -1
			}
		}
	}
	return false
}