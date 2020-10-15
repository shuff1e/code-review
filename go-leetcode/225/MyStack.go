package main

/*
225. 用队列实现栈
使用队列实现栈的下列操作：

push(x) -- 元素 x 入栈
pop() -- 移除栈顶元素
top() -- 获取栈顶元素
empty() -- 返回栈是否为空
注意:

你只能使用队列的基本操作-- 也就是 push to back, peek/pop from front, size, 和 is empty 这些操作是合法的。
你所使用的语言也许不支持队列。 你可以使用 list 或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
你可以假设所有操作都是有效的（例如, 对一个空的栈不会调用 pop 或者 top 操作）。
 */

// 为了满足栈的特性，即最后入栈的元素最先出栈，在使用队列实现栈时，应满足队列前端的元素是最后入栈的元素。
//
//可以使用两个队列实现栈的操作，其中 queue1 用于存储栈内的元素，queue2作为入栈操作的辅助队列。
//
//
//入栈操作时，
//首先将元素入队到 queue2，
//然后将 queue1的全部元素依次出队并入队到 queue2，
//此时 queue2的前端的元素即为新入栈的元素，
//
//
//再将 queue1和 queue2 互换，则queue1的元素即为栈内的元素，queue1的前端和后端分别对应栈顶和栈底。
//
//
//由于每次入栈操作都确保 queue1 的前端元素为栈顶元素，因此出栈操作和获得栈顶元素操作都可以简单实现。
//出栈操作只需要移除 queue1 的前端元素并返回即可，获得栈顶元素操作只需要获得 queue1 的前端元素并返回即可（不移除元素）。
//
//由于 queue1 用于存储栈内的元素，判断栈是否为空时，只需要判断 queue1 是否为空即可。

type MyStack struct {
	queue1 []int
	queue2 []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{queue1: make([]int,0),queue2: make([]int,0)}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue2 = append(this.queue2,x)
	for len(this.queue1) > 0 {
		this.queue2 = append(this.queue2,this.queue1[0])
		this.queue1 = this.queue1[1:]
	}
	this.queue1,this.queue2 = this.queue2,this.queue1
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	result := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return result
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue1[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// 方法一使用了两个队列实现栈的操作，也可以使用一个队列实现栈的操作。
//
//使用一个队列时，为了满足栈的特性，即最后入栈的元素最先出栈，同样需要满足队列前端的元素是最后入栈的元素。
//
//入栈操作时，首先获得入栈前的元素个数 n，然后将元素入队到队列，再将队列中的前 n 个元素（即除了新入栈的元素之外的全部元素）依次出队并入队到队列，此时队列的前端的元素即为新入栈的元素，且队列的前端和后端分别对应栈顶和栈底。
//
//由于每次入栈操作都确保队列的前端元素为栈顶元素，因此出栈操作和获得栈顶元素操作都可以简单实现。出栈操作只需要移除队列的前端元素并返回即可，获得栈顶元素操作只需要获得队列的前端元素并返回即可（不移除元素）。
//
//由于队列用于存储栈内的元素，判断栈是否为空时，只需要判断队列是否为空即可。

/*

type MyStack struct {
    queue []int
}

func Constructor() (s MyStack) {
	return
}

func (s *MyStack) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

func (s *MyStack) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

func (s *MyStack) Top() int {
	return s.queue[0]
}

func (s *MyStack) Empty() bool {
	return len(s.queue) == 0
}

 */
