package array

import (
	"github.com/shuff1e/code-review/util/math"
)

// 经典的01背包问题

// 将数组分为两部分，使得两部分的和最接近，返回两部分的差值。
// 例如：int[] array={1,0,1,7,2,4}，分为两部分为{1,0,1,2,4}，{7}，差值为1。

// 一种双核CPU的两个核能够同时的处理任务，现在有n个已知数据量的任务需要交给CPU处理，
// 假设已知CPU的每个核1秒可以处理1kb，每个核同时只能处理一项任务。n个任务可以按照任意顺序放入CPU进行处理，
// 现在需要设计一个方案让CPU处理完这批任务所需的时间最少，求这个最小的时间。


// dp[i][j]表示背包大小为j，当前为i的情况下，能放的最大重量
// dp[i][j] = max(dp[i-1][j],dp[i-1][j-arr[i]]+arr[i]) arr[i] < j
// 			  dp[i-1][j]
func GetDiff(arr []int) int {
	sum := 0
	for _,v := range arr {
		sum += v
	}

	dp := make([][]int,len(arr))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,sum/2+1)
	}

	for i := 0;i<len(arr);i++ {
		dp[i][0] = 0
	}

	for j := 1;j<=sum/2;j++ {
		if j > arr[0] {
			dp[0][j] = arr[0]
		} else {
			dp[0][j] = 0
		}
	}

	for i := 1;i<len(arr);i++ {
		for j := 1;j<=sum/2;j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i] {
				dp[i][j] = math.Max(dp[i][j],dp[i-1][j-arr[i]]+arr[i])
			}
		}
	}
	one := dp[len(arr)-1][sum/2]
	other := sum - one
	return math.Abs(one-other)
}