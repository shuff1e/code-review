package main

import "fmt"

/*
54. 螺旋矩阵
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
[ 1, 2, 3 ],
[ 4, 5, 6 ],
[ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
[1, 2, 3, 4],
[5, 6, 7, 8],
[9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

 */

// A：剑指第29

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9,10,11,12},
	}
	matrix = [][]int{
		{2,5},
		{8,4},
		{0,-1},
	}
	result := spiralOrder(matrix)
	fmt.Printf("%#v\n",result)
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]int,len(matrix)*len(matrix[0]))
	index := 0
	// [0][0->col-1]
	// [col-1][0->row-1]
	// [row-1][col-1->0]
	// [row-1->0][0]
	start := 0
	for (start+1)*2 <= len(matrix) && (start+1)*2 <= len(matrix[0]) {
		for j := start;j<len(matrix[0]) -1 - start;j++ {
			result[index] = matrix[start][j]
			index ++
		}
		for i := start;i<len(matrix)-1-start;i++ {
			result[index] = matrix[i][len(matrix[0])-1-start]
			index ++
		}
		for j := len(matrix[0]) - 1 -start;j > start;j-- {
			result[index] = matrix[len(matrix)-1-start][j]
			index ++
		}
		for i:=len(matrix)-1-start;i>start;i-- {
			result[index] = matrix[i][start]
			index ++
		}
		start++
	}

	if start*2 == len(matrix) ||  start*2 == len(matrix[0]) {
		return result
	}

	if len(matrix) >= len(matrix[0]) {
		for i:=start;i<len(matrix)-start;i++ {
			result[index] = matrix[i][len(matrix[0])-1-start]
			index ++
		}
	} else {
		for j := start;j<len(matrix[0])-start;j++ {
			result[index] = matrix[len(matrix)-1-start][j]
			index ++
		}
	}
	return result
}
