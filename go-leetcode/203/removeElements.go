package main

import "fmt"

/*
203. 移除链表元素
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
 */

func main() {
	node := (*ListNode)(nil)
	//node := &ListNode{Val: 1}
	//node.Next = &ListNode{Val: 1}
	//node.Next.Next = &ListNode{Val: 1}
	//node.Next.Next.Next = &ListNode{Val: 1}
	//node.Next.Next.Next.Next = &ListNode{Val: 1}
	//node.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head := removeElements(node,1)
	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	prev := dummy
	curr := dummy
	for curr != nil {
		curr = curr.Next
		for curr != nil && curr.Val == val {
			curr = curr.Next
		}
		prev.Next = curr
		prev = curr
	}
	return dummy.Next
}

/*

class Solution {
  public ListNode removeElements(ListNode head, int val) {
    ListNode sentinel = new ListNode(0);
    sentinel.next = head;

    ListNode prev = sentinel, curr = head;
    while (curr != null) {
      if (curr.val == val) prev.next = curr.next;
      else prev = curr;
      curr = curr.next;
    }
    return sentinel.next;
  }
}

 */