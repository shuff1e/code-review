package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 9：用两个栈实现队列
// 题目：用两个栈实现一个队列。队列的声明如下，请实现它的两个函数appendTail
// 和deleteHead，分别完成在队列尾部插入结点和在队列头部删除结点的功能。

type myQueue struct {
	stack1 *linkedliststack.Stack
	stack2 *linkedliststack.Stack
}

func New() *myQueue {
	return &myQueue{
		stack1: linkedliststack.New(),
		stack2: linkedliststack.New(),
	}
}
func (q *myQueue) appendTail(v interface{}) {
	q.stack1.Push(v)
}

func (q *myQueue) deleteHead() (interface{},bool) {
	if q.stack2.Empty() {
		for !q.stack1.Empty() {
			v,_ := q.stack1.Pop()
			q.stack2.Push(v)
		}
	}
	return q.stack2.Pop()
}

func main() {
	queue := New()
	queue.appendTail("a")
	queue.appendTail("b")
	queue.appendTail("c")

	fmt.Println(queue.deleteHead())
	fmt.Println(queue.deleteHead())
	queue.appendTail("d")
	queue.appendTail("e")
	fmt.Println(queue.deleteHead())
	fmt.Println(queue.deleteHead())
	queue.appendTail("f")
	fmt.Println(queue.deleteHead())
}