package main

/*
416. 分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.
 */

// 背包问题

// dp[i][volume] = Max(dp[i-1][volume-volume[i]] + value[i],dp[i-1][volume])

func canPartition(nums []int) bool {
	sum := 0
	for i := 0;i<len(nums);i ++ {
		sum += nums[i]
	}
	if sum % 2 != 0 {
		return false
	}

	sum = sum/2
	dp := make([][]int,len(nums))

	for i := 0;i<len(dp);i++ {
		dp[i] = make([]int,sum+1)
	}

	for j := 0;j<len(dp[0]);j++ {
		if j >= nums[0] {
			dp[0][j] = nums[0]
		}
	}

	for i := 1;i<len(dp);i++ {
		for j := 0;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i] {
				dp[i][j] = Max(dp[i][j],dp[i-1][j-nums[i]] + nums[i])
			}
		}
	}
	return dp[len(dp)-1][sum] == sum
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*
创建二维数组 dp，包含 n 行 target+1 列，其中 dp[i][j] 表示从数组的 [0,i] 下标范围内选取若干个正整数（可以是 0 个），
是否存在一种选取方案使得被选取的正整数的和等于 j。初始时，dp 中的全部元素都是 false。

在定义状态之后，需要考虑边界情况。以下两种情况都属于边界情况。

如果不选取任何正整数，则被选取的正整数等于 0。因此对于所有 0≤i<n，都有 dp[i][0]=true。

当 i==0 时，只有一个正整数 nums[0] 可以被选取，因此 dp[0][nums[0]]=true。

对于 i>0 且 j>0 的情况，如何确定 dp[i][j] 的值？需要分别考虑以下两种情况。

如果 j≥nums[i]，则对于当前的数字 nums[i]，可以选取也可以不选取，两种情况只要有一个为 true，
就有 dp[i][j]=true。

如果不选取 nums[i]，则 dp[i][j]=dp[i−1][j]；
如果选取 nums[i]，则 dp[i][j]=dp[i−1][j−nums[i]]。
如果 j<nums[i]，则在选取的数字的和等于 j 的情况下无法选取当前的数字 nums[i]，
因此有 dp[i][j]=dp[i−1][j]。

状态转移方程如下：

dp[i][j]={
dp[i−1][j] ∣ dp[i−1][j−nums[i]], j≥nums[i]
dp[i−1][j], j<nums[i]
​
最终得到 dp[n−1][target] 即为答案。
 */

func canPartition2(nums []int) bool {
	sum := 0
	for i := 0;i<len(nums);i ++ {
		sum += nums[i]
	}
	if sum % 2 != 0 {
		return false
	}

	sum = sum/2
	dp := make([][]bool,len(nums))

	for i := 0;i<len(dp);i++ {
		dp[i] = make([]bool,sum+1)
	}

	for i := 0;i<len(dp);i++ {
		dp[i][0] = true
	}

	if sum >= nums[0] {
		dp[0][nums[0]] = true
	}
	for i := 1;i<len(dp);i++ {
		for j := 1;j<len(dp[0]);j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[len(dp)-1][sum]
}