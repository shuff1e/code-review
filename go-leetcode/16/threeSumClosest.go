package main

import (
	"fmt"
	"sort"
)

/*
16. 最接近的三数之和
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

 */

// A：如果使用暴力的方法，一次确定一个数字，然后再分别计算另外两个数字
// 另外两个数字的计算，可以使用双指针

func main() {
	arr := []int{-1,2,1,-4}
	fmt.Println(threeSumClosest(arr,2))
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

func threeSumClosest(nums []int, target int) int {
	sort.Sort(data(nums))

	diff := 0x7ffffffff
	closetSum := 0

	outer:
	for first := 0;first<len(nums);first++ {
		for first > 0 && first < len(nums) && nums[first] == nums[first-1] {
			first ++
		}
		second := first + 1
		third := len(nums) - 1
		for second < third {
			for second > first + 1 && second < len(nums) && nums[second] == nums[second-1] {
				second ++
			}
			if second >= third {
				break
			}

			curSum := nums[first] + nums[second] + nums[third]
			if Abs(curSum-target) < diff {
				diff = Abs(curSum-target)
				closetSum = curSum
			}

			if curSum < target {
				second ++
			} else if curSum > target {
				third --
				for third >= 0 && nums[third] == nums[third+1] {
					third --
				}
			} else {
				break outer
			}
		}
	}
	return closetSum
}

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}