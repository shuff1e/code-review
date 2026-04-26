package main

import "fmt"

/*
739. 每日温度

给定一个整数数组 temperatures，其中 temperatures[i] 表示第 i 天的气温。
请返回一个数组 answer，answer[i] 表示第 i 天之后需要等待多少天才会出现更高的气温。
如果之后不会再出现更高的气温，则 answer[i] 为 0。

示例 1：
输入：temperatures = [73,74,75,71,69,72,76,73]
输出：[1,1,4,2,1,1,0,0]

示例 2：
输入：temperatures = [30,40,50,60]
输出：[1,1,1,0]

示例 3：
输入：temperatures = [30,60,90]
输出：[1,1,0]

提示：
1 <= temperatures.length <= 10^5
30 <= temperatures[i] <= 100
*/

// 单调栈
// 单调递减

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(arr)
	fmt.Println(result)

	arr = []int{30, 40, 50, 60}
	result = dailyTemperatures(arr)
	fmt.Println(result)

	arr = []int{30, 30, 30, 60}
	result = dailyTemperatures(arr)
	fmt.Println(result)
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return []int{}
	}

	result := make([]int, len(temperatures))

	stack := []int{}

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 {
			peek := stack[len(stack)-1]
			if temperatures[peek] >= temperatures[i] {
				break
			}
			stack = stack[:len(stack)-1]

			result[peek] = i - peek
		}
		stack = append(stack, i)
	}

	return result
}
