package main

import (
	"fmt"
)

/*
37. 解数独
编写一个程序，通过填充空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。

一个数独。

答案被标成红色。

提示：

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
 */

func main() {

	board := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	doItBetter(board)
	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			fmt.Printf("%c",board[i][j])
		}
		fmt.Println()
	}

}

func solveSudoku(board [][]byte)  {
	rows := make([]map[byte]int,9)
	cols :=  make([]map[byte]int,9)
	boxes :=  make([]map[byte]int,9)
	for i := 0;i<9;i++ {
		rows[i] = make(map[byte]int,0)
		cols[i] = make(map[byte]int,0)
		boxes[i] = make(map[byte]int,0)
	}

	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			if board[i][j] != '.' {
				rows[i][board[i][j]] = rows[i][board[i][j]] + 1
				cols[j][board[i][j]] = cols[j][board[i][j]] + 1
				boxes[(i/3)*3+j/3][board[i][j]] = boxes[(i/3)*3+j/3][board[i][j]] + 1
			}
		}
	}
	found := false
	result := make([][]byte,9)
	for i := 0;i<9;i++ {
		result[i] = make([]byte,9)
	}
	help(board,0,0,rows,cols,boxes,result,&found)
	board = result

	//for i := 0;i<len(board);i++ {
	//	for j := 0;j<len(board[0]);j++ {
	//		fmt.Printf("%c",result[i][j])
	//	}
	//	fmt.Println()
	//}
}

func help(matrix [][]byte,row ,col int,rows,cols,boxes []map[byte]int,result [][]byte,found *bool) {
	if *found {
		return
	}
	if row == 9 {
		*found = true
		//for i := 0;i<9;i++ {
		//	for j := 0;j<9;j++ {
		//		result[i][j] = matrix[i][j]
		//	}
		//}
		return
	}

	index := -1
	for j := col;j < 9;j++ {
		if matrix[row][j] == '.' {
			index = j
			break
		}
	}

	if index == -1 {
		help(matrix,row+1,0,rows,cols,boxes,result,found)
	} else {
		for num := byte(1) ;num <= 9;num ++ {
			if rows[row][num+'0'] == 0 &&
				cols[index][num+'0'] == 0 &&
				boxes[(row/3)*3+index/3][num+'0'] == 0 {
				matrix[row][index] = num + '0'
				rows[row][num+'0'] = 1
				cols[index][num+'0'] = 1
				boxes[(row/3)*3+index/3][num+'0'] = 1

				help(matrix,row,index+1,rows,cols,boxes,result,found)
				if !(*found) {
					matrix[row][index] = '.'
					rows[row][num+'0'] = 0
					cols[index][num+'0'] = 0
					boxes[(row/3)*3+index/3][num+'0'] = 0
				}
			}
		}
	}
}

// 二维其实也就是一纬，计算机内存只有一纬
func doIt(board [][]byte) {
	rows := make([]map[byte]bool,9)
	cols :=  make([]map[byte]bool,9)
	boxes :=  make([]map[byte]bool,9)
	for i := 0;i<9;i++ {
		rows[i] = make(map[byte]bool,0)
		cols[i] = make(map[byte]bool,0)
		boxes[i] = make(map[byte]bool,0)
	}
	// 填充spaces中的这些位置
	spaces := []int{}
	for i := 0;i<9;i++ {
		for j := 0;j<9;j++ {
			if board[i][j] == '.' {
				spaces = append(spaces,i*9+j)
			} else { // 已经填充的位置
				rows[i][board[i][j] - '0'] = true
				cols[j][board[i][j] - '0'] = true
				boxes[(i/3)*3+j/3][board[i][j] - '0'] = true
			}
		}
	}
	found := false
	helpDoIt(board,spaces,0,rows,cols,boxes,&found)

}

func helpDoIt(board [][]byte,spaces []int,index int,rows,cols,boxes []map[byte]bool,found *bool) {
	// 不能写在这个地方，因为该程序中还可能再循环下去
	//if *found {
	//	return
	//}
	if index == len(spaces) {
		*found = true
		return
	}
	// 对于每一个可以填充的位置，填充之，用1-9填充
	row := spaces[index]/9
	col := spaces[index]%9
	for i := byte(1);i<=9 && !*found;i++ {
		if !rows[row][i] &&
			!cols[col][i] &&
			!boxes[(row/3)*3+col/3][i] {

			rows[row][i] = true
			cols[col][i] = true
			boxes[(row/3)*3+col/3][i] = true

			board[row][col] = i + '0'
			helpDoIt(board,spaces,index+1,rows,cols,boxes,found)

			rows[row][i] = false
			cols[col][i] = false
			boxes[(row/3)*3+col/3][i] = false
		}
	}
}

func doItBetter(board [][]byte) {
	rows := make([]int,9)
	cols := make([]int,9)
	boxes := make([]int,9)

	spaces := []int{}
	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			if board[i][j] == '.' {
				spaces = append(spaces,i*9+j)
			} else {
				flip(i,j,int(board[i][j] - '0' -
					1),rows,cols,boxes)
			}
		}
	}
	found := false
	doItBetterHelp(board,spaces,0,rows,cols,boxes,&found)
}

const UPPER_LIMIT = 0x1ff

func doItBetterHelp(board [][]byte,spaces []int,index int,rows,cols,boxes []int,found *bool) {
	if index == len(spaces) {
		*found = true
		return
	}
	row := spaces[index]/9
	col := spaces[index]%9

	limit := (^( rows[row] | cols[col] | boxes[(row/3)*3+col/3] ))&UPPER_LIMIT
	for ;limit > 0 && !*found;limit = limit & (limit - 1) {
		// 得到最右边的1
		rightMost := limit & (-limit)
		//rightMost := limit & (^limit + 1)
		digit := getBitCount(rightMost-1)
		// 100
		// 011
		// 对应着3，所以是 digit + 1
		flip(row,col,digit,rows,cols,boxes)
		board[row][col] = '0' + 1 + byte(digit)
		doItBetterHelp(board,spaces,index+1,rows,cols,boxes,found)
		flip(row,col,digit,rows,cols,boxes)
	}
}

func flip(i, j, digit int,rows,cols,boxes []int) {
	// 表示该位置 哪些数字已经被取完了
	rows[i] ^= 1 << digit
	cols[j] ^= 1 << digit
	boxes[(i/3)*3+j/3] ^= 1 << digit
}

func getBitCount(x int) int {
	count := 0
	for x > 0 {
		x = x & (x-1)
		count ++
	}
	return count
}