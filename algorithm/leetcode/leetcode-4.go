package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		temp := nums2
		nums2 = nums1
		nums1 = temp
	}

	imin := 0
	imax := len(nums1)
	m = len(nums1)
	n = len(nums2)
	halfLen := (m+n+1)/2

	for imin <= imax {
		i := (imin+imax)/2
		j := halfLen - i

		if (i < imax && nums2[j-1] > nums1[i]) {
			imin = i + 1
		} else if (i>imin && nums1[i-1] > nums2[j]) {
			imax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j ==0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = myMax(nums1[i-1],nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if (j == n) {
				minRight = nums1[i]
			} else {
				minRight = myMin(nums2[j],nums1[i])
			}
			return float64(maxLeft + minRight)/2.0
		}
	}
	return 0.0
}

func myMax(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func myMin(x,y int) int {
	if x < y {
		return x
	}
	return y
}
