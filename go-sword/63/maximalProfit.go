package main

import "fmt"

// 63：股票的最大利润
// 题目：假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖交易该股
// 票可能获得的利润是多少？例如一只股票在某些时间节点的价格为{9, 11, 8, 5,
// 7, 12, 16, 14}。如果我们能在价格为5的时候买入并在价格为16时卖出，则能
// 收获最大的利润11。

// A：所求为差值的最大值
// 对于每个数，只要得到该数之前的最小值，就可以得到以这个价格卖出时的最大利润

func getDiff(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}

	minNumber := Min(arr[0],arr[1])
	diff := arr[1] - arr[0]

	for i := 2;i<len(arr);i++ {
		diff = Max(diff,arr[i]-minNumber)
		minNumber = Min(minNumber,arr[i])
	}
	return diff
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	Test("Test1",[]int{ 4, 1, 3, 2, 5 },4)
	Test("Test2",[]int{ 1, 2, 4, 7, 11, 16 },15)
	Test("Test3",[]int{ 16, 11, 7, 4, 2, 1 },-1)
	Test("Test4",[]int{ 16, 16, 16, 16, 16 },0)
	Test("Test5",[]int{ 9, 11, 5, 7, 16, 1, 4, 2 },11)
	Test("Test6",[]int{ 2, 4 },2)
	Test("Test7",[]int{ 4, 2 },-2)
	Test("Test8",nil,0)
}

func Test(name string,numbers []int,expected int) {
	fmt.Println(name)
	if getDiff(numbers) != expected {
		panic("fuck")
	}
}