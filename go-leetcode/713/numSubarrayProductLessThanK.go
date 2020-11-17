package main

import (
	"fmt"
	"math"
)

/*

713. 乘积小于K的子数组
给定一个正整数数组 nums。

找出该数组内乘积小于 k 的连续的子数组的个数。

示例 1:

输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8个乘积小于100的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
说明:

0 < nums.length <= 50000
0 < nums[i] < 1000
0 <= k < 10^6

 */

// 10,5,2,6
// [10]
// [5] [10,5]
// [2] [5,2]
// [6] [2,6] [5,2,6]

func main() {
	arr := []int{10,5,2,6}
	k := 100
	result := numSubarrayProductLessThanK(arr,k)
	fmt.Println(result)
	result2 := numSubarrayProductLessThanK2(arr,k)
	fmt.Println(result2)
}

/*

方法一：二分查找
分析

我们可以使用二分查找解决这道题目，即对于固定的 i，二分查找出最大的 j 满足 nums[i] 到 nums[j] 的乘积小于 k。但由于乘积可能会非常大（在最坏情况下会达到 1000^{50000}），会导致数值溢出，因此我们需要对 nums 数组取对数，将乘法转换为加法，即
​
 nums[i]=∑lognums[i]，这样就不会出现数值溢出的问题了。

算法

对 nums 中的每个数取对数后，我们存储它的前缀和 prefix，即 prefix[i+1]=∑x=0i nums[x]，这样在二分查找时，对于 i 和 j，我们可以用 prefix[j+1]−prefix[i] 得到 nums[i] 到 nums[j] 的乘积的对数。对于固定的 i，当找到最大的满足条件的 j 后，它会包含 j−i+1 个乘积小于 k 的连续子数组。

下面的代码和算法中下标的定义略有不同。

 */

func numSubarrayProductLessThanK(nums []int, k int) int {
	result := 0
	product := 1
	left,right := 0,0
	for ;right < len(nums);right ++ {
		product *= nums[right]
		for product >= k && left <= right {
			product /= nums[left]
			left ++
		}
		result += right - left + 1
	}
	return result
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	logk := math.Log(float64(k))
	// 前缀和
	prefix := make([]float64,len(nums) + 1 )
	for i := 0;i<len(nums);i++ {
		prefix[i+1] = prefix[i] + math.Log(float64(nums[i]))
	}

	result := 0
	for i := 0;i<len(prefix);i++ {
		l := i + 1
		r := len(prefix)
		for l < r {
			mid := (l + r)/2
			// i - > mid
			// prefix[mid] - prefix[i]
			// 相当于 nums[i] + nums[i+1] + ... + nums[mid-1]

			// 找到从左往右数
			// 第一个大于或者等于 logk的数
			if prefix[mid] < prefix[i] + logk - 1e-9 {
				l = mid + 1
			} else {
				r = mid
			}
			// (l-1) - (i-1) + 1
		}
		result += l - i - 1
	}
	return result
}

/*

func main() {
	arr := []int{1,1,2,2,3,4,5,6}
	k := 2
	result := help(arr,k)
	// 2
	fmt.Println(result)
}

func help(arr []int,k int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		mid := (l + r)/2
		if arr[mid] < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

 */