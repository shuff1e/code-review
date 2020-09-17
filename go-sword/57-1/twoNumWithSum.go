package main

import "fmt"

// 57（一）：和为s的两个数字
// 题目：输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们
// 的和正好是s。如果有多对数字的和等于s，输出任意一对即可。

// 1, 2, 4, 7, 11, 15
// 15

























// A：利用上数组递增的性质。可以对于每一个数字，二分查找互补的数字，这样复杂度是nlogn
// 双指针,p1指向开头，p2指向结尾
// 如果p1+p2>k，p2左移
// 如果p1+p2<k，p2右移













func getPair(arr []int,k int) (int,int,bool) {
	if len(arr) <= 1 {
		return -1,-1,false
	}
	p1 := 0
	p2 := len(arr) - 1
	for p1<p2 {
		sum := arr[p1] + arr[p2]
		if sum == k {
			return arr[p1],arr[p2],true
		} else if sum > k {
			p2 --
		} else {
			p1 ++
		}
	}
	return -1,-1,false
}

func main() {
	Test([]int{1, 2, 4, 7, 11, 15},15,true)
	Test([]int{1, 2, 4, 7, 11, 16},17,true)
	Test([]int{1, 2, 4, 7, 11, 16},10,false)
	Test([]int{},0,false)
}

func Test(arr []int,k int,expected bool) {
	_,_,ok := getPair(arr,k)
	fmt.Println(ok == expected)
}