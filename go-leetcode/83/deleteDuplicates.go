package main

import "fmt"

/*
83. 删除排序链表中的重复元素
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node := deleteDuplicates(node1)

	for node != nil {
		fmt.Print(node.Val,"->")
		node = node.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	return getAppearOnce(head)
}

func getAppearOnce(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for head != nil && head.Next != nil && head.Val == head.Next.Val {
		head = head.Next
	}
	head.Next = getAppearOnce(head.Next)
	return head
}