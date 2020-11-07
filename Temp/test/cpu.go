package main

import "fmt"

// 一种双核CPU的两个核能够同时的处理任务，现在有n个已知数据量的任务需要交给CPU处理，
// 假设已知CPU的每个核1秒可以处理1kb，每个核同时只能处理一项任务。n个任务可以按照任意顺序放入CPU进行处理，
// 现在需要设计一个方案让CPU处理完这批任务所需的时间最少，求这个最小的时间。

func main() {
	arr := []int{1,1,1,7,2,4}
	fmt.Println(getMinTime(arr))
}

func getMinTime(arr []int) int {
	sum := 0
	for i := 0;i<len(arr);i ++ {
		sum += arr[i]
	}
	result := help2(arr,sum/2)
	return sum - sum/2 - result
}


// dp[i][j] = Max(dp[i-1][j],dp[i-1][j-arr[i]] + arr[i])

func help2(arr []int,capacity int ) int {
	dp := make([][]int,len(arr))
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,capacity+1)
	}
	for i := 0;i<len(dp);i++ {
		dp[i][0] = 0
	}
	for j := 0;j<len(dp[0]);j++ {
		if j >= arr[0] {
			dp[0][j] = arr[0]
		}
	}

	for i := 1;i<len(dp);i++ {
		for j := 1;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i] {
				dp[i][j] = Max(dp[i][j],dp[i-1][j-arr[i]] + arr[i])
			}
		}
	}
	return dp[len(dp)-1][capacity]
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}