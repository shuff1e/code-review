package main

import "fmt"

func main() {
	ch := make(chan int)
	// 这个地方一定要传ch进去
	// 不然就是闭包，然后这个goroutine始终能访问到最新的ch
	go func(out chan int) {
		for i := 2;;i++ {
			out <- i
			//fmt.Println("Gen",i)
		}
	}(ch)

	for i := 0;i<100;i++ {
		v := <- ch
		fmt.Println(i+1,":",v)

		temp := make(chan int) // temp是out
		go func(in,out chan int,base int) {
			for value := range in {
				if value %base != 0 {
					out <- value
				}
			}
		}(ch,temp,v)

		ch = temp // temp变成in
	}
}

// Gen -> 2 ->
// Gen 	    -> 3 ->
// Gen
