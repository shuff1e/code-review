package main

import "fmt"

/*
86. 分隔链表
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。



示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

*/

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 5}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 2}
	node.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	node = partition2(node,3)
	for node != nil {
		fmt.Print(node.Val,"->")
		node = node.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func partition2(head *ListNode,x int) *ListNode {
	leftHead := &ListNode{Val: -1}
	rightHead := &ListNode{Val: -1}
	left := leftHead
	right := rightHead

	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}

	right.Next = nil // important
	left.Next = rightHead.Next
	return leftHead.Next
}

func partition(head *ListNode, x int) *ListNode {
	node := &ListNode{Val: x}
	node.Next = head
	head = node

	temp := head
	tempPre := (*ListNode)(nil)

	mark := head
	markPre := (*ListNode)(nil)


	for temp != nil {
		if temp.Val < x {
			markPre = mark
			mark = mark.Next
			a,b := swap(markPre,tempPre)
			markPre = a
			tempPre = b
			mark = markPre.Next
			temp = tempPre.Next
		}
		tempPre = temp
		temp = temp.Next
	}

	return node.Next
}

// 1->10->3->4->5->6
// 1,10
func swap(a,b *ListNode) (*ListNode,*ListNode) {
	if a == b {
		return a,b
	}
	if a.Next == b {
		nextNext := b.Next.Next
		next := b.Next
		a.Next = next
		a.Next.Next = b
		b.Next = nextNext
		return a,a.Next
	}
	aNext := a.Next
	aNextNext := a.Next.Next
	bNext := b.Next
	bNextNext := b.Next.Next

	a.Next = bNext
	bNext.Next = aNextNext

	b.Next = aNext
	aNext.Next = bNextNext
	return a,b
}
