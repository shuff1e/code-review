package main

import "fmt"

/*
79. 单词搜索
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母不允许被重复使用。

示例:

board =
[
['A','B','C','E'],
['S','F','C','S'],
['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：

board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
 */

func main() {
	board := [][]byte{
			{'A','B','C','E'},
	{'S','F','C','S'},
	{'A','D','E','E'},
}
	word := "ABCCED"
	//word = "ABCB"
	word = "SEEE"
	fmt.Println(exist(board,word))
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool,len(board))
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,len(board[0]))
	}
	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			if help(board,i,j,word,0,visited) {
				return true
			}
		}
	}
	return false
}

func help(matrix [][]byte,row,col int,word string,index int,visited [][]bool) bool {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || index >= len(word) {
		return false
	}
	if visited[row][col] {
		return false
	}

	visited[row][col] = true
	defer func() {
		visited[row][col] = false
	}()

	if matrix[row][col] != word[index] {
		return false
	}
	if index == len(word) - 1 {
		return true
	}


	return help(matrix,row+1,col,word,index+1,visited) ||
		help(matrix,row-1,col,word,index+1,visited) ||
		help(matrix,row,col+1,word,index+1,visited) ||
		help(matrix,row,col-1,word,index+1,visited)
}