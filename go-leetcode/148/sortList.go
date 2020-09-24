package main

import "fmt"

/*
148. 排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

 */

/*
解答一：归并排序（递归法）
题目要求时间空间复杂度分别为O(nlogn)O(nlogn)和O(1)O(1)，根据时间复杂度我们自然想到二分法，从而联想到归并排序；

对数组做归并排序的空间复杂度为 O(n)O(n)，分别由新开辟数组O(n)O(n)和递归函数调用O(logn)O(logn)组成，而根据链表特性：

数组额外空间：链表可以通过修改引用来更改节点顺序，无需像数组一样开辟额外空间；
递归额外空间：递归调用函数将带来O(logn)O(logn)的空间复杂度，因此若希望达到O(1)O(1)空间复杂度，则不能使用递归。
通过递归实现链表归并排序，有以下两个环节：

分割 cut 环节： 找到当前链表中点，并从中点将链表断开（以便在下次递归 cut 时，链表片段拥有正确边界）；
我们使用 fast,slow 快慢双指针法，奇数个节点找到中点，偶数个节点找到中心左边的节点。
找到中点 slow 后，执行 slow.next = None 将链表切断。
递归分割时，输入当前链表左端点 head 和中心节点 slow 的下一个节点 tmp(因为链表是从 slow 切断的)。
cut 递归终止条件： 当head.next == None时，说明只有一个节点了，直接返回此节点。
合并 merge 环节： 将两个排序链表合并，转化为一个排序链表。
双指针法合并，建立辅助ListNode h 作为头部。
设置两指针 left, right 分别指向两链表头部，比较两指针处节点值大小，由小到大加入合并链表头部，指针交替前进，直至添加完两个链表。
返回辅助ListNode h 作为头部的下个节点 h.next。
时间复杂度 O(l + r)，l, r 分别代表两个链表长度。
当题目输入的 head == None 时，直接返回None。
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	printList(head)
	head = sortList(head)
	printList(head)
}

func printList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val,"->")
		root = root.Next
	}
	fmt.Println("nil")
}

func sortList(head *ListNode) *ListNode {
	// boundary
	if head == nil || head.Next == nil {
		return head
	}

	// split
	slow,fast := head,head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil

	// sorted left and right
	left := sortList(head)
	right := sortList(next)

	// merge
	result := &ListNode{
		Val: 0,
	}
	temp := result
	for left != nil && right != nil {
		if left.Val <= right.Val {
			temp.Next = left
			left = left.Next
			temp = temp.Next
		} else {
			temp.Next = right
			right = right.Next
			temp = temp.Next
		}
	}

	if left != nil {
		temp.Next = left
	}
	if right != nil {
		temp.Next = right
	}
	return result.Next
}