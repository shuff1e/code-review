package main

import "fmt"

// Q：N个长度不一的数组，所有数组都是有序的，从大到小打印这N个数组整体最大的topK

type Node struct {
	arrayNum int
	index int
	Value int
}

// 堆操作
// heap index，下沉
// 用于初始化，给堆中插入数据
func heapify(heap []*Node, index, heapSize int ) {
	left := 2*index+1
	right := 2*index+2
	largest := index

	for left < heapSize {
		if heap[left].Value > heap[index].Value {
			largest = left
		}
		if right < heapSize && heap[right].Value > heap[largest].Value {
			largest = right
		}
		if index != largest {
			swap(heap,index,largest)
		} else {
			break
		}
		index = largest
		left = 2*index+1
		right = 2*index+2
	}
}

// 堆操作
// heap index 冒泡
// 用于pop出堆中的最大数据后，重建堆
func heapInsert(heap []*Node,index int) {
	parent := (index-1)/2
	for parent >= 0 {
		if heap[parent].Value < heap[index].Value {
			swap(heap,parent,index)
		} else {
			break
		}
		parent = (index-1)/2
	}
}

func swap(heap []*Node, left, right int) {
	temp := heap[left]
	heap[left] = heap[right]
	heap[right] = temp
}

func printTopK(matrix [][]int,topK int) {
	heapSize := len(matrix)
	heap := make([]*Node,heapSize)

	for i := 0;i<heapSize;i++ {
		length := len(matrix[i]) - 1
		node := &Node{i,length,matrix[i][length]}
		heap[i] = node
		heapInsert(heap,i)
	}

	for i := 0;i<topK;i++ {
		if heapSize < 0 {
			break
		}
		fmt.Println(heap[0].Value)
		if heap[0].index > 0 {
			arrayNum := heap[0].arrayNum
			index := heap[0].index - 1
			heap[0] = &Node{arrayNum,index,matrix[arrayNum][index]}
		} else {
			swap(heap,0,heapSize-1)
			heapSize --
		}
		heapify(heap,0,heapSize)
	}
}

func main() {
	matrix := [][]int {
		{1,2,30,40},
		{1,3,5,600},
		{7,8,9,32},
	}
	printTopK(matrix,4)
}