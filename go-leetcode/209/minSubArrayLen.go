package main

import "fmt"

/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 s ，
找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：

输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。


进阶：

如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 */

// A：滑动窗口

func main() {
	arr := []int{}
	fmt.Println(minSubArrayLen(3,arr))
}

func minSubArrayLen(s int, nums []int) int {
	left,right := 0,0
	result := 0x7fffffff
	sum := 0
	for left <= right && right < len(nums) {
		for sum < s && right < len(nums) {
			sum += nums[right]
			right ++
		}
		for sum >= s && left < right {
			result = Min(result,right-left)
			sum -= nums[left]
			left ++
		}
	}
	if result == 0x7fffffff {
		return 0
	}
	return result
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

/*

class Solution {
    public int minSubArrayLen(int s, int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return 0;
        }
        int ans = Integer.MAX_VALUE;
        int[] sums = new int[n + 1];
        // 为了方便计算，令 size = n + 1
        // sums[0] = 0 意味着前 0 个元素的前缀和为 0
        // sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
        // 以此类推
        for (int i = 1; i <= n; i++) {
            sums[i] = sums[i - 1] + nums[i - 1];
        }
        for (int i = 1; i <= n; i++) {
            int target = s + sums[i - 1];
            int bound = Arrays.binarySearch(sums, target);
            if (bound < 0) {
                bound = -bound - 1;
            }
            if (bound <= n) {
                ans = Math.min(ans, bound - (i - 1));
            }
        }
        return ans == Integer.MAX_VALUE ? 0 : ans;
    }
}


public static int binarySearch(int[] a,
               int key)

Searches the specified array of ints for the specified value using the binary search algorithm. The array must be sorted (as by the sort(int[]) method) prior to making this call. If it is not sorted, the results are undefined. If the array contains multiple elements with the specified value, there is no guarantee which one will be found.

Parameters:
a - the array to be searched
key - the value to be searched for

Returns:
index of the search key, if it is contained in the array; otherwise, (-(insertion point) - 1). The insertion point is defined as the point at which the key would be inserted into the array: the index of the first element greater than the key, or a.length if all elements in the array are less than the specified key. Note that this guarantees that the return value will be >= 0 if and only if the key is found.


⑵.binarySearch(object[ ], int fromIndex, int endIndex, object key);

如果要搜索的元素key在指定的范围内，则返回搜索键的索引；否则返回-1或者”-“(插入点)。

eg：

1.该搜索键在范围内，但不在数组中，由1开始计数；

2.该搜索键在范围内，且在数组中，由0开始计数；

3.该搜索键不在范围内，且小于范围内元素，由1开始计数；

4.该搜索键不在范围内，且大于范围内元素，返回-(endIndex + 1);（特列）



class Solution {
    public int minSubArrayLen(int s, int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return 0;
        }
        int ans = Integer.MAX_VALUE;
        int start = 0, end = 0;
        int sum = 0;
        while (end < n) {
            sum += nums[end];
            while (sum >= s) {
                ans = Math.min(ans, end - start + 1);
                sum -= nums[start];
                start++;
            }
            end++;
        }
        return ans == Integer.MAX_VALUE ? 0 : ans;
    }
}

 */