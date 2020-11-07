package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "LEET"
	fmt.Println(convert(str,3))
}

func convert(s string, numRows int) string {
	K := numRows
	str := s
	if K == 1 {
		return str
	}
	start := 0
	result := []byte{}
	for start < len(str) {
		result = append(result,str[start])
		start += 2*(K-1)
	}

	offset := 2*(K - 2)
	for i := 1;i<K-1;i++ {
		start := i
		for start < len(str) {
			result = append(result,str[start])
			if start + offset < len(str) {
				result = append(result,str[start + offset])
			}
			start += 2*(K-1)
		}
		offset -= 2
	}

	start = K - 1
	for start < len(str) {
		result = append(result,str[start])
		start += 2*(K-1)
	}
	return String(result)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}
