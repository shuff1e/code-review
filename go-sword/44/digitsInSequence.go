package main

import "fmt"

// 44：数字序列中某一位的数字
// 题目：数字以0123456789101112131415…的格式序列化到一个字符序列中。在这
// 个序列中，第5位（从0开始计数）是5，第13位是1，第19位是4，等等。请写一
// 个函数求任意位对应的数字。

// A：0 -> 9，位置是0 -> 9

// 10 -> 19,位置是10 -> 39
// 10 -> 99，位置是10 -> 189
// 10的位置是 9+2*1
// 11的位置是 9+2*2 = 13
// 总结，就是 9+ 2*(x-9)

// 100的位置是 9+2*(99-9) + 3*(100-99)

// 位置是14的话，数字是12
// 14-9 = 5
// 5/2 = 2 ，5%2 = 1
// 9+2 = 11 ，还有余数，因此是12

// 1位数的是10个
// 二位数的一共是90个
// 三位数的一共是100->999,一共是 900个

// 0 9 digit=1
// 9 90 digit=2
// 99 900 digit=3
// 999 9000

func getDigit(num int) int {
	//if num <= 9 {
	//	return num
	//}
	//num -= 9

	// 几位数
	digit := 1
	// 开始的数字
	start := 0
	// digit位的，一共有多少个数字
	count := 9
	for num > count {
		num -= count*digit
		count *= 10
		start = start*10+9
		digit ++
	}
	// 排第几
	index := num/digit
	rest := num%digit

	result := start + index
	if rest == 0 {
		return result%10
	} else {
		// 正数第rest位
		result ++
		temp := digit
		for temp != rest {
			result /= 10
			temp --
		}
		return result%10
	}
}

func Test(name string,k,expected int) {
	fmt.Println(getDigit(k))
	fmt.Println(name,getDigit(k) == expected)
}

func main()  {
	Test("Test1", 0, 0);
	Test("Test2", 1, 1);
	Test("Test3", 9, 9);
	Test("Test4", 10, 1);
	Test("Test5", 189, 9);  // 数字99的最后一位，9
	Test("Test6", 190, 1);  // 数字100的第一位，1
	Test("Test7", 1000, 3); // 数字370的第一位，3
	Test("Test8", 1001, 7); // 数字370的第二位，7
	Test("Test9", 1002, 0); // 数字370的第三位，0
}

// 改成递归的方式，而不是迭代的方式
