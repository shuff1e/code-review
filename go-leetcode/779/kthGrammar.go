package main

import "fmt"

/*

779. 第K个语法符号
在第一行我们写上一个 0。接下来的每一行，将前一行中的0替换为01，1替换为10。

给定行数 N 和序数 K，返回第 N 行中第 K个字符。（K从1开始）


例子:

输入: N = 1, K = 1
输出: 0

输入: N = 2, K = 1
输出: 0

输入: N = 2, K = 2
输出: 1

输入: N = 4, K = 5
输出: 1

解释:
第一行: 0
第二行: 01
第三行: 0110
第四行: 01101001
       01101001 10010110

0
1
6
6 <<4 + ^6

f(4) = f(3) << 2^(3-1) + ^f(3)

f(n) = f(n-1) << 2^(n-1) + ^f(n-1)

f(2) = 01

f(n) & ( 1 << (2^(n-1))-k) )

注意：

N 的范围 [1, 30].
K 的范围 [1, 2^(N-1)].

 */

func main() {
	N := 5
	K := 8
	fmt.Println(kthGrammar(N,K))
	fmt.Println(kthGrammar2(N,K))
}

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if N == 2 {
		if 1 & (1<< (Pow(2,N-1)-K) ) > 0 {
			return 1
		} else {
			return 0
		}
	}
	if N == 3 {
		if 6 & (1<< (Pow(2,N-1)-K) ) > 0 {
			return 1
		}
		return 0
	}
	if N == 4 {
		// 6 *2*2*2*2 + 1 + 8 = 96 + 9 = 105
		if 105 & (1<< (Pow(2,N-1)-K) ) > 0 {
			return 1
		}
		return 0
	}
	initial := []byte{105}
	temp := make([]byte,1)

	for i := 4;i<N;i++ {
		for j := 0;j<len(initial);j++ {
			temp[j] = ^initial[j]
		}
		initial = append(initial,temp...)
		temp = append(temp,temp...)
	}
	index := Pow(2,N-1)-K + 1

	// index = 9 = 8 + 1
	// index == 16
	count := index/8
	diff := index%8
	if diff == 0 {
		return int(initial[len(initial)-count] >> 7) & 1
	} else {
		return int(initial[len(initial)-count-1] >> (diff-1)) & 1
	}
}

// 2^3
// 11
func Pow(base,exp int) int {
	result := 1
	for exp > 0 {
		if exp & 1 > 0 {
			result *= base
		}
		base = base * base
		exp = exp >> 1
	}
	return result
}

// f(n) = f(n-1) << 2^(n-1) + ^f(n-1)
// f(n) & ( 1 << (2^(n-1))-k) )
func kthGrammar2(N int, K int) int {
	result := help(N,K)
	if result {
		return 1
	}
	return 0
}

func help(N int,K int) bool {
	if N == 1 {
		return false
	}
	// N层总的长度为pow(2,N-1)
	length := Pow(2,N-1)
	// 0 1 2 3 4 5 6 7
	if K <= length/2 {
		return help(N-1,K)
	}
	return !help(N-1,K-length/2)
}