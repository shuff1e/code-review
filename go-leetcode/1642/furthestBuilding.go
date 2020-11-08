package main

import "fmt"

/*

1642. 可以到达的最远建筑
给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。

你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。

当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始 ）时：

如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块
如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块
如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标 从 0 开始 ）。


示例 1：


输入：heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
输出：4
解释：从建筑物 0 出发，你可以按此方案完成旅程：
- 不使用砖块或梯子到达建筑物 1 ，因为 4 >= 2
- 使用 5 个砖块到达建筑物 2 。你必须使用砖块或梯子，因为 2 < 7
- 不使用砖块或梯子到达建筑物 3 ，因为 7 >= 6
- 使用唯一的梯子到达建筑物 4 。你必须使用砖块或梯子，因为 6 < 9
无法越过建筑物 4 ，因为没有更多砖块或梯子。
示例 2：

输入：heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
输出：7
示例 3：

输入：heights = [14,3,19,3], bricks = 17, ladders = 0
输出：3


提示：

1 <= heights.length <= 105
1 <= heights[i] <= 106
0 <= bricks <= 109
0 <= ladders <= heights.length

 */

func main() {
	heights := []int{1,5,1,2,3,4,10000}
	bricks := 4
	ladders := 1
	fmt.Println(furthestBuilding(heights,bricks,ladders))
}

/*

方法一：优先队列 + 贪心
思路与算法

在移动的过程中，我们会需要若干次需要使用砖块或者梯子的情况。
假设当前我们需要移动到下一建筑物，但必须使用 1 架梯子或者 Δh 个砖块，那么我们如何抉择是使用梯子还是砖块呢？

我们可以用贪心的思路来想这个问题。「梯子」相当于一次性的无限量砖块，那么我们一定是把梯子用在刀刃上。

也就是说，如果我们有 l 架梯子，那么我们会在 Δh 最大的那 l 次使用梯子，而在剩余的情况下使用砖块。

这样一来，我们就可以得到正确的算法了：我们使用优先队列实时维护不超过 l 个最大的 Δh，这些就是使用梯子的地方。
对于剩余的 Δh，我们需要使用砖块，因此需要对它们进行累加，
如果某一时刻这个累加值超过了砖块的数目 b，那么我们就再也无法移动了。

 */

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := NewMinHeap()
	diffSum := 0

	for i:=1;i<len(heights);i++ {
		diff := heights[i] - heights[i-1]
		if diff <= 0 {
			continue
		}

		h.AddLast(diff)
		if h.size > ladders {
			x := h.removeFirst()
			diffSum += x
			if diffSum > bricks {
				return i-1
			}
		}
	}
	return len(heights) - 1
}

type minHeap struct {
	data []int
	size int
}

func NewMinHeap() *minHeap {
	return &minHeap{
		data: make([]int,0),
		size: 0,
	}
}
func (h *minHeap) AddLast(x int) {
	h.size ++
	if len(h.data) < h.size {
		h.data = append(h.data,make([]int,h.size-len(h.data))...)
	}
	h.data[h.size-1] = x
	h.bubbleUp()
}

func (h *minHeap) removeFirst() int {
	result := h.data[0]
	h.data[0],h.data[h.size-1] = h.data[h.size-1],h.data[0]
	h.size --
	h.bubbleDown(0)
	return result
}

func (h *minHeap) bubbleUp() {
	index := h.size - 1
	for index > 0 {
		parent := (index - 1)/2
		if h.data[parent] > h.data[index] {
			h.data[parent],h.data[index] = h.data[index],h.data[parent]
		} else {
			break
		}
		index = parent
	}
}

func (h *minHeap) bubbleDown(index int) {
	for index < h.size {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < h.size && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < h.size && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index],h.data[smallest] = h.data[smallest],h.data[index]
		index = smallest
	}
}


/*

func furthestBuilding(heights []int, bricks int, ladders int) int {
	//i := 1
	//for ;i<len(heights);i++ {
	//	if heights[i] <= heights[i-1] {
	//		continue
	//	} else {
	//		if bricks >= (heights[i] - heights[i-1]) {
	//			bricks -= (heights[i] - heights[i-1])
	//		} else {
	//			if ladders > 0 {
	//				ladders --
	//			} else {
	//				break
	//			}
	//		}
	//	}
	//}
	//return i-1
	//memo := make([][][]int,len(heights))
	//
	//for i := 0;i<len(memo);i++ {
	//	memo[i] = make([][]int,bricks+1)
	//	for j := 0;j<len(memo[i]);j++ {
	//		memo[i][j] = make([]int,ladders+1)
	//		for k := 0;k<len(memo[i][j]);k ++ {
	//			memo[i][j][k] = -1
	//		}
	//	}
	//}

	return help(heights,1,bricks,ladders)
}

func help(heights []int,index int,bricks,ladders int) int {
	if index == len(heights) {
		return index - 1
	}

	//if memo[index][bricks][ladders] != -1 {
	//	return memo[index][bricks][ladders]
	//}

	if heights[index] <= heights[index - 1] {
		return help(heights,index+1,bricks,ladders)
	}

	diff := heights[index] - heights[index - 1]

	result := index - 1
	if diff <= bricks {
		result = help(heights,index+1,bricks - diff,ladders)
	}

	if ladders > 0 {
		result = Max(result,
			help(heights,index+1,bricks,ladders-1))
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

 */