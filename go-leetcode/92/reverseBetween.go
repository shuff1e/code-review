package main

import "fmt"

/*
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
 */

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	head := reverseBetween(node,1,3)
	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	pivot := &ListNode{
		Val: -1,
	}
	pivot.Next = head

	temp := pivot
	count := 1
	for count < m && temp != nil {
		temp = temp.Next
		count ++
	}
	if count < m {
		return pivot.Next
	}
	if temp == nil {
		return pivot.Next
	}

	//
	temp2 := temp.Next
	tail := temp2


	// (0,2) (1,3) (2,4) （3，5）
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL
	pre := (*ListNode)(nil)
	count = 0
	for count < n - m +1 && temp2 != nil{
		next := temp2.Next
		temp2.Next = pre
		pre = temp2
		temp2 = next
		count ++
	}
	temp.Next = pre
	tail.Next = temp2
	return pivot.Next
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	// 1 2 3 4 5
	//   2   4
	// 1 2 3 4
	// 1 2 3 4

	dummy := &ListNode{}
	dummy.Next = head

	index := 1
	temp := head
	prev := dummy

	for index < m {
		temp = temp.Next
		prev = prev.Next
		index ++
	}
	for index < n {
		index ++
		temp = temp.Next
	}

	next := temp.Next
	temp.Next = nil

	temp = prev.Next
	prev.Next = nil

	h,t := reverse(temp)
	prev.Next = h
	t.Next = next
	return dummy.Next
}

func reverse(node *ListNode) (*ListNode,*ListNode) {
	temp := node
	prev := (*ListNode)(nil)
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev,temp
}
