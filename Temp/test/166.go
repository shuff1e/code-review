package main

import (
	"fmt"
)

func main() {
	cache := Constructor(2)
	//cache.Put(2,1)
	//fmt.Println(cache.Get(2))

	cache.Put(1, 1);
	cache.Put(2, 2);
	fmt.Println(cache.Get(1));       // 返回  1
	cache.Put(3, 3);    // 该操作会使得关键字 2 作废
	fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
	cache.Put(4, 4);    // 该操作会使得关键字 1 作废
	fmt.Println(cache.Get(1));       // 返回 -1 (未找到)
	fmt.Println(cache.Get(3));       // 返回  3
	fmt.Println(cache.Get(4));       // 返回  4
}

type listNode struct {
	key int
	value int
	prev *listNode
	next *listNode
}

type LRUCache struct {
	head *listNode
	tail *listNode
	capacity int
	dict map[int]*listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head: nil,
		tail: nil,
		capacity: capacity,
		dict: make(map[int]*listNode,0),
	}
}

// 读不到，返回
// 读到，删除，并放到头部
func (this *LRUCache) Get(key int) int {
	if node,ok := this.dict[key];ok {
		result := node.value
		this.removeNode(node)
		this.insertAtHead(node)
		return result
	} else {
		return -1
	}
}

// 更新，
// 删除，放到头部

// 插入
// 空的，满的，正常
// 操作map
func (this *LRUCache) Put(key int, value int)  {
	if node,ok := this.dict[key];ok {
		node.value = value
		this.removeNode(node)
		this.insertAtHead(node)
	} else if this.head == nil {
		temp := &listNode{key: key,value: value}
		this.head,this.tail = temp,temp
		this.dict[key] = temp
	} else if len(this.dict) == this.capacity {
		temp := &listNode{key: key,value: value}

		delete(this.dict,this.tail.key)
		this.removeNode(this.tail)

		this.dict[key] = temp
		this.insertAtHead(temp)
	} else {
		temp := &listNode{key: key,value: value}
		this.dict[key] = temp
		this.insertAtHead(temp)
	}
}

func (this *LRUCache) removeNode(node *listNode) {
	if this.head == this.tail {
		this.head,this.tail = nil,nil
	} else if this.head == node {
		next := this.head.next
		this.head.next = nil
		next.prev = nil
		this.head = next
	} else if this.tail == node {
		prev := this.tail.prev

		prev.next = nil
		this.tail.prev = nil

		this.tail = prev
	} else {
		prev := node.prev
		next := node.next

		node.prev = nil
		node.next = nil

		prev.next = next
		next.prev = prev
	}
}

func (this *LRUCache) insertAtHead(node *listNode) {
	if this.head == nil {
		this.head,this.tail = node,node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
}
