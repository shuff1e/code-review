package main

import (
	"fmt"
	"sort"
)

/*
90. 子集 II
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
[2],
[1],
[1,2,2],
[2,2],
[1,2],
[]
]
 */

func main()  {
	arr := []int{1,2,2}
	result := subsetsWithDup(arr)
	for i :=0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

type data []int

func (d data) Len() int {
	return len(d)
}

func (d data) Less(i,j int) bool {
	return d[i] < d[j]
}

func (d data) Swap(i,j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func subsetsWithDup(nums []int) [][]int {
	sort.Sort(data(nums))
	temp := []int{}
	result := [][]int{}
	help(nums,0,&temp,&result)
	return result
}

// 长度为0的
// 长度为1的
// 长度为2的，以此类推
func help(arr []int,index int,temp *[]int,result *[][]int) {
	temp2 := make([]int,len(*temp))
	copy(temp2,*temp)
	*result = append(*result,temp2)
	for i := index;i<len(arr);i++ {
		if i != index && arr[i] == arr[i-1] {
			continue
		}

		*temp = append(*temp,arr[i])
		help(arr,i+1,temp,result)
		*temp = (*temp)[0:len(*temp)-1]
	}
}