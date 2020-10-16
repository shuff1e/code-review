package main

import "fmt"

/*
240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
示例:

现有矩阵 matrix 如下：

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。

给定 target = 20，返回 false。
 */

// 右上角，如果target > arr[i][j]，排除i行
// 如果target < arr[i][j] 排除j列

func main() {
	matrix := [][]int {
		{1,   4,  7, 11, 15},
	{2,   5,  8, 12, 19},
	{3,   6,  9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix,20))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m,n := len(matrix),len(matrix[0])
	i,j :=0,n-1
	for i < m-1 && j > 0 {
		if target > matrix[i][j] {
			i ++
		} else if target < matrix[i][j] {
			j --
		} else {
			return true
		}
	}
	if i == m-1 {
		left,right := 0,j
		for left <= right {
			mid := (left + right)/2
			if matrix[m-1][mid] > target {
				right = mid -1
			} else if matrix[m-1][mid] < target {
				left = mid + 1
			} else {
				return true
			}
		}
	}

	if j == 0 {
		left,right := i,m-1
		for left <= right {
			mid := (left + right)/2
			if matrix[mid][0] > target {
				right = mid -1
			} else if matrix[mid][0] < target {
				left = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

