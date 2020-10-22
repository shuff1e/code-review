package main

import (
	"fmt"
)

/*
1116. 打印零与奇偶数
假设有这么一个类：

class ZeroEvenOdd {
public ZeroEvenOdd(int n) { ... }      // 构造函数
public void zero(printNumber) { ... }  // 仅打印出 0
public void even(printNumber) { ... }  // 仅打印出 偶数
public void odd(printNumber) { ... }   // 仅打印出 奇数
}
相同的一个 ZeroEvenOdd 类实例将会传递给三个不同的线程：

线程 A 将调用 zero()，它只输出 0 。
线程 B 将调用 even()，它只输出偶数。
线程 C 将调用 odd()，它只输出奇数。
每个线程都有一个 printNumber 方法来输出一个整数。请修改给出的代码以输出整数序列 010203040506... ，其中序列的长度必须为 2n。



示例 1：

输入：n = 2
输出："0102"
说明：三条线程异步执行，其中一个调用 zero()，另一个线程调用 even()，最后一个线程调用odd()。正确的输出为 "0102"。
示例 2：

输入：n = 5
输出："0102030405"
*/

// zero->even->zero->odd

type Foo struct {
	index int
	n int
}

// 两个环状的
func (f *Foo) zero(left1,right1,left2,right2,exit chan struct{}) {
	for i := 0;i<f.n;i++ {
			fmt.Print("0->")
			right1 <- struct{}{}
			<- left1

			fmt.Print("0->")
			right2 <- struct{}{}
			<- left2
	}
	close(exit)
}

func (f *Foo) even(left,right chan struct{}) {
	for i := 0;i<f.n;i++ {
		<-left
		fmt.Print(2*(i+1)-1,"->")
		right <- struct{}{}
	}
}

func (f *Foo) odd(left,right chan struct{}) {
	for i := 0;i<f.n;i++ {
		<- left
		fmt.Print(2*(i+1),"->")
		right <- struct {}{}
	}
}

func main() {
	f := &Foo{index: 0,
		n: 5}
	exit := make(chan struct{})

	left1,right1,left2,right2 := make(chan struct{}),make(chan struct{}),make(chan struct{}),make(chan struct{})

	go f.zero(left1,right1,left2,right2,exit)

	go f.even(right1,left1)

	go f.odd(right2,left2)

	<- exit
}