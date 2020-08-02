package Golang

import (
	"fmt"
	"time"
)

func Producer(out chan<- int) {
	for i := 0; ; i++ {
		out <- i
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
		time.Sleep(time.Second)
	}
}