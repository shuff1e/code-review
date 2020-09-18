package main

import (
	"fmt"
	"sort"
)

// 61：扑克牌的顺子
// 题目：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
// 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王可以看成任意数字。



// A：大小王假设为0
// 对数组排序，如果空隙的个数大于0的个数，则是invalid

type myStruct []int

func (s myStruct) Len() int {
	return len(s)
}

func (s myStruct) Less(i,j int) bool {
	return s[i] < s[j]
}

func (s myStruct) Swap(i,j int) {
	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}

func isValid(arr myStruct) bool {
	if len(arr) == 0 {
		return false
	}

	sort.Sort(arr)

	index := len(arr)
	// 第一个不为0的数
	for i,v := range arr {
		if v != 0 {
			index = i
			break
		}
	}
	zeroNumber := index - 0
	gap := 0
	for i := index ;i< len(arr)-1;i++ {
		// 两个数相等，有对子，不可能是顺子
		if arr[i+1] == arr[i] {
			return false
		}
		gap += arr[i+1] - arr[i] - 1
	}
	return gap <= zeroNumber
}

func Test(name string,arr []int,expected bool) {
	fmt.Println(name)
	fmt.Println(isValid(myStruct(arr)) == expected)
}


















func main() {
	Test("Test1",[]int{ 1, 3, 2, 5, 4 },true)
	Test("Test2",[]int{ 1, 3, 2, 6, 4 },false)
	Test("Test3",[]int{ 0, 3, 2, 6, 4 },true)
	Test("Test4",[]int{ 0, 3, 1, 6, 4 },false)
	Test("Test5",[]int{ 1, 3, 0, 5, 0 },true)
	Test("Test6",[]int{ 1, 3, 0, 7, 0 },false)
	Test("Test7",[]int{ 1, 0, 0, 5, 0 },true)
	Test("Test8",[]int{ 1, 0, 0, 7, 0 },false)
	Test("Test9",[]int{ 3, 0, 0, 0, 0 },true)
	Test("Test10",[]int{ 0, 0, 0, 0, 0 },true)
	Test("Test11",[]int{ 1, 0, 0, 1, 0 },false)
	Test("Test12",[]int{},false)
}