package main

import (
	"fmt"
)

// k个goroutine依次顺序打印n个数字

// daisy chain 组成一个环

func main() {
	k := 4
	n := 25
	leftMost := make(chan int)
	left,right := leftMost,leftMost

	exit := make(chan struct{})

	for i := 0;i<k-1;i++ {
		right = make(chan int)
		go f(i+1,left,right,exit,n)
		left = right
	}
	go f(k,left,leftMost,exit,n)

	leftMost <- 1
	<- exit
}

func f(index int,left ,right chan int,exit chan struct{},n int) {
	for {
		i := <- left
		if i <= n {
			fmt.Println("index",index,"---print--->",i)
			i++
			right <- i
		} else {
			fmt.Println("done at index",index)
			close(exit)
		}
	}
}

/*

package main

import "fmt"

func first(left,right chan struct{},n int) {
	for i := 0;i<n;i++{
		<- left
		fmt.Println("first",i)
		right <- struct{}{}
	}
}

func second(left,right chan struct{},n int) {
	for i := 0;i<n;i++{
		<- left
		fmt.Println("second",i)
		right <- struct{}{}
	}

}

func third(left,right chan struct{},n int,exit chan struct{}) {
	for i := 0;i<n;i++{
		if i < n-1 {
			<- left
			fmt.Println("third",i)
			right <- struct{}{}
		} else {
			<- left
			fmt.Println("third",i)
			close(exit)
		}
	}
}

func main() {
	leftMost := make(chan struct{})
	left,right := leftMost,leftMost

	n := 5
	arr := []func(chan struct{},chan struct{},int){first,second}
	for i := 0;i<len(arr);i++ {
		right = make(chan struct{})
		go arr[i](left,right,n)
		left = right
	}

	exit := make(chan struct{})
	go third(right,leftMost,n,exit)

	go func() {leftMost <- struct{}{}}()
	<- exit
}

 */
