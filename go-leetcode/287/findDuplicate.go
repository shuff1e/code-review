package main

import "fmt"

/*

287. 寻找重复数
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2
示例 2:

输入: [3,1,3,4,2]
输出: 3
说明：

不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。

 */

// 将每个数放在自己该放的位置上
// nums[i] = i+ 1

// 3,1,3,4,2

// 2,1,3,3,4

func main() {
	arr := []int{1,3,4,2,2}
	arr = []int{3,1,3,4,2}
	fmt.Println(findDuplicate2(arr))
}

func findDuplicate(nums []int) int {
	for i := 0;i<len(nums);i++ {
		for nums[i] != i + 1 {
			if nums[nums[i]-1] == nums[i] {
				return nums[i]
			}
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		}
	}
	return -1
}

func findDuplicate2(nums []int) int {
	start,end := 0,len(nums)-1
	for {
		p1,p2 := partition3(nums,start,end,nums[start])
		if p1 != p2 {
			return nums[p1]
		} else if nums[p1] < p1 + 1 {
			// 在左边
			end = p1 -1
		} else {
			start = p1 + 1
		}
	}
}

func partition3(arr []int,start,end int,pivot int) (int,int) {
	less := start - 1
	more := end + 1
	left := start

	for left < more {
		if arr[left] < pivot {
			less ++
			arr[less],arr[left] = arr[left],arr[less]
			left ++
		} else if arr[left] > pivot {
			more --
			arr[left],arr[more] = arr[more],arr[left]
		} else {
			left ++
		}
	}
	return less + 1,more-1
}