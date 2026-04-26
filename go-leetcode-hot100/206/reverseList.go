package main

import "fmt"

/*
206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:

你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 */

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 5}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 2}
	node.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	result := reverseList2(node)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result,_ := help(head)
	return result
}

func help(node *ListNode) (*ListNode,*ListNode) {
	if node.Next == nil {
		return node,node
	}

	head,prev := help(node.Next)
	node.Next = nil
	prev.Next = node
	return head,node
}

/*

class Solution {
    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode p = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return p;
    }
}

 */