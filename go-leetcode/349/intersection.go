package main

import (
	"fmt"
	"sort"
)

/*

349. 两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]

说明：

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

 */

func main() {
	arr1 := []int{1,2,2,1}
	arr2 := []int{2,2}
	fmt.Println(intersection(arr1,arr2))
	fmt.Println(intersection2(arr1,arr2))
}

func intersection(nums1 []int, nums2 []int) []int {
	dict1 := map[int]struct{}{}
	dict2 := map[int]struct{}{}

	for _,v := range nums1 {
		dict1[v] = struct{}{}
	}

	for _,v := range nums2 {
		dict2[v] = struct{}{}
	}

	result := make([]int,0)
	for k,_ := range dict1 {
		if _,ok := dict2[k];ok {
			result = append(result,k)
		}
	}

	return result
}

// 排序，双指针

// 1 1 2 2
//     2 2

func intersection2(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	result := []int{}
	p1,p2 := 0,0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result,nums1[p1])
			p1 ++
			p2 ++
			for p1 < len(nums1) && nums1[p1] == nums1[p1-1] {
				p1 ++
			}
			for p2 < len(nums2) && nums2[p2] == nums2[p2-1] {
				p2 ++
			}
		} else if nums1[p1] < nums2[p2] {
			p1 ++
		} else if nums1[p1] > nums2[p2] {
			p2 ++
		}
	}
	return result
}