package main

import "fmt"

/*
216. 组合总和 III
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：

所有数字都是正整数。
解集不能包含重复的组合。
示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
 */

// 每个数字都有选和不选两种选择

func main() {
	result := combinationSum3(3,3)
	for i :=0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

func combinationSum3(k int, n int) [][]int {
	temp := []int{}
	result := [][]int{}
	help(k,n,1,9,&temp,&result)
	return result
}

func help(k,target ,curr,n int ,temp *[]int,result *[][]int) {
	// temp增长了
	if len(*temp) == k {
		if checkSum(*temp,target) {
			temp2 := make([]int,len(*temp))
			copy(temp2,*temp)
			*result = append(*result,temp2)
		}
		return
	}
	// temp没有增长，需要k-len(*temp)，可以提供的n - curr + 1，两者比较
	//
	if k - len(*temp) == n - curr + 1 {
		mark := curr
		for ;curr <= n;curr++ {
			*temp = append(*temp,curr)
		}
		if checkSum(*temp,target) {
			temp2 := make([]int,len(*temp))
			copy(temp2,*temp)
			*result = append(*result,temp2)
		}
		curr = mark
		for ;curr <= n;curr++ {
			*temp = (*temp)[:len(*temp)-1]
		}
		return
	}
	if k - len(*temp) > n - curr + 1 {
		return
	}
	if curr > n {
		return
	}

	// 选
	*temp = append(*temp,curr)
	help(k,target,curr+1,n,temp,result)
	*temp = (*temp)[:len(*temp)-1]
	// 不选
	help(k,target,curr+1,n,temp,result)
}

func checkSum(arr []int,target int) bool {
	sum := 0
	for i := 0;i<len(arr);i++ {
		sum += arr[i]
	}
	return sum == target
}