package main

/*
135. 分发糖果
老师想给孩子们分发糖果，有 N 个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。

你需要按照以下要求，帮助老师给这些孩子分发糖果：

每个孩子至少分配到 1 个糖果。
相邻的孩子中，评分高的孩子必须获得更多的糖果。
那么这样下来，老师至少需要准备多少颗糖果呢？

示例 1:

输入: [1,0,2]
输出: 5
解释: 你可以分别给这三个孩子分发 2、1、2 颗糖果。
示例 2:

输入: [1,2,2]
输出: 4
解释: 你可以分别给这三个孩子分发 1、2、1 颗糖果。
第三个孩子只得到 1 颗糖果，这已满足上述两个条件。
 */

/*

在这种方法中，我们使用两个一维数组 left2right 和 right2left 。
数组 left2right 用来存储每名学生只与左边邻居有关的所需糖果数。
也就是假设规则为：如果一名学生评分比他左边学生高，那么他应该比他左边学生得到更多糖果。
类似的，right2left 数组用来保存只与右边邻居有关的所需糖果数。
也就是假设规则为：如果一名学生评分比他右边学生高，那么他应该比他右边学生得到更多糖果。

首先，我们在 left2right 和 right2left 中，给每个学生 1 个糖果。
然后，我们从左向右遍历整个数组，只要当前学生评分比他左邻居高，
我们在 left2right 数组中更新当前学生的糖果数 left2right[i]=left2right[i−1]+1，
这是因为在每次更新前，当前学生的糖果数一定小于等于他左邻居的糖果数。

在从左到右扫描后，我们用同样的方法从右到左只要当前学生的评分比他右边（第 (i+1) 个）学生高，
就更新 right2left[i] 为 right2left[i]=right2left[i+1]+1 。

现在，对于数组中第 i 个学生，为了满足题中条件，我们需要给他 max(left2right[i],right2left[i]) 个糖果。
因此，最后我们得到最少糖果数：

minimum_candies=
i=0
∑ ​max(left2right[i],right2left[i])
n−1


其中， n 是评分数组的长度。

 */
func candy(ratings []int) int {
	left,right := make([]int,len(ratings)),make([]int,len(ratings))
	// 每个孩子至少分配到 1 个糖果。
	for i := 0 ;i<len(ratings);i++ {
		left[i],right[i] = 1,1
	}
	for i := 1;i<len(ratings);i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	// 相邻的孩子中，评分高的孩子必须获得更多的糖果。
	for i := len(ratings)-2;i>=0;i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	sum := 0
	for i := 0;i<len(ratings);i++ {
		sum += Max(left[i],right[i])
	}
	return sum
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func candy2(ratings []int) int {
	candies := make([]int,len(ratings))
	for i := 0;i<len(ratings);i++ {
		candies[i] = 1
	}
	for i := 1;i<len(ratings);i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	sum := candies[len(ratings)-1]
	for i := len(ratings)-2;i>=0;i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = Max(candies[i],candies[i+1] + 1)
		}
		sum += candies[i]
	}
	return sum
}