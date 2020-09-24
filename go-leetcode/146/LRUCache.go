package main

import "fmt"

/*
146. LRU缓存机制
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 Get 和 写入数据 Put 。

获取数据 Get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
写入数据 Put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。
当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache( 2  **缓存容量**  );

cache.Put(1, 1);
cache.Put(2, 2);
cache.Get(1);       // 返回  1
cache.Put(3, 3);    // 该操作会使得关键字 2 作废
cache.Get(2);       // 返回 -1 (未找到)
cache.Put(4, 4);    // 该操作会使得关键字 1 作废
cache.Get(1);       // 返回 -1 (未找到)
cache.Get(3);       // 返回  3
cache.Get(4);       // 返回  4
*/

// A：
// Get操作返回数据，可以使用map
// 并把数据放在链表头部，需要删除节点，再添加到头部

// set操作更新数据，删除节点，再添加到头部
// set操作插入数据到头部，或者删除尾部数据之后，插入到头部

// map + 双向链表

func main() {
	cache := Constructor(1)
	cache.Put(2,1)
	fmt.Println(cache.Get(2))

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
	m map[int]*listNode
	capacity int
	head *listNode
	tail *listNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m: make(map[int]*listNode,capacity),
		capacity: capacity,
		head : nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if node,ok := this.m[key];!ok {
		return -1
	} else {
		this.removeNode(node)
		this.insertAtHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int)  {
	// 同时操作map和list

	// 是否有该key，是更新key还是插入key
	// 插入的时候是空的还是满的

	// 更新还是插入
	if node,ok := this.m[key];ok {
		node.value = value
		this.removeNode(node)
		this.insertAtHead(node)
	} else if this.head == nil {
		// 插入空的
		node := &listNode{
			key: key,
			value: value,
		}
		this.head = node
		this.tail = node
		this.m[key] = node
		return
	} else if len(this.m) == this.capacity {
		// 插入满的
		// 先删除尾部
		delete(this.m,this.tail.key)
		this.removeNode(this.tail)

		//再插入
		node := &listNode{
			key: key,
			value: value,
		}
		this.m[key] = node
		this.insertAtHead(node)
	} else {
		node := &listNode{
			key: key,
			value: value,
		}
		this.m[key] = node
		this.insertAtHead(node)
	}
}

func (this *LRUCache) removeNode(node *listNode) {
	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else if this.head == node {
		next := this.head.next
		this.head.next = nil
		next.prev = nil
		this.head = next
	} else if this.tail == node {
		prev := this.tail.prev
		this.tail.prev = nil
		prev.next = nil
		this.tail = prev
	} else {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev

		node.prev = nil
		node.next = nil
	}
}

func (this *LRUCache) insertAtHead(node *listNode) {
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
}