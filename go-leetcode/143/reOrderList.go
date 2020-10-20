package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 例子一
// 给定链表 1->2->3->4, 重新排列为 1->4->2->3.

// 例子二
// 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	stack := prepareStack(head)
	helper(head,stack)
}

func helper(node *ListNode,stack *linkedliststack.Stack) *ListNode {
	nextHead := node.Next
	temp,ok := stack.Pop()
	if !ok {
		panic("fuck")
	}
	v := temp.(*ListNode)
	if v == node {
		node.Next = nil
		return node
	}

	if v == nextHead {
		nextHead.Next = nil
		return node
	}
	node.Next = v
	v.Next = helper(nextHead,stack)
	return node
}

func prepareStack(head *ListNode) *linkedliststack.Stack {
	stack := linkedliststack.New()
	for head != nil {
		stack.Push(head)
		head = head.Next
	}
	return stack
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}

func Test5() {
	head := (*ListNode)(nil)
	printList(head)
	reorderListBetter(head)
	printList(head)
}

func Test4() {
	head := &ListNode{
		Val: 1,
	}
	printList(head)
	reorderListBetter(head)
	printList(head)

}

func Test3() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	printList(head)
	reorderListBetter(head)
	printList(head)
}

func Test2() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}
	printList(head)
	reorderListBetter(head)
	printList(head)
}
func Test1() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}
	printList(head)
	reorderListBetter(head)
	printList(head)
}

func Test6() {
	head := &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	printList(head)
	reorderListBetter(head)
	printList(head)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
	fmt.Println()
}

// 此题目为2019年计算机统考408真题
//
// 先观察 L：L0→L1→…→Ln-1→Ln ，
// 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
// 发现是由L摘取第一个元素，再摘取倒数第一个元素....依次合并成的。
// 为了方便链表后半段取元素，需要先将L后半段原地逆置，否则每取最后一个节点都需要遍历一次链表。
// 【408真题题目要求空间复杂度为O（1），不能借助栈】
//
// 步骤
//
//1、先找处链表L的中间结点，为此设置两个指针p和q，
// 指针p每次走一步，指针q每次走两步，当指针q到达链尾时，指针p正好在链表的中间结点；
//2、然后将L的后半段结点原地逆置；
//3、从单链表前后两段中 依次各取一个结点，按要求重排；
//

func reorderListBetter(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	mid := findMid(head)
	other := reverseList(mid.Next)
	mid.Next = nil

	for head != nil && other != nil {
		headNext := head.Next
		otherNext := other.Next

		head.Next = other
		other.Next = headNext

		head = headNext
		other = otherNext
	}
}

func findMid(head *ListNode) *ListNode {
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
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
	// one more line
}