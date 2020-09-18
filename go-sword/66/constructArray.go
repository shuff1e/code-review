package main

import "fmt"

// 66：构建乘积数组
// 题目：给定一个数组A[0, 1, …, n-1]，请构建一个数组B[0, 1, …, n-1]，其
// 中B中的元素B[i] =A[0]×A[1]×… ×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

// A：B[0] = A[1] * ... * A[n-1]
// 一种直观的方法，是A中n-1个数相乘，这样时间复杂度为n^2

// 可以分成两部分，B[i] = A[0]*...*A[i-1] * A[i+1]*...*A[n-1] = C[i] * D[i]
// C和D分别可以求出来
// C[i] = A[0]*...*A[i-1] = C[i-1]*A[i-1] ,C[0] = 1
// D[i] = A[i+1]*...*A[n-1] = D[i+1]*A[i+1], D[n-1] = 0

func constructArray(arr []int) []int {
	C := make([]int,len(arr))
	D := make([]int,len(arr))

	C[0] = 1
	for i := 1;i<len(arr);i++ {
		C[i] = C[i-1] * arr[i-1]
	}

	D[len(arr)-1] = 1
	for i := len(arr) -2 ;i>=0;i-- {
		D[i] = D[i+1]*arr[i+1]
	}

	result := make([]int,len(arr))
	for i := 0;i<len(result);i++ {
		result[i] = C[i] * D[i]
	}

	return result
}

func constructBetter(arr []int) []int {
	result := make([]int,len(arr))
	result[0] = 1
	for i := 1;i<len(arr);i++ {
		 result[i] = result[i-1] * arr[i-1]
	}

	temp := 1
	for i := len(arr) - 2;i>=0;i-- {
		temp = temp * arr[i+1]
		result[i] *= temp
	}
	return result
}

func theSame(arr1,arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i,v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

func Test(arr,expected []int) {
	result := constructArray(arr)
	fmt.Printf("%#v\n",result)
	if !theSame(expected,result) {
		panic("fuck")
	}
	result  = constructBetter(arr)
	if !theSame(expected,result) {
		panic("fuck")
	}
}

func main() {
	Test([]int{ 1, 2, 3, 4, 5 },[]int{ 120, 60, 40, 30, 24 })
	Test([]int{ 1, 2, 0, 4, 5 },[]int{ 0, 0, 40, 0, 0 })
	Test([]int{ 1, 2, 0, 4, 0 },[]int{ 0, 0, 0, 0, 0 })
	Test([]int{ 1, -2, 3, -4, 5 },[]int{ 120, -60, 40, -30, 24 })
	Test([]int{ 1, -2 },[]int{ -2, 1 })
}