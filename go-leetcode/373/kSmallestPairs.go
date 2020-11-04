package main

import "container/heap"

/*

373. 查找和最小的K对数字
给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。

找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。

示例 1:

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:

输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
示例 3:

输入: nums1 = [1,2], nums2 = [3], k = 3
输出: [1,3],[2,3]
解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

 */

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	k = Min(k,len(nums1)*len(nums2))
	result := [][]int{}

	// ptrs[i] = j 表示 nums1[i]此时的候选元素是nums2[j]
	ptrs := make([]int,len(nums1))

	for k > 0 {
		k--
		minIndex := -1
		min := 0x7fffffff
		for i := 0;i<len(nums1);i++ {
			if ptrs[i] < len(nums2) && nums1[i] + nums2[ptrs[i]] < min {
				minIndex = i
				min = nums1[i] + nums2[ptrs[i]]
			}
		}
		result = append(result,[]int{nums1[minIndex],nums2[ptrs[minIndex]]})
		ptrs[minIndex] ++
	}
	return result
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

type pair struct {
	// nums1的元素值
	v int
	// v对应的在nums2中的候选值的索引
	index int
}

func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {
	k = Min(k,len(nums1)*len(nums2))
	if k == 0 {
		return nil
	}
	result := [][]int{}

	h := &minHeap{
		data: make([]pair,0),
		nums2: nums2,
	}

	for i := 0;i<len(nums1);i++ {
		heap.Push(h,pair{nums1[i],0})
	}

	for k > 0 {
		k--
		result = append(result,[]int{h.Peek().v, nums2[h.Peek().index]})
		tempPair := heap.Pop(h).(pair)
		tempPair.index ++
		if tempPair.index < len(nums2) {
			heap.Push(h,tempPair)
		}
	}
	return result
}

type minHeap struct {
	data []pair
	nums2 []int
}

func (h minHeap) Len() int {
	return len(h.data)
}

func (h minHeap) Less(i,j int) bool {
	return h.data[i].v + h.nums2[h.data[i].index] < h.data[j].v + h.nums2[h.data[j].index]
}

func (h minHeap) Swap(i,j int) {
	h.data[i],h.data[j] = h.data[j],h.data[i]
}

func (h *minHeap) Push(x interface{}) {
	h.data = append(h.data,x.(pair))
}

func (h *minHeap) Pop() interface{} {
	result := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return result
}

func (h *minHeap) Peek() pair {
	return h.data[0]
}