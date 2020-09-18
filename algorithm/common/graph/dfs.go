package main

import (
	"algorithm/common/help"
	"fmt"
)

// 给出图的临接矩阵

const size = 8
var points []string = []string{"A","B","C","D","E","F","G","H"}
var edge [][]int = [][]int{
	{0,1,1,0,0,0,0,0},
	{1,0,0,1,1,0,0,0},
	{1,0,0,0,0,0,0,0},
	{0,1,0,0,0,0,0,0},
	{0,1,0,0,0,0,0,0},
	{0,0,0,0,0,0,1,1},
	{0,0,0,0,0,1,0,0},
	{0,0,0,0,0,1,0,0},
}

var flag []bool

func DFS() {
	flag = make([]bool,size)
	stack := help.NewStack()
	index := 0
	for index < size {
		if !flag[index] {
			flag[index] = true
			fmt.Println(points[index])
			stack.Push(index)
			for stack.Length() > 0 {
				temp,_ := stack.Peek()
				help := Helper(temp)
				if help >= 0 {
					stack.Push(help)
				} else{
					stack.Pop()
				}
			}
		}
		index ++
	}
}

func Helper(index int) int {
	for i := 0;i<size;i++ {
		if !flag[i] && edge[index][i] >0 {
			flag[i] = true
			fmt.Println(points[i])
			return i
		}
	}
	return -1
}

func BFS() {
	flag = make([]bool,size)
	queue := help.NewMyQueue()
	index := 0
	for index < size {
		if !flag[index] {
			flag[index] = true
			queue.Add(index)
			fmt.Println(points[index])
			for queue.Length() > 0 {
				temp,_ := queue.Peek()
				for {
					help := Helper(temp)
					if help >= 0 {
						queue.Add(help)
					} else {
						break
					}
				}
				queue.Poll()
			}
		}
		index ++
	}
}

func main() {
	DFS()
	fmt.Println()
	BFS()
}