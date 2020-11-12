package main

import "fmt"

/*

312. 戳气球
有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。

求所能获得硬币的最大数量。

说明:

你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
示例:

输入: [3,1,5,8]
输出: 167
解释: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
     coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

 */

func main() {
	arr := []int{}
	fmt.Println(maxCoins(arr))
}

func maxCoins(nums []int) int {
	n := len(nums)
	temp := make([]int,n+2)
	temp[0] = 1
	temp[n+1] = 1

	for i := 0;i<len(nums);i++ {
		temp[i+1] = nums[i]
	}

	dp := make([][]int,n+2)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,n+2)
	}

	// 长度
	for length := 3;length<=n+2;length ++ {
		for i := 0;i<=n+2-length;i++ {
			res := 0
			for k := i + 1;k<i+length-1;k++ {
				left := dp[i][k]
				right := dp[k][i+length-1]
				res = Max(res,left + right + temp[i] * temp[k] * temp[i+length-1])
			}
			dp[i][i+length-1] = res
		}
	}
	return dp[0][n+1]
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
