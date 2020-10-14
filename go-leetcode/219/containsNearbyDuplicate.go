package main

import "fmt"

/*
219. 存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

示例 1:

输入: nums = [1,2,3,1], k = 3
输出: true
示例 2:

输入: nums = [1,0,1,1], k = 1
输出: true
示例 3:

输入: nums = [1,2,3,1,2,3], k = 2
输出: false
 */


func main() {
	arr := []int{1}
	k := 1
	fmt.Println(containsNearbyDuplicate(arr,k))
}

// 滑动窗口
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	if k == 0 {
		return false
	}
	dict := map[int]struct{}{}

	for i := 0;i<k+1 && i <len(nums);i++ {
		if _,ok := dict[nums[i]];ok {
			return true
		} else {
			dict[nums[i]] = struct{}{}
		}
	}

	for i := 1;i<=len(nums)-k-1;i++ {
		delete(dict,nums[i-1])
		if _,ok := dict[nums[i+k]];ok {
			return true
		}
		dict[nums[i+k]] = struct{}{}

	}
	return false
}

/*


public boolean containsNearbyDuplicate(int[] nums, int k) {
    Set<Integer> set = new HashSet<>();
    for (int i = 0; i < nums.length; ++i) {
        if (set.contains(nums[i])) return true;
        set.add(nums[i]);
        if (set.size() > k) {
            set.remove(nums[i - k]);
        }
    }
    return false;
}

 */