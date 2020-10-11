package main

/*
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1:

输入: [3,2,3]
输出: 3
示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2
 */

// 是之前的元素，+1
// 不是减一
// 如果为0 重置

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	count := 1
	prev := nums[0]
	for i := 1;i<len(nums);i++ {
		if count == 0 {
			count = 1
			prev = nums[i]
			continue
		}
		if nums[i] == prev {
			count ++
		} else {
			count --
		}
	}
	return prev
}