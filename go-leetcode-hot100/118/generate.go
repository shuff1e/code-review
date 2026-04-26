package main

import "fmt"

/*
118. 杨辉三角
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
      [1],
     [1,1],
    [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
[1,5,10,10,5,1]
[1,6,15,20,15,6,1]
]
 */

func main() {
	matrix := generate(5)
	for i := 0;i<len(matrix);i++ {
		fmt.Printf("%#v\n",matrix[i])
	}
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1},{1,1}}
	}
	result := [][]int{{1},{1,1}}
	for i := 3;i<=numRows;i++ {
		temp := []int{1}
		for j := 1;j<i-1;j++ {
			temp = append(temp,result[i-2][j-1] + result[i-2][j])
		}
		temp = append(temp,1)
		result = append(result,temp)
	}
	return result
}
