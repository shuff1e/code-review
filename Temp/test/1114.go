package main

import (
	"fmt"
)

const n=1

func zero(left1,right1 chan struct{},left2,right2 chan struct{}) {
	for i := 0; i < n; i++ {
		fmt.Print("0")
		right1 <- struct{}{}
		<-left1
		fmt.Print("0")
		right2 <- struct{}{}
		<-left2
	}
}

func one(left ,right chan struct{}) {
	for i:=0;i<n;i++ {
		<- left
		fmt.Print("1")
		right <- struct{}{}
	}
}

func two(left,right,exit chan struct{}) {
	defer func() {
		close(exit)
	}()
	for i :=0;i<n;i++ {
		<- left
		fmt.Print("2")
		right <- struct{}{}
	}
}

func main() {
	left1,right1,left2,right2 := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})
	exit := make(chan struct{})
	go zero(left1,right1,left2,right2)
	go one(right1,left1)
	go two(right2,left2,exit)

	//left1 <- struct{}{}
	//left2 <- struct{}{}
	<- exit
}
