package main

import "fmt"

/*
215. 数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:

输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
说明:

你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 */

// 剑指 40
func main() {
	arr := []int{3,2,3,1,2,4,5,5,6}
	k := 4
	arr = []int{3,2,1,5,6,4}
	k = 2
	fmt.Println(findKthLargest2(arr,k))
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) || len(nums) == 0 {
		return -1
	}
	for {
		index := partition(nums)
		if index + 1 == k {
			return nums[index]
		} else if index + 1 < k {
			nums = nums[index+1:]
			k = k - (index + 1)
		} else {
			nums = nums[:index]
		}
	}
	return -1
}

func partition(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	pivot := arr[0]
	mark := 0
	for i := 0;i<len(arr);i++ {
		if arr[i] > pivot {
			mark ++
			swap(arr,i,mark)
		}
	}
	swap(arr,0,mark)
	return mark
}

func swap(arr []int,x,y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}


func findKthLargest2(nums []int, k int) int {
	if k > len(nums) || len(nums) == 0 {
		return -1
	}
	heap := NewMinHeap(k)
	for i := 0;i<k;i++ {
		heap.Add(nums[i])
	}
	for i:=k;i<len(nums);i++ {
		if heap.Peek() < nums[i] {
			heap.RemoveFirst()
			heap.Add(nums[i])
		}
	}
	return heap.Peek()
}

type minHeap struct {
	data []int
	size int
}

func NewMinHeap(k int) *minHeap {
	return &minHeap{
		data: make([]int,k),
		size: 0,
	}
}

func (heap *minHeap) Add(x int) {
	heap.data[heap.size] = x
	heap.size ++
	heap.BubbleUp()
}

func (heap *minHeap) RemoveFirst() int {
	result := heap.data[0]
	heap.swap(0,heap.size-1)
	heap.size --
	heap.BubbleDown(0)
	return result
}

func (heap *minHeap) Peek() int {
	return heap.data[0]
}

func (heap *minHeap) BubbleDown(index int) {
	for index < heap.size {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < heap.size && heap.data[left] < heap.data[index] {
			smallest = left
		}
		if right < heap.size && heap.data[right] < heap.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		heap.swap(index,smallest)
		index = smallest
	}
}

func (heap *minHeap) BubbleUp() {
	index := heap.size - 1
	for index >= 0 {
		parent := (index-1)/2
		if heap.data[parent] > heap.data[index] {
			heap.swap(index,parent)
			index = parent
		} else {
			break
		}
	}
}

func (heap *minHeap) swap(x,y int) {
	temp := heap.data[x]
	heap.data[x] = heap.data[y]
	heap.data[y] = temp
}