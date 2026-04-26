package main

import "fmt"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例1
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//
//示例2
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5

// A：如果两个数组的个数和n为奇数
// 中位数为下标为n/2
// n为偶数，中位数下标为n/2-1,n/2

//可以归并，然后找到第n/2个数

func findMedianSortedArraysBad(nums1 []int, nums2 []int) float64 {
	arr := merge(nums1,nums2)
	if len(arr) == 0 {
		return 0
	}
	if len(arr) % 2 ==1 {
		return float64(arr[len(arr)/2])
	} else {
		return (float64(arr[len(arr)/2-1]) + float64(arr[len(arr)/2]))/2
	}
}

func findIndex(k int,arr []int) int {
	start,end := 0,len(arr)-1
	for {
		index := partition(start,end,arr)
		if index == k {
			return arr[k]
		} else if index < k {
			start = index + 1
		} else {
			end = index - 1
		}
	}
	return -1
}

func partition(start ,end int,arr []int) int {
	pivot := arr[start]
	mark := start
	for i := start;i<=end;i++ {
		if arr[i] < pivot {
			mark ++
			swap(i,mark,arr)
		}
	}
	swap(start,mark,arr)
	return mark
}

func swap(i,j int,arr []int) {
	if i == j {
		return
	}
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func merge(nums1 []int, nums2 []int) []int {
	arr := make([]int,len(nums1) + len(nums2))
	index := len(arr)-1
	index1,index2 := len(nums1)-1,len(nums2)-1

	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			arr[index] = nums1[index1]
			index --
			index1 --
		} else {
			arr[index] = nums2[index2]
			index --
			index2 --
		}
	}
	for index1 >= 0 {
		arr[index] = nums1[index1]
		index --
		index1 --
	}
	for index2 >= 0 {
		arr[index] = nums2[index2]
		index --
		index2 --
	}
	return arr
}

func main() {
	arr := findMedianSortedArrays([]int{1},[]int{2,3,4,5,6,7})
	fmt.Println(arr)
	//fmt.Printf("%#v\n",arr)
	//arr = []int{3,1,4,5,5,6,2}
	//fmt.Println(findIndex(1,arr))
}

// A：类似25匹马，5个赛道赛跑的问题，看每次能排除哪些元素
// 二分法
//
// 假设数组总的长度为m,n，则中位数排第(m+n+1)/2的数字，或者是(m+n+1)/2+1，假设为k = (m+n+1)/2
//
// 两个数组中，均以所以为k/2-1的数为基准

// 如果 A[k/2-1] <= B[k/2-1]，就算A[0]大于B[k/2-2]，
// A[k/2-1]左边加上自己，也只有k/2+k/2-1个数字，因为可以排除A[k/2-1]左边的部分
// B[k/2-1]左边的部分不可能排除，因为B[0]可能大于A[len(A)-1]

// 如果k/2-1越界，就以最后一个数字为基准

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1,p2:=0,0
	count := len(nums1) - p1 + len(nums2) - p2
	k := (count+1)/2

	for p1 < len(nums1) && p2 < len(nums2) {
		if k == 1 {
			if (len(nums1) + len(nums2)) % 2 == 1 {
				return float64(Min(nums1[p1],nums2[p2]))
			} else {
				if p1+1 < len(nums1) && nums1[p1+1] < nums2[p2] {
					return (float64(nums1[p1]) + float64(nums1[p1+1]))/2
				} else if p2+1 < len(nums2) && nums2[p2+1] < nums1[p1] {
					return (float64(nums2[p2]) + float64(nums2[p2+1]))/2
				} else {
					return (float64(nums1[p1]) + float64(nums2[p2]))/2
				}
			}
		}
		if p1 + k/2 - 1 < len(nums1) && p2 + k/2 -1 < len(nums2) {
			if nums1[p1 + k/2 - 1] <= nums2[p2 + k/2 - 1] {
				p1 = p1 + k/2
				count -= k/2
				k -= k/2
			} else {
				p2 = p2 + k/2
				count -= k/2
				k -= k/2
			}
		} else if p1 + k/2 - 1 < len(nums1) {
			if nums1[p1 + k/2 - 1] <= nums2[len(nums2)-1] {
				p1 = p1 + k/2
				count -= k/2
				k -= k/2
			} else {
				count -= len(nums2)-p2
				k -= len(nums2)-p2
				p2 = len(nums2)
			}
		} else if p2 + k/2 -1 < len(nums2) {
			if nums1[len(nums1)-1] <= nums2[p2 + k/2 -1] {
				count -= len(nums1) - p1
				k -= len(nums1) - p1
				p1 = len(nums1)
			} else {
				p2 = p2 + k/2
				count -= k/2
				k -= k/2
			}
		} else {
			if nums1[len(nums1)-1] <= nums2[len(nums2)-1] {
				count -= len(nums1) - p1
				k -= len(nums1) - p1
				p1 = len(nums1)
			} else {
				count -= len(nums2)-p2
				k -= len(nums2)-p2
				p2 = len(nums2)
			}
		}
	}

	if p1 == len(nums1) {
		if (len(nums1) + len(nums2)) % 2 == 1 {
			return float64(nums2[p2 + k - 1])
		} else {
			return (float64(nums2[p2 + k - 1]) + float64(nums2[p2 + k]))/2
		}
	}

	if p2 == len(nums2) {
		if (len(nums1) + len(nums2)) % 2 == 1 {
			return float64(nums1[p1 + k - 1])
		} else {
			return (float64(nums1[p1 + k -1]) + float64(nums1[p1 + k]))/2
		}
	}
	return 0
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

