package main

import "sort"

/*
57. 插入区间
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。


注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。
 */


// A：

type data [][]int

func (d data) Len() int {
	return len(d)
}

func (d data) Less(i,j int) bool {
	return d[i][0] < d[j][0]
}

func (d data) Swap(i,j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func insert(intervals [][]int, newInterval []int) [][]int {
	sort.Sort(data(intervals))
	merged := [][]int{}
	start,end := -1,-1

	i := 0
	for i = 0;i<len(intervals);i++ {
		// 有交集
		if !(intervals[i][0] > newInterval[1] || intervals[i][1] < newInterval[0] ) {
			if start == -1 {
				start = i
			}
			end = i
			// 右边的部分
		} else if intervals[i][0] > newInterval[1] {
			break
		} else {
			// 最开始的部分
			merged = append(merged,intervals[i])
		}
	}

	if start != -1 {
		merged = append(merged,[]int{
			Min(newInterval[0],intervals[start][0]),
			Max(newInterval[1],intervals[end][1]),
		})
	} else {
		merged = append(merged,newInterval)
	}

	for ;i<len(intervals) ;i++ {
		merged = append(merged,intervals[i])
	}
	return merged
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}
