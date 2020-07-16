package util

import (
	"math/rand"
	"time"
)

func Swap(arr []int,x,y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

// Generates a slice of size, size filled with random numbers
func GenerateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}