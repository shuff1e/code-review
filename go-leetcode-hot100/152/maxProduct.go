package main

import "fmt"

/*
152. 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:

输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
示例 2:

输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
 */

// 连续子数组

func main() {
	arr := []int{-1,-2,3,4}
	arr = []int{-2,3,-4}
	//arr = []int{2,3,-2,4}
	//arr = []int{5,6,-3,4,-3}
	fmt.Println(maxProduct(arr))
}

// 我们可以根据正负性进行分类讨论。
//
// 考虑当前位置如果是一个负数的话，那么我们希望以它前一个位置结尾的某个段的积也是个负数，
// 这样就可以负负得正，并且我们希望这个积尽可能「负得更多」，即尽可能小。
// 如果当前位置是一个正数的话，我们更希望以它前一个位置结尾的某个段的积也是个正数，并且希望它尽可能地大。
// 于是这里我们可以再维护一个 fmin (i)，它表示以第 i 个元素结尾的乘积最小子数组的乘积，那么我们可以得到这样的动态规划转移方程：
// 		dpMax[i] = Max(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i],nums[i])
//		dpMin[i] = Min(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i],nums[i])

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dpMax := make([]int,len(nums))
	dpMin := make([]int,len(nums))

	dpMax[0] = nums[0]
	dpMin[0] = nums[0]
	result := dpMax[0]
	for i := 1;i<len(nums);i++ {
		dpMax[i] = Max(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i],nums[i])
		dpMin[i] = Min(dpMax[i-1]*nums[i],dpMin[i-1]*nums[i],nums[i])
		if dpMax[i] > result {
			result = dpMax[i]
		}
	}
	return result
}

func Max(x,y,z int) int {
	if x < y {
		x = y
	}
	if x < z {
		x = z
	}
	return x
}

func Min(x,y,z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}