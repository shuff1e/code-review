package main

import (
	"fmt"
	"sort"
)

/*
47. 全排列 II
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
[1,1,2],
[1,2,1],
[2,1,1]
]
 */

// A：需要让重复的元素依次被选取
// i > 0 && arr[i] == arr[i-1] && !visited[i-1] ; continue

// 回溯法 ：一种通过探索所有可能的候选解来找出所有的解的算法。
// 如果候选解被确认不是一个解的话（或者至少不是最后一个解），回溯算法会通过在上一步进行一些变化抛弃该解，即回溯并且再次尝试。
//
// 这个问题可以看作有 n 个排列成一行的空格，我们需要从左往右依此填入题目给定的 n 个数，每个数只能使用一次。
// 那么很直接的可以想到一种穷举的算法，即从左往右每一个位置都依此尝试填入一个数，看能不能填完这 n 个空格，
// 在程序中我们可以用「回溯法」来模拟这个过程。
//
// 很容易想到的一个处理手段是我们定义一个标记数组 vis[] 来标记已经填过的数，那么在填第 first 个数的时候我们遍历题目给定的 n 个数，
// 如果这个数没有被标记过，我们就尝试填入，并将其标记，继续尝试填下一个位置
// 搜索回溯的时候要撤销这一个位置填的数以及标记，并继续尝试其他没被标记过的数。
//
// 我们可以将题目给定的 n 个数的数组 nums[] 划分成左右两个部分，左边的表示已经填过的数，右边表示待填的数，
// 我们在递归搜索的时候只要动态维护这个数组即可。
//
// 举个简单的例子，假设我们有 [2, 5, 8, 9, 10] 这 5 个数要填入，
// 已经填到第 3 个位置，已经填了 [8,9] 两个数，那么这个数组目前为 [8, 9 | 2, 5, 10] 这样的状态，
// 分隔符区分了左右两个部分。假设这个位置我们要填 10 这个数，为了维护数组，我们将 2 和 10 交换，
// 即能使得数组继续保持分隔符左边的数已经填过，右边的待填 [8, 9, 10 | 2, 5] 。
//
// 如果题目要求按字典序输出，那么请还是用标记数组或者其他方法。

func main() {
	nums := []int{1,3,3,0,2}
	result := permuteUnique(nums)
	for i := 0;i<len(result);i++ {
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

func permuteUnique(nums []int) [][]int {
	sort.Sort(data(nums))
	result := [][]int{}
	visited := make([]bool,len(nums))
	temp := []int{}
	help(nums,0,&result,visited,&temp)
	return result
}

func help(arr []int, level int,result *[][]int,visited []bool,temp *[]int) {
	if level == len(arr) {
		temp2 := make([]int,len(arr))
		copy(temp2,*temp)
		*result = append(*result,temp2)
	}
	for i := 0;i<len(arr);i++ {
		if visited[i] || (i > 0 && arr[i] == arr[i-1] && !visited[i-1]) {
			continue
		}
		visited[i] = true
		*temp = append(*temp,arr[i])
		help(arr,level+1,result,visited,temp)
		visited[i] = false
		*temp = (*temp)[:len(*temp)-1]
	}
}

