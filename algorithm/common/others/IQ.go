package main

import "fmt"

func getNumberOfZero(n int) int {
	sum := 0
	for n > 0 {
		sum += n/5
		n/=5
	}
	return sum
}

func main() {
	n := 10
	sum := 1
	for i :=n;i>0;i-- {
		sum *= i
	}
	fmt.Println(sum)
	fmt.Println(getNumberOfZero(n))
}