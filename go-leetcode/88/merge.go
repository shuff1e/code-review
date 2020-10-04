package main

import "fmt"

/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。


示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]
 */

// A：从后往前，两个指针

func main() {
	arr1 := []int{1,2,3,0,0,0}
	arr2 := []int{2,5,6}
	merge(arr1,3,arr2,len(arr2))
	fmt.Printf("%#v\n",arr1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	index := len(nums1)-1
	index1 := m-1
	index2 := n-1

	for index >= 0 && index1 >= 0 && index2 >=0 {
		if nums1[index1] < nums2[index2] {
			nums1[index] = nums2[index2]
			index --
			index2 --
		} else {
			nums1[index] = nums1[index1]
			index--
			index1--
		}
	}
	for index1 >= 0 {
		nums1[index] = nums1[index1]
		index--
		index1--
	}
	for index2 >= 0 {
		nums1[index] = nums2[index2]
		index --
		index2 --
	}
}