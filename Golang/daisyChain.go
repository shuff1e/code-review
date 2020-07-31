package Golang

import "fmt"

// DaisyChain

func f(left chan int,right chan int,x int) {
	temp := <-left
	fmt.Println(x)
	right <- temp
}

func DaisyChain(arr []int) {
	leftMost := make(chan int)
	left := leftMost
	right := leftMost
	for _,v := range arr {
		right = make(chan int)
		go f(left,right,v)
		left = right
	}
	go func(c chan int){c <- 0}(leftMost)
	fmt.Println(<- right)
}