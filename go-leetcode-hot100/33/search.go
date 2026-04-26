package main

/*
33. 搜索旋转排序数组
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
 */

//
// 在常规二分搜索的时候查看当前 mid 为分割位置分割出来的两个部分 [l, mid] 和 [mid + 1, r] 哪个部分是有序的，
// **并根据有序的那个部分确定我们该如何改变二分搜索的上下界，因为我们能够根据有序的那部分判断出 target 在不在这个部分**

// 如果 [l, mid - 1] 是有序数组，且 target 的大小满足 [nums[l],nums[mid]，
// 则我们应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找。
// 如果 [mid, r] 是有序数组，且 target 的大小满足 nums[mid+1],nums[r]，
// 则我们应该将搜索范围缩小至 [mid + 1, r]，否则在 [l, mid - 1] 中寻找。

func search(nums []int, target int) int {
	left,right := 0,len(nums) - 1

	for left <= right {
		mid := (left+right)/2
		if nums[mid] == target {
			return mid
			// 如果只有2个元素，mid就是left,这里要注意
		} else if nums[mid] >= nums[left] { // mid落在左边
			// target 在 nums[left],nums[mid] 范围内
			if nums[left] <= target && target < nums[mid] {
				right = mid -1
			} else {
				left = mid + 1
			}
		} else {  // mid落在右边
			// target 在nums[right] ,nums[mid] 范围内
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

