package main

import "fmt"

/*

给一个链表，翻转 L 到 R的部分

例如链表，1->2->3->4->5->6->7->8
L为 1
R为 3

翻转后为 3->2->1->4->5->6->7->8

 */

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	node.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	node.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	node.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	node.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	printList(node)

	result := reverseLR(node,2,14)
	printList(result)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val,"->")
		node = node.Next
	}
	fmt.Println()
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseLR(head *ListNode,l,r int) *ListNode {
	if l >= r {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// dummy的index为0
	index := 0
	temp := dummy

	for index < l - 1 && temp != nil {
		index ++
		temp = temp.Next
	}

	if index < l - 1 {
		return head
	}

	if temp == nil || temp.Next == nil || temp.Next.Next == nil {
		return head
	}

	p1 := temp
	help := p1.Next
	p1.Next = nil

	index ++
	temp = help

	for index < r && temp != nil {
		index ++
		temp = temp.Next
	}

	if index < r || temp == nil {
		a,_ := reverse(help)
		p1.Next = a
	} else {
		p2 := temp.Next
		temp.Next = nil

		a,b := reverse(help)

		p1.Next = a
		b.Next = p2
	}
	return dummy.Next
}

func reverse(node *ListNode) (*ListNode,*ListNode) {
	tail := node
	prev := (*ListNode)(nil)
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev,tail
}