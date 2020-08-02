package Golang

import (
	"fmt"
	"sync"
	"time"
)

// 多个reader等待共享资源ready的场景

var sharedRsc = false

func Test() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := sync.Mutex{}
	c := sync.NewCond(&m)
	go func() {
		c.L.Lock()
		defer c.L.Unlock()
		defer wg.Done()
		for sharedRsc == false {
			fmt.Println("goroutine 1 wait")
			c.Wait()
		}
	}()

	go func() {
		c.L.Lock()
		defer c.L.Unlock()
		defer wg.Done()
		for sharedRsc == false {
			fmt.Println("goroutine 2 wait")
			c.Wait()
		}
	}()


	time.Sleep(time.Second * 2)
	c.L.Lock()
	fmt.Println("main ready")
	sharedRsc = true
	c.Broadcast()
	fmt.Println("main broadcast")
	c.L.Unlock()
	wg.Wait()
}