package main

/*
66. 加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
 */

func plusOne(digits []int) []int {
	carry := 0
	for i := len(digits)-1;i>=0;i-- {
		sum := 0
		if i == len(digits) - 1 {
			sum = 1 + carry + digits[i]
		} else {
			sum = carry + digits[i]
		}
		carry = sum/10
		digits[i] = sum%10
	}
	if carry > 0 {
		result := []int{1}
		result = append(result,digits...)
		return result
	}
	return digits
}