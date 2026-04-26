package main

import "fmt"

/*
994. 腐烂的橘子

给定一个 m x n 的网格 grid，每个格子可能有三种值：
0 表示空格子；
1 表示新鲜橘子；
2 表示腐烂橘子。

每经过一分钟，腐烂橘子会使上下左右相邻的新鲜橘子腐烂。
返回直到没有新鲜橘子为止所需的最少分钟数；如果不可能让所有新鲜橘子腐烂，返回 -1。

示例 1：
输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4

示例 2：
输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1

示例 3：
输入：grid = [[0,2]]
输出：0

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] 的值为 0、1 或 2。
*/

func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1}}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1}}
	fmt.Println(orangesRotting(grid))

	grid = [][]int{
		{0, 2},
	}
	fmt.Println(orangesRotting(grid))
}

func orangesRotting(grid [][]int) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rotten := [][2]int{}
	freshCount := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				rotten = append(rotten, [2]int{i, j})
			}
			if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	edges := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	count := 0
	for {
		tempRotten := [][2]int{}
		for i := 0; i < len(rotten); i++ {
			for j := 0; j < len(edges); j++ {
				row := rotten[i][0] + edges[j][0]
				col := rotten[i][1] + edges[j][1]

				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) {
					if grid[row][col] == 1 {
						tempRotten = append(tempRotten, [2]int{row, col})
						freshCount--
						grid[row][col] = 2
					}
				}
			}
		}
		if len(tempRotten) == 0 {
			break
		}
		count++
		rotten = tempRotten
	}

	if freshCount > 0 {
		return -1
	}

	return count
}
