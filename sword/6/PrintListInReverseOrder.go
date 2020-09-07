package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 6：从尾到头打印链表
// 题目：输入一个链表的头结点，从尾到头反过来打印出每个结点的值。

type Node struct {
	value interface{}
	next  *Node
}

func printIteratively(n *Node) {
	stack := linkedliststack.New()
	root := n
	for root != nil {
		stack.Push(root.value)
		root = root.next
	}
	for !stack.Empty() {
		ele,_ := stack.Pop()
		fmt.Println(ele.(string))
	}
}

func printRecursively(n *Node) {
	if n == nil {
		return
	}
	printRecursively(n.next)
	fmt.Println(n.value.(string))
}

func main() {
	root := &Node{value: "1"}
	root.next = &Node{value: "2"}
	root.next.next = &Node{value: "3"}
	root.next.next.next = &Node{value: "4"}
	root.next.next.next.next = &Node{value: "5"}

	printIteratively(root)
	printRecursively(root)
}