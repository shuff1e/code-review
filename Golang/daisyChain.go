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

// 优先 channel
// 其次 atomic
// 其次 mutex，cond
// https://www.bwangel.me/2019/04/13/go-sync-channel/