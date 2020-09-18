package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Node struct {
	key string
	value string
	prev *Node
	next *Node
}

type Cache struct {
	mp map[string]*Node
	mutex *sync.Mutex
	cap int
	head *Node
	tail *Node
}

func NewLRUCache(cap int) *Cache {
	return &Cache{make(map[string]*Node), &sync.Mutex{}, cap, nil, nil}
}

func (c *Cache) Get(key string) (value string,ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	temp,ok := c.mp[key]
	if !ok {
		return "",false
	}
	c.removeAndInsert(temp)
	return temp.value ,true
}

func (c *Cache) Put(key,value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if len(c.mp) == 0 {
		n := &Node{
			key:   key,
			value: value,
			prev:  nil,
			next:  nil,
		}
		c.mp[key] = n
		c.head = n
		c.tail = n
		return
	}

	if temp,ok := c.mp[key];ok {
		temp.value = value
		c.removeAndInsert(temp)
	} else {
		if len(c.mp) >= c.cap {
			c.delTail()
		}
		n := &Node{
			key:   key,
			value: value,
			prev:  nil,
			next:  nil,
		}
		c.mp[key] = n
		n.next = c.head
		c.head.prev = n
		c.head = n
	}
}

func (c *Cache) removeAndInsert(n *Node) {
	// remove
	if n == c.head {
		return
	} else if n == c.tail {
		c.tail = n.prev
		c.tail.next = nil
		n.prev = nil
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
		n.next = nil
		n.prev = nil
	}
	// insert
	n.next = c.head
	c.head.prev = n
	c.head = n
}

// put的时候超过了capacity
func (c *Cache) delTail() {
	temp := c.tail
	c.tail = temp.prev
	c.tail.next = nil
	temp.prev = nil
	delete(c.mp,temp.key)
}

func main() {
	cache := NewLRUCache(3)
	wg := sync.WaitGroup{}
	for i :=0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Put(strconv.Itoa(i),strconv.Itoa(i))
			fmt.Println(cache.Get(strconv.Itoa(i+2)))
			cache.Put(strconv.Itoa(i+1),strconv.Itoa(i+1))
			fmt.Println(cache.Get(strconv.Itoa(i+1)))
			cache.Put(strconv.Itoa(i+2),strconv.Itoa(i+2))
			fmt.Println(cache.Get(strconv.Itoa(i+2)))
			cache.Put(strconv.Itoa(i+3),strconv.Itoa(i+3))
			fmt.Println(cache.Get(strconv.Itoa(i+2)))
		}(i)
	}
	wg.Wait()
}