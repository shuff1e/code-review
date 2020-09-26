package main

/*
7. 整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

 */

func reverse(x int) int {
	var result int64 = 0

	var flag int64 = 1
	if x < 0 {
		flag = -1
		x = -x
	}

	for x > 0 {
		result = result*10 + int64(x)%10
		if (flag == 1 && result > 0x7fffffff) || (flag == -1 && result > 0x80000000){
			return 0
		}
		x = x/10
	}
	return int(flag*result)
}