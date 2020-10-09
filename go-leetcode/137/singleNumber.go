package main

import "fmt"

/*
137. 只出现一次的数字 II
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,3,2]
输出: 3
示例 2:

输入: [0,1,0,1,0,1,99]
输出: 99
 */

func main() {
	arr := []int{1,1,1,3}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	once,twice := 0,0
	for _,x := range nums {
		// 出现两次，现在又出现
		triple := twice & x
		// 之前出现一次，这次又出现
		// 之前出现两次，这次没出现
		twice = (once & x) | (twice & ^x)
		// 之前出现了一次，这次又出现的需要删除掉
		// 出现三次的需要删除掉
		once = once ^ x
		once = once & (^triple)
	}
	return once
}
