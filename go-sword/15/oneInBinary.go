package main

import "fmt"

// 15：二进制中1的个数
// 题目：请实现一个函数，输入一个整数，输出该数二进制表示中1的个数。例如
// 把9表示成二进制是1001，有2位是1。因此如果输入9，该函数输出2。

// A：n和n-1与能够消除掉末尾的1
func getNumberOfOne(n int) int {
	count := 0
	for n >0 {
		n = n & (n-1)
		count ++
	}
	return count
}

func main() {
	Test(0, 0);
	Test(1, 1);
	Test(10, 2);
	Test(0x7FFFFFFF, 31);
	Test(0xFFFFFFFF, 32);
	Test(0x80000000, 1);
}

func Test(n ,expected int) {
	fmt.Println(getNumberOfOne(n) == expected)
}