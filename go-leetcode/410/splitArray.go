package main

import "fmt"

/*

410. 分割数组的最大值
给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。
设计一个算法使得这 m 个子数组各自和的最大值最小。

注意:
数组长度 n 满足以下条件:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
示例:

输入:
nums = [7,2,5,10,8]
m = 2

输出:
18

解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。

 */

/*

方法一：动态规划
思路与算法

「将数组分割为 m 段，求……」是动态规划题目常见的问法。

本题中，我们可以令 f[i][j] 表示将数组的前 i 个数分割为 j 段所能得到的最大连续子数组和的最小值。
在进行状态转移时，我们可以考虑第 j 段的具体范围，即我们可以枚举 k，其中前 k 个数被分割为 j−1 段，
而第 k+1 到第 i 个数为第 j 段。此时，这 j 段子数组中和的最大值，
就等于 f[k][j−1] 与 sub(k+1,i) 中的较大值，其中 sub(i,j) 表示数组 nums 中下标落在区间 [i,j] 内的数的和。

由于我们要使得子数组中和的最大值最小，因此可以列出如下的状态转移方程：

f[i][j]=

k=0
min{max(f[k][j−1],sub(k+1,i))}
i−1

​


对于状态 f[i][j]，由于我们不能分出空的子数组，因此合法的状态必须有 i≥j。
对于不合法（i<j）的状态，由于我们的目标是求出最小值，因此可以将这些状态全部初始化为一个很大的数。
在上述的状态转移方程中，一旦我们尝试从不合法的状态 f[k][j−1] 进行转移，那么 max(⋯) 将会是一个很大的数，就不会对最外层的 min{⋯} 产生任何影响。

此外，我们还需要将 f[0][0] 的值初始化为 0。在上述的状态转移方程中，当 j=1 时，唯一的可能性就是前 i 个数被分成了一段。
如果枚举的 k=0，那么就代表着这种情况；如果 k !=0，对应的状态 f[k][0] 是一个不合法的状态，无法进行转移。
因此我们需要令 f[0][0]=0。

最终的答案即为 f[n][m]。

 */

// 数组的前 i 个数分割为 j 段
// sumArr[i] 表示 前i个数的和

// f[i][j] = min{max(f[k][j−1],sub(k+1,i))}

func main() {
	arr := []int{7,2,5,10,8}
	result := splitArray2(arr,2)
	fmt.Println(result)
}

func splitArray(nums []int, m int) int {
	if len(nums) == 0 {
		return 0
	}
	const INTEGER_MAX = 0x7fffffff

	sumArr := make([]int,len(nums)+1)
	for i := 0;i<len(nums);i++ {
		sumArr[i+1] = sumArr[i] + nums[i]
	}
	// [i,j] sumArr[j+1]-sumArr[i]

	dp := make([][]int,len(nums) + 1)
	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,m + 1)
		for j := 0;j<len(dp[i]);j++ {
			dp[i][j] = INTEGER_MAX
		}
	}

	dp[0][0] = 0

	dp[1][1] = nums[0]
	// dp[1][1] = Min(dp[1][1],Max(dp[0][0],sumArr[1]-sumArr[0]))

	// dp[2][1]   dp[2][2]

	for i:=2;i<=len(nums);i++ {
		for j := 1;j<=Min(i,m);j++ {
			for k := 0;k<i;k++ {
				dp[i][j] = Min(dp[i][j],Max(dp[k][j-1],sumArr[i] - sumArr[k]))
			}
		}
	}
	return dp[len(nums)][m]
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func splitArray2(nums []int, m int) int {
	// left和right是可能的结果的最大值和最小值
	left := 0
	right := 0
	for i := 0;i<len(nums);i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}

	for left < right {
		mid := (left + right)/2
		if checkValid(nums,mid,m) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 是否能分成m段，每段小于等于x
// 如果能，说明x取大了
// 如果不能，说明x取小了

func checkValid(arr []int,x,m int) bool {
	count := 1
	sum := 0
	for i := 0;i<len(arr);i++ {
		if sum + arr[i] > x {
			count ++
			sum = arr[i]
		} else {
			sum += arr[i]
		}
	}
	return count <= m
}