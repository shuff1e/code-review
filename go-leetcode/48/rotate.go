package main

import "fmt"

/*
48. 旋转图像
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix =
[
[1,2,3],
[4,5,6],
[7,8,9]
],


[0,1] -> [1,0]

原地旋转输入矩阵，使其变为:
[
[7,4,1],
[8,5,2],
[9,6,3]
]

示例 2:

给定 matrix =
[
[ 5, 1, 9,11],
[ 2, 4, 8,10],
[13, 3, 6, 7],
[15,14,12,16]
],

15  1  9  5
2   4  8  10
13  3  6   7
16  14 12  11

15  13  9   5
2    4  8   1
12   3  6    7
16  14  10  11

原地旋转输入矩阵，使其变为:
[
[15,13, 2, 5],
[14, 3, 4, 1],
[12, 6, 8, 9],
[16, 7,10,11]
]
 */

// 1  2  3  4
// 5  6  7  8
// 9 10  11 12
// 13 14 15 16

// 3 -> 9

// 1 4 7
// 2 5 8
// 3 6 9
//
// 7 4 1
// 8 5 2
// 9 6 3

func main() {
	matrix := [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}
	rotate2(matrix)
	for i := 0;i<len(matrix);i++ {
		fmt.Printf("%#v\n",matrix[i])
	}
}

func rotate(matrix [][]int)  {
	// transpose matrix
	// 沿着 1 5 9 翻转
	for i := 0;i < len(matrix);i ++ {
		for j := i;j<len(matrix[0]);j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0])/2;j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix[0])-1-j]
			matrix[i][len(matrix[0])-1-j] = temp
		}
	}
}

// 1  2  3  4
// 5  6  7  8
// 9 10  11 12
// 13 14 15 16

// 13 2  3  1
// 5  6  7  8
// 9  10 11 12
// 16 14 15  4

func rotate2(matrix [][]int)  {
	for i := 0;i<(len(matrix)+1)/2;i++ {
		for j := i;j<len(matrix) - i - 1;j++ {
			// 交换 1 4 16 13 四个元素
			temp := make([]int,4)
			row := i
			col := j
			for k := 0;k<4;k++ {
				temp[k] = matrix[row][col]
				x := row
				row = col
				col = len(matrix) - 1 - x
				// 0,0 -> 0,3
				// 0,1 -> 1,3
				// 0,3 -> 3,3

				// 0,1 -> 1,2
			}
			for k := 0;k<4;k++ {
				matrix[row][col] = temp[(k+3)%4]
				x := row
				row = col
				col = len(matrix) - 1 - x
			}
		}
	}

}