package main

/*
128. 最长连续序列
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

 */

// A：如果数组排序过，则很容易找到最长连续序列
// 但是排序的时间复杂度一般都是O(nlogn)

// 换一种思路
// 对于每个数x，枚举x+1,x+2一直到x+n是否存在
// 当x-1存在时，x是否存在不用计算了，因为x-1会覆盖这种情况

// 只用遍历一次，时间复杂度时O(n)
// 空间复杂度是O(n)

func longestConsecutive(nums []int) int {
	mmp := make(map[int]int,len(nums))
	for _,v := range nums {
		mmp[v] = v
	}
	result := 0
	for k,_ := range mmp {
		if _,ok := mmp[k-1];ok {
			continue
		}
		cur := k+1
		length := 1
		for {
			if _,ok := mmp[cur];ok {
				cur ++
				length ++
			} else {
				break
			}
		}
		result = Max(result,length)
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}