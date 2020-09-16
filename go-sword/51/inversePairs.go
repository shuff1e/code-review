package main

import "fmt"

// 51：数组中的逆序对
// 题目：在数组中的两个数字如果前面一个数字大于后面的数字，则这两个数字组
// 成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。

// 在数组{7,5,6,4}中，一共存在5个逆序对
// [7,5],[7,6],[7,4],[5,4],[6,4]

// A：将数组分为两部分，如果两部分都是排序过的，
// 左边数组的指针为P1，右边数组的指针为P2
// 如果P1>P2，则可以一次得到两个数组之间的逆序对

// 类似归并排序的过程

func mergeSort(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	count := 0
	mid := len(arr)/2
	count1 := mergeSort(arr[0:mid])
	count2 := mergeSort(arr[mid:])
	help := make([]int,len(arr))
	helpIndex := len(help)-1
	for p1,p2 := mid-1,len(arr)-1;p1>=0 || p2>=mid; {
		if p1 == -1 {
			for p2 >=mid {
				help[helpIndex] = arr[p2]
				p2 --
				helpIndex --
			}
			break
		}
		if p2 == mid-1 {
			for p1 >= 0 {
				help[helpIndex] = arr[p1]
				p1 --
				helpIndex --
			}
			break
		}

		if arr[p1] > arr[p2] {
			help[helpIndex] = arr[p1]
			// 更新count
			count += p2 - mid + 1
			p1 --
			helpIndex--
		} else {
			help[helpIndex] = arr[p2]
			p2 --
			helpIndex --
		}
	}
	for i := 0;i<len(help);i++ {
		arr[i] = help[i]
	}
	return count + count1 + count2
}

func main() {
	Test([]int{ 1, 2, 3, 4, 7, 6, 5 },3)
	Test([]int{ 6, 5, 4, 3, 2, 1 },15)
	Test([]int{ 1, 2, 3, 4, 5, 6 },0)
	Test([]int{ 1 },0)
	Test([]int{ 1, 2 },0)
	Test([]int{ 2, 1 },1)
	Test([]int{ 1, 2, 1, 2, 1 },3)
	Test([]int{},0)
}

func Test(arr []int,expected int) {
	count := mergeSort(arr)
	fmt.Printf("%#v\n",arr)
	fmt.Println(count,expected)
	if count != expected {
		panic("fuck")
	}
}