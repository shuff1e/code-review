package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return -1
	}

	sort.Ints(nums)

	ans := nums[0] + nums[1] + nums[2]
	for i := 0;i<len(nums);i++ {
		L := i + 1
		R := len(nums) -1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == target {
				return sum
			}
			if abs(target,ans) > abs(target,sum) {
				ans = sum
			}

			if sum > target {
				R --
			} else if sum < target {
					L ++
			}

		}
	}

	return ans
}

func abs(x,y int) int{
	if x > y {
		return x -y
	} else {
		return y -x
	}
}

func min16(x,y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	result := threeSumClosest(nums,target)
	println(result)
}