package main
/*
155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。
 */

// 两个栈
type MinStack struct {
	data []int
	min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min: []int{},
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.data) == 0 {
		this.data = append(this.data,x)
		this.min = append(this.min,x)
	} else {
		top := this.GetMin()
		if x < top {
			this.min = append(this.min,x)
		} else {
			this.min = append(this.min,top)
		}
		this.data = append(this.data,x)
	}
}


func (this *MinStack) Pop()  {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]
}


func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
