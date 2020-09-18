package main

import "fmt"

var matrix [][]int

func he(arr []int,start,end int) int {
	if matrix[start][end] != 0 {
		return matrix[start][end]
	}

	if start == end {
		return arr[start]
	}
	if start + 1 == end && arr[start] == arr[end] {
		return 2*arr[start]
	} else if arr[start] == arr[end] {
		matrix[start+1][end-1] = he(arr,start+1,end-1)
		return matrix[start+1][end-1] + 2*arr[start]
	} else {
		matrix[start+1][end] = he(arr,start+1,end)
		left := matrix[start+1][end] + 2 * arr[start]

		matrix[start][end-1] = he(arr,start,end-1)
		right := matrix[start][end-1] + 2*arr[end]
		if left < right {
			return left
		} else {
			return right
		}
	}
}

func getMinCost(arr []int) int {
	length := len(arr)
	matrix = make([][]int,length)
	for i:=0;i<length;i++ {
		matrix[i] = make([]int,length)
	}
	return he(arr,0,length-1)
}

func main() {
	arr := []int{1,2,3,4,5,6,7,2,3,8}
	cost := getMinCost(arr)
	fmt.Println(cost)
	cost = dpWay(arr)
	fmt.Println(cost)
}

func dpWay(arr []int) int {
	length := len(arr)

	dp := make([][]int,length)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int,length)
	}

	for i := 0;i<length;i++ {
		dp[i][i] = arr[i]
	}

	for i :=0;i+1<length;i++ {
		if arr[i] == arr[i+1] {
			dp[i][i+1] = 2*arr[i]
		} else if arr[i] < arr[i+1] {
			dp[i][i+1] = 2*arr[i] + arr[i+1]
		} else {
			dp[i][i+1] = 2*arr[i+1] + arr[i]
		}
	}

	for diff := 2;diff<length;diff++ {
		for i := 0;i+diff<length;i ++ {
			if arr[i] == arr[i+diff] {
				dp[i][i+diff] = dp[i+1][i+diff-1] + 2*arr[i]
			} else {
				dp[i][i+diff] = minCostMin(dp[i+1][i+diff]+2*arr[i],dp[i][i+diff-1]+2*arr[i+diff])
			}
		}
	}
	return dp[0][length-1]
	// dp[start][end] = min(dp[start+1][end]+2*arr[start],dp[start][end-1]+arr[end])
}

func minCostMin(x,y int) int {
	if x > y {
		return y
	}
	return x
}