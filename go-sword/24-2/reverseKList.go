package main

import "fmt"

// Q：输入一个链表，以k个为一部分反转，不足k的不反转
// 123 456 789 10
// 翻转成 321 654 987 10

// A：找到第k个，反转，返回root
// root.next = 递归
// 终止条件，不足k个，直接返回
























type Node struct {
	value int
	next *Node
}

func reverseKList(root *Node,k int) *Node {
	if root == nil {
		return nil
	}

	count := 1
	temp := root
	for count < k && temp != nil && temp.next != nil {
		temp = temp.next
		count += 1
	}
	if count < k {
		return root
	}

	next := temp.next
	temp.next = nil
	head,tail := reverseList(root)
	tail.next = reverseKList(next,k)
	return head
}

// 反转链表，并返回新的头节点和尾节点
func reverseList(root *Node) (head,tail *Node) {
	if root == nil {
		return nil,nil
	}
	prev := (*Node)(nil)
	cur := root
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev,root
}

func main() {
	Test(3,1,2,3,4,5,6,7)
	Test(3,1,2)
	Test(3,1,2,3)
	Test(3,1)
	Test(1,1)
	Test(1)
}

func Test(k int,args ...int) {
	node := createList(args...)
	printList(node)
	temp := reverseKList(node,k)
	printList(temp)
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