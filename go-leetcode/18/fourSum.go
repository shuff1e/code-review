package main

import "sort"

/*
18. 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，
使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
[-1,  0, 0, 1],
[-2, -1, 1, 2],
[-2,  0, 0, 2]
]

 */

type data []int

// Len is the number of elements in the collection.
func (d data) Len() int{
	return len(d)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (d data) Less(i, j int) bool {
	return d[i] < d[j]
}
// Swap swaps the elements with indexes i and j.
func (d data) Swap(i, j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func fourSum(nums []int, target int) [][]int {
	sort.Sort(data(nums))
	result := [][]int{}
	for first := 0;first<len(nums);first ++ {
		for first > 0 && first < len(nums) && nums[first] == nums[first-1] {
			first ++
		}
		for second := first + 1;second<len(nums);second ++ {
			for second > first + 1 && second < len(nums) && nums[second] == nums[second-1] {
				second ++
			}
			third := second + 1
			forth := len(nums) - 1
			for third < forth {
				for third > second + 1 && third < len(nums) && nums[third] == nums[third-1] {
					third ++
				}
				if third >= forth {
					break
				}
				if nums[first] + nums[second] + nums[third] + nums[forth] == target {
					result = append(result,[]int{nums[first],nums[second],nums[third],nums[forth]})
					third ++
					forth --
					for forth < len(nums) - 1 && forth >= 0 && nums[forth] == nums[forth+1] {
						forth --
					}
				} else if nums[first] + nums[second] + nums[third] + nums[forth] < target {
					third ++
				} else {
					forth --
					for forth < len(nums) - 1 && forth >= 0 && nums[forth] == nums[forth+1] {
						forth --
					}
				}
			}
		}
	}
	return result
}