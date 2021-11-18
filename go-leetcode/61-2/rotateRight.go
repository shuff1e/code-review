package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

//61. 旋转链表
//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

// 先整体旋转链表
// 再从k的位置砍成两半
// 分别旋转这两部分

// 1 2 3 4 5
// k = 2
// 4 5 1 2 3

// 5 4 3 2 1
// 4 5 1 2 3

func main() {
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    head.Next.Next = &ListNode{Val: 3}
    head.Next.Next.Next = &ListNode{Val: 4}
    head.Next.Next.Next.Next = &ListNode{Val: 5}
    head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
    head = rotateRight(head,20)
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }

}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    head,_,size := reverseList(head)
    temp := head
    k = k % size
    if k == 0 {
        head,_,size = reverseList(head)
        return head
    }

    for i:= 0;i<k-1;i++ {
        temp = temp.Next
    }

    next := temp.Next
    temp.Next = nil

    h1,t1,_ := reverseList(head)
    h2,_,_ := reverseList(next)
    t1.Next = h2
    return h1

}

func reverseList(node *ListNode) (head,tail *ListNode,size int) {
    tail = node
    prev := (*ListNode)(nil)

    for node != nil {
        size ++
        next := node.Next
        node.Next = prev

        prev = node
        node = next
    }
    return prev,tail,size
}
