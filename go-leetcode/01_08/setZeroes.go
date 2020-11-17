package main

/*

面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。



示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

 */

/*

先记录第一行第一列是否有0，然后利用第一行第一列作为标记，先清除非第一行或非第一列的数据，最后根据第一列第一列是否有0来清除第一行第一列。

 */

func setZeroes(matrix [][]int)  {
	firstRowZero := false
	firstColZero := false

	for i := 0;i<len(matrix);i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for j := 0;j<len(matrix[0]);j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 1;i<len(matrix);i++ {
		for j := 1;j<len(matrix[0]);j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1;i<len(matrix);i++ {
		for j := 1;j<len(matrix[0]);j++ {
			if matrix[i][0] == 0 ||
				matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0;j<len(matrix[0]);j++ {
			matrix[0][j] = 0
		}
	}

	if firstColZero {
		for i := 0;i<len(matrix);i++ {
			matrix[i][0] = 0
		}
	}

}