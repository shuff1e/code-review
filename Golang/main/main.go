package main

func main() {
	//arr := []int{1,2,3,4,5,6,7}
	//Golang.DaisyChain(arr)
	//ch := make(chan int,100)
	//go Golang.Producer(ch)
	//go Golang.Consumer(ch)
	//done := make(chan struct{})
	//Golang.Test2()
	//<-done
	//p := Golang.NewPublisher(100*time.Millisecond,10)
	//docker := p.SubscribeTopic(func(v interface{}) bool {
	//	if key,ok := v.(string);ok {
	//		if strings.HasPrefix(key,"docker") {
	//			return true
	//		}
	//	}
	//	return false
	//})
	//
	//go p.Publish("docker 1234")
	//
	//time.Sleep(time.Second*1)
	//
	//go func() {
	//	fmt.Println(<-docker)
	//}()
	//go Golang.FillToken()
	//go func() {
	//	for {
	//		result := Golang.TakeAvailable(false)
	//		fmt.Println(result)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//threadNum := 5
	//for i := 0;i<threadNum;i++ {
	//	Golang.Wg.Add(1)
	//	go Golang.IncCounter(i)
	//}
	//Golang.Wg.Wait()
	//fmt.Println(Golang.Counter)
	//Golang.Test5()
	//time.Sleep(time.Second*3)
	//var wg sync.WaitGroup
	//for i := 0;i<10;i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		Golang.Incr()
	//	}()
	//}
	//wg.Wait()
	//done := make(chan struct{})
	//Golang.Test6()
	//indexes := make([]int,10)
	//for i := 0;i<len(indexes);i++ {
	//	indexes[i] = i
	//}
	//Golang.Shuffle(indexes)
	//fmt.Printf("%#v\n",indexes)
	//
	//<- done
	//	pool := Golang.NewPool(100)
	//	println(runtime.NumGoroutine())
	//	for i:=0;i<1000;i++ {
	//		pool.Add(1)
	//		go func() {
	//			time.Sleep(time.Second)
	//			println(runtime.NumGoroutine())
	//			pool.Done()
	//		}()
	//	}
	//	pool.Wait()
	//	println(runtime.NumGoroutine())
	//Golang.Test7()
	//mutex := sync.Mutex{}
	//mutex.Unlock()
	//c := Golang.Boring("boring")
	//c := Golang.FanIn(Golang.Boring("Joe"),Golang.Boring("Ann"))
	//for i :=0;i<10;i++ {
	//	msg1 := <-c;fmt.Println(msg1.Str)
	//	msg2 := <-c;fmt.Println(msg2.Str)
		//fmt.Printf("you sat %q\n", <-c)
		//msg1.Wait <- true
		//msg2.Wait <- true
	//}
	//fmt.Println("you are boring,i am leaving")
	//
	//ch := make(chan int)
	//close(ch)
	//select {
	//case <- ch:
	//	fmt.Println("fuck")
	//}
	//rand.Seed(time.Now().UnixNano())
	//start := time.Now()
	//results := Golang.Google("golang")
	//elapsed := time.Since(start)
	//fmt.Println(results)
	//fmt.Println(elapsed)
	//arr := []int{}
	//arr = append(arr,0)
	//fmt.Printf("%#v\n",arr)
}
