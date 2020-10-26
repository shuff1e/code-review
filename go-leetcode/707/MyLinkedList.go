package main

import "fmt"

/*

707. 设计链表
设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针/引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。

在链表类中实现这些功能：

get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。


示例：

MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
linkedList.get(1);            //返回2
linkedList.deleteAtIndex(1);  //现在链表是1-> 3
linkedList.get(1);            //返回3


提示：

所有val值都在 [1, 1000] 之内。
操作次数将在  [1, 1000] 之内。
请不要使用内置的 LinkedList 库。

 */

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1,2)   //链表变为1-> 2-> 3
	fmt.Println(linkedList.Get(1))            //返回2
	linkedList.DeleteAtIndex(1)  //现在链表是1-> 3
	fmt.Println(linkedList.Get(1))            //返回3
}

type ListNode struct {
	Val int
	Next *ListNode
	Prev *ListNode
}

type MyLinkedList struct {
	head *ListNode
	tail *ListNode
	size int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index > this.size - 1 {
		return -1
	}
	if index > this.size / 2 {
		// 正数index
		// 0 1 2 3 4 5
		// index = 2
		// 6-2=4
		// 1 2 3 4
		// 5 4 3 2
		temp := this.tail
		count := 1

		for count < this.size - index {
			temp = temp.Prev
			count ++
		}
		return temp.Val
	}
	count := 0
	temp := this.head
	for count < index {
		temp = temp.Next
		count ++
	}
	return temp.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	temp := &ListNode{Val: val}

	if this.size == 0 {
		this.head = temp
		this.tail = temp
	} else {
		temp.Next = this.head
		this.head.Prev = temp

		this.head = temp
	}
	this.size ++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	temp := &ListNode{Val: val}
	if this.size == 0 {
		this.head = temp
		this.tail = temp
	} else {
		this.tail.Next = temp
		temp.Prev = this.tail

		this.tail = temp
	}
	this.size ++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.size {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}

	if index > this.size / 2 {
		// 正数index
		// 0 1 2 3 4 5
		// index = 2
		// 6-2=4
		// 1 2 3 4
		// 5 4 3 2
		temp := this.tail
		count := 1

		for count < this.size - index {
			temp = temp.Prev
			count ++
		}

		prev := temp.Prev
		tempNode := &ListNode{Val: val}

		prev.Next  = tempNode
		tempNode.Prev = prev

		tempNode.Next = temp
		temp.Prev = tempNode

		this.size ++
		return
	}
	count := 0
	temp := this.head
	for count < index {
		temp = temp.Next
		count ++
	}
	tempNode := &ListNode{Val: val}
	prev := temp.Prev

	prev.Next  = tempNode
	tempNode.Prev = prev

	tempNode.Next = temp
	temp.Prev = tempNode

	this.size ++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.size {
		return
	}
	if this.size == 1 {
		this.head = nil
		this.tail = nil
		this.size = 0
		return
	}

	if index == this.size - 1 {
		prev := this.tail.Prev
		prev.Next = nil
		this.tail.Prev = nil
		this.tail = prev

		this.size --
		return
	}

	if index == 0 {
		next := this.head.Next
		next.Prev = nil
		this.head.Next = nil
		this.head = next

		this.size --
		return
	}

	if index > this.size / 2 {
		temp := this.tail
		count := 1

		for count < this.size - index {
			temp = temp.Prev
			count ++
		}
		prev := temp.Prev
		next := temp.Next
		prev.Next = next
		next.Prev = prev
		this.size --
		return
	}
	count := 0
	temp := this.head
	for count < index {
		temp = temp.Next
		count ++
	}
	prev := temp.Prev
	next := temp.Next
	prev.Next = next
	next.Prev = prev
	this.size --
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */