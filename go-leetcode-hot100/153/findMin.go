package main

import "fmt"

/*
153. 寻找旋转排序数组中的最小值
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

你可以假设数组中不存在重复元素。

示例 1:

输入: [3,4,5,1,2]
输出: 1
示例 2:

输入: [4,5,6,7,0,1,2]
输出: 0
 */

// arr[mid] > arr[left],最小值在右边
// arr[mid] < arr[left],最小值在左边

// 中间元素mid
// 如果mid>最右边的元素，说明最小值肯定在mid右边
// 如果mid<最右边的元素，说明最小值在mid左边
// 如果mid=最右边的元素，最小值可能在mid右边，也可能在mid左边

func main() {
	arr := []int{4,5,6,7,0,1,2}
	arr = []int{3,4,5,1,2}
	arr = []int{3,1,2}
	fmt.Println(findMin(arr))
}

func findMin(nums []int) int {
	left,right := 0,len(nums)-1
	for left < right {
		mid := (left + right)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

/*

在这个改进版本的二分搜索算法中，我们需要找到这个点。下面是关于变化点的特点：
所有变化点左侧元素 > 数组第一个元素
所有变化点右侧元素 < 数组第一个元素

算法
找到数组的中间元素 mid。
如果中间元素 > 数组第一个元素，我们需要在 mid 右边搜索变化点。
如果中间元素 < 数组第一个元素，我们需要在 mid 左边搜索变化点。


当我们找到变化点时停止搜索，当以下条件满足任意一个即可：
nums[mid] > nums[mid + 1]，因此 mid+1 是最小值。
nums[mid - 1] > nums[mid]，因此 mid 是最小值。

class Solution {
  public int findMin(int[] nums) {
    // If the list has just one element then return that element.
    if (nums.length == 1) {
      return nums[0];
    }

    // initializing left and right pointers.
    int left = 0, right = nums.length - 1;

    // if the last element is greater than the first element then there is no rotation.
    // e.g. 1 < 2 < 3 < 4 < 5 < 7. Already sorted array.
    // Hence the smallest element is first element. A[0]
    if (nums[right] > nums[0]) {
      return nums[0];
    }

    // Binary search way
    while (right >= left) {
      // Find the mid element
      int mid = left + (right - left) / 2;

      // if the mid element is greater than its next element then mid+1 element is the smallest
      // This point would be the point of change. From higher to lower value.
      if (nums[mid] > nums[mid + 1]) {
        return nums[mid + 1];
      }

      // if the mid element is lesser than its previous element then mid element is the smallest
      if (nums[mid - 1] > nums[mid]) {
        return nums[mid];
      }

      // if the mid elements value is greater than the 0th element this means
      // the least value is still somewhere to the right as we are still dealing with elements
      // greater than nums[0]
      if (nums[mid] > nums[0]) {
        left = mid + 1;
      } else {
        // if nums[0] is greater than the mid value then this means the smallest value is somewhere to
        // the left
        right = mid - 1;
      }
    }
    return -1;
  }
}

 */