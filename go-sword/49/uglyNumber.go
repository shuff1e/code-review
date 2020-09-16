package main

import "fmt"

// 49：丑数
// 题目：我们把只包含因子2、3和5的数称作丑数（Ugly Number）。求按从小到
// 大的顺序的第1500个丑数。例如6、8都是丑数，但14不是，因为它包含因子7。
// 习惯上我们把1当做第一个丑数。

// A：只包含因子2，3，5
// 一直除以，2，3，5

func isUglyNumber(num int) bool {
	if num == 0 {
		return false
	}
	for num %2 == 0 {
		num = num/2
	}
	for num %3 == 0 {
		num = num/3
	}
	for num % 5 == 0 {
		num = num/5
	}
	return num == 1
}

func getUglyNumber1(index int) int {
	if index <= 0 {
		return 0
	}
	count := 0
	num := 0
	for num < index {
		if isUglyNumber(count) {
			num ++
		}
		count ++
	}
	return count - 1
}

func main() {
	Test(1, 1);

	Test(2, 2);
	Test(3, 3);
	Test(4, 4);
	Test(5, 5);
	Test(6, 6);
	Test(7, 8);
	Test(8, 9);
	Test(9, 10);
	Test(10, 12);
	Test(11, 15);

	Test(1500, 859963392);

	Test(0, 0);
}

func Test(index,expected int) {
	fmt.Println(getUglyNumber1(index) == getUglyNumber2(index))
	if getUglyNumber2(index) != expected {
		panic("fuck")
	}
}

// 用递归和动态规划的方法思考
// 下一个丑数肯定是，之前某个位置的丑数*2，之前某个位置的丑数*3，之前某个位置的丑数*5，中最小的一个
// 第一个丑数是arr[0] = 1，第二个丑数肯定是1*2，1*3，1*5中最小的那个
// 因为选取了1*2，arr[1] = 2，那说明arr[0]*2肯定不会是下次的丑数，arr[1]*2将可能会是下次的丑数
// 用这样的方法，每次都能找到下一个丑数

func getUglyNumber2(index int) int {
	if index <= 0 {
		return 0
	}
	arr := make([]int,index)
	count := 1
	arr[0] = 1
	position := 1
	multiply2,multiply3,multiply5 := 0,0,0

	for count < index {
		min := Min(arr[multiply2]*2,arr[multiply3]*3,arr[multiply5]*5)
		arr[position] = min
		position ++
		count ++
		if min == arr[multiply2]*2 {
			multiply2 += 1
		}
		if min == arr[multiply3]*3 {
			multiply3 ++
		}
		if min == arr[multiply5]*5 {
			multiply5 ++
		}
	}
	return arr[len(arr)-1]
}

func Min(x,y,z int) int {
	if x > y {
		x = y
	}
	if x > z {
		x = z
	}
	return x
}