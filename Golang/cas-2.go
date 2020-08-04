package Golang

import (
	"fmt"
	"sync/atomic"
	"time"
)

// https://blog.csdn.net/weixin_30739595/article/details/97934430?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-2.channel_param

func Test3() {
	var count uint32

	trigger := func(i uint32,fn func()) {
		for {
			// lock
			if n := atomic.LoadUint32(&count);n ==i {
				fn()
				// release
				atomic.AddUint32(&count,1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0);i<10;i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i,fn)
		}(i)
	}

	//trigger(10,func() {})
}
