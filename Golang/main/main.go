package main

import (
	"fmt"
	"github.com/shuff1e/code-review/Golang"
	"time"
)

func main() {
	//arr := []int{1,2,3,4,5,6,7}
	//Golang.DaisyChain(arr)
	//ch := make(chan int,100)
	//go Golang.Producer(ch)
	//go Golang.Consumer(ch)
	//done := make(chan struct{})
	//Golang.Test2()
	//<-done
	//p := Golang.NewPublisher(100*time.Millisecond,10)
	//docker := p.SubscribeTopic(func(v interface{}) bool {
	//	if key,ok := v.(string);ok {
	//		if strings.HasPrefix(key,"docker") {
	//			return true
	//		}
	//	}
	//	return false
	//})
	//
	//go p.Publish("docker 1234")
	//
	//time.Sleep(time.Second*1)
	//
	//go func() {
	//	fmt.Println(<-docker)
	//}()
	go Golang.FillToken()
	go func() {
		for {
			result := Golang.TakeAvailable(false)
			fmt.Println(result)
			time.Sleep(time.Second)
		}
	}()

	<- make(chan bool)
}
