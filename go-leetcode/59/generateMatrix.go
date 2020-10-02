package main

import "fmt"

/*
59. 螺旋矩阵 II
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
[ 1, 2, 3 ],
[ 8, 9, 4 ],
[ 7, 6, 5 ]
]
 */

// A：lc54

func main() {
	result := generateMatrix(4)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

func generateMatrix(n int) [][]int {
	result := make([][]int,n)
	for i := 0;i<len(result);i++ {
		result[i] = make([]int,n)
	}

	rowStart,rowEnd,colStart,colEnd := 0,n-1,0,n-1
	index := 1
	for rowEnd > rowStart && colEnd > colStart {
		// [0][0->col-1]
		// [col-1][0->row-1]
		// [row-1][col-1->0]
		// [row-1->0][0]
		for j:=colStart;j<colEnd;j++ {
			result[rowStart][j] = index
			index ++
		}
		for i:=rowStart;i<rowEnd;i++ {
			result[i][colEnd] = index
			index ++
		}
		for j := colEnd;j>colStart;j-- {
			result[rowEnd][j] = index
			index++
		}
		for i := rowEnd;i>rowStart;i-- {
			result[i][colStart] = index
			index ++
		}
		rowStart ++
		colStart ++
		rowEnd --
		colEnd --
	}
	if rowStart == rowEnd {
		for j := colStart;j<=colEnd;j++ {
			result[rowStart][j] = index
			index ++
		}
	} else {
		for i := rowStart;i<=rowEnd;i++ {
			result[i][colStart] = index
			index ++
		}
	}
	return result
}

// 1 1 1 1 1
// 1 1 1 1 1
// 1 1 1 1 1