package main

import "fmt"

// 57（二）：为s的连续正数序列
// 题目：输入一个正数s，打印出所有和为s的连续正数序列（至少含有两个数）。
// 例如输入15，由于1+2+3+4+5=4+5+6=7+8=15，所以结果打印出3个连续序列1～5、
// 4～6和7～8。

// A：从(s+1)/2开始
// 从8开始，p1指向7，p2都指向8
// 7+8=15，打印

// 然后p1，p2左移动
// p1到6，p2到7，6+7小于15,p1左移,5+6+7=18>15，
// p2 右移，sum-8=11，小于15，左移，4+5+6=15，打印

// p1到3，p2到5

// 大于，p2左移，小于，p1左移
// ==，都左移动


















func printAllSequence(k int) {
	p2 := (k+1)/2
	p1 := p2-1
	if p1 <= 0 {
		return
	}

	sum := p1+p2
	for p1 > 0 {
		if sum == k {
			printArr(p1,p2)
			sum = sum - p2
			p2 --

			p1 --
			sum = sum + p1
		} else if sum < k {
			p1 --
			sum += p1
		} else {
			sum -= p2
			p2 --
		}
	}
}

func printArr(left,right int) {
	sum := 0
	for i := left;i<=right;i++ {
		fmt.Print(i," ")
		sum += i
	}
	fmt.Println("--->",sum)
}

func Test(name string,k int) {
	fmt.Println(name)
	printAllSequence(k)
}

func main() {
	Test("test1", 1);
	Test("test2", 3);
	Test("test3", 4);
	Test("test4", 9);
	Test("test5", 15);
	Test("test6", 100);
}