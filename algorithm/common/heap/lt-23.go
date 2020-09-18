package main

import "fmt"

// 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。


//输入:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//输出: 1->1->2->3->4->4->5->6

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func heapInsert2(heap []*ListNode,index int) {
	for index > 0 {
		parent := (index-1)/2
		if heap[parent].Val > heap[index].Val {
			swap2(heap,index,parent)
			index = parent
		} else {
			break
		}
	}
}

func heapify2(heap []*ListNode,index ,heapSize int ) {
	left := 2*index +1
	right := 2*index +2
	smallest := index
	for left < heapSize {
		if heap[left].Val < heap[index].Val {
			smallest = left
		}
		if right < heapSize && heap[right].Val < heap[smallest].Val {
			smallest = right
		}
		if smallest != index {
			swap2(heap,smallest,index)
		} else {
			break
		}
		index = smallest
		left = 2*smallest + 1
		right = 2*smallest + 2
	}
}

func swap2(heap []*ListNode,a,b int) {
	temp := heap[a]
	heap[a] = heap[b]
	heap[b] = temp
}

func isEmpty(lists []*ListNode) (result []*ListNode,isEmpty bool) {
	isEmpty = true
	for _,list := range lists {
		if list != nil {
			isEmpty = false
			result = append(result,list)
		}
	}
	return
}

func mergeKLists(lists []*ListNode) *ListNode {
	result,empty := isEmpty(lists)
	if empty{
		return nil
	}
	heapSize := len(result)
	heap := make([]*ListNode,heapSize)

	for i :=0;i<heapSize;i++ {
		heap[i] = result[i]
		heapInsert2(heap,i)
	}

	head := &ListNode{}
	temp := head
	for heapSize > 0 {
		temp.Next = heap[0]
		temp = temp.Next
		if heap[0].Next != nil {
			heap[0] = heap[0].Next
		} else {
			swap2(heap,heapSize-1,0)
			heapSize --
		}
		heapify2(heap,0,heapSize)
	}
	return head.Next
}

func generateList(array []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _,v := range array {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	return head.Next
}

func generateLists(matrix [][]int) (result []*ListNode) {
	for _,v := range matrix {
		result = append(result,generateList(v))
	}
	return
}

func main() {
	matrix := [][]int {
		{-8,-7,-7,-5,1,1,3,4},
		{-2},
		{-10,-10,-7,0,1,3},
		{2},
	}
	lists := generateLists(matrix)

	node := mergeKLists(lists)
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}