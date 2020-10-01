package main

import "fmt"

/*
45. 跳跃游戏 II
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
说明:

假设你总是可以到达数组的最后一个位置。

 */

// A：dp[i] = min(dp[i+1],dp[i+2])+1
// dp[len(nums)-1] = 0

// [2,3,1,2,4,2,3]
// max = 2 4
// end = 2 4
// step = 1

func main() {
	arr := []int{2,3,0,1,4}
	fmt.Println(jump3(arr))
}

func jump(nums []int) int {
	const MAX = 0x7fffffff
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int,len(nums))
	for i := 0;i<len(nums);i++ {
		dp[i] = MAX
	}
	dp[len(nums) - 1] = 0

	for i := len(nums) - 2;i>=0;i-- {
		for step := 1;step<=nums[i];step++ {
			if i + step < len(nums) {
				dp[i] = Min(dp[i],dp[i+step]+1)
			}
		}
	}
	return dp[0]
}

func Min(x,y int) int {
	if x > y {
		return y
	}
	return x
}

// 贪心，每次都选离最后一个位置最远的，能跳到最后一个位置的index

func jump2(nums []int) int {
	end := len(nums)-1
	steps := 0
	for end > 0 {
		for i := 0 ;i < len(nums) - 1;i++ {
			if i + nums[i] >= end {
				end = i
				steps ++
				break
			}
		}
	}
	return steps
}

// 正向计算，每次都选择能跳的最远的位置

func jump3(nums []int) int {
	maxPosition := 0
	end := 0
	steps := 0
	for i := 0;i<len(nums) && end < len(nums) - 1;i++ {
		maxPosition = Max(maxPosition,i+nums[i])
		if i == end {
			end = maxPosition
			steps ++
		}
	}
	return steps
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}