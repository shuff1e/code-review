package main

import "fmt"

/*

445. 两数相加 II
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。
它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



进阶：

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。



示例：

输入：(7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 8 -> 0 -> 7

 */

func main() {
	head := (*ListNode)(nil)
	//head := &ListNode{Val: 9}
	//head.Next = &ListNode{Val: 9}
	//head.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next.Next = &ListNode{Val: 5}
	//head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	node := (*ListNode)(nil)
	//node := &ListNode{Val: 5}
	//node.Next = &ListNode{Val: 6}
	//node.Next.Next = &ListNode{Val: 4}

	result := addTwoNumbers(head,node)
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
	stack1 := []int{}
	stack2 := []int{}
	for l1 != nil {
		stack1 = append(stack1,l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2,l2.Val)
		l2 = l2.Next
	}

	carry := 0
	result := (*ListNode)(nil)
	for len(stack1) > 0 && len(stack2) > 0 {
		sum := stack1[len(stack1)-1] + stack2[len(stack2)-1] + carry
		carry = sum/10
		temp := &ListNode{Val: sum%10}
		temp.Next = result
		result = temp

		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
	}

	for len(stack1) > 0 {
		sum :=stack1[len(stack1)-1] + carry
		carry = sum/10
		temp := &ListNode{Val: sum%10}
		temp.Next = result
		result = temp

		stack1 = stack1[:len(stack1)-1]
	}

	for len(stack2) > 0 {
		sum :=stack2[len(stack2)-1] + carry
		carry = sum/10
		temp := &ListNode{Val: sum%10}
		temp.Next = result
		result = temp

		stack2 = stack2[:len(stack2)-1]
	}

	if carry != 0 {
		temp := &ListNode{Val: carry}
		temp.Next = result
		result = temp
	}

	return result
}