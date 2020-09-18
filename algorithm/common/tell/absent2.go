package main

import "fmt"

func findAbsent2(array []int,n int) (x,y int){
	a,b,c,d := 0,0,0,0
	for i:=0;i<n;i++ {
		a += i
		b += i*i
	}

	for _,v := range array {
		c += v
		d += v*v
	}

	diff := a - c
	diff2 := b - d
	fmt.Println(b,d)
	fmt.Println(diff,diff2)

	for i:=0;i<n;i++ {
		for j:=i+1;j<n;j++ {
			if i + j == diff && i*i + j*j == diff2 {
				return i,j
			}
		}
	}
	return -1,-1
}

func main() {
	array := []int{0,1,2,3,5,6,7,8}
	length := 10
	a,b := findAbsent2(array,length)
	fmt.Println(a,b)
}
