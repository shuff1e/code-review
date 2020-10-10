package main

import "fmt"

/*
147. 对链表进行插入排序
对链表进行插入排序。

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5
 */

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 5}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 2}
	node.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	node = insertionSortList(node)
	for node != nil {
		fmt.Print(node.Val,"->")
		 node = node.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	curr := head.Next
	dummy := &ListNode{Val: -0x80000000}
	dummy.Next = head
	for curr != nil {
		// 两个pre(pre+dummy)
		if pre.Val > curr.Val {
			temp := dummy
			for temp.Next.Val < curr.Val {
				temp = temp.Next
			}
			tempNext := temp.Next
			temp.Next = curr
			pre.Next = curr.Next
			curr.Next = tempNext

			curr = pre.Next
		} else {
			curr = curr.Next
			pre = pre.Next
		}
	}
	return dummy.Next
}
