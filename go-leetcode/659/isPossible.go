package main

import (
	"container/heap"
	"fmt"
)

/*

659. 分割数组为连续子序列
给你一个按升序排序的整数数组 num（可能包含重复数字），
请你将它们分割成一个或多个子序列，其中每个子序列都由连续整数组成且长度至少为 3 。

如果可以完成上述分割，则返回 true ；否则，返回 false 。



示例 1：

输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3
3, 4, 5


示例 2：

输入: [1,2,3,3,4,4,5,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3, 4, 5
3, 4, 5


示例 3：

输入: [1,2,3,4,4,5]
输出: False


提示：

输入的数组长度范围为 [1, 10000]

 */

func main() {
	m := &minHeap{}
	heap.Push(m,10)
	heap.Push(m,1)
	heap.Push(m,5)
	heap.Push(m,7)
	result := heap.Pop(m)
	fmt.Println(result)
	result = m.Peek()
	fmt.Println(result)

	// 以4结尾的有1,2,3,4
	// 以及4
	// 5会添加到4
	ok := isPossible2([]int{1,2,3,4,4,5})
	fmt.Println(ok)
	// 1,2,3,3,4,4,5,5
	// 1,2,3,4
	// 3,4,5
}

// 由于需要将数组分割成一个或多个由连续整数组成的子序列，因此只要知道子序列的最后一个数字和子序列的长度，就能确定子序列。
//
// 当 x 在数组中时，如果存在一个子序列以 x−1 结尾，长度为 k，则可以将 x 加入该子序列中，得到长度为 k+1 的子序列。如果不存在以 x−1 结尾的子序列，则必须新建一个只包含 x 的子序列，长度为 1。
//
// 当 x 在数组中时，如果存在多个子序列以 x−1 结尾，应该将 x 加入其中的哪一个子序列？由于题目要求每个子序列的长度至少为 3，显然应该让最短的子序列尽可能长，因此应该将 x 加入其中最短的子序列。
//
// 基于上述分析，可以使用哈希表和最小堆进行实现。
//
// 哈希表的键为子序列的最后一个数字，值为最小堆，用于存储所有的子序列长度，最小堆满足堆顶的元素是最小的，因此堆顶的元素即为最小的子序列长度。
//
// 遍历数组，当遍历到元素 x 时，可以得到一个以 x 结尾的子序列。
func isPossible(nums []int) bool {
	dict := make(map[int]*minHeap)
	for _,x := range nums {
		// 初始化x
		if _,ok := dict[x];!ok {
			temp := &minHeap{}
			dict[x] = temp
		}

		if _,ok := dict[x-1];ok {
			prev := heap.Pop(dict[x-1])
			if dict[x-1].Len() == 0 {
				delete(dict,x-1)
			}
			heap.Push(dict[x],prev.(int)+1)
		} else {
			heap.Push(dict[x],1)
		}
	}

	for _,v := range dict {
		if v.Peek() < 3 {
			return false
		}
	}
	return true
}

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i,j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i,j int) {
	m[i],m[j] = m[j],m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m,x.(int))
}

func (m *minHeap) Pop() interface{} {
	result := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return result
}

func (m *minHeap) Peek() int {
	return (*m)[0]
}

// 从方法一可以看到，对于数组中的元素 x，如果存在一个子序列以 x−1 结尾，则可以将 x 加入该子序列中。将 x 加入已有的子序列总是比新建一个只包含 x 的子序列更优，因为前者可以将一个已有的子序列的长度增加 1，而后者新建一个长度为 1 的子序列，而题目要求分割成的子序列的长度都不小于 3，因此应该尽量避免新建短的子序列。
//
// 基于此，可以通过贪心的方法判断是否可以完成分割。
//
// 使用两个哈希表，第一个哈希表存储数组中的每个数字的剩余次数，第二个哈希表存储数组中的每个数字作为结尾的子序列的数量。
func isPossible2(nums []int) bool {
	countMap := make(map[int]int)
	endMap := make(map[int]int)

	for _,x := range nums {
		countMap[x] ++
	}

	for _,x := range nums {
		count := countMap[x]
		if count == 0 {
			continue
		}
		prev := endMap[x-1]
		if prev > 0 {
			countMap[x] --
			endMap[x-1] --
			endMap[x] ++
		} else {
			count1 := countMap[x+1]
			count2 := countMap[x+2]
			if count1 > 0 && count2 > 0 {
				countMap[x] --
				countMap[x+1] --
				countMap[x+2] --
				endMap[x+2] ++
			} else {
				return false
			}
		}
	}
	return true
}