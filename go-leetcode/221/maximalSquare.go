package main

import "fmt"

/*
221. 最大正方形
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
 */

/*

可以使用动态规划降低时间复杂度。
我们用 dp(i,j) 表示以 (i,j) 为右下角，且只包含 1 的正方形的边长最大值。如果我们能计算出所有 dp(i,j) 的值，
那么其中的最大值即为矩阵中只包含 1 的正方形的边长最大值，其平方即为最大正方形的面积。

那么如何计算 dp 中的每个元素值呢？对于每个位置 (i,j)，检查在矩阵中该位置的值：

如果该位置的值是 0，则 dp(i, j) = 0，因为当前位置不可能在由 1 组成的正方形中；

如果该位置的值是 1，则 dp(i,j) 的值由其上方、左方和左上方的三个相邻位置的 dp 值决定。
具体而言，当前位置的元素值等于三个相邻位置的元素中的最小值加 1，状态转移方程如下：

dp(i,j)=min(dp(i−1,j),dp(i−1,j−1),dp(i,j−1))+1

如果读者对这个状态转移方程感到不解，可以参考 1277. 统计全为 1 的正方形子矩阵的官方题解，其中给出了详细的证明。

此外，还需要考虑边界条件。如果 i 和 j 中至少有一个为 0，则以位置 (i,j) 为右下角的最大正方形的边长只能是 1，因此 dp(i,j)=1。

以下用一个例子具体说明。原始矩阵如下。


0 1 1 1 0
1 1 1 1 0
0 1 1 1 1
0 1 1 1 1
0 0 1 1 1

对应的 dp 值如下。


0 1 1 1 0
1 1 2 2 0
0 1 2 3 1
0 1 2 3 2
0 0 1 2 3

 */

func main() {
	matrix := [][]byte{
		{'1','0','1','0','0'},
		{'1','0','1','1','1'},
		{'1','1','1','1','1'},
		{'1','0','0','1','0'}}
	matrix = [][]byte{{'1'}}
	fmt.Println(maximalSquare2(matrix))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int,m)
	for i := 0;i<m;i++ {
		dp[i] = make([]int,n)
	}
	max := 0

	for i := 0;i<m;i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			max = Max(max,1)
		}
	}

	for j := 0;j<n;j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			max = Max(max,1)
		}
	}

	for i := 1;i<m;i++ {
		for j := 1;j<n;j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = Min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1
				max = Max(max,dp[i][j])
			}
		}
	}
	return max*max
}

func Min(args ...int) int {
	x := args[0]
	for i :=1;i<len(args);i++ {
		if x > args[i] {
			x = args[i]
		}
	}
	return x
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	return getDP(matrix)
}

func getDP(matrix [][]byte) int {
	result := 0
	dp := make([]int,len(matrix[0]))
	for i := 0;i<len(matrix);i++ {
		for j := 0;j<len(matrix[0]);j++ {
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		result = Max(result,getArea(dp))
	}
	return result
}

func getArea(arr []int) int {
	// 单调栈
	// for i := range arr
	//     for len(stack) > 0 && arr[stack.peek()] > arr[i]
	//		    temp := stack.pop()
	//          left := stack.empty()?-1:stack.peek()
	//          right := i - 1
	//     stack.push(i)
	arr = append(arr,0)
	max := 0
	stack := []int{}
	for i := 0;i<len(arr);i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}
			right := i-1
			// 底和高中最小的一个作为边
			min := Min(arr[temp],right-left)
			max = Max(max,min*min)
		}
		stack = append(stack,i)
	}
	return max
}

func Min2(x,y int) int {
	if x < y {
		return x
	}
	return y
}