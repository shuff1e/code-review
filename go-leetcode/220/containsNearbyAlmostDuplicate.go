package main

import "fmt"

/*
220. 存在重复元素 III
在整数数组 nums 中，是否存在两个下标 i 和 j，使得 nums [i] 和 nums [j] 的差的绝对值小于等于 t ，
且满足 i 和 j 的差的绝对值也小于等于 ķ 。

如果存在则返回 true，不存在返回 false。

示例 1:

输入: nums = [1,2,3,1], k = 3, t = 0
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1, t = 2
输出: true
示例 3:

输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false

 */


func main() {
	arr := []int{1,5,9,1,5,9}
	k := 2
	t := 3
	fmt.Println(containsNearbyAlmostDuplicate(arr,k,t))
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}

	dict := map[int]int{}
	w := t+1
	for i := 0;i<len(nums);i++ {
		id := getID(nums[i],w)
		if _,ok := dict[id];ok {
			return true
		}
		if v,ok := dict[id+1];ok && Abs(v-nums[i]) < w {
			return true
		}
		if v,ok := dict[id-1];ok && Abs(v-nums[i]) < w {
			return true
		}
		dict[id] = nums[i]
		if i >= k {
			delete(dict,getID(nums[i-k],w))
		}
	}
	return false
}

// 1,5,9,1,5,9
//k=2,t=3
//
//w=4
//0,1,2,3放一个桶 x/w
//
//-1,-2,-3,-4放一个桶
//(x+1)/w - 1

func getID(x ,w int) int {
	if x < 0 {
		return (x+1)/w - 1
	}
	return x/w
}

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
