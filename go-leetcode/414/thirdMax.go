package main

import "fmt"

/*

414. 第三大的数
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。

示例 1:

输入: [3, 2, 1]

输出: 1

解释: 第三大的数是 1.
示例 2:

输入: [1, 2]

输出: 2

解释: 第三大的数不存在, 所以返回最大的数 2 .
示例 3:

输入: [2, 2, 3, 1]

输出: 1

解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
存在两个值为2的数，它们都排第二。

 */

func main() {
	arr := []int{2,2,3,1}
	arr = []int{1,2}
	fmt.Println(thirdMax(arr))
}

func thirdMax(nums []int) int {
	first := -0x800000000
	second := -0x800000000
	third := -0x800000000

	for i := 0;i<len(nums);i++ {
		if nums[i] == first ||
			nums[i] == second ||
			nums[i] == third {
			continue
		}
		if nums[i] > first {
			third = second
			second = first
			first = nums[i]
		} else if nums[i] > second {
			third = second
			second = nums[i]
		} else if nums[i] > third {
			third = nums[i]
		}
	}

	if third != -0x800000000 {
		return third
	}
	return first
}
