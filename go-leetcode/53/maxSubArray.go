package main

import "fmt"

/*
53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
 */

// A：dp[i]表示以arr[i]为结尾的最大子数组和

func main() {
	arr := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(arr))
}

func maxSubArray(nums []int) int {
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1;i<len(nums);i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		result = Max(result,dp[i])
	}
	return result
}


// 线段树
// https://zhuanlan.zhihu.com/p/34150142
// 对于每一个区间，维护四个变量
//
// lSum 表示 [l, r] 内以 l 为左端点的最大子段和
// rSum 表示 [l,r] 内以 r 为右端点的最大子段和
// mSum 表示 [l,r] 内的最大子段和
// iSum 表示 [l,r] 的区间和
//
// merge的时候，
// 整体的的lSum = max(left.lSum,left.iSum + right.lSum)
// 整体的的rSum = max(right.rSum,right.iSum+ left.rSum)
// 整体的的mSum = max(left.mSum,right.mSum, left.rSum + right.lSum)
// 整体的的lSum = left.iSum + right.iSum

// 它不仅可以解决区间 [0,n−1]，
// 还可以用于解决任意的子区间 [l,r] 的问题。
// 如果我们把 [0,n−1] 分治下去出现的所有子区间的信息都用堆式存储的方式记忆化下来，
// 即建成一颗真正的树之后，我们就可以在 O(logn) 的时间内求到任意区间内的答案，
// 我们甚至可以修改序列中的值，做一些简单的维护，
// 之后仍然可以在 O(logn) 的时间内求到任意区间内的答案，
// 对于大规模查询的情况下，这种方法的优势便体现了出来

func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

type status struct {
	lSum int
	rSum int
	mSum int
	iSum int
}

func pushUp(left,right status) status {
	lSum := Max(left.lSum,left.iSum+right.lSum)
	rSum := Max(right.rSum,right.iSum+left.rSum)
	mSum := Max(Max(left.mSum,right.mSum),left.rSum + right.lSum)
	iSum := left.iSum + right.iSum
	return status{
		lSum: lSum,
		rSum: rSum,
		mSum: mSum,
		iSum: iSum,
	}
}

func get(nums []int,left,right int) status {
	if left == right {
		return status{
			nums[left],
			nums[left],
			nums[left],
			nums[left],
		}
	}
	mid := (left+right) >> 1
	leftStatus := get(nums,left,mid)
	rightStatus := get(nums,mid+1,right)
	return pushUp(leftStatus,rightStatus)
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
