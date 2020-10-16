package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

/*
347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]

提示：

你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
你可以按任意顺序返回答案。
 */

func main() {
	arr := []int{1,1,1,2,2,3}
	arr = []int{}
	k := 2
	fmt.Println(topKFrequent(arr,k))
	fmt.Println(topKFrequent2(arr,k))
}

func topKFrequent(nums []int, k int) []int {
	frequency := map[int]int{}
	for i := 0;i<len(nums);i++ {
		frequency[nums[i]] ++
	}
	h := &minHeap{}
	heap.Init(h)

	for key,value := range frequency {
		heap.Push(h,[2]int{key,value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	result := make([]int,k)
	for i := k-1;i >= 0;i-- {
		result[i] = heap.Pop(h).([2]int)[0]
	}
	return result
}

type minHeap [][2]int

func (h minHeap) Len() int {return len(h)}

func (h minHeap) Less(i,j int) bool {return h[i][1] < h[j][1] }

func (h minHeap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h,x.([2]int))
}

func (h *minHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}



func topKFrequent2(nums []int, k int) []int {
	frequency := map[int]int{}
	for i := 0;i<len(nums);i++ {
		frequency[nums[i]] ++
	}
	arr := [][2]int{}
	for key,value := range frequency {
		arr = append(arr,[2]int{key,value})
	}

	start,end := 0,len(arr)-1
	index := partition(arr,start,end)
	for index != k - 1 {
		if index > (k-1) {
			end = index - 1
		}
		if index < k-1 {
			start = index + 1
		}
		index = partition(arr,start,end)
	}

	result := []int{}
	for i := 0;i<k;i++ {
		result = append(result,arr[i][0])
	}
	return result
}

func partition(arr [][2]int,start,end int) int {
	if start >= end {
		return start
	}
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int() % (end - start + 1) + start
	arr[picked],arr[start] = arr[start],arr[picked]

	pivot := arr[start]
	mark := start
	for i := start;i<=end;i++ {
		if arr[i][1] > pivot[1] {
			mark ++
			arr[i],arr[mark] = arr[mark],arr[i]
		}
	}
	arr[start],arr[mark] = arr[mark],arr[start]
	return mark
}