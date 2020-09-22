package main

// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// 例一：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6

// 例二：
// 输入：lists = []
// 输出：[]

// 例三：
// 输入：lists = [[]]
// 输出：[]

// k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4
//

// A：类似于合并K个排序的数组
// 关键在于，将每次找到最小数字的时间降低
// 可以使用小根堆

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := New(len(lists))
	for _,v := range lists {
		if v != nil {
			heap.Add(v)
		}
	}

	result := heap.Remove()
	if result != nil && result.Next != nil {
		heap.Add(result.Next)
	}

	prev := result

	for heap.Size() > 0 {
		temp := heap.Remove()
		if temp.Next != nil {
			heap.Add(temp.Next)
		}
		prev.Next = temp
		prev = temp
	}
	return result
}

type minHeap struct {
	data []*ListNode
	size int
}

func New(n int) *minHeap {
	return &minHeap{
		data: make([]*ListNode,n),
		size: 0,
	}
}

func (heap *minHeap) Size() int {
	return heap.size
}

func (heap *minHeap) Add(nodes ...*ListNode) {
	if len(nodes) == 0 {
		return
	} else if len(nodes) == 1 {
		heap.size ++
		heap.data[heap.size-1]= nodes[0]
		heap.bubbleUp()
	} else {
		for _,v := range nodes {
			heap.size ++
			heap.data[heap.size-1] = v
		}
		for i := (heap.size+1)/2;i>=0;i-- {
			heap.bubbleDown(i)
		}
	}
}

func (heap *minHeap) Remove() *ListNode {
	if heap.size == 0 {
		return nil
	}
	result := heap.data[0]
	heap.swap(0,heap.size-1)
	heap.size --
	heap.bubbleDown(0)
	return result
}

func (heap *minHeap) bubbleDown(index int) {
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index
		if left < heap.size && heap.data[left].Val < heap.data[smallest].Val {
			smallest = left
		}
		if right < heap.size && heap.data[right].Val < heap.data[smallest].Val {
			smallest = right
		}
		if smallest == index {
			break
		}
		heap.swap(smallest,index)
		index = smallest
	}
}

func (heap *minHeap) bubbleUp() {
	index := heap.size - 1
	for index > 0 {
		parent := (index-1)/2
		if heap.data[parent].Val > heap.data[index].Val {
			heap.swap(parent,index)
		} else {
			break
		}
		index = parent
	}
}

func (heap *minHeap) swap(i,j int) {
	temp := heap.data[i]
	heap.data[i] = heap.data[j]
	heap.data[j] = temp
}