/*

// arr1,arr2
// 两个数组中整体排名K的

// arr1 1 2 3 4 5
// arr2 6 7 8 9 10
// K 为 5，两边排名K/2=5/2=2的
// 2小于7，所以1，2不可能是排名第5的

// arr1 3 4 5
// arr2 6 7 8 9 10
// K 为 3,3/2=1
// 3 < 6，3不可能是排名第3的

// arr1 4 5
// arr2 6 7 8 9 10
// K为2 2/2=1
// 4 < 6，4不可能是排第二的

// arr1 5
// arr2 6 7 8 9 10
// K为1
// 返回5

// arr1 3
// arr2 1 2 5 6 7 8
// K = 5

func main() {
	// 1,1,2,3,3,4,5,6,8,9,10,12
	arr1 := []int{}
	arr2 := []int{1,3,8,9,10,12}
	result := findMedianSortedArrays(arr1,arr2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	K := len(nums1) + len(nums2)
	if K % 2 == 0 {
		result1 := help(nums1,nums2,0,0,K/2)
		result2 := help(nums1,nums2,0,0,K/2+1)
		return (float64(result1) + float64(result2))/2
	}
	// 8 排第四,第五
	// 9 排第五
	result := help(nums1,nums2,0,0,(K+1)/2)
	return float64(result)
}

func help(nums1,nums2 []int,p1,p2 int,K int) int {
	for p1 < len(nums1) && p2 < len(nums2) && K > 1 {
		if p1 + K/2 - 1 < len(nums1) && p2 + K/2 - 1 < len(nums2) {
			if nums1[p1 + K/2 -1] <= nums2[p2 + K/2 -1] {
				p1 = p1 + K/2
				K -= K/2
			} else {
				p2 = p2 + K/2
				K -= K/2
			}
		} else if p1 + K/2 - 1 < len(nums1) {
			if nums1[p1 + K/2 -1] <= nums2[len(nums2)-1] {
				p1 = p1 + K/2
				K -= K/2
			} else {
				K = K - (len(nums2)-p2)
				p2 = len(nums2)
			}
		} else if p2 + K/2 - 1 < len(nums2) {
			if nums2[p2 + K/2 - 1] <= nums1[len(nums1)-1] {
				p2 = p2 + K/2
				K -= K/2
			} else {
				K = K -(len(nums1)-p1)
				p1 = len(nums1)
			}
		} else {
			if nums1[len(nums1)-1] <= nums2[len(nums2)-1] {
				K = K - (len(nums1)-p1)
				p1 = len(nums1)
			} else {
				K = K - (len(nums2)-p2)
				p2 = len(nums2)
			}
		}
	}

	if K == 1 {
		if p1 < len(nums1) && p2 < len(nums2) {
			return Min(nums1[p1],nums2[p2])
		} else if p1 < len(nums1) {
			return nums1[p1]
		} else {
			return nums2[p2]
		}
	}

	if p1 < len(nums1) {
		return nums1[p1 + K - 1]
	}

	return nums2[p2 + K - 1]
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

 */