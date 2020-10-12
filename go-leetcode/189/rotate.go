package main

import "fmt"

/*
189. 旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
 */

// 类似字符串旋转
// 7,6,5,4,3,2,1

func main() {
	arr := []int{1,2,3,4,5,6,7}
	rotate1(arr,6)
	fmt.Printf("%#v\n",arr)
}

func rotate(nums []int, k int)  {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	left,right := 0,len(arr)-1
	for left < right {
		temp := arr[left]
		arr[left] = arr[right]
		arr[right] = temp
		left ++
		right --
	}
}

// 1,2,3,4,5,6,7
// 7,1,2,3,4,5,6,

// 5,6,7,1,2,3,4

// 暴力，每次旋转一个数
func rotate1(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	for i := 0;i<k;i++ {
		previous := nums[len(nums)-1]
		for j := 0;j<len(nums);j++ {
			temp := nums[j]
			nums[j] = previous
			previous = temp
		}
	}
}

func rotate2(nums []int, k int) {
	arr := make([]int,len(nums))
	for i := 0;i<len(nums);i++ {
		arr[(i+k)%len(nums)] = nums[i]
	}
	for i := 0;i<len(nums);i++ {
		nums[i] = arr[i]
	}
}

//把元素看做同学，把下标看做座位，大家换座位。
//
//第一个同学离开座位去第k+1个座位，第k+1个座位的同学被挤出去了，他就去坐他后k个座位，如此反复。
//但是会出现一种情况，就是其中一个同学被挤开之后，坐到了第一个同学的位置（空位置，没人被挤出来），但是此时还有人没有调换位置，这样就顺着让第二个同学换位置。
//
//那么什么时候就可以保证每个同学都换完了呢？n个同学，换n次，所以用一个count来计数即可。

func rotate3(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	count := 0
	for start :=0;count < len(nums);start ++ {
		curr := start
		prev := nums[start]
		for  {
			next := (curr+k)%len(nums)
			temp := nums[next]

			nums[next] = prev
			curr = next
			prev = temp

			count ++
			if curr == start {
				break
			}
		}
	}
}