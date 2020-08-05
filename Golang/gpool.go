package Golang

import "sync"

// 控制并发数的wg

type Pool struct {
	queue chan struct{}
	wg *sync.WaitGroup
}

func NewPool(size int) *Pool {
	return &Pool{
		queue: make(chan struct{},size),
		wg: &sync.WaitGroup{},
	}
}

func (p *Pool) Add(delta int) {
	for i := 0;i<delta;i++ {
		p.queue <- struct{}{}
	}
	for i := 0;i>delta;i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<- p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}