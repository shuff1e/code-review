package main

import "fmt"

/*
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例：

输入：4
输出：[
[".Q..",  // 解法 1
"...Q",
"Q...",
"..Q."],

["..Q.",  // 解法 2
"Q...",
"...Q",
".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。

提示：

皇后彼此不能相互攻击，也就是说：任何两个皇后都不能处于同一条横行、纵行或斜线上。
 */

func main() {
	result := solveNQueens(0)
	for _,v := range result {
		for _,s := range v {
			fmt.Println(s)
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	temp,result := [][]int{},[][]string{}
	help(n,0,&temp,&result)
	return result
}

func help(n int,level int,temp *[][]int, result *[][]string) {
	if level == n {
		if len(*temp) > 0 {
			strResult := []string{}
			for _,v := range *temp {
				strResult = append(strResult,makeStr(v,n))
			}
			*result = append(*result,strResult)
		}
		return
	}
	for i := 0;i<n;i++ {
		if checkValid(level,i,*temp) {
			*temp = append(*temp,[]int{level,i})
			help(n,level+1,temp,result)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}

func makeStr(point []int,n int) string {
	result := ""
	for j:=0;j<n;j++ {
		if j == point[1] {
			result = result + "Q"
		} else {
			result = result + "."
		}
	}
	return result
}

func checkValid(i,j int, temp [][]int) bool {
	for _,v := range temp {
		if Abs(v[0] - i) == Abs(v[1]-j) || v[1] == j {
			return false
		}
	}
	return true
}

func Abs(x int ) int {
	if x > 0 {
		return x
	}
	return -x
}