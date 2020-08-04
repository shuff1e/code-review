package Golang

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	Counter int32
	Wg sync.WaitGroup
)

func IncCounter(index int) {
	defer Wg.Done()
	spinNUm := 0
	for {
		old := Counter
		ok := atomic.CompareAndSwapInt32(&Counter,old,old+1)
		if ok {
			break
		} else {
			spinNUm ++
		}
	}
	fmt.Printf("thread,%d,spinnum,%d\n",index,spinNUm)
}
