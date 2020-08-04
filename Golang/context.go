package Golang

import (
	"context"
	"fmt"
	"time"
)

var key = "name"

func Test2() {
	ctx,cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx,key,"监控1")
	go watch(valueCtx)

	time.Sleep(5*time.Second)
	cancel()
	time.Sleep(5*time.Second)

}

func watch(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Value(key),"监控退出，停止了")
			return
		default:
			fmt.Println(ctx.Value(key),"goroutine 监控中")
			time.Sleep(2*time.Second)
		}
	}
}