package main

import "fmt"

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//

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

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next = &ListNode{Val: 5}
	//head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	printList(head)
	head = swapPairs(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := &ListNode{}
	result := prev
	for head != nil && head.Next != nil {
		nextNext := head.Next.Next
		next := head.Next

		next.Next = head
		head.Next = nil

		prev.Next = next
		prev = head
		head = nextNext
	}
	prev.Next = head
	return result.Next
}

