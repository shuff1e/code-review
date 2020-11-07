package main

import "fmt"

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