package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 18（二）：删除链表中重复的结点
// 题目：在一个排序的链表中，如何删除重复的结点？例如，在图3.4（a）中重复
// 结点被删除之后，链表如图3.4（b）所示。
// (a) 1->2->3->3->5
// (b) 1->2->5

type Node struct {
	value int
	next *Node
}

// 直接在链表上操作
func deleteDup(root *Node) *Node {
	return doit(root)
}

// 获得以root开头的不重复的链表
func doit(root *Node) *Node {
	if root == nil {
		return nil
	}
	// 获得以root开头的节点不重复的第一个元素
	node := findFirstSingle(root)
	if node == nil {
		return nil
	}
	// node.next 为 node.next开头的第一个不重复的元素
	node.next = doit(node.next)
	return node
}

// 发现链表中第一个不重复的元素
func findFirstSingle(root *Node) *Node {
	if root == nil || root.next == nil {
		return root
	}
	prev := (*Node)(nil)
	for {
		if prev == nil {
			if root.value == root.next.value {
				prev = root
				root = root.next
			} else {
				return root
			}
		} else if root.next == nil {
			if root.value == prev.value {
				return nil
			} else {
				return root
			}
		} else if prev.value == root.value {
			root = root.next
		} else if root.value == root.next.value {
			prev = root
			root = root.next
		} else {
			return root
		}
	}
}

// 使用stack
func deleteDuplication(root *Node) *Node {
	stack := linkedliststack.New()
	for root != nil {
		stack.Push(root)
		root = root.next
	}
	prev := (*Node)(nil)
	findDupli := false
	for !stack.Empty() {
		temp,_ := stack.Pop()
		cur := temp.(*Node)
		if prev == nil || cur.value != prev.value {
			if findDupli {
				cur.next = prev.next
				prev = cur
				findDupli = false
			} else {
				cur.next = prev
				prev = cur
			}
		} else {
			findDupli = true
			continue
		}
	}
	if findDupli {
		return prev.next
	}
	return prev
}

func main() {
	Test(1,2,3,3,4,4,5)
	Test(1,2,3,4,5,6,7)
	Test(1,1,1,1,1,1,2)
	Test(1,1,1,1,1,1,1)
	Test(1,1,2,2,3,3,4,4)
	Test(1,1,2,3,3,4,5,5)
	Test(1,2)
	Test(1)
	Test(1,1)
	Test()
}

func Test(args ...int) {
	node := createList(args...)
	temp := deleteDuplication(node)
	printList(temp)
	//node := createList(args...)
	//temp := deleteDup(node)
	//printList(temp)
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