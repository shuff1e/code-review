package main

import "fmt"

/*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入:
[
[1,1,1],
[1,0,1],
[1,1,1]
]
输出:
[
[1,0,1],
[0,0,0],
[1,0,1]
]
示例 2:

输入:
[
[0,1,2,0],
[3,4,5,2],
[1,3,1,5]
]

1001


输出:
[
[0,0,0,0],
[0,4,5,0],
[0,3,1,0]
]
进阶:

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？
 */

// A： 我们可以用每行和每列的第一个元素作为标记，这个标记用来表示这一行或者这一列是否需要赋零。
// 这意味着对于每个节点不需要访问 M+N 个格子而是只需要对标记点的两个格子赋值。
//
//if cell[i][j] == 0 {
//    cell[i][0] = 0
//    cell[0][j] = 0
//}
// 这些标签用于之后对矩阵的更新，如果某行的第一个元素为零就将整行置零，如果某列的第一个元素为零就将整列置零。
//
// 算法
//
// 遍历整个矩阵，如果 cell[i][j] == 0 就将第 i 行和第 j 列的第一个元素标记。
// 第一行和第一列的标记是相同的，都是 cell[0][0]，所以需要一个额外的变量告知第一列是否被标记，同时用 cell[0][0] 继续表示第一行的标记。
// 然后，从第二行第二列的元素开始遍历，如果第 r 行或者第 c 列被标记了，那么就将 cell[r][c] 设为 0。这里第一行和第一列的作用就相当于方法一中的 row_set 和 column_set 。
// 然后我们检查是否 cell[0][0] == 0 ，如果是则赋值第一行的元素为零。
// 然后检查第一列是否被标记，如果是则赋值第一列的元素为零。

func main() {
	matrix := [][]int{
		{0,1,2,0},
	{3,4,5,2},
	{1,3,1,5},
}
	setZeroesBetter(matrix)
	for i := 0;i<len(matrix);i++ {
		fmt.Printf("%#v\n",matrix[i])
	}
}

func setZeroesBetter(matrix [][]int) {
	isCol := false
	for i := 0;i<len(matrix);i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1;j<len(matrix[0]);j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	for i := 1;i<len(matrix);i++ {
		for j := 1;j<len(matrix[0]);j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if matrix[0][0] == 0 {
		for j := 0;j<len(matrix[0]);j++ {
			matrix[0][j] = 0
		}
	}
	if isCol {
		for i := 0;i<len(matrix);i++ {
			matrix[i][0] = 0
		}
	}
}

func setZeroes(matrix [][]int)  {
	colLimit := 0
	rowLimit := 0
	colLength := len(matrix[0])

	for i := 0;i<len(matrix);i++ {
		foundZero := false
		for j := 0;j<len(matrix[0]);j++ {
			if matrix[i][j] == 0 {
				foundZero = true
				colLimit |= 1 << (colLength-1-j)
			}
		}
		if foundZero {
			rowLimit |= (1 << i)
			setZeroForRow(matrix,i)
		}
	}

	for i := 0;i<len(matrix);i++ {
		if rowLimit & (1<<i) > 0 {
			continue
		}
		for j := 0;j<len(matrix[0]);j++ {
			if colLimit & (1 << (colLength-1-j)) > 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZeroForRow(matrix [][]int,row int) {
	for j := 0;j<len(matrix[0]);j++ {
		matrix[row][j] = 0
	}
}