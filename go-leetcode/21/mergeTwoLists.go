package main

/*
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: 0,
	}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
			temp = temp.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
			temp = temp.Next
		}
	}

	if l1 != nil {
		temp.Next = l1
	}

	if l2 != nil {
		temp.Next = l2
	}

	return head.Next
}
