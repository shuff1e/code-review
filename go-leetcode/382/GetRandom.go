package main

import (
	"math/rand"
	"time"
)

/*

382. 链表随机节点
给定一个单链表，随机选择链表的一个节点，并返回相应的节点值。保证每个节点被选的概率一样。

进阶:
如果链表十分大且长度未知，如何解决这个问题？你能否使用常数级空间复杂度实现？

示例:

// 初始化一个单链表 [1,2,3].
ListNode head = new ListNode(1);
head.next = new ListNode(2);
head.next.next = new ListNode(3);
Solution solution = new Solution(head);

// getRandom()方法应随机返回1,2,3中的一个，保证每个元素被返回的概率相等。
solution.getRandom();

 */

/*

蓄水池算法

遍历到i时，

进去的概率为 m/i

当i < m，留下的概率为

(1-m/(m+1) * 1/m) * ( 1 - m/(m+2) * 1/m)

当i > m，留下的概率为
如果i这个数字被换出，首先i+1需要换入，换入的概率为 m / (i+1)，
然后i需要被换出，概率为1/m
1-换出的概率

(1- m/(i+1) * 1/m ) * (1- m/(i+2) * 1/m )

i当时换入的概率为 m/i
m/i * i/(i+1)

 */

type ListNode struct {
    Val int
    Next *ListNode
}

type Solution struct {
	head *ListNode
}


/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	if this.head == nil {
		return -1
	}
	rand.Seed(time.Now().UnixNano())

	cur := this.head
	result := cur.Val
	i := 1

	cur = cur.Next
	for cur != nil {
		i++
		if rand.Int63n(int64(i)) == 0 {
			result = cur.Val
		}
		cur = cur.Next
	}
	return result
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */