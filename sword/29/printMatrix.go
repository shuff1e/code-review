package main

import (
	"fmt"
)

// 29：顺时针打印矩阵
// 题目：输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

// 输入为：
// 1   2   3   4
// 5   6   7   8
// 9   10  11  12
// 13  14  15  16

// 输出为：1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10







// [0][0->col-1)
// [0->row-1)[col-1]
// [row-1][col-1->0)
// [row-1->0)[0]

func printMatrix(arr [][]int) {
	rowCount := len(arr)
	rowStart := 0
	colStart := 0
	colCount := len(arr[0])
	for rowCount-1 > rowStart && colCount-1 > colStart {
		for j := colStart;j<colCount-1;j++ {
			fmt.Print(arr[rowStart][j]," ")
		}
		for i := rowStart;i<rowCount-1;i++ {
			fmt.Print(arr[i][colCount-1]," ")
		}
		for j := colCount - 1;j>colStart;j-- {
			fmt.Print(arr[rowCount-1][j]," ")
		}
		for i := rowCount - 1;i>rowStart;i-- {
			fmt.Print(arr[i][colStart]," ")
		}
		rowStart ++
		colStart ++
		rowCount --
		colCount --
	}
	if rowStart == rowCount - 1 {
		for j := colStart;j<colCount;j++ {
			fmt.Print(arr[rowStart][j]," ")
		}
	} else if colStart == colCount - 1 {
		for i := rowStart;i<rowCount;i++ {
			fmt.Print(arr[i][colStart]," ")
		}
	}
}

func main() {
	/*
	   1
	*/
	//Test(1, 1);

	/*
	   1    2
	   3    4
	*/
	//Test(2, 2);

	/*
	   1    2    3    4
	   5    6    7    8
	   9    10   11   12
	   13   14   15   16
	*/
	//Test(4, 4);

	/*
	   1    2    3    4    5
	   6    7    8    9    10
	   11   12   13   14   15
	   16   17   18   19   20
	   21   22   23   24   25
	*/
	//Test(5, 5);

	/*
	   1
	   2
	   3
	   4
	   5
	*/
	//Test(1, 5);

	/*
	   1    2
	   3    4
	   5    6
	   7    8
	   9    10
	*/
	Test(2, 5);

	/*
	   1    2    3
	   4    5    6
	   7    8    9
	   10   11   12
	   13   14   15
	*/
	Test(3, 5);

	/*
	   1    2    3    4
	   5    6    7    8
	   9    10   11   12
	   13   14   15   16
	   17   18   19   20
	*/
	Test(4, 5);

	/*
	   1    2    3    4    5
	*/
	Test(5, 1);

	/*
	   1    2    3    4    5
	   6    7    8    9    10
	*/
	Test(5, 2);

	/*
	   1    2    3    4    5
	   6    7    8    9    10
	   11   12   13   14    15
	*/
	Test(5, 3);

	/*
	   1    2    3    4    5
	   6    7    8    9    10
	   11   12   13   14   15
	   16   17   18   19   20
	*/
	Test(5, 4);
}

func Test(col,row int) {
	if row < 1 || col < 1 {
		return
	}

	arr := make([][]int,row)
	for i := 0;i<row;i++ {
		arr[i] = make([]int,col)
	}

	for i := 0;i<row;i++ {
		for j :=0;j<col;j++ {
			arr[i][j] = i*col + j + 1
		}
	}
	//for i := 0;i<row;i++ {
	//	fmt.Println(arr[i])
	//}

	printMatrix(arr)
	fmt.Println()
}