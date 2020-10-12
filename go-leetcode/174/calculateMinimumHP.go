package main

import "fmt"

/*
174. 地下城游戏
一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。
我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。

骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。

有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；
其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。

为了尽快到达公主，骑士决定每次只向右或向下移动一步。



编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。

例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)


说明:

骑士的健康点数没有上限。

任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
 */


/*
0   -2  3
-1  0   0
-3  -3  -2

从(0,0)到(1,2)

路径一，0,-2,3,0 路径和为1，需要的最小初始值为3
路径二，0,-1,0,0 路径和为-1，需要的最小初始值为2

到终点-2，应该选择路径一

但是如果终点是0，应该选择路径二

因为有两个重要程度相同的参数同时影响后续的决策。也就是说，这样的动态规划是不满足「无后效性」的。

简单来讲，无后效性指的是当前状态确定后，之后的状态转移与之前的状态/决策无关。



于是我们考虑从右下往左上进行动态规划。令 dp[i][j] 表示从坐标 (i,j) 到终点所需的最小初始值。
换句话说，当我们到达坐标 (i,j) 时，如果此时我们的路径和不小于 dp[i][j]，我们就能到达终点。

这样一来，我们就无需担心路径和的问题，只需要关注最小初始值。对于 dp[i][j]，我们只要关心 dp[i][j+1] 和 dp[i+1][j] 的最小值 min。记当前格子的值为
dungeon(i,j)，那么在坐标 (i,j) 的初始值只要达到 min−dungeon(i,j) 即可。同时，初始值还必须大于等于 11。这样我们就可以得到状态转移方程：

dp[i][j]=max(min(dp[i+1][j],dp[i][j+1])−dungeon(i,j),1)

最终答案即为 dp[0][0]。



3   3   1
4   3   3
9   6   3


 */

func main() {
	matrix := [][]int{{0,-2,3},
		{-1,0,0},
	{-3,-3,-2},
	}
	matrix = [][]int{
		{-2,-3,3},
		{-5,-10,1},
		{10,30,-5},
	}
	dp := calculateMinimumHP(matrix)
	path := getPath(dp)
	fmt.Println(dp[0][0])
	for i := 0;i<len(path);i++ {
		fmt.Print(matrix[path[i][0]][path[i][1]],"->")
	}
	fmt.Println()
}

func getPath(matrix [][]int) [][]int {
	result := [][]int{}
	result = append(result,[]int{0,0})
	i,j := 0,0
	for ;i<len(matrix)-1;i++ {
		for ;j<len(matrix[0])-1;j++ {
			if matrix[i+1][j] < matrix[i][j+1] {
				result = append(result,[]int{i+1,j})
			} else {
				result = append(result,[]int{i,j+1})
			}
		}
	}
	if i == len(matrix)-1 {
		for ;j < len(matrix[0]);j++ {
			result = append(result,[]int{i,j})
		}
	}
	if j == len(matrix[0]) - 1 {
		for ;i < len(matrix);i++ {
			result = append(result,[]int{i,j})
		}
	}
	return result
}

func calculateMinimumHP(dungeon [][]int) [][]int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		//return 0
		return nil
	}
	dp := make([][]int,len(dungeon))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,len(dungeon[0]))
	}
	rowCount := len(dungeon)
	colCount := len(dungeon[0])

	dp[rowCount-1][colCount-1] = Max(1-dungeon[rowCount-1][colCount-1],1)
	for i := rowCount - 2;i>=0;i-- {
		dp[i][colCount-1] = Max(dp[i+1][colCount-1] - dungeon[i][colCount-1],1)
	}
	for j := colCount-2;j>=0;j-- {
		dp[rowCount-1][j] = Max(dp[rowCount-1][j+1] - dungeon[rowCount-1][j],1)
	}
	for i := rowCount-2;i>=0;i-- {
		for j := colCount-2;j>=0;j-- {
			min := Min(dp[i+1][j],dp[i][j+1])
			dp[i][j] = Max(min-dungeon[i][j],1)
		}
	}
	//return dp[0][0]
	return dp
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}