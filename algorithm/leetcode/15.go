package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := make([][]int,0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i := 0;i<len(nums);i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i+1
		R := len(nums) - 1
		for L < R {
			if nums[i]+ nums[L] + nums[R] == 0 {
				result = append(result,[]int{nums[i],nums[L],nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L ++
				}
				for L < R && nums[R] == nums[R-1] {
					R --
				}
				L ++
				R --
			} else if nums[i]+ nums[L] + nums[R] > 0 {
				R --
			} else {
					L ++
			}
		}
	}
	return result
}

func main() {
	input := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(input))
}
