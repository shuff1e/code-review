package main

import "fmt"

/*
78. 子集
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
[3],
[1],
[2],
[1,2,3],
[1,3],
[2,3],
[1,2],
[]
]
 */

func main() {
	arr := []int{1,2,3}
	result := subsets(arr)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

func subsets(nums []int) [][]int {
	result := [][]int{}
	temp := []int{}
	help(nums,0,&temp,&result)
	return result
}

func help(arr []int,level int,temp *[]int,result *[][]int) {
	if level == len(arr) {
		temp2 := make([]int,len(*temp))
		copy(temp2,*temp)
		*result = append(*result,temp2)
		return
	}

	// not choose
	help(arr,level+1,temp,result)

	// choose
	*temp = append(*temp,arr[level])
	help(arr,level+1,temp,result)
	*temp = (*temp)[0:len(*temp)-1]
}
