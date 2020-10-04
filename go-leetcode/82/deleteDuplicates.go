package main

import "fmt"

/*
82. 删除排序链表中的重复元素 II
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5

1 2 3 4 5 6

示例 2:

输入: 1->1->1->2->3
输出: 2->3
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 3}

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
	return getFirstAppearingOnce(head)
}

func getFirstAppearingOnce(head *ListNode) *ListNode {
	temp := head
	if temp == nil {
		return temp
	}

	pre := (*ListNode)(nil)
	for (temp != nil && temp.Next != nil && temp.Val == temp.Next.Val) ||
		(temp != nil && pre != nil && temp.Val == pre.Val ){
		pre = temp
		temp = temp.Next
	}

	if temp != nil {
		temp.Next = getFirstAppearingOnce(temp.Next)
	}

	return temp
}