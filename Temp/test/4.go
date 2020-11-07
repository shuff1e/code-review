package main

import "fmt"

/*

给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000

 */





















// arr1,arr2
// 两个数组中整体排名K的

// arr1 1 2 3 4 5
// arr2 6 7 8 9 10
// K 为 5，两边排名K/2=5/2=2的
// 2小于7，所以1，2不可能是排名第5的

// arr1 3 4 5
// arr2 6 7 8 9 10
// K 为 3,3/2=1
// 3 < 6，3不可能是排名第3的

// arr1 4 5
// arr2 6 7 8 9 10
// K为2 2/2=1
// 4 < 6，4不可能是排第二的

// arr1 5
// arr2 6 7 8 9 10
// K为1
// 返回5

// arr1 3
// arr2 1 2 5 6 7 8
// K = 5

func main() {
	// 1,1,2,3,3,4,5,6,8,9,10,12
	arr1 := []int{}
	arr2 := []int{1,3,8,9,10,12}
	result := findMedianSortedArrays(arr1,arr2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	K := len(nums1) + len(nums2)
	if K % 2 == 0 {
		result1 := help(nums1,nums2,0,0,K/2)
		result2 := help(nums1,nums2,0,0,K/2+1)
		return (float64(result1) + float64(result2))/2
	}
	// 8 排第四,第五
	// 9 排第五
	result := help(nums1,nums2,0,0,(K+1)/2)
	return float64(result)
}

func help(nums1,nums2 []int,p1,p2 int,K int) int {
	for p1 < len(nums1) && p2 < len(nums2) && K > 1 {
		if p1 + K/2 - 1 < len(nums1) && p2 + K/2 - 1 < len(nums2) {
			if nums1[p1 + K/2 -1] <= nums2[p2 + K/2 -1] {
				p1 = p1 + K/2
				K -= K/2
			} else {
				p2 = p2 + K/2
				K -= K/2
			}
		} else if p1 + K/2 - 1 < len(nums1) {
			if nums1[p1 + K/2 -1] <= nums2[len(nums2)-1] {
				p1 = p1 + K/2
				K -= K/2
			} else {
				K = K - (len(nums2)-p2)
				p2 = len(nums2)
			}
		} else if p2 + K/2 - 1 < len(nums2) {
			if nums2[p2 + K/2 - 1] <= nums1[len(nums1)-1] {
				p2 = p2 + K/2
				K -= K/2
			} else {
				K = K -(len(nums1)-p1)
				p1 = len(nums1)
			}
		} else {
			if nums1[len(nums1)-1] <= nums2[len(nums2)-1] {
				K = K - (len(nums1)-p1)
				p1 = len(nums1)
			} else {
				K = K - (len(nums2)-p2)
				p2 = len(nums2)
			}
		}
	}

	if K == 1 {
		if p1 < len(nums1) && p2 < len(nums2) {
			return Min(nums1[p1],nums2[p2])
		} else if p1 < len(nums1) {
			return nums1[p1]
		} else {
			return nums2[p2]
		}
	}

	if p1 < len(nums1) {
		return nums1[p1 + K - 1]
	}

	return nums2[p2 + K - 1]
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}