package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	leftMost := make(chan struct{})
	left,right := leftMost,leftMost

	for i := 0;i<10;i++ {
		right = make(chan struct{}) // right 是 out
		go func(left,right chan struct{},index int) {
			// do something
			time.Sleep(time.Second*time.Duration(rand.Intn(10)))
			<-left
			fmt.Println(index,"done")
			right <- struct{}{}
		}(left,right,i)
		left = right // right 变成in
	}
	go func(){
		leftMost <- struct{}{}
	}()
	<- right
}