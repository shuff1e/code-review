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
