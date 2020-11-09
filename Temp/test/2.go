package main

import "fmt"

/*

2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

 */

func main() {
	node1 := &ListNode{Val: 2}
	node1.Next = &ListNode{Val: 3}
	node1.Next.Next = &ListNode{Val: 4}
	node1.Next.Next.Next = &ListNode{Val: 5}
	node1.Next.Next.Next.Next = &ListNode{Val: 6}

	node2 := &ListNode{Val: 5}
	node2.Next = &ListNode{Val: 7}

	result := addTwoNumbers(node1,node2)

	for result != nil {
		fmt.Print(result.Val,"->")
		result = result.Next
	}

}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	root := dummy
	carry := 0

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + carry
		dummy.Next = &ListNode{
			Val: sum%10,
		}
		carry = sum/10

		l1 = l1.Next
		l2 = l2.Next
		dummy = dummy.Next
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		sum := l1.Val + carry
		dummy.Next = &ListNode{Val: sum%10}
		carry = sum/10

		l1 = l1.Next
		dummy = dummy.Next
	}

	if carry != 0 {
		dummy.Next = &ListNode{Val: carry}
	}
	return root.Next
}