package main

import "fmt"

// 12：矩阵中的路径
// 题目：请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有
// 字符的路径。路径可以从矩阵中任意一格开始，每一步可以在矩阵中向左、右、
// 上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入
// 该格子。例如在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字
// 母用下划线标出）。但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个
// 字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
// A B T G
// C F C S
// J D E H

func search(matrix [][]byte,path string) bool {
	visited := make([][]bool,len(matrix))
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,len(matrix[0]))
	}

	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0]);j++ {
			if find(matrix,i,j,visited,path,0) {
				return true
			}
		}
	}
	return false
}

func find(matrix [][]byte,i,j int , visited [][]bool,path string,index int) bool {
	// 超过边界
	if i < 0 || i >= len(matrix) || j < 0 || j >=len(matrix[0]) {
		return false
	}
	// 已经访问过
	if visited[i][j] {
		return false
	}
	visited[i][j] = true
	defer func() {visited[i][j] = false}()
	// 如果不匹配
	if matrix[i][j] != path[index] {
		return false
	}
	// 如果匹配而且到了字符串的最后
	if index == len(path) -1 {
		return true
	}
	// 没到字符串的最后，继续搜索
	return find(matrix,i+1,j,visited,path,index+1) ||
		find(matrix,i,j+1,visited,path,index+1) ||
		find(matrix,i-1,j,visited,path,index+1) ||
		find(matrix,i,j-1,visited,path,index+1)
}

func main() {
	//ABTG
	//CFCS
	//JDEH

	//BFCE
	matrix := [][]byte{
		{'A','B','T','G'},
		{'C','F','C','S'},
		{'J','D','E','H'},
	}
	path := "BFCE"
	fmt.Println(search(matrix,path))
	//ABCE
	//SFCS
	//ADEE

	//SEE
	matrix = [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	path = "SEE"
	fmt.Println(search(matrix,path))
	//ABTG
	//CFCS
	//JDEH

	//ABFB
	matrix = [][]byte{
		{'A','B','T','G'},
		{'C','F','C','S'},
		{'J','D','E','H'},
	}
	path = "ABFB"
	fmt.Println(search(matrix,path))

	//ABCEHJIG
	//SFCSLOPQ
	//ADEEMNOE
	//ADIDEJFM
	//VCEIFGGS

	//SLHECCEIDEJFGGFIE
	matrix = [][]byte{
		{'A','B','C','E','H','J','I','G'},
		{'S','F','C','S','L','O','P','Q'},
		{'A','D','E','E','M','N','O','E'},
		{'A','D','I','D','E','J','F','M'},
		{'V','C','E','I','F','G','G','S'},
	}
	path = "SLHECCEIDEJFGGFIE"
	fmt.Println(search(matrix,path))

	//A

	//A
	matrix = [][]byte{
		{'A'},
	}
	path = "A"
	fmt.Println(search(matrix,path))
}