package main

import "fmt"

// V(i-1,W) = Max(V(i-1,W),V(i-1,W-wi)+vi)
func beibao(weight []int,value []int,capacity int) []int {
	dp := make([][]int,len(weight)+1)

	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,capacity+1)
	}

	for i := 0;i<len(dp);i++ {
		dp[i][0] = 0
	}

	for j := 0;j<capacity+1;j++ {
		dp[0][j] = 0
	}

	for i := 1;i<len(dp);i++ {
		for j:=1;j<capacity+1;j++ {
			if weight[i-1] <= j {
				dp[i][j] = Max(dp[i-1][j],dp[i-1][j-weight[i-1]]+value[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	result := []int{}

	i,j := len(dp)-1,capacity

	for ;i > 0;i-- {
		if dp[i][j] == dp[i-1][j] {
			continue
		} else {
			result = append(result,i-1)
			j -= weight[i-1]
		}
	}

	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	weight := []int{1,2,5,6,7}
	value := []int{1,6,18,22,28}
	capacity := 11
	result := beibao(weight,value,capacity)
	fmt.Println(result)

}