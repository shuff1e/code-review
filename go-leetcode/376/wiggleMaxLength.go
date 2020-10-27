package main

import (
	"fmt"
)

/*

376. 摆动序列
如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。

例如， [1,7,4,9,2,5] 是一个摆动序列，因为差值 (6,-3,5,-7,3) 是正负交替出现的。相反, [1,4,7,2,5] 和 [1,7,4,5,5] 不是摆动序列，第一个序列是因为它的前两个差值都是正数，第二个序列是因为它的最后一个差值为零。

给定一个整数序列，返回作为摆动序列的最长子序列的长度。 通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。

示例 1:

输入: [1,7,4,9,2,5]
输出: 6
解释: 整个序列均为摆动序列。
示例 2:

输入: [1,17,5,10,13,15,10,5,16,8]
输出: 7
解释: 这个序列包含几个长度为 7 摆动序列，其中一个可为[1,17,10,13,10,16,8]。
示例 3:

输入: [1,2,3,4,5,6,7,8,9]
输出: 2
进阶:
你能否用 O(n) 时间复杂度完成此题?

*/

func main() {

	// 1 17 5 10
	// 13 15 10
	// 10 5 16 8 19
	arr := []int{1,17, 25,10,10, 15, 10,5,1}
	arr = []int{1,-1}
	fmt.Println(wiggleMaxLength2(arr))
}

func wiggleMaxLengthBad(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	return 1 + Max(helpBad(nums,0,true),helpBad(nums,0,false))
}

func helpBad(arr []int,index int,isUp bool) int {
	count := 0
	for i := index+1;i<len(arr);i++ {
		if (isUp && arr[i] < arr[index]) || (!isUp && arr[i] > arr[index]) {
			count = Max(count,1 + helpBad(arr,i,!isUp))
		}
	}
	return count
}

func wiggleMaxLengthBetter(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	up := make([]int,len(nums))
	down := make([]int,len(nums))

	for i := 1;i<len(nums);i++ {
		for j := 0;j<i;j++ {
			if nums[i] > nums[j] {
				up[i] = Max(up[i],down[j]+1)
			} else if nums[i] < nums[j] {
				down[i] = Max(down[i],up[j] + 1)
			}
		}
	}
	return 1 + Max(up[len(nums)-1],down[len(nums)-1])
}

func wiggleMaxLengthBest(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	up := make([]int,len(nums))
	down := make([]int,len(nums))

	up[0],down[0] = 1,1

	for i := 1;i<len(nums);i++ {
		if nums[i] > nums[i-1] {
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = up[i-1] + 1
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return Max(up[len(nums)-1],down[len(nums)-1])
}

// 连续子数组
func wiggleMaxLength2(nums []int) int {
	const UP = 1
	const DOWN = 0
	if len(nums) <= 1 {
		return len(nums)
	}

	status := -1
	ans := 0
	curr := 0

	for i := 1 ;i<len(nums);i++ {
		if nums[i] > nums[i-1] {
			if status == DOWN {
				curr ++
				ans = Max(ans,curr)
				status = UP
			} else {
				curr = 2
				ans = Max(ans,curr)
				status = UP
			}
		} else if nums[i] < nums[i-1] {
			if status == UP {
				curr ++
				ans = Max(ans,curr)
				status = DOWN
			} else {
				curr = 2
				ans = Max(ans,curr)
				status = DOWN
			}
		} else {
			curr = 0
			status = -1
		}
	}
	return ans
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*

public class Solution {
    public int wiggleMaxLength(int[] nums) {
        if (nums.length < 2)
            return nums.length;
        int down = 1, up = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] > nums[i - 1])
                up = down + 1;
            else if (nums[i] < nums[i - 1])
                down = up + 1;
        }
        return Math.max(down, up);
    }
}

 */