package main

import "fmt"

/*

315. 计算右侧小于当前元素的个数
给定一个整数数组 nums，按要求返回一个新数组 counts。
数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例：

输入：nums = [5,2,6,1]
输出：[2,1,1,0]
解释：
5 的右侧有 2 个更小的元素 (2 和 1)
2 的右侧仅有 1 个更小的元素 (1)
6 的右侧有 1 个更小的元素 (1)
1 的右侧有 0 个更小的元素


提示：

0 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4

 */

// 5 2 6 1

// 5 2 6 1
// 1 0 1 0
// 2 1 1 0

// 2 5 -> 1 6

func main() {
	arr := []int{5}
	result := countSmaller(arr)
	fmt.Printf("%#v\n",result)
	fmt.Printf("%#v\n",arr)
	arr = []int{1}
	fmt.Println(binarySearch(arr,0,len(arr)-1,1))
}

func countSmaller(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := make([]int,len(nums))
	helpArr := make([]int,len(nums))

	help(nums,0,len(nums)-1,result,helpArr)
	return result
}

func help(arr []int,left,right int,result []int,helpArr []int) {
	if left == right {
		helpArr[left] = arr[left]
		return
	}
	mid := (left + right) /2
	help(arr,left,mid,result,helpArr)
	help(arr,mid + 1,right,result,helpArr)

	for i := left;i<=mid;i++ {
		result[i] += binarySearch(helpArr,mid+1,right,arr[i])
	}

	tempArr := make([]int,right-left + 1)
	helpIndex := 0
	p1,p2 := left,mid+1
	for p1 <= mid && p2 <= right {
		if helpArr[p1] < helpArr[p2] {
			tempArr[helpIndex] = helpArr[p1]
			p1 ++
		} else {
			tempArr[helpIndex] = helpArr[p2]
			p2 ++
		}
		helpIndex ++
	}
	for p1 <= mid {
		tempArr[helpIndex] = helpArr[p1]
		helpIndex ++
		p1 ++
	}
	for p2 <= right {
		tempArr[helpIndex] = helpArr[p2]
		helpIndex ++
		p2 ++
	}

	for i := 0;i<len(tempArr);i++ {
		helpArr[i+left] = tempArr[i]
	}
}

// 1 5 2 6 1

// 1 2  1 5 6
// 7 -> r为2
// 1 -> r为-1
// 2 -> r为0

func binarySearch(arr []int,left,right int,target int) int {
	start := left
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right - start + 1
}