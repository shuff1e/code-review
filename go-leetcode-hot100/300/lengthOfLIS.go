package main

import "fmt"

/*

300. 最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:

可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
你算法的时间复杂度应该为 O(n2) 。
进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?

 */

func main() {
	arr := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS2(arr))
}

// 动态规划
// dp[i]表示以i位置结尾的，最长的递增子序列的长度
// dp[i] = Max(dp[j] + 1)  arr[j] < arr[i]

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	dp := make([]int,len(nums))
	for i := 0;i<len(nums);i++ {
		dp[i] = 1
		for j := i-1;j>=0;j-- {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i],dp[j]+1)
			}
		}
		result = Max(result,dp[i])
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

// dp[i] 表示长度为i的递增子序列中，结尾的数中，最小的
func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	length := 1
	dp := make([]int,len(nums) + 1)
	dp[length] = nums[0]

	for i := 1;i<len(nums);i++ {
		if nums[i] > dp[length] {
			length ++
			dp[length] = nums[i]
		} else {
			l,r := 0,length
			pos := 0

			for l <= r {
				mid := (l + r)/2
				if dp[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			dp[pos + 1] = nums[i]
		}
	}
	return length
}