package main

import "fmt"

/*

310. 最小高度树
对于一个具有树特征的无向图，我们可选择任何一个节点作为根。
图因此可以成为树，在所有可能的树中，具有最小高度的树被称为最小高度树。
给出这样的一个图，写出一个函数找到所有的最小高度树并返回他们的根节点。

格式

该图包含 n 个节点，标记为 0 到 n - 1。给定数字 n 和一个无向边 edges 列表（每一个边都是一对标签）。

你可以假设没有重复的边会出现在 edges 中。由于所有的边都是无向边， [0, 1]和 [1, 0] 是相同的，
因此不会同时出现在 edges 里。

示例 1:

输入: n = 4, edges = [[1, 0], [1, 2], [1, 3]]

        0
        |
        1
       / \
      2   3

输出: [1]
示例 2:

输入: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]

     0  1  2
      \ | /
        3
        |
        4
        |
        5

输出: [3, 4]
说明:

根据树的定义，树是一个无向图，其中任何两个顶点只通过一条路径连接。
换句话说，一个任何没有简单环路的连通图都是一棵树。
树的高度是指根节点和叶子节点之间最长向下路径上边的数量。

 */

func main() {
	edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}
	result := findMinHeightTrees(6,edges)
	fmt.Println(result)
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	}
	result := []int{}
	if n == 1 {
		result = append(result,0)
		return result
	}

	// 建立图关系
	slice := make([][]int,n)
	for i := 0;i<n;i++ {
		slice[i] = make([]int,0)
	}
	// 出度表
	degree := make([]int,n)

	for i := 0 ;i<len(edges);i++ {
		degree[edges[i][0]] ++
		degree[edges[i][1]] ++
		slice[edges[i][0]] = append(slice[edges[i][0]],edges[i][1])
		slice[edges[i][1]] = append(slice[edges[i][1]],edges[i][0])
	}

	queue := []int{}
	for i := 0;i<n;i++ {
		if degree[i] == 1 {
			queue = append(queue,i)
		}
	}

	for len(queue) > 0 {
		result = []int{}
		size := len(queue)
		for i := 0;i<size;i++ {
			curr := queue[0]
			queue = queue[1:]
			// queue 中都是出度为1的
			result = append(result,curr)
			neighbors := slice[curr]
			for _,neighbor := range neighbors {
				degree[neighbor] --
				if degree[neighbor] == 1 {
					queue = append(queue,neighbor)
				}
			}
		}
	}
	return result
}