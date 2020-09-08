package array

import (
	"fmt"
	"github.com/shuff1e/code-review/util/math"
)

// Q:给定数组arr，返回arr的最长递增子序列
// arr = [2,1,5,3.1,6,4,8,9,7]，返回的最长递增子序列为[1,3.1,4,8,9]

// A:dp[i] = Max(dp[j]) + 1 (j<i,arr[j]<arr[i])
// N2
func getDP1(arr []int) []int {
	dp := make([]int,len(arr))
	for i := 0;i<len(arr);i++ {
		dp[i] = 1
		for j := i-1;j>=0;j-- {
			if arr[j] < arr[i] {
				dp[i] = math.Max(dp[i],dp[j]+1)
			}
		}
	}
	return dp
}

func getSubArr(arr []int,dp []int) []int {
	result := []int{}
	maxDP := 0
	maxIndex := 0
	for i := 0;i<len(dp);i++ {
		if dp[i] > maxDP {
			maxDP = dp[i]
			maxIndex = i
		}
	}
	result = append(result,arr[maxIndex])
	for i := maxIndex;i>=0;i--{
		if arr[i] < arr[maxIndex] && dp[maxIndex] == dp[i] + 1 {
			maxIndex = i
			result = append(result,arr[i])
		}
	}
	return result
}

func SolveArray01(arr []int) []int {
	dp := getDP2(arr)
	fmt.Printf("%#v\n",dp)
	result := getSubArr(arr,dp)
	return result
}

// dp[i] = Max(dp[j]) + 1 (j<i,arr[j]<arr[i])
// 找小于arr[i]且dp最大的j的过程变成二分查找
// NlogN
func getDP2(arr []int) []int {
	// right表示ends数组的界限，0到right的部分是有效的
	// 假设ends数组中index=1的地方为元素3，
	// 则表示目前为止，长度为index+1的子序列中，结尾元素最小的是3
	// 例如 arr为[2,1,5,3.1,6,4,8,9,7]
	// 则ends为[1,3.1]表示长度为2的子序列中，以3为结尾是最小的

	dp := make([]int,len(arr))
	ends := make([]int,len(arr))

	dp[0]=1
	ends[0] = arr[0]
	right,l,m,r := 0,0,0,0
	for i := 0;i<len(arr);i++ {
		l = 0
		r = right
		for l <= r {
			m = (l+r)/2
			if arr[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		ends[l] = arr[i]
		right = math.Max(right,l)
		dp[i] = l + 1
	}
	return dp
}