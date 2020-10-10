package main

import (
	"fmt"
)

/*
149. 直线上最多的点数
给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。

示例 1:

输入: [[1,1],[2,2],[3,3]]
输出: 3
解释:
^
|
|        o
|     o
|  o
+------------->
0  1  2  3  4
示例 2:

输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
输出: 4
解释:
^
|
|  o
|     o        o
|        o
|  o        o
+------------------->
0  1  2  3  4  5  6
 */

func main() {
	points := [][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}}
	//points = [][]int{{0,0},{94911150,94911151},{94911151,94911152}}
	fmt.Println(maxPoints(points))
}

// 最多多少个点在一条直线上
// 相当于计算，过某个点的直线，最多经过多少个点
func maxPoints(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	maxCount := 1
	for i := 0;i<len(points)-1;i++ {
		maxCount = Max(maxCount,getMaxPointsOfLine(points,i))
	}
	return maxCount
}

// 计算过某个点的直线，最多经过多少个点
func getMaxPointsOfLine(points [][]int,start int) int {
	// lines的default 是 1
	lines := map[kb]int{}
	horizone := 1
	duplicates := 0
	for i := start+1;i<len(points);i++ {
		addLine(points,start,i,&horizone,&duplicates,lines)
	}

	max := horizone
	for  _,v := range lines {
		max = Max(max,v)
	}
	return max + duplicates
}

func addLine(points [][]int,x,y int,hori,dup *int,lines map[kb]int) {
	// x和y的三种情况
	// 重合
	// 同一水平线
	// slope
	if points[x][0] == points[y][0] && points[x][1] == points[y][1] {
		*dup = *dup + 1
	} else if points[x][1] == points[y][1] {
		*hori = *hori + 1
	} else {
		// 如果直接是 slope = float64(points[x][0] - points[y][0])/float64(points[x][1]-points[y][1])
		// 浮点数计算，[][]int{{0,0},{94911150,94911151},{94911151,94911152}}的case不能通过
		a,b := points[x][0] - points[y][0],points[x][1]-points[y][1]
		// d是最大公约数
		d := gcd(a,b)
		slope := kb{a/d,b/d}

		if _,ok := lines[slope];ok {
			lines[slope] = lines[slope] + 1
		} else {
			lines[slope] = 2
		}
	}
}

type kb struct {
	v1 int
	v2 int
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

// a和b的最大公约数
func gcd(a,b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a,a)
}