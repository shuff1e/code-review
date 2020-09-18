package main

import "fmt"

// 123 456 789 10
// 翻转成 321 654 987 10

type ListNode struct {
	data int
	next *ListNode
}

func reverseKList(node *ListNode,k int) *ListNode {
	if node == nil {
		return nil
	}
	head := node
	temp := k-1
	for temp > 0 && node !=nil {
		temp--
		node = node.next
	}
	if temp > 0 || node == nil {
		return head
	}
	next := node.next
	node.next = nil
	left,right := reverseList(head)
	right.next = reverseKList(next,k)
	return left
}

func reverseList(node *ListNode) (*ListNode,*ListNode) {
	tail := node
	cur := node
	var prev *ListNode
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev,tail
}
func main() {
	head := &ListNode{data: 1}
	head.next = &ListNode{data: 2}
	head.next.next = &ListNode{data: 3}
	head.next.next.next = &ListNode{data: 4}
	head.next.next.next.next = &ListNode{data: 5}
	head.next.next.next.next.next = &ListNode{data: 6}
	//head.next.next.next.next.next.next = &ListNode{data: 7}
	head = reverseKList(head,2)
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
