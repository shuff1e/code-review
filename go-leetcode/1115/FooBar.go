package main

import "fmt"

/*
1115. 交替打印FooBar
我们提供一个类：

class FooBar {
  public void foo() {
    for (int i = 0; i < n; i++) {
      print("foo");
    }
  }

public void bar() {
  for (int i = 0; i < n; i++) {
      print("bar");
    }
  }

}
两个不同的线程将会共用一个 FooBar 实例。其中一个线程将会调用 foo() 方法，另一个线程将会调用 bar() 方法。

请设计修改程序，以确保 "foobar" 被输出 n 次。



示例 1:

输入: n = 1
输出: "foobar"
解释: 这里有两个线程被异步启动。其中一个调用 foo() 方法, 另一个调用 bar() 方法，"foobar" 将被输出一次。
示例 2:

输入: n = 2
输出: "foobarfoobar"
解释: "foobar" 将被输出两次。
 */


// https://studygolang.com/articles/12242

// ping pong 模型

// daisy chain -> daisy chain 组成环 -> daisy chain 组成多个环

func foo(in,out chan struct{},n int) {
	for i:= 0;i<n;i++ {
		<- in
		fmt.Print("foo->")
		out <- struct{}{}
	}
}

func bar(in,out,Exit chan struct{},n int) {
	defer func() {
		close(Exit)
	}()
	for i:=0;i<n;i++{
		<- in
		fmt.Print("bar->")
		out <- struct{}{}
	}
}

func Test1(n int) {
	// 最终球会落在一个人的手里，所以in要有一个buffer
	in := make(chan struct{},1)
	out := make(chan struct{})
	Exit := make(chan struct{})

	go foo(in,out,n)
	go bar(out,in,Exit,n)

	// 裁判发球
	in <- struct{}{}

	// 阻塞主线程
	<- Exit
}

func main() {
	n := 5
	Test1(n)
	fmt.Println()
	Test2(n)
}

func Test2(n int) {
	ch := make(chan struct{})
	exit := make(chan struct{})
	go foo2(ch,n)
	go bar2(ch,exit,n)
	<- exit
}

func foo2(ch chan struct{},n int) {
	for i := 0;i<n;i++ {
		<- ch
		fmt.Print("foo->")
		ch <- struct{}{}
	}
}

func bar2(ch,exit chan struct{},n int) {
	defer func() {
		close(ch)
		close(exit)
	}()
	for i := 0;i<n;i++ {
		ch <- struct{}{}
		<- ch
		fmt.Print("bar->")
	}
}

/*

func printA(n int,left,right chan struct{}) {
	for i:=0;i<n;i++ {
		<- left
		fmt.Println("A")
		right <- struct{}{}
	}
}

func printB(n int,left,right chan struct{},done chan struct{}) {
	for i := 0;i<n;i++ {
		<- left
		fmt.Println("B")
		if i != n - 1 {
			right <- struct{}{}
		}
	}
	close(done)
}

func main() {

	// A->B->C->D

	n := 5
	leftMost := make(chan struct{},1)
	left,right := leftMost,leftMost

	right  = make(chan struct{},1)
	go printA(n,left,right)

	left = right

	done := make(chan struct{})
	go printB(n,left,leftMost,done)

	go func(ch chan struct{}) {
		leftMost <- struct{}{}
	}(leftMost)

	<- done

}

func daisyChain() {
	leftMost := make(chan int)
	left,right := leftMost,leftMost

	n := 10000
	for i := 0;i<n;i++ {
		right = make(chan int)
		go f(left,right)
		left = right
	}

	go func(ch chan int) {
		ch <- 1
	}(leftMost)

	<- right
}

func f(left,right chan int) {
	temp := <-left

	// do sth
	temp ++

	right <- temp
}

 */

