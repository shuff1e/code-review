package main

/*
61. 旋转链表
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	length,tail := getLength(head)
	k = k%length
	if k == 0 {
		return head
	}
	pre := help(head,k)
	temp := pre.Next
	pre.Next = nil
	tail.Next = head
	return temp
}

// 1->2->3->nil
// count 2 3
// head  2 3
func getLength(head *ListNode) (int,*ListNode) {
	count := 1
	for head.Next != nil {
		head = head.Next
		count ++
	}
	return count,head
}

// 输入: 1->2->3->4->5->6->NULL, k = 2
// 输出: 5->6->1->2->3->4->NULL

// 1 2 3 4 5 6

// count  0 1 2
// fast   1 2 3

// fast 3 4 5 6
// slow 1 2 3 4

func help(head *ListNode,k int) (Kth *ListNode) {
	// 得到倒数第2个
	count := 0
	fast := head
	for count < k && fast != nil {
		fast = fast.Next
		count ++
	}
	if count < k {
		return nil
	}
	if fast == nil {
		return nil
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
