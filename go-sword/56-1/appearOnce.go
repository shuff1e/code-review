package main

import "fmt"

// 56（一）：数组中只出现一次的两个数字
// 题目：一个整型数组里除了两个数字之外，其他的数字都出现了两次。请写程序
// 找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
























// A：其他数都是出现两次，相同的数字异或之后，为0。那么将这些数字异或之后，只剩下需要的2个数字的异或结果result
// 而且result中有一位不为0
// 可以根据这一位是否为0，将数组再分为两部分
// 再对两部分分别异或







func getUniqueNumber(arr []int) (int,int) {
	result := 0
	for _,v := range arr {
		result = result ^ v
	}
	firstOneBit := result & (^result + 1)

	result1,result2 := 0,0
	for _, v := range arr {
		if v & firstOneBit > 0 {
			result1 = result1 ^ v
		} else {
			result2 = result2 ^ v
		}
	}
	if result1 < result2 {
		return result1,result2
	}
	return result2,result1
}

func Test(name string, arr []int,expected1,expected2 int) {
	fmt.Println(name)
	v1,v2 := getUniqueNumber(arr)
	fmt.Println(v1,v2,expected1,expected2)
	if v1 != expected1 || v2 != expected2 {
		panic("err")
	}
}

func main() {
	Test("Test1",[]int{ 2, 4, 3, 6, 3, 2, 5, 5 },4,6)
	Test("Test2",[]int{ 4, 6 },4,6)
	Test("Test3",[]int{ 4, 6, 1, 1, 1, 1 },4,6)
}