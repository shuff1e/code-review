package main

import (
	"algorithm/common/help"
	"fmt"
)

// 用两个stack实现一个queue
// add poll peek

type myQueue struct {
	stackPush *help.Stack
	stackPop *help.Stack
}

func NewMyQueue() *myQueue {
	return &myQueue{
		help.NewStack(),
		help.NewStack(),
	}
}

func (q *myQueue) add(v int) {
	q.stackPush.Push(v)
}

func (q *myQueue) poll() (int,error) {
	if q.stackPop.Length() == 0 {
		for q.stackPush.Length() != 0 {
			v,_ := q.stackPush.Pop()
			q.stackPop.Push(v)
		}
	}
	return q.stackPop.Pop()
}

func (q *myQueue) peek() (int,error) {
	if q.stackPop.Length() == 0 {
		for q.stackPush.Length() != 0 {
			v,_ := q.stackPush.Pop()
			q.stackPop.Push(v)
		}
	}
	return q.stackPop.Peek()
}

func main() {
	q := NewMyQueue()

	slice := help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)

	for _,v := range slice {
		q.add(v)
	}

	for i := 0;i<10;i++ {
		fmt.Println(q.poll())
	}

	slice = help.GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)

	for _,v := range slice {
		q.add(v)
	}
	for i := 0;i<30;i++ {
		fmt.Println(q.poll())
	}
}
