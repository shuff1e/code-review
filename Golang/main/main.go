package main

import "github.com/shuff1e/code-review/Golang"

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
	done := make(chan struct{})
	Golang.Test6()
	<- done

}
