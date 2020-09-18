package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	result := [][]int{}

	for i := 0;i<len(nums)-3;i++ {

		if i >0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i+1;j<len(nums)-2;j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			L := j+1
			R := len(nums) -1
			for L < R {
				if nums[i] + nums[j] + nums[L] + nums[R] == target {
					result = append(result,[]int{nums[i],nums[j],nums[L],nums[R]})

					for L<R && nums[L] == nums[L+1] {
						L ++
					}
					for L<R && nums[R] == nums[R-1] {
						R --
					}
					L ++
					R --

				} else if nums[i] + nums[j] + nums[L] + nums[R] < target {
					for L < R &&nums[L] == nums[L+1] {
						L ++
					}
					L ++
				} else {
					for L < R &&nums[R] == nums[R-1] {
						R --
					}
					R --
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums,target)
	for _,v := range result {
		fmt.Println(v)
	}
}