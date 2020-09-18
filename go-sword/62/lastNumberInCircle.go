package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

// 62：圆圈中最后剩下的数字
// 题目：0, 1, …, n-1这n个数字排成一个圆圈，从数字0开始每次从这个圆圈里
// 删除第m个数字。求出这个圆圈里剩下的最后一个数字。

// A：约瑟夫环问题
// 0，1，2，3，4 组成的一个环，每次删除第三个
// 删除的数字依次为  2 0 4 1
// 最后剩下3





















func getRestNumber(n ,k int) int {
	if n < 1 || k < 1 {
		return -1
	}
	list := arraylist.New()
	for i := 0;i<=n-1;i++ {
		list.Add(i)
	}

	index := 0
	for list.Size() > 1 {
		for i := 1;i<k;i++ {
			index ++
			//if index == list.Size() {
			//	index = 0
			//}
		}
		index = index % list.Size()

		list.Remove(index)
		//if index == list.Size() {
		//	index = 0
		//}
		index = index % list.Size()
	}

	v,_ := list.Get(0)
	return v.(int)
}


func main() {
	Test("Test1", 5, 3, 3);
	Test("Test2", 5, 2, 2);
	Test("Test3", 6, 7, 4);
	Test("Test4", 6, 6, 3);
	Test("Test5", 0, 0, -1);
	Test("Test6", 4000, 997, 1027);
}

// f(n,m)表示n个人中，选m个剩下的那个人
// f(n-1,m)表示n-1个人中，选m个剩下的那个人
// 但是每杀掉一个人，下一个人成为头，相当于把数组的位置向前移动m位
// 比如第一次是从0开始，数了三个数，杀掉了2，那么下一次开头的是3
// 相当于开头从0变到了3
// 考虑到溢出 (f(n-1,m)+m)%n

// f(n,m)在n=1的时候，直接返回0

func getNumberBetter(n,k int) int {
	if n < 1 || k < 1 {
		return -1
	}
	// f(1,m) = 0
	result := 0
	for i := 2;i<=n;i++ {
		result = (result + k)%i
	}
	return result
}

func getNumber3(n,k int) int {
	if n < 1 || k < 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	return (getNumber3(n-1,k) + k) % n
}

func Test(name string,n,k,expected int) {
	fmt.Println(name)
	if getRestNumber(n,k) != expected {
		panic("fuck")
	}
	if getNumberBetter(n,k) != expected {
		panic("fuck")
	}
	if getNumber3(n,k) != expected {
		panic("fuck")
	}
}