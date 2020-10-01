package main

import (
	"fmt"
)

/*
41. 缺失的第一个正数
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1


提示：

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。

 */

// A：可将所有的数放入hash表中
// 然后从1开始枚举，看该数字是否在哈希表中
// 这样时间和空间复杂度都是O(n)

// 但是要求常数空间复杂度
// 可以将原数组改造成一个hash表

// 对于原数组中的元素i,可以将数组中的i-1位置打上标记（设置为负数）
// 原来为负数的设置为len(arr)+1

// 未出现的最小正整数肯定在[1,len(arr)+1]范围内

//例如，-1,3,1,7,4，未出现的为5

//
// -1,1,3,7,4
// 1,-1,3,7,4

// -1,1,3,4,7

// nums[nums[i] - 1] != nums[i])
// i = 3
// nums[3]=5
// nums[4] = nums[3] = 5
// 1,1,1,5,5

func main() {
	arr := []int{7,8,9,11,12}
	arr = []int{3,4,-1,1}
	fmt.Println(firstMissingPositive(arr))
}

func firstMissingPositive(nums []int) int {
	for i := 0;i<len(nums);i++ {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for i := 0;i<len(nums);i++ {
		abs := Abs(nums[i])
		if abs <= len(nums) && nums[abs-1] > 0 {
			nums[abs-1] = -nums[abs-1]
		}
	}
	for i := 0;i<len(nums);i++ {
		if nums[i] > 0 {
			return i+1
		}
	}
	return len(nums) + 1
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// A：置换，将x放在index为x-1的位置上
// 由于每次的交换操作都会使得某一个数交换到正确的位置，因此交换的次数最多为 N，整个方法的时间复杂度为 O(N)

func firstMissingPositive2(nums []int) int {
	for i := 0 ;i<len(nums);i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1]{
			swap(nums,i,nums[i]-1)
		}
	}
	for i := 0;i<len(nums);i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func swap(nums []int,i,j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}