package main

import "fmt"

/*

328. 奇偶链表
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。

请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:

输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL
示例 2:

输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

 */

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}
	node.Next.Next.Next = &ListNode{Val: 4}
	node.Next.Next.Next.Next = &ListNode{Val: 5}
	//node.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	result := oddEvenList2(node)
	fmt.Println(result)
}

 type ListNode struct {
     Val int
     Next *ListNode
 }

func oddEvenList(head *ListNode) *ListNode {
	// 1 2 3 4 5
	// 1 3 5 2 4
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	prev1 := dummy1
	prev2 := dummy2

	temp := head
	for temp != nil && temp.Next != nil {
		prev1.Next = temp
		prev2.Next = temp.Next

		prev1 = prev1.Next
		prev2 = prev2.Next

		temp = temp.Next.Next
	}

	if temp != nil {
		prev1.Next = temp
		prev1 = prev1.Next
	}

	prev1.Next = nil
	prev2.Next = nil

	prev1.Next = dummy2.Next
	return dummy1.Next

}

func oddEvenList2(head *ListNode) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	temp1 := dummy1
	temp2 := dummy2

	temp := head

	for temp != nil {
		temp1.Next = temp
		temp1 = temp1.Next

		temp = temp.Next
		if temp != nil {
			temp2.Next = temp
			temp2 = temp2.Next
			temp = temp.Next
		}
	}

	temp1.Next = nil
	temp2.Next = nil

	temp1.Next = dummy2.Next
	return dummy1.Next
}