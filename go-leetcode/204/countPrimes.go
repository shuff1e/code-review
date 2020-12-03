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
	fmt.Println(countPrimes5(979))
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

// 暴力1
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

// 暴力2
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

// 埃氏筛
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

// 埃氏筛，用位运算加速
// 每32个数，占用help的一个元素的位置，
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

// 线性筛
// 对于45这个数，它同时会被3和5两个数标记
// 优化的目的，就是让每个数都只被标记一次。这样复杂度即能保证为O(n)

// 相较于埃氏筛，我们多维护一个 primes 数组表示当前得到的质数集合。我们从小到大遍历，如果当前的数 x 是质数，就将其加入 primes 数组。

// 另一点与埃氏筛不同的是，「标记过程」不再仅当 x 为质数时才进行，而是对每个整数 x 都进行。
// 对于整数 x，我们不再标记其所有的倍数 x*x,x*(x+1),…，
// 而是只标记质数集合中的数与 x 相乘的数，即 x*primes0,x*primes1,…，
// 且在发现 x mod primes[i] == 0 的时候结束当前标记。
//
// 核心点在于：如果 x 可以被 primes[i]整除，那么对于合数 y=x*primes[i+1] 而言，
// 它一定在后面遍历到 x/primes[i]*primes[i+1]
//​ 这个数的时候会被标记，
// 其他同理，这保证了每个合数只会被其「最小的质因数」筛去，即每个合数被标记一次。
//
// 例如 对于45，
// 5*9=45
// 3*15=45
//
// 当i为9的时候，9%3==0，所以9不会乘以5去标记45
// 因为 9/3*5=15
// 15肯定比9更大，会由15来标记45
func countPrimes5(n int) int {
	primes := []int{}
	isPrime := make([]int,n)
	for i := 2;i<n;i++ {
		if isPrime[i] == 0 {
			primes = append(primes,i)
		}
		for j := 0;j<len(primes) && i*primes[j] < n;j++ {
			isPrime[i*primes[j]] = 1
			if i % primes[j] == 0 {
				break
			}
		}
	}
	return len(primes)
}