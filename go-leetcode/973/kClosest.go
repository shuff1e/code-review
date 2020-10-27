package main

import (
	"container/heap"
	"fmt"
)

/*

973. 最接近原点的 K 个点
我们有一个由平面上的点组成的列表 points。需要从中找出 K 个距离原点 (0, 0) 最近的点。

（这里，平面上两点之间的距离是欧几里德距离。）

你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

示例 1：

输入：points = [[1,3],[-2,2]], K = 1
输出：[[-2,2]]
解释：
(1, 3) 和原点之间的距离为 sqrt(10)，
(-2, 2) 和原点之间的距离为 sqrt(8)，
由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
示例 2：

输入：points = [[3,3],[5,-1],[-2,4]], K = 2
输出：[[3,3],[-2,4]]
（答案 [[-2,4],[3,3]] 也会被接受。）


提示：

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000

 */









func main() {
	points := [][]int{{3,3},{5,-1},{-2,4}}
	result := kClosest(points,2)
	for _,v := range result {
		fmt.Println(v)
	}
}

func kClosest(points [][]int, K int) [][]int {
	if len(points) <= K {
		return points
	}

	h := &maxHeap{}
	index := 0
	for ;index < K;index ++ {
		heap.Push(h,points[index])
	}

	for ;index < len(points);index++ {
		temp := h.Peek()
		if getDistance(points[index]) < getDistance(temp) {
			heap.Pop(h)
			heap.Push(h,points[index])
		}
	}

	result := [][]int{}
	for h.Len() > 0 {
		result = append(result,h.Pop().([]int))
	}
	return result
}

func getDistance(arr []int) int {
	return arr[0]*arr[0] + arr[1] * arr[1]
}

type maxHeap [][]int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i,j int) bool {
	// maxHeap就是大于
	return getDistance(h[i]) > getDistance(h[j])
}

func (h maxHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h,x.([]int))
}

func (h *maxHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}

// 注意这里的peek
func (h *maxHeap) Peek() []int {
	return (*h)[0]
}
