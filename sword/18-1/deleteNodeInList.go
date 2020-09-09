package main

import "fmt"

// 18（一）：在O(1)时间删除链表结点
// 题目：给定单向链表的头指针和一个结点指针，定义一个函数在O(1)时间删除该
// 结点。

// A：如果不是尾节点，用后一个节点的值覆盖该节点的值，并删除后一个节点
// 如果是尾节点，还是要找到该节点的前驱
// 总的时间复杂度为 (n + (n-1)*1)/n=2-1/n

type Node struct {
	value int
	next *Node
}

func deleteNode(root,target *Node) *Node {
	if root == nil {
		return nil
	}
	if target == nil {
		return root
	}

	if root == target {
		return root.next
	}

	// 是尾节点
	if target.next == nil {
		temp := root
		for temp.next != target {
			temp = temp.next
		}
		temp.next = nil
		return root
	}

	target.value = target.next.value
	target.next = target.next.next

	return root
}

func printList(root *Node) {
	for root != nil {
		fmt.Print(root.value,"->")
		root = root.next
	}
	fmt.Println("nil")
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
}

func Test1() {
	node1 := &Node{value: 1}
	node1.next = &Node{value: 2}
	node1.next.next = &Node{value: 3}
	node1.next.next.next = &Node{value: 4}
	node1.next.next.next.next = &Node{value: 5}
	node1.next.next.next.next.next = &Node{value: 6}
	printList(deleteNode(node1,node1.next.next.next.next.next))
}

func Test2() {
	node1 := &Node{value: 1}
	node1.next = &Node{value: 2}
	node1.next.next = &Node{value: 3}
	node1.next.next.next = &Node{value: 4}
	node1.next.next.next.next = &Node{value: 5}
	node1.next.next.next.next.next = &Node{value: 6}
	printList(deleteNode(node1,node1))
}

func Test3() {
	node1 := &Node{value: 1}
	node1.next = &Node{value: 2}
	node1.next.next = &Node{value: 3}
	node1.next.next.next = &Node{value: 4}
	node1.next.next.next.next = &Node{value: 5}
	node1.next.next.next.next.next = &Node{value: 6}
	printList(deleteNode(node1,node1.next.next))
}

func Test4() {
	node1 := &Node{value: 1}
	printList(deleteNode(node1,node1))
}
