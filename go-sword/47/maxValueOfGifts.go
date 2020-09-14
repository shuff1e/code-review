package main

import "fmt"

// 47：礼物的最大价值
// 题目：在一个m×n的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值
// （价值大于0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或
// 者向下移动一格直到到达棋盘的右下角。给定一个棋盘及其上面的礼物，请计
// 算你最多能拿到多少价值的礼物？

// A：dp[i][j] = max(dp[i+1][j],dp[i][j+1]) + cur

// for i := row-2;i>=0;i--
//      for j := col - 2;i>=0;j--
//         if dp[i+1][j] < dp[i][j+1];dp[i][j] = matrix[i][j] + dp[i][j+1]

// 或者用一行的话，if dp[j] < dp[j+1];dp[j] = matrix[i][j] + dp[j+1]


// 可以增加memo，减少重复计算
func getMaxValue(matrix [][]int,i,j int) int {
	if i >= len(matrix) || j >= len(matrix[0]) {
		return 0
	}
	// 向下
	v1 := getMaxValue(matrix,i+1,j)
	// 向右
	v2 := getMaxValue(matrix,i,j+1)
	return matrix[i][j] + Max(v1,v2)
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	Test("Test1",[][]int{
		{ 1, 2, 3 },
		{ 4, 5, 6 },
		{ 7, 8, 9 },
	},29)
	Test("Test2",[][]int{
		{ 1, 10, 3, 8 },
		{ 12, 2, 9, 6 },
		{ 5, 7, 4, 11 },
		{ 3, 7, 16, 5 },
	},53)
	Test("Test3",[][]int{
		{ 1, 10, 3, 8 },
	},22)
	Test("Test4",[][]int{
		{ 1 },
		{ 12 },
		{ 5 },
		{ 3 },
	},21)
	Test("Test5",[][]int{
		{ 3 },
	},3)
	Test("Test6",nil,0)
}

func Test(name string,matrix [][]int,expected int) {
	fmt.Println(name)
	if getMaxValue(matrix,0,0) != expected {
		panic("fuck")
	}
}