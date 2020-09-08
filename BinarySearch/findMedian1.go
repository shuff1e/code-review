package BinarySearch

import "github.com/shuff1e/code-review/util/math"

//给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
//
//请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
// 例如给定两个数组
// A: 1 3.1 4 9
// B: 1 2 3.1 4 5 6 7 8 9
// 返回中位数 4

// arr1 A的长度为m,arr2 B的长度为n
// 如果(m+n)%2==1，中位数就是找两个数组中排第(m+n)/2的数
// 如果(m+n)%2==0，中位数就是找两个数组中排第(m+n)/2和(m+n)/2+1的平均值

// 将问题转化为寻找两个排序数组中排第k位的元素
// 情况一：如果A[k/2-1]<=B[k/2-1]，那么A[k/2-1]前面最多有k/2-1+k/2-1=k-2个元素，加上A[k/2-1]一共k-1个元素
// 那么A[0]一直到A[k/2-1]都不可能是排第k位的元素
// 排除A[0]一直到A[k/2-1]，相当于求剩余数组中排第k-k/2的元素

// 情况二：如果B[k/2-1]<A[k/2-1]，同理

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1,length2 := len(nums1),len(nums2)
	totalLength := length1 + length2
	if totalLength % 2 == 1 {
		return float64(getKthElement(nums1,nums2,totalLength/2 + 1))
	} else {
		return (float64(getKthElement(nums1,nums2,totalLength/2)) +
			float64(getKthElement(nums1,nums2,totalLength/2 + 1)))/2
	}
}

func getKthElement(arr1 []int,arr2 []int,k int) int {
	index1,index2 := 0,0
	length1,length2 := len(arr1),len(arr2)
	for {
		// 边界情况
		if index1 == length1 {
			return arr2[index2+k-1]
		}
		if index2 == length2 {
			return arr1[index1+k-1]
		}
		if k == 1 {
			return math.Min(arr1[index1],arr2[index2])
		}
		// 正常情况
		half := k/2
		newIndex1 := math.Min(index1+half,length1) - 1
		newIndex2 := math.Min(index2+half,length2) - 1
		pivot1,pivot2 := arr1[newIndex1],arr2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 -index2 + 1)
			index2 = newIndex2 + 1
		}
	}
}

// 假设A[0]一直到A[i-1]，以及B[0]一直到B[j-1]为左部分
// A[i]一直到A[len(A)-1]，以及B[j]到B[len(B)-1]为右部分
// 假设len(A)=m，len(B)=n

// 则需要满足len(left) == len(right) 或者 len(left) == len(right+1)
// 且max(left)<=min(right)

// 即 i+j = m-i+n-j(m+n是偶数） 或者 i+j = m-i+n-j+1（m+n是奇数）
// i+j = (m+n+1)/2
// j =  (m+n+1)/2 - i
// 这里假设m<n，否则j可能出现负数

// 且A[i-1]<=B[j-1]，B[j-1]<A[i]
// 由于随着j减小，A增大，B减小
// A[i-1]<=B[j-1]可以推出B[j-1]<A[i]

// 因此我们可以对 i 在 [0, m] 的区间上进行二分搜索，
// 找到最大的满足 A[i-1] <= B[j] 的 i 值，就得到了划分的方法。
// 此时，划分前一部分元素中的最大值，以及划分后一部分元素中的最小值，才可能作为就是这两个数组的中位数。
//
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}
	m,n := len(nums1),len(nums2)
	left,right := 0,m
	median1,median2 := 0,0
	for left <= right {
		i := (left + right) / 2;
		j := (m + n + 1) / 2 - i;

		// nums_im1, nums_i, nums_jm1, nums_j 分别表示 nums1[i-1], nums1[i], nums2[j-1], nums2[j]
		nums_im1 := 0
		if i ==0 {
			nums_im1 = math.INTEGER_MIN_VALUE
		} else {
			nums_im1 = nums1[i-1]
		}
		nums_i := 0
		if i ==m {
			nums_i = math.INTEGER_MAX_VALUE
		} else {
			nums_i = nums1[i]
		}

		nums_jm1 := 0
		if j ==0 {
			nums_jm1 = math.INTEGER_MIN_VALUE
		} else {
			nums_jm1 = nums2[j-1]
		}
		nums_j := 0
		if j == n {
			nums_j = math.INTEGER_MAX_VALUE
		} else {
			nums_j = nums2[j]
		}

		if (nums_im1 <= nums_j) {
			median1 = math.Max(nums_im1, nums_jm1);
			median2 = math.Min(nums_i, nums_j);
			left = i + 1;
		} else {
			right = i - 1;
		}

	}
	if (m+n)%2 == 0 {
		return (float64(median1) + float64(median2)) / 2.0
	} else {
		return float64(median1)
	}
}


