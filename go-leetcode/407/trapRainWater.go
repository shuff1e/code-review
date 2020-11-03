package main

import (
	"container/heap"
	"fmt"
)

/*

407. 接雨水 II
给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。

示例：

给出如下 3x6 的高度图:
[
  [1,4,3,1,3,2],
  [3,2,1,3,2,4],
  [2,3,3,2,3,1]
]

返回 4 。

如上图所示，这是下雨前的高度图[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]] 的状态。

下雨后，雨水将会被存储在这些方块中。总的接雨水量是4。

提示：

1 <= m, n <= 110
0 <= heightMap[i][j] <= 20000

 */

func main() {
	matrix := [][]int{
		{1,4,3,1,3,2},
		{3,2,1,3,2,4},
		{2,3,3,2,3,1},
	}
	result := trapRainWater(matrix)
	fmt.Println(result)
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
	}

	visited := make([][]bool,len(heightMap))
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,len(heightMap[0]))
	}

	h := &minHeap{}

	for i := 0;i<len(heightMap);i++ {
		for j := 0;j<len(heightMap[0]);j++ {
			if i == 0 ||
				i == len(heightMap) - 1 ||
				j == 0 ||
				j == len(heightMap[0]) - 1 {
				heap.Push(h,[3]int{i,j,heightMap[i][j]})
				visited[i][j] = true
			}
		}
	}

	res := 0
	dirs := [][2]int{
		// 上
		{0,1},
		// 下
		{0,-1},
		// 右边
		{1,0},
		// 左边
		{-1,0},
	}

	for h.Len() > 0 {
		poll := heap.Pop(h).([3]int)

		for k := 0;k<len(dirs);k++ {
			nx := poll[0] + dirs[k][0]
			ny := poll[1] + dirs[k][1]
			if nx >= 0 &&
				nx < len(heightMap) &&
				ny >= 0 &&
				ny < len(heightMap[0]) &&
				!visited[nx][ny] {
				if poll[2] > heightMap[nx][ny] {
					res += poll[2] - heightMap[nx][ny]
				}
				heap.Push(h,[3]int{nx,ny,Max(poll[2],heightMap[nx][ny])})
				visited[nx][ny] = true
			}
		}
	}
	return res
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

type minHeap [][3]int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i,j int) bool {
	return h[i][2] < h[j][2]
}

func (h minHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h,x.([3]int))
}

func (h *minHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}