package Golang

import (
	"sync"
	"sync/atomic"
)

type singleton struct {}

var (
	instance *singleton
	initialized uint32
	mu sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized,1)
		instance = &singleton{}
	}
	return instance
}

type Once struct {
	done uint32
	m sync.Mutex
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done,1)
		f()
	}
}
