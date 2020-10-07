package main

import "fmt"

/*
119. 杨辉三角 II
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

输入: 5
输出:
[
      [1],
     [1,1],
    [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]  5
[1,5,10,10,5,1] 6
[1,6,15,20,15,6,1]
]
 */

// 每一项都是C(n,i)
// C(4,1) = 4
// C(4,2) = 4*3/2=6

// C(n,i) = C(n,i-1) * (n-i+1)/i

// C(5,2) = 5*4/2=10

func main() {
	result := getRow(4)
	fmt.Printf("%#v\n",result)
}

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	result := []int{}
	result = append(result,1)
	for i := 1;i<=rowIndex-1;i++ {
		result = append(result,result[i-1]*(rowIndex-i+1)/i)
	}
	result = append(result,1)
	return result
}