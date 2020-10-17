package main

import "fmt"

/*
260. 只出现一次的数字 III
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

示例 :

输入: [1,2,1,3,2,5]
输出: [3,5]
注意：

结果输出的顺序并不重要，对于上面的例子， [5, 3] 也是正确答案。
你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？
 */

// 异或，result就是a和b的异或结果
// 比如result为1100
// 然后根据某一位是否为1，将数组分成两部分
// 这2个数字分别在这2部分中

func main() {
	arr := []int{1,2,1,3,2,5,4,4,3,6}
	arr = []int{1,2}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) []int {
	if len(nums) <= 1 {
		return nil
	}

	result := 0
	for i := 0;i<len(nums);i++ {
		result = result ^ nums[i]
	}

	// 得到最右边的1
	rightMost := result & (^result + 1)
	one,two := 0,0
	for i := 0;i<len(nums);i++ {
		if nums[i] & rightMost != 0 {
			one = one ^ nums[i]
		} else {
			two = two ^ nums[i]
		}
	}
	return []int{one,result^one}
}