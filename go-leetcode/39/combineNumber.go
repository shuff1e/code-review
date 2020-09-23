package main

import (
	"fmt"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。 
示例 1：

输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
[7],
[2,2,3]
]
示例 2：

输入：candidates = [2,3,5], target = 8,
所求解集为：
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
 

提示：

1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/

// A：类似换钱的方法数
// 每个数字都可以选或者不选


func main() {
	slice := []int{2,3,5}
	target := 8
	result := combinationSum(slice,target)
	//result := [][]int{}
	//Test(&result)
	fmt.Printf("%#v\n",result)
}

func Test(arr *[][]int) {
	*arr = append(*arr,[]int{1,2,3})
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	temp := []int{}
	// 使用指针的指针！
	combine(candidates,target,0,temp,&result)
	return result
}

func combine(arr []int,target,level int,temp []int,result *[][]int) {
	if level == len(arr) {
		if target == 0 {
			temp2 := make([]int,len(temp))
			copy(temp2,temp)
			*result = append(*result,temp2)
		}
		return
	}
	if target == 0 && len(temp) > 0 {
		temp2 := make([]int,len(temp))
		copy(temp2,temp)
		*result = append(*result,temp2)
		return
	}
	if target < 0 {
		return
	}

	// 选
	temp = append(temp,arr[level])
	combine(arr,target-arr[level],level,temp,result)
	// 不选
	temp = temp[0:len(temp)-1]
	combine(arr,target,level+1,temp,result)
}