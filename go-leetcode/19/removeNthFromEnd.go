package main

/*
19. 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

 */

// fast先走2步
// fast 3 4 5
// slow 1 2 3

type ListNode struct {
	Val int
	Next *ListNode
}

// 1 2 3 4 5

// 先说正常情况
// 倒数第三个，我们希望fast走到4
// 这样
// fast 4 5
// slow 1 2

// 1 2 3 4 5
// 0 1 2 3 4
//       在这个位置跳出循环
// 如果 4 == nil，temp就是nil，直接返回head.next
// 如果 3 == nil，直接在index=2时跳出循环

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	index := 0
	for index < n && fast != nil {
		index ++
		fast = fast.Next
	}

	if index < n {
		return nil
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil && slow != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}