package main

import (
	"fmt"
)

/*

321. 拼接最大数
给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，
表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，
要求从同一个数组中取出的数字保持其在原数组中的相对顺序。

求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。

说明: 请尽可能地优化你算法的时间和空间复杂度。

示例 1:

输入:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
输出:
[9, 8, 6, 5, 3]
示例 2:

输入:
nums1 = [6, 7]
nums2 = [6, 0, 4]
k = 5
输出:
[6, 7, 6, 0, 4]
示例 3:

输入:
nums1 = [3, 9]
nums2 = [8, 9]
k = 3
输出:
[9, 8, 9]

 */

func main() {
	arr1 := []int{3, 4, 6, 5}
	arr2 := []int{9, 1, 2, 5, 8, 3}
	k := 5
	fmt.Println(help([]int{8,6,9},1))
	//fmt.Println(maxNumber(arr1,arr2,k))
	arr1 = []int{8,6,9}
	arr2 = []int{1,7,5}
	fmt.Println(merge(arr1,arr2))
	k = 3
	fmt.Println(maxNumber(arr1,arr2,k))
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	result := make([]int,k)
	// 从 nums1中取i个
	// 从 num2中取k-i个
	for i := 0;i<=k;i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			arr1 := help(nums1,i)
			arr2 := help(nums2,k-i)
			tempArr := merge(arr1,arr2)
			temp := Max(result,tempArr)
			if temp == 2 {
				result = tempArr
			}
		}
	}
	return result
}

func merge(arr1 ,arr2 []int) []int {
	result := []int{}
	for len(arr1) > 0 && len(arr2) > 0 {
		max_black := Max(arr1,arr2)
		if max_black == 1 {
			result = append(result,arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result,arr2[0])
			arr2 = arr2[1:]
		}
	}
	if len(arr1) > 0 {
		result = append(result,arr1...)
	}
	if len(arr2) > 0 {
		result = append(result,arr2...)
	}
	return result
}

// 单调栈
//316
//321
//402
//1081
// 从arr中选取k个数字，组成一个最大的
func help(arr []int,k int) []int {
	drop := len(arr) - k
	stack := []int{}
	for i :=0 ;i<len(arr);i++ {
		for len(stack) > 0 && drop > 0 && stack[len(stack)-1] < arr[i] {
			drop --
			stack = stack[:len(stack)-1]
		}
		stack = append(stack,arr[i])
	}
	for drop > 0 {
		drop --
		stack = stack[:len(stack)-1]
	}
	return stack
}

func Max(x,y []int) int {
	length := Min(len(x),len(y))
	for i := 0;i<length;i ++ {
		if x[i] < y[i] {
			return 2
		} else if x[i] > y[i] {
			return 1
		}
	}
	if len(x) < len(y) {
		return 2
	}
	return 1
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}