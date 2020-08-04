package Golang

import (
	"fmt"
)

// A,B 两个goroutine交替打印1，2，3，4和A，B，C，D

// https://blog.csdn.net/qianghaohao/article/details/97007270
func Test4() {
	A := make(chan struct{},1)
	B := make(chan struct{})

	go func() {
		s := []int{1,2,3,4}
		for i:=0;i<len(s);i++ {
			<- A
			fmt.Println(s[i])
			B <- struct{}{}
		}
	}()

	go func() {
		s := []rune{'A','B','C','D'}
		for i := 0;i<len(s);i++ {
			<- B
			fmt.Printf("%c\n",s[i])
			A <- struct{}{}
		}
	}()

	A <- struct{}{}
}

// A,B两个goroutine交替输出奇数，偶数

func Test6() {
	A := make(chan struct{},1)
	B := make(chan struct{})

	go func() {
		for i := 1;i<10;i++ {
			<- A
			fmt.Println("A= ",2*i-1)
			B <- struct{}{}
		}
	}()

	go func() {
		for i := 1;i<10;i++ {
			<- B
			fmt.Println("B= ",2*i)
			A <- struct{}{}
		}
	}()
	A <- struct{}{}
}