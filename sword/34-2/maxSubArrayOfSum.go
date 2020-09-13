package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
)

// Q：给定一个无序数组arr，其中元素可正，可负，可0，给定一个整数k
// 求arr所有的子数组中累加和为k的最长子数组长度


// A：记录每个数字和出现的最早的位置
func getMaxLength(arr []int,k int) int {
	mmp := hashmap.New()
	mmp.Put(0,-1)
	sum := 0
	maxLength := 0

	for i := 0;i<len(arr);i++ {
		sum += arr[i]
		if _,ok := mmp.Get(sum);!ok {
			mmp.Put(sum,i)
		}
		if temp,ok := mmp.Get(sum-k);ok {
			tempLen := i - temp.(int)
			if tempLen > maxLength {
				maxLength = tempLen
			}
		}
	}
	return maxLength
}

func main() {
	Test("Test1",[]int{1,2,3,4,5,-1,-2,3},12,7)
}

func Test(name string,arr []int,k,expected int) {
	fmt.Println(name,getMaxLength(arr,k) == expected)
}