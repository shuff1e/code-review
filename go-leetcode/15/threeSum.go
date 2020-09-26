package main

import (
	"fmt"
	"sort"
)

/*
15. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
[-1, 0, 1],
[-1, -1, 2]
]

 */

// -4,-1,-1,0,1,2,
func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("%#v\n",threeSum(arr))
}

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

func threeSum(nums []int) [][]int {
	sort.Sort(data(nums))

	result := [][]int{}

	for first := 0;first<len(nums);first++ {
		for first > 0 && first < len(nums) && nums[first] == nums[first-1]  { // 相同的元素只会取一次
			first ++
		}

		second := first + 1
		third := len(nums) - 1
		for second < third {
			for second > first +1 && second < len(nums) && nums[second] == nums[second-1] { // 相同的元素只会取一次
				second ++
			}
			// 注意指针的位置，以及是否越界
			if second >= third {
				break
			}
			if nums[second] + nums[third] == -nums[first] {
				result = append(result,[]int{nums[first],nums[second],nums[third]})
				second ++
				third --
				for third >= 0 && nums[third] == nums[third+1] { // 相同的元素只会取一次
					third--
				}
			} else if nums[second] + nums[third] < -nums[first] {
				second ++
			} else if nums[second] + nums[third] > -nums[first] {
				third --
				for third >= 0 && nums[third] == nums[third+1] { // 相同的元素只会取一次
					third--
				}
			}

		}
	}
	return result
}

