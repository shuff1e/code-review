package main

import (
	"fmt"
)

// 40：最小的k个数
// 题目：输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8
// 这8个数字，则最小的4个数字是1、2、3、4。

// A：快速排序的partition

func getKLeast1(arr []int,k int) []int{
	if len(arr) < k || len(arr) == 0 {
		return []int{}
	}
	start := 0
	end := len(arr) - 1
	index := partition(arr,start,end)
	for index != k {
		if index > k {
			end = index -1
		}
		if index < k {
			start = index + 1
		}
		index = partition(arr,start,end)
	}
	return arr[0:k]
}

func partition(arr []int,start,end int) int {
	if start >= end {
		return start
	}
	pivot := arr[start]
	mark := start
	for i := start;i<=end;i++ {
		if arr[i] < pivot {
			mark ++
			swap(arr,i,mark)
		}
	}
	swap(arr,start,mark)
	return mark
}

func swap(arr []int,i,j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func main() {
	Test("Test1",[]int{4, 5, 1, 6, 2, 7, 3, 8},[]int{1, 2, 3, 4})
	Test("Test2",[]int{4, 5, 1, 6, 2, 7, 3, 8},[]int{1, 2, 3, 4, 5, 6, 7, 8})
	Test("Test3",[]int{4, 5, 1, 6, 2, 7, 3, 8},make([]int,10))
	Test("Test4",[]int{4, 5, 1, 6, 2, 7, 3, 8},[]int{1})
	Test("Tesst5",[]int{4, 5, 1, 6, 2, 7, 3, 8},nil)
	Test("Test6",[]int{4, 5, 1, 6, 2, 7, 2, 8},[]int{1,2})
	Test("Test7",nil,nil)
}

func Test(name string,data []int,expected []int) {
	fmt.Println(name)

	printArray(expected)

	result := getKLeast2(data,len(expected))
	printArray(result)

}

func printArray(arr []int) {
	for _,v := range arr {
		fmt.Print(v," ")
	}
	fmt.Println()
}

func getKLeast2(arr []int,k int) []int{
	if len(arr) == 0 || len(arr) < k {
		return []int{}
	}

	heap := New(k)
	heap.AddLast(arr[0:k]...)

	for i := k;i<len(arr);i++ {
		v,_ := heap.GetFirst()
		if arr[i] < v {
			heap.RemoveFirst()
			heap.AddLast(arr[i])
		}
	}

	result := make([]int,k)
	for i := 0;i<k;i++ {
		v,_ := heap.RemoveFirst()
		result[i] = v
	}
	return result
}

// 最小的k个数，使用大根堆
type maxHeap struct {
	data []int
	size int
}

func New(k int) *maxHeap {
	return &maxHeap{
		data: make([]int,k),
		size: 0,
	}
}

func (heap *maxHeap) AddLast(args ...int) bool {
	if heap.size == len(heap.data) {
		return false
	}
	if len(args) == 0 {
		return false
	} else if len(args) == 1 {
		// 只添加一个元素
		heap.size++
		heap.data[heap.size-1] = args[0]
		// 这个地方不能用append，因为slice的size和我们记录的size不一样
		// 这里需要我们相当于自己实现slice的cap和len等机制
		heap.bubbleUp()
	} else {
		// 添加多个元素，从中间 bubbleDown
		for _,v := range args {
			heap.size ++
			heap.data[heap.size-1] = v
		}
		index := heap.size/2 +1
		for i := index;i>=0;i-- {
			heap.bubbleDown(i)
		}
	}
	return true
}

func (heap *maxHeap) RemoveFirst() (int,bool) {
	if heap.size == 0{
		return -1,false
	}
	result := heap.data[0]
	heap.swap(0,heap.size-1)
	heap.size --
	if heap.size > 0 {
		heap.bubbleDown(0)
	}
	return result,true
}

func (heap *maxHeap) GetFirst() (int,bool) {
	if heap.size == 0 {
		return -1,false
	}
	return heap.data[0],true
}

//         0
//     /      \
//    1        2
//  /   \    /   \
// 3     4   5   6

func (heap *maxHeap) bubbleUp() {
	index := heap.size - 1
	for index != 0 {
		parent := (index-1)/2
		if heap.data[parent] < heap.data[index] {
			swap(heap.data,parent,index)
			index = parent
		} else {
			break
		}
	}
}

func (heap *maxHeap) bubbleDown(index int) {
	for index < heap.size {
		largest := index
		left := 2*index + 1
		right := 2 *index + 2
		if left < heap.size && heap.data[left] > heap.data[index] {
			largest = left
		}
		if right < heap.size && heap.data[right] > heap.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}
		heap.swap(index,largest)
		index = largest
	}
}

func (heap *maxHeap) swap(i,j int) {
	temp := heap.data[i]
	heap.data[i] = heap.data[j]
	heap.data[j] = temp
}

