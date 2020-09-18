package main

import (
	"algorithm/common/help"
	"fmt"
)

// 删除单链表中的倒数第K个节点

type myNode struct {
	value int
	next *myNode
}

func delK(node *myNode,k int) *myNode {
	fast := node
	// 对应k>0
	for i:=0;i<k;i++ {
		if fast == nil {
			return node
		}
		fast = fast.next
	}
	// 对应 k==0
	if fast == nil {
		return node.next
	}
	// 对应k<0
	slow := node
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next
	return node
}

func generateNList(n int) *myNode {
	slice := help.GenerateSlice(n)
	head := &myNode{}
	temp := head
	for _,v := range slice {
		temp.next = &myNode{value: v}
		temp = temp.next
	}
	return head.next
}

func main() {
	node := generateNList(4)
	temp := node
	for temp != nil {
		fmt.Print(temp.value)
		fmt.Print(" ")
		temp=temp.next
	}
	fmt.Println("------")
	//result := delK(node,5)
	result := removeLastKthNode(node,5)
	temp = result
	for temp != nil {
		fmt.Print(temp.value)
		fmt.Print(" ")
		temp=temp.next
	}
}

func removeLastKthNode(node *myNode,k int) *myNode {
	if node == nil || k < 1 {
		return node
	}
	cur := node
	for cur!=nil{
		cur = cur.next
		k--
	}
	// k=0
	if k == 0 {
		return node.next
	}
	if k < 0 {
		cur = node
		k++
		for k!=0 {
			cur = cur.next
			k++
		}
		cur.next = cur.next.next
	}
	return node
}