package main

import "fmt"

/*
31. 下一个排列
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

 */

// A：
// 1 3 5 4 2
// 下一个数字是
// 1 4 2 3 5

// 首先找的5的位置，没有比5，4，2更大的序列了
// 然后5,4,2中第一个比3大的是4
// 那说明3的位置应该换为4了

// 然后变成
// 1,4,5,3,2
// 再 reverse 5,3,2

// 1 2 3 4 5
// 下一个是
// 1 2 3 5 4

func main() {
	arr := []int{2,3,1}
	nextPermutation(arr)
	fmt.Printf("%#v\n",arr)
}

func nextPermutation(nums []int)  {
	if len(nums) <= 1 {
		return
	}
	index := 0
	for i := len(nums) - 1;i>0;i-- {
		if nums[i] > nums[i - 1] {
			index = i
			break
		}
	}
	if index > 0 {
		for i := len(nums) - 1;i>= index;i-- {
			if nums[i] > nums[index-1] {
				swap(nums,i,index-1)
				break
			}
		}
	}
	reverse(nums,index,len(nums)-1)
}

func swap(nums []int,i,j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int,left ,right int) {
	for left < right {
		swap(nums,left,right)
		left++
		right--
	}
}
