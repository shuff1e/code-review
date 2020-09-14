package main

import "fmt"

// 42：连续子数组的最大和
// 题目：输入一个整型数组，数组里有正数也有负数。数组中一个或连续的多个整
// 数组成一个子数组。求所有子数组的和的最大值。要求时间复杂度为O(n)。

// A：以某个index为结尾的子数组的最大和为sum[index]，
// arr[index] + sum[index-1] ， 当sum[index-1] > 0
// arr[index] ， 当sum[index-1] <= 0

// 例如arr为 1,-3,2,3
// 则sum为   1,-2,2,5

func getMaxSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	_,maxSum := getMaxSumHelper(arr,len(arr)-1)
	return maxSum
}

func getMaxSumHelper(arr []int,index int) (sum ,maxSum int) {
	if index == 0 {
		return arr[index],arr[index]
	}
	temp,tempSum := getMaxSumHelper(arr,index-1)
	if temp <= 0 {
		curSum := arr[index]
		return curSum,max(tempSum,arr[index])
	} else {
		curSum := arr[index] + temp
		return curSum,max(tempSum,curSum)
	}
}

func max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	Test("Test1",[]int{1, -2, 3, 10, -4, 7, 2, -5},18)
	Test("Test2",[]int{-2, -8, -1, -5, -9},-1)
	Test("Test3",[]int{2, 8, 1, 5, 9},25)
	Test("Test4",nil,0)
}

func Test(name string,data []int,expected int) {
	fmt.Println(name)
	fmt.Println(getMaxSum(data) == expected)
}