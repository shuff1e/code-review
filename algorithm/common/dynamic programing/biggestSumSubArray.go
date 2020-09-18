package main

import (
	"algorithm/common/help"
	"fmt"
)

func getDP(arr []int) []int {
	dp := make([]int,len(arr))
	for i := 0;i<len(arr);i++ {
		dp[i] = 1
		for j := i-1;j>=0;j-- {
			if arr[j] < arr[i] {
				dp[i] = help.Max(dp[j] + 1,dp[i])
			}
		}
	}
	return dp
}

func getList(dp []int,arr []int) []int {
	maxDP := 0
	maxIndex := 0

	for i := 0;i<len(dp);i++ {
		if dp[i] > maxDP {
			maxDP = dp[i]
			maxIndex = i
		}
	}

	result := []int{}

	for i := maxIndex;i >= 0;i-- {
		if arr[i] < arr[maxIndex] && dp[i] == dp[maxIndex] -1 {
			result = append(result,arr[i])
			maxIndex = i
		}
	}
	return result
}

func main() {
	arr := []int{1,3,5,9,7,11,6}
	dp := getDP2(arr)
	fmt.Printf("%#v\n",dp)
	list := getList(dp,arr)
	fmt.Printf("%#v\n",list)
}

func getDP2(arr []int) []int {
	dp := make([]int,len(arr))
	ends := make([]int,len(dp))
	ends[0] = arr[0]
	dp[0] = 1
	right,l,r,m := 0,0,0,0
	for i := 1;i<len(arr);i++ {
		l = 0
		r = right
		for l <= r {
			m = (l+r)/2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m -1
			}
		}
		ends[l] = arr[i]
		right = help.Max(right,l)
		dp[i] = l + 1
	}
	return dp
}


