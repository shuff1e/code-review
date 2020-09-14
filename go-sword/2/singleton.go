package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 面试题2：实现Singleton模式
// 题目：设计一个类，我们只能生成该类的一个实例。

// A：实现一个单例类，即实现 sync.Once

// atomic 提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作，
// 善用 atomic 能够避免程序中出现大量的锁操作。

// 载入操作能够保证原子的读变量的值，当读取的时候，任何其他CPU操作都无法对该变量进行读写，
// 其实现机制受到底层硬件的支持。

// 优先 channel
// 其次 atomic
// 其次 mutex，cond
// https://www.bwangel.me/2019/04/13/go-sync-channel/

type Once struct {
	done uint32
	mutex sync.Mutex
}

func (o *Once) Do(f func()) {
	// Note: Here is an incorrect implementation of Do:
	//
	//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//		f()
	//	}
	//
	// Do guarantees that when it returns, f has finished.
	// This implementation would not implement that guarantee:
	// given two simultaneous calls, the winner of the cas would
	// call f, and the second would return immediately, without
	// waiting for the first's call to f to complete.
	// This is why the slow path falls back to a mutex, and why
	// the atomic.StoreUint32 must be delayed until after f returns.

	// https://stackoverflow.com/questions/56423412/sync-once-implementation

	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	if o.done == 0 {
		//defer func() { o.done = 1}()
		// StoreUint32 is necessary because the field is read concurrently, without holding the lock, by LoadUint32, right there in the same function.
		// https://stackoverflow.com/questions/55964014/why-sync-once-using-atomic-in-doslow
		defer atomic.StoreUint32(&o.done,1)
		f()
	}
}


// test code here

var a = 0
func assign() {
	a = a+1
}

func main() {
	o := Once{}
	wg := sync.WaitGroup{}

	for i := 0;i<5;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			o.Do(assign)
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
