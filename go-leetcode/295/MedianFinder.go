package main

import (
	"container/heap"
	"fmt"
)

/*
295. 数据流的中位数
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。

例如，

[2,3,4] 的中位数是 3

[2,3] 的中位数是 (2 + 3) / 2 = 2.5

设计一个支持以下两种操作的数据结构：

void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
示例：

addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:

如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 */

// 小根堆
// 大根堆

// index为偶数的放大根堆，如果arr[index]大于小根堆的top，将小根堆pop出来放到大根堆
// index为奇数的放小根堆，如果arr[index]小于大根堆的top，将大根堆pop出来放到小根堆

func main() {
	obj := Constructor()
	obj.AddNum(-1)
	obj.AddNum(-2)
	obj.AddNum(-3)
	obj.AddNum(-4)
	obj.AddNum(-5)
	// 1 1 2 3 4
	fmt.Println(obj.FindMedian())
}

type myHeap struct {
	data []int
	comparator func(x,y int) bool
}

func (h myHeap) Len() int { return len(h.data) }
func (h myHeap) Less(i,j int) bool { return h.comparator(h.data[i],h.data[j]) }
func (h myHeap) Swap(i,j int) { h.data[i],h.data[j] = h.data[j],h.data[i] }

func (h *myHeap) Push(x interface{}) {
	(*h).data = append((*h).data,x.(int))
}

func (h *myHeap) Pop() interface{} {
	result := ((*h).data)[len((*h).data)-1]
	(*h).data = ((*h).data)[:len((*h).data)-1]
	return result
}

func (h *myHeap) Peek() int {
	return (*h).data[0]
}

type MedianFinder struct {
	index int
	minHeap *myHeap
	maxHeap *myHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		index: 0,
		minHeap: &myHeap{
			data: make([]int,0),
			comparator: func(x, y int) bool {
				return x < y
			},
		},
		maxHeap: &myHeap{
			data: make([]int,0),
			comparator: func(x, y int) bool {
				return x > y
			},
		},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	if this.index % 2 == 0 {
		if this.minHeap.Len() > 0 && this.minHeap.Peek() < num {
			heap.Push(this.maxHeap,heap.Pop(this.minHeap))
			heap.Push(this.minHeap,num)
		} else {
			heap.Push(this.maxHeap,num)
		}
	} else {
		if this.maxHeap.Len() > 0 && this.maxHeap.Peek() > num {
			heap.Push(this.minHeap,heap.Pop(this.maxHeap))
			heap.Push(this.maxHeap,num)
		} else {
			heap.Push(this.minHeap,num)
		}
	}
	this.index ++
}


func (this *MedianFinder) FindMedian() float64 {
	if this.index == 0 {
		return 0
	}

	if this.index % 2 == 0 {
		return (float64(this.maxHeap.Peek()) + float64(this.minHeap.Peek()))/2
	} else {
		return float64(this.maxHeap.Peek())
	}
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
