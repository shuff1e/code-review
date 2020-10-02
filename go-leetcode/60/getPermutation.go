package main

import (
	"fmt"
	"strconv"
)

/*
60. 第k个排列
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
 */

// 1 2 3 4
// 2 1 3 4 --> 7
// 2 1 4 3 --> 8

// 2 3 1 4 --> 9
// 2 3 4 1
// 2 4 1 3
// 2 4 3 1

// 3 1 2 4
// 3 1 4 2

// A：k=12
// 3

func main() {
	fmt.Println(getPermutation(1,1))
}

// n=4
// k=9,n=4,"1234"后面三位一共有3!个数字,
// 9/3!=1,9%3!=3,说明9的第一位是[1,2,3,4]中排在第二位的数字，因此第一位是2
// 然后问题转换为k=3,n=3的问题，3/2!=1,3%2!=1,因此此处也是排在第二位的数字，因为2已经被拿走了，因此是3
// 然后问题转换为k=1,n=2的问题，1/1!=1,1%1!=0，因此此处是排在第一位的数字，是1

// 对于k=12，n=4,12/3!=2,12%3!=0，因此此处是该序列中最大的数，则将2取走后，剩余的序列按照从后往前的顺序取

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	chosen := []int{}
	temp := n

	fac := factor(temp-1)
	for k > 0 {
		a := k/fac
		k = k%fac
		if k > 0 {
			chosen = append(chosen,a+1)
		} else {
			chosen = append(chosen,a)
		}
		fac/=(temp-1)
		temp --
	}

	str := ""
	help := make(arrayList,n)
	for i :=0;i<n;i++ {
		help[i] = i+1
	}

	for _,v := range chosen {
		// 排第几位的数字
		str = str + strconv.Itoa(help.get(v-1))
	}
	for i:=len(help)-1;i>=0;i-- {
		if help[i] != -1 {
			str = str+strconv.Itoa(help[i])
		}
	}
	return str
}

func factor(n int) int {
	result := 1
	for n > 0 {
		result = result*n
		n = n-1
	}
	return result
}

type arrayList []int

func (a *arrayList) get(index int) int {
	result := (*a)[index]
	*a = append((*a)[0:index],(*a)[index+1:]...)
	return result
}