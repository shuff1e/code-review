package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

/*
1117. H2O 生成
现在有两种线程，氧 oxygen 和氢 hydrogen，你的目标是组织这两种线程来产生水分子。

存在一个屏障（barrier）使得每个线程必须等候直到一个完整水分子能够被产生出来。

氢和氧线程会被分别给予 releaseHydrogen 和 releaseOxygen 方法来允许它们突破屏障。

这些线程应该三三成组突破屏障并能立即组合产生一个水分子。

你必须保证产生一个水分子所需线程的结合必须发生在下一个水分子产生之前。

换句话说:

如果一个氧线程到达屏障时没有氢线程到达，它必须等候直到两个氢线程到达。
如果一个氢线程到达屏障时没有其它线程到达，它必须等候直到一个氧线程和另一个氢线程到达。
书写满足这些限制条件的氢、氧线程同步代码。



示例 1:

输入: "HOH"
输出: "HHO"
解释: "HOH" 和 "OHH" 依然都是有效解。
示例 2:

输入: "OOHHHH"
输出: "HHOHHO"
解释: "HOHHHO", "OHHHHO", "HHOHOH", "HOHHOH", "OHHHOH", "HHOOHH", "HOHOHH" 和 "OHHOHH" 依然都是有效解。


提示：

输入字符串的总长将会是 3n, 1 ≤ n ≤ 50；
输入字符串中的 “H” 总数将会是 2n 。
输入字符串中的 “O” 总数将会是 n 。
 */

/*

Semaphore(int permits)
Creates a Semaphore with the given number of permits and nonfair fairness setting.

void    acquire()
Acquires a permit from this semaphore, blocking until one is available, or the thread is interrupted.
void    acquire(int permits)
Acquires the given number of permits from this semaphore, blocking until all are available, or the thread is interrupted.

void    release()
Releases a permit, returning it to the semaphore.
void    release(int permits)
Releases the given number of permits, returning them to the semaphore.


import java.util.concurrent.Semaphore;
import java.util.concurrent.CountDownLatch;

class H2O {
    private Semaphore h = new Semaphore(2);
    private Semaphore o = new Semaphore(0);

    public H2O() {
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
        h.acquire();
        // releaseHydrogen.run() outputs "H". Do not change or remove this line.
        releaseHydrogen.run();
        o.release();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
        o.acquire(2);
        // releaseOxygen.run() outputs "O". Do not change or remove this line.
        releaseOxygen.run();
        h.release(2);
    }
}

 */

var h = semaphore.NewWeighted(2)
var o = semaphore.NewWeighted(2)
var ctx = context.Background()

func hydrogen() {
	for {
		if err := h.Acquire(ctx,1);err != nil {
			fmt.Println(err)
		}
		fmt.Print("H")
		time.Sleep(time.Millisecond)
		o.Release(1)
	}
}

func oxygen() {
	for {
		if err := o.Acquire(ctx,2);err != nil {
			fmt.Println(err)
		}
		fmt.Print("O")
		time.Sleep(time.Millisecond)
		h.Release(2)
	}
}

func main() {
	o.Acquire(ctx,1)
	o.Acquire(ctx,1)

	go hydrogen()
	go oxygen()
	time.Sleep(time.Second*10)
}