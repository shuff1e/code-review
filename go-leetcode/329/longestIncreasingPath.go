package main

/*

329. 矩阵中的最长递增路径
给定一个整数矩阵，找出最长递增路径的长度。

对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

示例 1:

输入: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
输出: 4
解释: 最长递增路径为 [1, 2, 6, 9]。
示例 2:

输入: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
输出: 4
解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。

 */

func longestIncreasingPath(matrix [][]int) int {
	memo := make([][]int,len(matrix))
	for i := 0;i<len(memo);i++ {
		memo[i] = make([]int,len(matrix[0]))
	}

	dirs := [][]int{
		{0,1},
		{0,-1},
		{1,0},
		{-1,0},
	}

	result := 0
	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0]);j++ {
			result = Max(result,help(matrix,i,j,memo,dirs))
		}
	}
	return result
}

func help(matrix [][]int,row,col int,memo [][]int,dirs [][]int) int {
	if memo[row][col] != 0 {
		return memo[row][col]
	}

	memo[row][col] ++
	for i := 0;i<len(dirs);i++ {
		newRow := row + dirs[i][0]
		newCol := col + dirs[i][1]

		if newRow >= 0 && newRow < len(matrix) &&
			newCol >= 0 && newCol < len(matrix[0]) &&
			matrix[row][col] < matrix[newRow][newCol] {
			memo[row][col] = Max(memo[row][col],help(matrix,newRow,newCol,memo,dirs) + 1)
		}
	}
	return memo[row][col]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*


方法二：拓扑排序
从方法一可以看到，每个单元格对应的最长递增路径的结果只和相邻单元格的结果有关，那么是否可以使用动态规划求解？

根据方法一的分析，动态规划的状态定义和状态转移方程都很容易得到。方法一中使用的缓存矩阵 memo 即为状态值，
状态转移方程如下：

memo[i][j] = max(memo[x][y]) + 1
其中(x, y)与(i, j)在矩阵中相邻，并且 matrix[x][y] > matrix[i][j]
​
动态规划除了状态定义和状态转移方程，还需要考虑边界情况。这里的边界情况是什么呢？

如果一个单元格的值比它的所有相邻单元格的值都要大，那么这个单元格对应的最长递增路径是 1，这就是边界条件。
这个边界条件并不直观，而是需要根据矩阵中的每个单元格的值找到作为边界条件的单元格。

仍然使用方法一的思想，将矩阵看成一个有向图，计算每个单元格对应的出度，即有多少条边从该单元格出发。
对于作为边界条件的单元格，该单元格的值比所有的相邻单元格的值都要大，因此作为边界条件的单元格的出度都是 0。

基于出度的概念，可以使用拓扑排序求解。从所有出度为 0的单元格开始广度优先搜索，
每一轮搜索都会遍历当前层的所有单元格，更新其余单元格的出度，并将出度变为 0的单元格加入下一层搜索。
当搜索结束时，搜索的总层数即为矩阵中的最长递增路径的长度。

 */

/*

class Solution {
    public int[][] dirs = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    public int rows, columns;

    public int longestIncreasingPath(int[][] matrix) {
        if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
            return 0;
        }
        rows = matrix.length;
        columns = matrix[0].length;
        int[][] outdegrees = new int[rows][columns];
        for (int i = 0; i < rows; ++i) {
            for (int j = 0; j < columns; ++j) {
                for (int[] dir : dirs) {
                    int newRow = i + dir[0], newColumn = j + dir[1];
                    if (newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j]) {
                        ++outdegrees[i][j];
                    }
                }
            }
        }



        Queue<int[]> queue = new LinkedList<int[]>();
        for (int i = 0; i < rows; ++i) {
            for (int j = 0; j < columns; ++j) {
                if (outdegrees[i][j] == 0) {
                    queue.offer(new int[]{i, j});
                }
            }
        }



        int ans = 0;
        while (!queue.isEmpty()) {
            ++ans;
            int size = queue.size();
            for (int i = 0; i < size; ++i) {
                int[] cell = queue.poll();
                int row = cell[0], column = cell[1];
                for (int[] dir : dirs) {
                    int newRow = row + dir[0], newColumn = column + dir[1];
                    if (newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[row][column]) {
                        --outdegrees[newRow][newColumn];
                        if (outdegrees[newRow][newColumn] == 0) {
                            queue.offer(new int[]{newRow, newColumn});
                        }
                    }
                }
            }
        }

        return ans;



    }
}

 */