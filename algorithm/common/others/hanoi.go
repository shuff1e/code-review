package main

import "fmt"

func move(n int,from string,help string, to string) {
	fmt.Println("moving "+ fmt.Sprintf("%d",n) + " from "+ from + " to "+ to)
}

func hanoi(n int ,from string,help string,to string) {
	if n == 1 {
		move(n,from ,help,to)
		return
	}
	hanoi(n-1,from,to,help)
	move(n,from,help,to)
	hanoi(n-1,help,from,to)
}

func main() {
	from := "A"
	help := "B"
	to := "C"
	hanoi(4,from,help,to)
}