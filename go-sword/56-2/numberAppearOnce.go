package main

import (
	"fmt"
)

// 56（二）：数组中唯一只出现一次的数字
// 题目：在一个数组中除了一个数字只出现一次之外，其他数字都出现了三次。请
// 找出那个出现一次的数字。

// A：将这些数字每一位的bit相加，如果能被3整除，说明没有要求的数字
// 如果不能被3整除，说明这一位是要求的数字

func getNumber(arr []int) int32 {
	count := [32]int{}
	for _,v := range arr {
		bitMask := 1
		for j := 31;j>=0;j-- {
			if v &bitMask > 0 {
				count[j] ++
			}
			bitMask = bitMask << 1
		}
	}
	// 原码
	// 反码
	// 补码：反码加1称为补码
	// 需要指定是	int32
	// 不然不会按负数计算
	var result int32 = 0
	for _,v := range count {
		result = result << 1
		if v %3 != 0 {
			result += 1
		}
	}
	//fmt.Println(str)
	return result
}

func main() {
	Test("Test1",[]int{ 1, 1, 2, 2, 2, 1, 3 },3)
	Test("Test2",[]int{ 4, 3, 3, 2, 2, 2, 3 },4)
	Test("Test3",[]int{ 4, 4, 1, 1, 1, 7, 4 },7)
	Test("Test4",[]int{ -10, 214, 214, 214 },-10)
	Test("Test5",[]int{ -209, 3467, -209, -209 },3467)
	Test("Test6",[]int{ 1024, -1025, 1024, -1025, 1024, -1025, 1023 },1023)
	Test("Test7",[]int{ -1024, -1024, -1024, -1023 },-1023)
	Test("Test8",[]int{ -23, 0, 214, -23, 214, -23, 214 },0)
	Test("Test9",[]int{ 0, 3467, 0, 0, 0, 0, 0, 0 },3467)
}

func Test(name string,arr []int,expected int32) {
	fmt.Println(name,getNumber(arr),expected)
	if getNumber(arr) != expected {
		panic("fuck")
	}
}