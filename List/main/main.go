package main

import (
	"fmt"
	"github.com/shuff1e/code-review/List"
)

func main() {
	head := &List.ListNode{Data: 1}
	head.Next = &List.ListNode{Data: 2}
	head.Next.Next = &List.ListNode{Data: 3}
	head.Next.Next.Next = &List.ListNode{Data: 4}
	head.Next.Next.Next.Next = &List.ListNode{Data: 5}
	head.Next.Next.Next.Next.Next = &List.ListNode{Data: 6}
	//head.Next.Next.Next.Next.Next.next = &List.ListNode{Data: 7}
	head = List.ReverseKList(head,2)
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}
