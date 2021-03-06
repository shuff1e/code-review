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

/*

抽取 m 个
抽取 m 个如何保证等概率，看第一的题解分析，我这里就不讲了


//功能：实现从 [1, n] 中等概率的抽取 m 个数字

public static int[] h(int[] data, int m) {
    int[] reservoir = new int[m];
    // 1. 初始化：选取 data 前 m 个元素放入蓄水池 res 中
    for (int i = 0; i < m; i++) {
        res[i] = data[i];
    }
    Random random = new Random();
    // 2. 以 m/k 的概率选择 第 k 个数字
    for (int i = m; i < data.length; i++) {
        int d = random.nextInt(i + 1);
        //
        //    3. 如果随机整数落在 [0, m-1] 范围内，则替换蓄水池中的元素
        //    对于当前元素来说，它有可能替换掉蓄水池内部的任意一个位置的元素

        //    因为 random.nextInt(i + 1) 表示从 [0, i + 1] 这些位置中选择一个位置来作为当前元素的替换位置，
        //    如果选中蓄水池的位置，表示当前元素能够进入蓄水池，如果没选中，表示当前元素失去了进入蓄水池的机会
        //    所以蓄水池每个位置被选中的概率为 1 / i+1，而总共有 m 个位置，
        //    所以对于当前元素来说，它有 m / i+1 的概率能够替换到蓄水池中
        //    当前节点抽到的位置为 d，如果 d < m，表示当前位置能够进入蓄水池，并且替换掉 d 位置的元素

        if (d < m)
            res[d] = data[i];
    }
    return res;
}

 */