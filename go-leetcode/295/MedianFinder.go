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

/*

// 数据流的中位数

// index是0的时候，放大顶堆

// 如果小顶堆是空，直接放大顶堆
// 如果小顶堆peek < x,从小顶堆pop出来，放入大顶堆，将x放到小顶堆
// 如果小顶堆peek >= x，直接放到大顶堆

// index是1的时候，放小顶堆

func main() {
	m := NewMedianOperator()
	m.Push(1)
	fmt.Println(m.Get())
	m.Push(100)
	fmt.Println(m.Get())
	m.Push(2)
	m.Push(2)
	fmt.Println(m.Get())
	m.Push(3)
	m.Push(4)
	m.Push(5)
	fmt.Println(m.Get())
	m.Push(5)
	m.Push(5)

	m.Push(5)
	m.Push(5)
	m.Push(5)
	m.Push(5)
	m.Push(5)
	m.Push(5)
	m.Push(5)
	fmt.Println(m.Get())
}

func NewMedianOperator() *medianOperator {
	minHeap := &binaryHeap{data: make([]int,0),
		size: 0,
	more: func(i, j int) bool {
		return i > j
	}}
	maxHeap := &binaryHeap{data: make([]int,0),
		size: 0,
		more: func(i, j int) bool {
			return i < j
		}}
	m := &medianOperator{
		minHeap: minHeap,
		maxHeap: maxHeap,
		index: 0,
	}
	return m
}

type medianOperator struct {
	minHeap *binaryHeap
	maxHeap *binaryHeap
	index int
}

func (m *medianOperator) Push(x int) {
	if m.index == 0 {
		if m.minHeap.Len() == 0 {
			m.maxHeap.Push(x)
		} else {
			if m.minHeap.Peek() < x {
				m.maxHeap.Push(m.minHeap.Pop())
				m.minHeap.Push(x)
				// 这种情况 不要忘记了
			} else {
				m.maxHeap.Push(x)
			}
		}
	} else {
		if m.maxHeap.Len() == 0 {
			m.minHeap.Push(x)
		} else {
			if m.maxHeap.Peek() > x {
				m.minHeap.Push(m.maxHeap.Pop())
				m.maxHeap.Push(x)
				// 这种情况 不要忘记了
			} else {
				m.minHeap.Push(x)
			}
		}
	}
	m.index = 1 - m.index
}

func (m *medianOperator) Get() float64 {
	if (m.minHeap.Len() + m.maxHeap.Len() )%2 == 1 {
		return float64(m.maxHeap.Peek())
	}
	// 这里也要注意是float64
	return float64(m.maxHeap.Peek() + m.minHeap.Peek()) / 2
}

type binaryHeap struct {
	data []int
	size int
	more func(i,j int) bool
}

func (h *binaryHeap) Len() int {
	return h.size
}

func (h *binaryHeap) Peek() int {
	return h.data[0]
}

func (h *binaryHeap) Push(x int) {
	h.size ++
	// 这里注意越界
	if len(h.data) < h.size {
		h.data = append(h.data,make([]int,h.size - len(h.data))...)
	}
	// 这里注意不能append，一定要用 h.size - 1
	h.data[h.size-1] = x
	h.bubbleUp()
}

func (h *binaryHeap) Pop() int {
	result := h.data[0]
	h.data[0],h.data[h.size-1] = h.data[h.size-1],h.data[0]
	h.size --
	h.bubbleDown(0)
	return result
}

// 0 1 2 3,4,5,6
func (h *binaryHeap) bubbleUp() {
	index := h.size - 1
	for index > 0 {
		parent := (index-1)/2
		if h.more(h.data[parent],h.data[index]) {
			h.data[parent],h.data[index] = h.data[index],h.data[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *binaryHeap) bubbleDown(index int) {
	for index < h.size {
		left := 2*index + 1
		right := 2*index + 2
		largest := index
		if left < h.size && h.more(h.data[index],h.data[left]) {
			// 注意这里不要写错了
			largest = left
		}
		if right < h.size && h.more(h.data[index],h.data[largest]) {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index],h.data[largest] = h.data[largest],h.data[index]
		index = largest
	}
}

 */