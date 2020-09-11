package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 30：包含min函数的栈
// 题目：定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的min
// 函数。在该栈中，调用min、push及pop的时间复杂度都是O(1)。

type myStack struct {
	// stack1 存放数据
	stack1 *linkedliststack.Stack
	// stack2 存放迄今为止，的最小值
	stack2 *linkedliststack.Stack
	comparator Comparator
}

func New() *myStack {
	return &myStack{
		stack1: linkedliststack.New(),
		stack2: linkedliststack.New(),
		comparator: compareFunc(func(o1, o2 interface{}) int {
			temp1 := o1.(int)
			temp2 := o2.(int)
			return temp1 - temp2
		}),
	}
}

// like http.HandleFunc
// compareFunc实现了Comparator接口
// 一般的函数再转换为compareFunc
type compareFunc func(o1,o2 interface{}) int

func (f compareFunc) compare(o1,o2 interface{}) int {
	return f(o1,o2)
}

type Comparator interface {
	// result >= 0,o1 >= o2
	// result < 0, o1 < o2
	compare(o1,o2 interface{}) (result int)
}

func (stack *myStack) push(v interface{}) {
	if stack.stack1.Empty() {
		stack.stack1.Push(v)
		stack.stack2.Push(v)
		return
	}
	temp,_ := stack.stack2.Peek()
	if stack.comparator.compare(temp,v) > 0 {
		stack.stack1.Push(v)
		stack.stack2.Push(v)
	} else {
		stack.stack1.Push(v)
		stack.stack2.Push(temp)
	}
}

func (stack *myStack) pop() (interface{},bool) {
	stack.stack2.Pop()
	return stack.stack1.Pop()
}

func (stack *myStack) min() (interface{},bool) {
	return stack.stack2.Peek()
}

func main() {
	stack := New()
	stack.push(3);
	Test("Test1", stack, 3);

	stack.push(4);
	Test("Test2", stack, 3);

	stack.push(2);
	Test("Test3", stack, 2);

	stack.push(3);
	Test("Test4", stack, 2);

	stack.pop();
	Test("Test5", stack, 2);

	stack.pop();
	Test("Test6", stack, 3);

	stack.pop();
	Test("Test7", stack, 3);

	stack.push(0);
	Test("Test8", stack, 0);
}

func Test(name string,stack *myStack,expected int) {
	v,ok := stack.min()
	if ok {
		fmt.Println(name,v == expected)
	} else {
		fmt.Println("invalid")
	}
}