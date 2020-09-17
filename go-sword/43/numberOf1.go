package main

import "fmt"

// 43：从1到n整数中1出现的次数
// 题目：输入一个整数n，求从1到n这n个整数的十进制表示中1出现的次数。例如
// 输入12，从1到12这些整数中包含1 的数字有1，10，11和12，1一共出现了5次。

// A：n为21345的话，分成两段，1->1345，1346->21345
// 1->1345可以递归处理

// 先分析1346->21345
// 第五位上的1，10000->11345，11346->21345
// 一共10000个(如果是11345，则是1346个)

// 然后分析后四位上的1
// 1346->21345
// 1346->9999，10000->11345,11346->21345
// 后四位，选定一位为1，其他任意三位可以是0-9
// 0999可以补到10999,因此上述3段的1为2*10^3*4

// 总结，最高位上的1，10^4*或者是1345+1
// 剩下位置上的1，10^3*4*2

func getNumberOfOnes(num int) int {
	if num == 0 {
		return 0
	}

	// 最高位
	temp := num
	count := 0
	for temp > 0 {
		temp /= 10
		count ++
	}
	if count == 1 {
		return 1
	}
	// 233
	// 33->133,134->233

	// 23
	// 3->13,14->23

	// 55-1
	// 5->15,15->25
	highest := num /pow(10,count-1)
	rest := num%pow(10,count-1)
	highestNumber := 0
	if highest > 1 {
		highestNumber = pow(10,count-1)
	} else {
		highestNumber = rest + 1
	}
	restNumber := pow(10,count-2)*(count-1)*highest
	return highestNumber + restNumber + getNumberOfOnes(rest)

}

func main() {
	Test("Test1",1, 1)
	Test("Test2",5, 1)
	Test("Test3",10, 2)
	Test("Test4",55, 16)
	Test("Test5",99, 20)
	Test("Test6",10000, 4001)
	Test("Test7",21345, 18821);
	Test("Test8",0,0)
}

func Test(name string,k,expected int) {
	fmt.Println(name)
	if getNumberOfOnes(k) != expected {
		panic("fuck")
	}
}

func pow(base,exp int) int {
	// base > 0,exp > 0
	// 101
	result := 1
	for exp > 0 {
		if exp & 1 == 1 {
			result *= base
		}
		base = base * base
		exp >>= 1
	}
	return result
}