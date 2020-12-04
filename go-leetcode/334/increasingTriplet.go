package main

import "fmt"

/*

334. 递增的三元子序列
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

数学表达式如下:

如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

示例 1:

输入: [1,2,3,4,5]
输出: true
示例 2:

输入: [5,4,3,2,1]
输出: false

 */

// 寻找最长递增子序列
// dp[i] = dp[j] + 1

func main() {
	arr := []int{1,2,3,4,5}
	arr = []int{5,4,3,2,1}
	fmt.Println(increasingTriplet3(arr))
}

func increasingTriplet(nums []int) bool {
	max := 0
	dp := make([]int,len(nums))
	for i := 0;i<len(nums);i++ {
		dp[i] = 1
		for j := 0;j<i;j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i],dp[j] + 1)
				max = Max(dp[i],max)
				if max == 3 {
					return true
				}
			}
		}
	}
	return false
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func increasingTriplet2(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	dp := make([]int,len(nums))
	ends := make([]int,len(nums))

	dp[0] = 1
	ends[0] = nums[0]

	right,l,m,r := 0,0,0,0

	for i := 0;i<len(nums);i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r)/2
			if nums[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		ends[l] = nums[i]
		right = Max(right,l)
		dp[i] = l + 1
		if dp[i] == 3 {
			return true
		}
	}
	return false
}

func increasingTriplet3(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	small,mid := 0x7fffffff,0x7fffffff
	for _,x := range nums {
		if x <= small {
			small = x
		} else if x < mid {
			mid = x
		} else if x > mid {
			return true
		}
	}
	return false
}