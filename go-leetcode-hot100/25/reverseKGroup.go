package main

import "fmt"

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
//k 是一个正整数，它的值小于或等于链表的长度。
//
//如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 
//
//示例：
//
//给你这个链表：1->2->3->4->5
//
//当 k = 2 时，应当返回: 2->1->4->3->5
//
//当 k = 3 时，应当返回: 3->2->1->4->5
//
// 
//
//说明：
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := (*ListNode)(nil)
	//head := &ListNode{Val: 1}
	//head.Next = &ListNode{Val: 2}
	//head.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next = &ListNode{Val: 4}
	//head.Next.Next.Next.Next = &ListNode{Val: 5}
	//head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	printList(head)
	head = reverseKGroup(head,4)
	//head,_ = reverseList(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
	fmt.Println()
}

// 砍下长度为k的一段，然后递归
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 || head == nil {
		return head
	}
	temp := head
	for i:=1;i<k;i++ {
		temp = temp.Next
		// 注意nil
		if temp == nil {
			return head
		}
	}

	next := temp.Next
	temp.Next = nil
	head,tail := reverseList(head)
	tail.Next = reverseKGroup(next,k)
	return head
}

func reverseList(head *ListNode) (*ListNode,*ListNode) {
	tail := head
	prev := (*ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev,tail
}