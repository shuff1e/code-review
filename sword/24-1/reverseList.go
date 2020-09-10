package main

import "fmt"

// 24-1：反转链表
// 题目：定义一个函数，输入一个链表的头结点，反转该链表并输出反转后链表的
// 头结点。

// A：1->2->3->4->5->6
// prev
// next = cur.next
// cur.next = prev
// prev = cur
// cur = next

type Node struct {
	value int
	next *Node
}

func reverseList(root *Node) *Node {
	if root == nil {
		return nil
	}
	prev := (*Node)(nil)
	cur := root
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	Test1()
	Test2()
	Test3()
}

func Test1() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	printList(root)
	root = reverseList(root)
	printList(root)
}

func Test2()  {
	root := &Node{value: 1}
	printList(root)
	root = reverseList(root)
	printList(root)
}

func Test3() {
	root := (*Node)(nil)
	printList(root)
	root = reverseList(root)
	printList(root)
}

func printList(root *Node) {
	for root != nil {
		fmt.Print(root.value,"->")
		root = root.next
	}
	fmt.Println("nil")
}