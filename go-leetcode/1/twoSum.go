package main

import "fmt"

/*
1. 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

 */

// 类似 128 longestConsecutive 问题
// 使用map，一遍过

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(arr,target))
}

func twoSum(nums []int, target int) []int {
	mmp := make(map[int]int,0)
	for i,v := range nums {
		if index,ok := mmp[target - v];ok {
			return []int{i,index}
		}
		mmp[v] = i
	}
	return nil
}