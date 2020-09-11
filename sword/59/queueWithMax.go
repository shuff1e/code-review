package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

// 59（二）：队列的最大值
// 题目：给定一个数组和滑动窗口的大小，请找出所有滑动窗口里的最大值。例如，
// 如果输入数组{2, 3, 4, 2, 6, 2, 5, 1}及滑动窗口的大小3，那么一共存在6个
// 滑动窗口，它们的最大值分别为{4, 4, 6, 6, 6, 5}，

// A：两个queue，一个存放数据，
// 一个存放迄今为止的最大值
// 例如，1，3，2，1
// 1 入queue1时，最小值是1
// 3 入queue1时，1不可能是最小值，将1remove出queue2去
// 2 入queue1时，2可能是最小值，将2入queue2
// 3出queue时，queue2的头部节点如果是3，则将3也出queue2

type myQueue struct {
	queue1 *arraylist.List
	queue2 *arraylist.List
	comparator Comparator
}

func New() *myQueue {
	return &myQueue{
		queue1: arraylist.New(),
		queue2: arraylist.New(),
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
	// result > 0,o1 > o2
	// result < 0, o1 < o2
	// result = 0,o1 == o2
	compare(o1,o2 interface{}) (result int)
}

func (queue *myQueue) AddLast(v interface{}) {
	queue.queue1.Add(v)
	for !queue.queue2.Empty() {
		temp,_ := queue.queue2.Get(queue.queue2.Size()-1)
		if queue.comparator.compare(v,temp) > 0 {
			queue.queue2.Remove(queue.queue2.Size()-1)
		} else {
			queue.queue2.Add(v)
			break
		}
	}
	if queue.queue2.Empty() {
		queue.queue2.Add(v)
	}
}

func (queue *myQueue) Max() (interface{},bool) {
	return queue.queue2.Get(0)
}

func (queue *myQueue) RemoveFirst() (interface{},bool) {
	v1,ok1 := queue.queue1.Get(0)
	v2,ok2 := queue.queue2.Get(0)
	if ok1 && ok2 {
		if queue.comparator.compare(v1,v2) == 0 {
			queue.queue2.Remove(0)
		}
	}
	queue.queue1.Remove(0)
	return v1,ok1
}

func FindMax(arr []int,k int) []int {
	if k == 0 || len(arr) == 0{
		return []int{}
	}
	if len(arr) < k {
		return []int{}
	}

	result := make([]int,len(arr)-k+1)
	queue := New()
	for i := 0;i<k;i++ {
		queue.AddLast(arr[i])
	}
	temp,_ := queue.Max()
	max := temp.(int)
	resultIndex := 0
	result[resultIndex] = max
	resultIndex ++

	for i := k;i<len(arr);i++ {
		queue.AddLast(arr[i])
		queue.RemoveFirst()
		temp,_ := queue.Max()
		max := temp.(int)
		result[resultIndex] = max
		resultIndex ++
	}

	return result
}

func Test(name string,queue *myQueue,expected int) {
	max,ok := queue.Max()
	if ok {
		fmt.Println(name,max.(int) == expected)
	} else {
		fmt.Println("invalid")
	}
}

func main() {
	// Test queue
	queue := New()
	// {2}
	queue.AddLast(2);
	Test("Test1", queue, 2);

	// {2, 3}
	queue.AddLast(3);
	Test("Test2", queue, 3);

	// {2, 3, 4}
	queue.AddLast(4);
	Test("Test3", queue, 4);

	// {2, 3, 4, 2}
	queue.AddLast(2);
	Test("Test4", queue, 4);

	// {3, 4, 2}
	queue.RemoveFirst();
	Test("Test5", queue, 4);

	// {4, 2}
	queue.RemoveFirst();
	Test("Test6", queue, 4);

	// {2}
	queue.RemoveFirst();
	Test("Test7", queue, 2);

	// {2, 6}
	queue.AddLast(6);
	Test("Test8", queue, 6);

	// {2, 6, 2}
	queue.AddLast(2);
	Test("Test9", queue, 6);

	// {2, 6, 2, 5}
	queue.AddLast(5);
	Test("Test9", queue, 6);

	// {6, 2, 5}
	queue.RemoveFirst();
	Test("Test10", queue, 6);

	// {2, 5}
	queue.RemoveFirst();
	Test("Test11", queue, 5);

	// {5}
	queue.RemoveFirst();
	Test("Test12", queue, 5);

	// {5, 1}
	queue.AddLast(1);
	Test("Test13", queue, 5);


	// Test Find
	Test2("Test1",[]int{ 2, 3, 4, 2, 6, 2, 5, 1 },[]int{ 4, 4, 6, 6, 6, 5 },3)
	Test2("Test2",[]int{ 1, 3, -1, -3, 5, 3, 6, 7 },[]int{ 3, 3, 5, 5, 6, 7 },3)
	Test2("Test3",[]int{ 1, 3, 5, 7, 9, 11, 13, 15 },[]int{ 7, 9, 11, 13, 15 },4)
	Test2("Test4",[]int{ 16, 14, 12, 10, 8, 6, 4 },[]int{ 16, 14, 12 },5)
	Test2("Test5",[]int{ 10, 14, 12, 11 },[]int{ 10, 14, 12, 11 },1)
	Test2("Test6",[]int{ 10, 14, 12, 11 },[]int{ 14 },4)
	Test2("Test7",[]int{ 10, 14, 12, 11 },[]int{},0)
	Test2("Test8",[]int{ 10, 14, 12, 11 },[]int{},5)
	Test2("Test9",[]int{},[]int{},5)
}

func Test2(name string,arr,expected []int,k int) {
	result := FindMax(arr,k)
	//fmt.Println(result)
	if len(result) != len(expected) {
		fmt.Println(name,false)
		return
	}
	for i := 0;i<len(result);i++ {
		if result[i] != expected[i] {
			fmt.Println(name,false)
			return
		}
	}
	fmt.Println(name,true)
}