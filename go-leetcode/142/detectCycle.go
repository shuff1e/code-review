package main

import "fmt"

/*
142. 环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

3->2->0->-4
   ^      |
   |------|
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

1->2
^  |
|--|
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。




进阶：
你是否可以不用额外空间解决此题？

 */

// A：快慢指针，如果碰撞，说明有环

// 3距离2 一共k步
// slow进入2的时候，fast走了2k步

// fast在环中走了k步，slow放进入环
// fast落后slow loop_size - k步

// 每走一步，fast追上一步，因此碰撞时，slow走了loop_size - k步

// fast 放回头节点，一次走一步，slow继续走，两者都在走k步后在环的入口处碰撞

//3->1->2->0->4->5
//      ^        |
//      |--------|

// slow 3 1 2 0 4 5
// fast 3 2 4 2 4

// slow 4 5 2
// fast 3 1 2


func main() {
	head := &ListNode{Data: 1}
	head.Next = &ListNode{Data: 2}
	head.Next.Next = &ListNode{Data: 3}
	head.Next.Next.Next = &ListNode{Data: 4}
	head.Next.Next.Next.Next = &ListNode{Data: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Data: 6}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next.Next
	head = (*ListNode)(nil)

	fmt.Println(detectCycle(head))
}

type ListNode struct {
	Data int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	slow = getCollision(&slow,&fast)
	if slow == nil {
		return nil
	}

	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

func getCollision(slow,fast **ListNode) *ListNode {
	for (*fast) != nil && (*fast).Next != nil {
		*fast = (*fast).Next.Next
		*slow = (*slow).Next
		if *fast == *slow {
			return *fast
		}
	}
	return nil
}
