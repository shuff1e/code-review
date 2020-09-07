package main

import "fmt"

// 4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。

// A：将整数与右上角的元素n比较
// 如果n < 整数 ，那么与n同一行的可以全部去除掉
// 如果n > 整数，那么与n同一列的可以全部去除掉
// 这样一次去掉一行或者一列，直到 n == 整数

func findInParitiallySortedMatrix(matrix [][]int,n int) bool {
	for row,col := 0,len(matrix[0])-1;row<=len(matrix)-1 && col >=0;{
		if matrix[row][col] < n {
			row += 1
		} else if matrix[row][col] > n {
			col -= 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix1 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix1,7))

	matrix2 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix2,5))

	matrix3 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix3,1))

	matrix4 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix4,15))

	matrix5 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix5,0))

	matrix6 := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(findInParitiallySortedMatrix(matrix6,16))
}

