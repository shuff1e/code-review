package main

import (
	"fmt"
)

/*
204. 计数质数
统计所有小于非负整数 n 的质数的数量。

示例 1：

输入：n = 10
输出：4
解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
示例 2：

输入：n = 0
输出：0
示例 3：

输入：n = 1
输出：0


提示：

0 <= n <= 5 * 106
 */

func main() {
	fmt.Println(countPrimes(979))
	fmt.Println(countPrimes1(979))
	fmt.Println(countPrimes2(979))
	fmt.Println(countPrimes3(979))
	fmt.Println(countPrimes4(979))
}

func countPrimes(n int) int {
	count := 0
	ch := GenerateNatural(n)

	// 这个地方不能for range ch
	// range会复制channel
	for {
		prime,ok := <-ch
		if !ok {
			break
		}
		count ++
		ch = PrimeFilter(ch,prime) // 基于新素数构造的过滤器
	}
	return count
}

func GenerateNatural(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 2;i<n;i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func PrimeFilter(in <-chan int,prime int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			if i % prime != 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func countPrimes1(n int) int {
	count := 0
	for i := 2;i<n;i++ {
		flag := false
		for j := 2;j<i;j++ {
			if i % j ==0 {
				flag = true
				break
			}
		}
		if !flag {
			count ++
		}
	}
	return count
}

func countPrimes2(n int) int {
	count := 1
	for i := 3;i<n;i++ {
		if i & 1 == 0 {
			continue
		}
		flag := false
		for j := 3;j*j<=i;j+=2 {
			if i % j ==0 {
				flag = true
				break
			}
		}
		if !flag {
			count ++
		}
	}
	return count
}

func countPrimes3(n int) int {
	count := 0
	help := make([]bool,n)
	for i := 2;i<n;i++ {
		if !help[i] {
			count++
			for j := i+i;j<n;j+=i {
				help[j] = true
			}
		}
	}
	return count
}
// 5%4 = 1
// 5 101
// 4 100

func countPrimes4(n int) int {
	count := 0
	help := make([]int,n/32+1)
	for i := 2;i<n;i++ {
		//if (help[i/32] & (1<<(i%32))) == 0 {
		if (help[i/32] & (1<<(i&31))) == 0 {
			count ++
			for j := i+i;j<n;j+=i {
				//help[j/32] = help[j/32] | (1<<(j%32))
				help[j/32] = help[j/32] | (1<<(j&31))
			}
		}
	}
	return count
}