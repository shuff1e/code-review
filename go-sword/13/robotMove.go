package main

import "fmt"

// 13：机器人的运动范围
// 题目：地上有一个m行n列的方格。一个机器人从坐标(0, 0)的格子开始移动，它
// 每一次可以向左、右、上、下移动一格，但不能进入行坐标和列坐标的数位之和
// 大于k的格子。例如，当k为18时，机器人能够进入方格(35, 37)，因为3+5+3+7=18。
// 但它不能进入方格(35, 38)，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

func maxReached(m,n,k int) int {
	visited := make([][]bool,m)
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,n)
	}
	return help(0,0,m,n,k,visited)
}

func help(x,y,m,n,k int,visited [][]bool) int {
	if x < 0 || x >=m || y < 0 || y >= n ||
		visited[x][y] ||
		getDigitSum(x) + getDigitSum(y) > k {
		return 0
	}
	visited[x][y] = true
	return 1 + help(x+1,y,m,n,k,visited) +
		help(x-1,y,m,n,k,visited) +
		help(x,y+1,m,n,k,visited) +
		help(x,y-1,m,n,k,visited)
}

func getDigitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x%10
		x /= 10
	}
	return sum
}

func main() {
	fmt.Println(test(5, 10, 10, 21))
	fmt.Println(test(15, 20, 20, 359))
	fmt.Println(test(10, 1, 100, 29))
	fmt.Println(test(10, 1, 10, 10))
	fmt.Println(test(15, 100, 1, 79))
	fmt.Println(test(-10, 10, 10, 0))
}

func test(k,row,col,expected int) bool {
	return maxReached(row,col,k) == expected
}