package main

/*
11. 盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。
    |   |
    |   |
    |   |
    |   |
    |   |
|   |   |   |
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。



示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49

 */

// A：双指针
// 关键是两个指针怎么移动
// arr[i] < arr[j] 的话，面积更大的话，要i变成i+1
// 因为如果j变成j-1，高最多还是arr[i]，但是底就减少了1，这样面积就变小了

// 如果arr[i] == arr[j]，也可以i变成i+1

func maxArea(height []int) int {
	left,right := 0,len(height) - 1
	result := 0

	for left < right {
		result = Max(result,(right - left)*Min(height[left],height[right]))
		if height[left] <= height[right] {
			left ++
		} else {
			right --
		}
	}

	return result
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x,y int) int {
	if x > y{
		return x
	}
	return y
}