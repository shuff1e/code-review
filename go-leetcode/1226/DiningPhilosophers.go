package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

/*
1226. 哲学家进餐
5 个沉默寡言的哲学家围坐在圆桌前，每人面前一盘意面。叉子放在哲学家之间的桌面上。（5 个哲学家，5 根叉子）

所有的哲学家都只会在思考和进餐两种行为间交替。哲学家只有同时拿到左边和右边的叉子才能吃到面，而同一根叉子在同一时间只能被一个哲学家使用。每个哲学家吃完面后都需要把叉子放回桌面以供其他哲学家吃面。只要条件允许，哲学家可以拿起左边或者右边的叉子，但在没有同时拿到左右叉子时不能进食。

假设面的数量没有限制，哲学家也能随便吃，不需要考虑吃不吃得下。

设计一个进餐规则（并行算法）使得每个哲学家都不会挨饿；也就是说，在没有人知道别人什么时候想吃东西或思考的情况下，每个哲学家都可以在吃饭和思考之间一直交替下去。



问题描述和图片来自维基百科 wikipedia.org



哲学家从 0 到 4 按 顺时针 编号。请实现函数 void wantsToEat(philosopher, pickLeftFork, pickRightFork, eat, putLeftFork, putRightFork)：

philosopher 哲学家的编号。
pickLeftFork 和 pickRightFork 表示拿起左边或右边的叉子。
eat 表示吃面。
putLeftFork 和 putRightFork 表示放下左边或右边的叉子。
由于哲学家不是在吃面就是在想着啥时候吃面，所以思考这个方法没有对应的回调。
给你 5 个线程，每个都代表一个哲学家，请你使用类的同一个对象来模拟这个过程。在最后一次调用结束之前，可能会为同一个哲学家多次调用该函数。



示例：

输入：n = 1
输出：[[4,2,1],[4,1,1],[0,1,1],[2,2,1],[2,1,1],[2,0,3],[2,1,2],[2,2,2],[4,0,3],[4,1,2],[0,2,1],[4,2,2],[3,2,1],[3,1,1],[0,0,3],[0,1,2],[0,2,2],[1,2,1],[1,1,1],[3,0,3],[3,1,2],[3,2,2],[1,0,3],[1,1,2],[1,2,2]]
解释:
n 表示每个哲学家需要进餐的次数。
输出数组描述了叉子的控制和进餐的调用，它的格式如下：
output[i] = [a, b, c] (3个整数)
- a 哲学家编号。
- b 指定叉子：{1 : 左边, 2 : 右边}.
- c 指定行为：{1 : 拿起, 2 : 放下, 3 : 吃面}。
如 [4,2,1] 表示 4 号哲学家拿起了右边的叉子。


提示：

1 <= n <= 60
 */

// 5个叉子，5个锁
var lockList = []sync.Mutex{
	sync.Mutex{},
	sync.Mutex{},
	sync.Mutex{},
	sync.Mutex{},
	sync.Mutex{},
}

// 限制，最多只有4个哲学家去持有叉子
var eatLimit = semaphore.NewWeighted(4)
 var ctx = context.Background()

func wantToEat(index int) {
	left := (index+1)%5
	right := index

	eatLimit.Acquire(ctx,1)

	lockList[left].Lock()
	lockList[right].Lock()

	fmt.Println("I pick up left fork")
	fmt.Println("I pick up right fork")

	fmt.Println("I am eating")

	fmt.Println("I put down left fork")
	fmt.Println("I put down right fork")

	lockList[left].Unlock()
	lockList[right].Unlock()

	eatLimit.Release(1)
}

