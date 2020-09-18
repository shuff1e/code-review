package main

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	carry := 0
	x := 0
	y := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}

		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum := x + y + carry
		carry = sum / 10
		result.Next = &ListNode{Val:sum%10}
		result = result.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		result.Next=&ListNode{Val:1}
	}
	return head.Next
}