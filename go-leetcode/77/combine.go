package main

import "fmt"

/*
77. 组合
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
[2,4],
[3,4],
[2,3],
[1,2],
[1,3],
[1,4],
]
 */

func main() {
	result := combine(4,2)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}

}

func combine(n int, k int) [][]int {
	result := [][]int{}
	temp := []int{}
	help(n,k,1,&temp,&result)
	return result
}

func help(n,k ,cur int,temp *[]int,result *[][]int) {
	if n - cur + 1< k - len(*temp) {
		return
	}
	if len(*temp) == k {
		temp2 := make([]int,len(*temp))
		copy(temp2,*temp)
		*result = append(*result,temp2)
		return
	}
	if cur == n + 1 {
		return
	}
	// 不选
	help(n,k,cur+1,temp,result)
	// 选
	*temp = append(*temp,cur)
	help(n,k,cur+1,temp,result)
	*temp = (*temp)[0:len(*temp)-1]
}