func main()  {
	go func(i int) {
		for {
			wantToEat(i)
		}
	}(0)

	go func(i int) {
		for {
			wantToEat(i)
		}
	}(1)

	go func(i int) {
		for {
			wantToEat(i)
		}
	}(2)

	go func(i int) {
		for {
			wantToEat(i)
		}
	}(3)

	go func(i int) {
		for {
			wantToEat(i)
		}
	}(4)

	select {

	}
}

 /*


 这道题本质上其实是想考察如何避免死锁。
 易知：当 5 个哲学家都拿着其左边(或右边)的叉子时，会进入死锁。

 PS：死锁的 4 个必要条件：

 互斥条件：一个资源每次只能被一个进程使用，即在一段时间内某 资源仅为一个进程所占有。此时若有其他进程请求该资源，则请求进程只能等待。
 请求与保持条件：进程已经保持了至少一个资源，但又提出了新的资源请求，而该资源 已被其他进程占有，此时请求进程被阻塞，但对自己已获得的资源保持不放。
 不可剥夺条件:进程所获得的资源在未使用完毕之前，不能被其他进程强行夺走，即只能 由获得该资源的进程自己来释放（只能是主动释放)。
 循环等待条件: 若干进程间形成首尾相接循环等待资源的关系。
 故最多只允许 4 个哲学家去持有叉子，可保证至少有 1 个哲学家能吃上意大利面（即获得到 2 个叉子）。
 因为最差情况下是：4 个哲学家都各自持有1个叉子，此时还 剩余 1 个叉子 可供使用，这 4 个哲学家中必然有1人能获取到这个 剩余的 1 个叉子，从而手持 2个叉子，可以吃意大利面。
 即：4 个人中，1 个人有 2 个叉子，3 个人各持 1 个叉子，共计 5 个叉子。

 既然最多只允许4个哲学家去持有叉子，那么如果只允许3个哲学家去持有叉子是否可行呢？

 当然可行，3个哲学家可以先都各自持有11把叉子，此时还剩余22把叉子；

 当这3个哲学家刚好都相邻(比如：编号为图中的0, 1, 2)，可能会造成只有11个哲学家能吃到意面的情况，具体而言即0号哲学家拿到了其左侧的叉子(编号为1)，1号哲学家也拿到了其左侧的叉子(编号为2)，2号哲学家也拿到了其左侧的叉子(编号为3)，此时只有0号哲学家能拿到其右侧的叉子(编号为0)，因此只有0号哲学家能吃到意面。
 而其余情况下，3个哲学家中都能有2人吃到意面。
 即：3 个人中，2 个人各自持有 2 个叉子，1 个人持有 1 个叉子，共计 5 个叉子。

 并且仔细想想，叉子的数目是固定的(个数为5)，直觉上来讲3个人去抢5个叉子 比 4个人去抢5个叉子效率高。

 用Semaphore去实现上述的限制：Semaphore eatLimit = new Semaphore(4);
 一共有5个叉子，视为5个ReentrantLock，并将它们全放入1个数组中。

 给叉子编号 0,1,2,3,4（对应数组下标）。
 具体编号我是如下图这般设计的：



 有了这些思路，代码实现就变得清晰起来。

 class DiningPhilosophers {
     //1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
     private final ReentrantLock[] lockList = {new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock()};

     //限制 最多只有4个哲学家去持有叉子
     private Semaphore eatLimit = new Semaphore(4);

     public DiningPhilosophers() {

     }

     // call the run() method of any runnable to execute its code
     public void wantsToEat(int philosopher,
                            Runnable pickLeftFork,
                            Runnable pickRightFork,
                            Runnable eat,
                            Runnable putLeftFork,
                            Runnable putRightFork) throws InterruptedException {

         int leftFork = (philosopher + 1) % 5;    //左边的叉子 的编号
         int rightFork = philosopher;    //右边的叉子 的编号

         eatLimit.acquire();    //限制的人数 -1

         lockList[leftFork].lock();    //拿起左边的叉子
         lockList[rightFork].lock();    //拿起右边的叉子

         pickLeftFork.run();    //拿起左边的叉子 的具体执行
         pickRightFork.run();    //拿起右边的叉子 的具体执行

         eat.run();    //吃意大利面 的具体执行

         putLeftFork.run();    //放下左边的叉子 的具体执行
         putRightFork.run();    //放下右边的叉子 的具体执行

         lockList[leftFork].unlock();    //放下左边的叉子
         lockList[rightFork].unlock();    //放下右边的叉子

         eatLimit.release();//限制的人数 +1
     }
 }


 方法 2：
 设置 1 个临界区以实现 1 个哲学家 “同时”拿起左右 2 把叉子的效果。
 即进入临界区之后，保证成功获取到左右 2 把叉子 并 执行相关代码后，才退出临界区。

 评论区看到有题友说方法2就是“只让1个哲学家就餐”的思路，无需将叉子视为ReentrantLock。

 下面我也给出了“只允许1个哲学家就餐”的代码。

 但是2者之间还是有细微的差别：
 方法2是在成功拿起左右叉子之后就退出临界区，而“只让1个哲学家就餐”是在拿起左右叉子 + 吃意面 + 放下左右叉子 一套流程走完之后才退出临界区。

 前者的情况可大概分为2种，举具体例子说明(可参照上面给出的图片)：

 1号哲学家拿起左右叉子(1号叉子 + 2号叉子)后就退出临界区，此时4号哲学家成功挤进临界区，他也成功拿起了左右叉子(0号叉子和4号叉子)，然后就退出临界区。
 1号哲学家拿起左右叉子(1号叉子 + 2号叉子)后就退出临界区，此时2号哲学家成功挤进临界区，他需要拿起2号叉子和3号叉子，但2号叉子有一定的概率还被1号哲学家持有(1号哲学家意面还没吃完)，因此2号哲学家进入临界区后还需要等待2号叉子。至于3号叉子，根本没其他人跟2号哲学家争夺，因此可以将该种情况视为“2号哲学家只拿起了1只叉子，在等待另1只叉子”的情况。
 总之，第1种情况即先后进入临界区的2位哲学家的左右叉子不存在竞争情况，因此先后进入临界区的2位哲学家进入临界区后都不用等待叉子，直接就餐。此时可视为2个哲学家在同时就餐(当然前1个哲学家有可能已经吃完了，但姑且当作是2个人同时就餐)。

 第2种情况即先后进入临界区的2位哲学家的左右叉子存在竞争情况(说明这2位哲学家的编号相邻)，因此后进入临界区的哲学家还需要等待1只叉子，才能就餐。此时可视为只有1个哲学家在就餐。

 至于“只允许1个哲学家就餐”的代码，很好理解，每次严格地只让1个哲学家就餐，由于过于严格，以至于都不需要将叉子视为ReentrantLock。

 方法2有一定的概率是“并行”，“只允许1个哲学家就餐”是严格的“串行”。


 class DiningPhilosophers {
     //1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
     private final ReentrantLock[] lockList = {new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock()};

     //让 1个哲学家可以 “同时”拿起2个叉子(搞个临界区)
     private ReentrantLock pickBothForks = new ReentrantLock();

     public DiningPhilosophers() {

     }

     // call the run() method of any runnable to execute its code
     public void wantsToEat(int philosopher,
                            Runnable pickLeftFork,
                            Runnable pickRightFork,
                            Runnable eat,
                            Runnable putLeftFork,
                            Runnable putRightFork) throws InterruptedException {

         int leftFork = (philosopher + 1) % 5;    //左边的叉子 的编号
         int rightFork = philosopher;    //右边的叉子 的编号

         pickBothForks.lock();    //进入临界区

         lockList[leftFork].lock();    //拿起左边的叉子
         lockList[rightFork].lock();    //拿起右边的叉子

         pickLeftFork.run();    //拿起左边的叉子 的具体执行
         pickRightFork.run();    //拿起右边的叉子 的具体执行

         pickBothForks.unlock();    //退出临界区

         eat.run();    //吃意大利面 的具体执行

         putLeftFork.run();    //放下左边的叉子 的具体执行
         putRightFork.run();    //放下右边的叉子 的具体执行

         lockList[leftFork].unlock();    //放下左边的叉子
         lockList[rightFork].unlock();    //放下右边的叉子
     }
 }

 方法 3：
 前面说过，该题的本质是考察 如何避免死锁。
 而当5个哲学家都左手持有其左边的叉子 或 当5个哲学家都右手持有其右边的叉子时，会发生死锁。
 故只需设计1个避免发生上述情况发生的策略即可。

 即可以让一部分哲学家优先去获取其左边的叉子，再去获取其右边的叉子；再让剩余哲学家优先去获取其右边的叉子，再去获取其左边的叉子。


 class DiningPhilosophers {
     //1个Fork视为1个ReentrantLock，5个叉子即5个ReentrantLock，将其都放入数组中
     private final ReentrantLock[] lockList = {new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock(),
             new ReentrantLock()};

     public DiningPhilosophers() {

     }

     // call the run() method of any runnable to execute its code
     public void wantsToEat(int philosopher,
                            Runnable pickLeftFork,
                            Runnable pickRightFork,
                            Runnable eat,
                            Runnable putLeftFork,
                            Runnable putRightFork) throws InterruptedException {

         int leftFork = (philosopher + 1) % 5;    //左边的叉子 的编号
         int rightFork = philosopher;    //右边的叉子 的编号

         //编号为偶数的哲学家，优先拿起左边的叉子，再拿起右边的叉子
         if (philosopher % 2 == 0) {
             lockList[leftFork].lock();    //拿起左边的叉子
             lockList[rightFork].lock();    //拿起右边的叉子
         }
         //编号为奇数的哲学家，优先拿起右边的叉子，再拿起左边的叉子
         else {
             lockList[rightFork].lock();    //拿起右边的叉子
             lockList[leftFork].lock();    //拿起左边的叉子
         }

         pickLeftFork.run();    //拿起左边的叉子 的具体执行
         pickRightFork.run();    //拿起右边的叉子 的具体执行

         eat.run();    //吃意大利面 的具体执行

         putLeftFork.run();    //放下左边的叉子 的具体执行
         putRightFork.run();    //放下右边的叉子 的具体执行

         lockList[leftFork].unlock();    //放下左边的叉子
         lockList[rightFork].unlock();    //放下右边的叉子
     }
 }


  */
