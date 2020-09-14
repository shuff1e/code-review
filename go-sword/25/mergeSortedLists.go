package main

import (
	"fmt"
)

// 25：合并两个排序的链表
// 题目：输入两个递增排序的链表，合并这两个链表并使新链表中的结点仍然是按
// 照递增排序的。例如输入图3.11中的链表1和链表2，则合并之后的升序链表如链
// 表3所示。
// 1->3->5->7
// 2->4->6->8
// 1->2->3->4->5->6->7->8

type Node struct {
	value int
	next *Node
}

func mergeTwoSortedLists(node1,node2 *Node) *Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	if node1.value <= node2.value {
		node1.next = mergeTwoSortedLists(node1.next,node2)
		return node1
	} else {
		node2.next = mergeTwoSortedLists(node1,node2.next)
		return node2
	}
}

func main() {
	Test([]int{1,3,5},[]int{2,4,6})
	Test([]int{1,3,5},[]int{1,3,5})
	Test([]int{1},[]int{2})
	Test([]int{1,3,5},nil)
	Test([]int{},[]int{})
}

func Test(arr1,arr2 []int) {
	node1 := createList(arr1...)
	printList(node1)
	node2 := createList(arr2...)
	printList(node2)
	node3 := mergeTwoSortedLists(node1,node2)
	printList(node3)
}

func createList(arr ...int) *Node {
	tail := (*Node)(nil)
	for i := len(arr)-1;i>=0;i-- {
		node := &Node{value: arr[i]}
		node.next = tail
		tail = node
	}
	return tail
}

func printList(root *Node) {
	for root != nil {
		fmt.Print(root.value,"->")
		root = root.next
	}
	fmt.Println("nil")
}