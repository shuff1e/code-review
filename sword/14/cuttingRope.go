package main

import "fmt"

// 14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。

// A：递归，同时加上memo
//
func cut1(n int,memo []int,first int) int {
	// 如果是整段砍下，这里直接返回1
	if n == 0 {
		return 1
	}
	// 如果已经计算过
	if memo[n] != 0 {
		return memo[n]
	}
	max := 0
	// 如果是第一次砍，不能直接砍成一段
	for i := 1;i<= n-first;i++ {
		temp := i*cut1(n-i,memo,0)
		if temp > max {
			max = temp
		}
	}
	memo[n] = max
	return max
}

func cut1Mem(n int) int {
	memo := make([]int,n+1)
	return cut1(n,memo,1)
}

func main() {
	test(1,0)
	test(2,1)
	test(3,2)
	test(4,4)
	test(5,6)
	test(6,9)
	test(7,12)
	test(8,18)
	test(9,27)
	test(10,36)
	test(50,86093442)
}

func test(n ,expected int) {
	fmt.Println(cut1Mem(n) == expected)
}