package Golang

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator:function that returns a channel
// channels as a handle on a service

// mulltiplexing
// Fan-in

// restoring sequence



func Boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0;;i++ {
			c <- Message{fmt.Sprintf("%s %d",msg,i),waitForIt}
			time.Sleep(time.Duration(rand.Intn(3))*time.Second)
			<- waitForIt
		}
	}()
	return c
}

func FanIn(input1,input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1: c <- s
			case s := <- input2: c <- s
			}
		}
	}()

	//go func() { for { c <- <- input1}}()
	//go func() { for { c <- <- input2}}()
	return c
}

type Message struct {
	Str string
	Wait chan bool
}


var (
	Web1 = fakeSearch("web1")
	Web2 = fakeSearch("web2")
	Image1 = fakeSearch("image1")
	Image2 = fakeSearch("image2")
	Video1 = fakeSearch("video1")
	Video2 = fakeSearch("video2")
)
type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		return Result(fmt.Sprintf("%s result for %q\n",kind,query))
	}
}

func First(query string,replicas ...Search) Result {
	c := make(chan Result)
	searchReplicas := func(i int) { c <- replicas[i](query)}
	for i := range replicas {
		go searchReplicas(i)
	}
	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query,Web1,Web2)}()
	go func() { c <- First(query,Image1,Image2)}()
	go func() { c <- First(query,Video1,Video2)}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0;i<3;i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <- timeout:
			fmt.Println("time out")
			return
		}
	}
	return
}

