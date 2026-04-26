package main

import "fmt"

/*
46. 全排列
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
[1,2,3],
[1,3,2],
[2,1,3],
[2,3,1],
[3,1,2],
[3,2,1]
]
 */

func main() {
	nums := []int{1,2,3}
	result := permute(nums)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

func permute(nums []int) [][]int {
	result := [][]int{}
	help(nums,0,&result)
	return result
}

func help(arr []int, start int,result *[][]int) {
	if start == len(arr) {
		temp := make([]int,len(arr))
		copy(temp,arr)
		*result = append(*result,temp)
	}
	for i := start;i<len(arr);i++ {
		swap(arr,start,i)
		help(arr,start+1,result)
		swap(arr,start,i)
	}
}

func swap(arr []int,i,j int) {
	 temp := arr[i]
	 arr[i] = arr[j]
	 arr[j] = temp
}