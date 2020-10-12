package main

import "fmt"

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1:

输入:
[
['1','1','1','1','0'],
['1','1','0','1','0'],
['1','1','0','0','0'],
['0','0','0','0','0']
]
输出: 1
示例 2:

输入:
[
['1','1','0','0','0'],
['1','1','0','0','0'],
['0','0','1','0','0'],
['0','0','0','1','1']
]
输出: 3
解释: 每座岛屿只能由水平和/或竖直方向上相邻的陆地连接而成。
 */

func main() {
	grid := [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	count := 0

	visited := make([][]bool,len(grid))
	for i := 0;i<len(grid);i++ {
		visited[i] = make([]bool,len(grid[0]))
	}

	for i :=0;i<len(grid);i++ {
		for j := 0;j<len(grid[0]);j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count ++
				help(i,j,grid,visited)
			}
		}
	}
	return count
}

func help(row,col int,grid [][]byte,visited [][]bool) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return
	}
	// 陆地无法延伸
	if grid[row][col] == '0' {
		return
	}
	if visited[row][col] {
		return
	}
	visited[row][col] = true
	help(row+1,col,grid,visited)
	help(row-1,col,grid,visited)
	help(row,col+1,grid,visited)
	help(row,col-1,grid,visited)
}

/*
BFS

class Solution {
    public int numIslands(char[][] grid) {
        if (grid == null || grid.length == 0) {
            return 0;
        }

        int nr = grid.length;
        int nc = grid[0].length;
        int num_islands = 0;

        for (int r = 0; r < nr; ++r) {
            for (int c = 0; c < nc; ++c) {
                if (grid[r][c] == '1') {
                    ++num_islands;
                    grid[r][c] = '0';
                    Queue<Integer> neighbors = new LinkedList<>();
                    neighbors.add(r * nc + c);
                    while (!neighbors.isEmpty()) {
                        int id = neighbors.remove();
                        int row = id / nc;
                        int col = id % nc;
                        if (row - 1 >= 0 && grid[row-1][col] == '1') {
                            neighbors.add((row-1) * nc + col);
                            grid[row-1][col] = '0';
                        }
                        if (row + 1 < nr && grid[row+1][col] == '1') {
                            neighbors.add((row+1) * nc + col);
                            grid[row+1][col] = '0';
                        }
                        if (col - 1 >= 0 && grid[row][col-1] == '1') {
                            neighbors.add(row * nc + col-1);
                            grid[row][col-1] = '0';
                        }
                        if (col + 1 < nc && grid[row][col+1] == '1') {
                            neighbors.add(row * nc + col+1);
                            grid[row][col+1] = '0';
                        }
                    }
                }
            }
        }

        return num_islands;
    }
}

 */

/*

我们用一个数组rank[]记录每个根节点对应的树的深度（如果不是根节点，其rank相当于以它作为根节点的子树的深度）。
一开始，把所有元素的rank（秩）设为1。合并时比较两个根节点，把rank较小者往较大者上合并。
路径压缩和按秩合并如果一起使用，时间复杂度接近O(n) ，但是很可能会破坏rank的准确性。

 */

type UnionFind struct {
	parent []int
	rank []int
	count int
}

// 返回 i的根节点
func (u *UnionFind) Find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.Find(u.parent[i])
	}
	return u.parent[i]
}

func (u *UnionFind) Union(x, y int) {
	px := u.Find(x)
	py := u.Find(y)
	if px == py {
		return
	}
	if u.rank[px] < u.rank[py] {
		u.parent[px] = py
	} else if u.rank[px] > u.rank[py] {
		u.parent[py] = px
	} else {
		u.parent[py] = px
		u.rank[px] = u.rank[px] + 1
	}
	u.count --
}

func (u *UnionFind) GetCount() int {
	return u.count
}

func NewUnionFind(grid [][]byte) *UnionFind {
	count := 0
	m := len(grid)
	n := len(grid[0])
	parent := make([]int,m*n)
	rank := make([]int,m*n)

	for i := 0;i<m;i++ {
		for j := 0;j<n;j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n+j
				count ++
			}
			rank[i*n+j] = 0
		}
	}
	return &UnionFind{
		count: count,
		parent: parent,
		rank: rank,
	}
}

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0{
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	u := NewUnionFind(grid)
	for i := 0;i<m;i++ {
		for j := 0;j<n;j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1>=0 && grid[i-1][j] == '1' {
					u.Union(i*n+j,(i-1)*n+j)
				}
				if i+1<m && grid[i+1][j] == '1' {
					u.Union(i*n+j,(i+1)*n+j)
				}
				if j-1 >=0 && grid[i][j-1] == '1' {
					u.Union(i*n+j,i*n+j-1)
				}
				if j+1 <n && grid[i][j+1] == '1' {
					u.Union(i*n+j,i*n+j+1)
				}
			}
		}
	}
	return u.GetCount()
}