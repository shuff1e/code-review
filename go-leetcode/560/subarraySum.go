package main

import "fmt"

/*
560. 和为 K 的子数组

给定一个整数数组 nums 和一个整数 k，统计并返回数组中和为 k 的连续子数组的个数。

子数组必须是原数组中连续的一段，长度至少为 1。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2

示例 2：
输入：nums = [1,2,3], k = 3
输出：2

提示：
1 <= nums.length <= 2 * 10^4
-1000 <= nums[i] <= 1000
-10^7 <= k <= 10^7
*/

func main() {
	nums := []int{1, 1, 1}
	k := 2
	ret := subarraySum(nums, k)
	fmt.Println(ret)

	nums = []int{1, 2, 3}
	k = 3
	ret = subarraySum(nums, k)
	fmt.Println(ret)

	nums = []int{1, -1, 1, -1}
	k = 0
	ret = subarraySum(nums, k)
	fmt.Println(ret)
}

func subarraySum(nums []int, k int) int {
	dict := map[int]int{
		0: 1,
	}
	result := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result += dict[sum-k]
		dict[sum]++
	}

	return result
}
