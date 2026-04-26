package main

import "sort"

/*
56. 合并区间
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: intervals = [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

提示：
intervals[i][0] <= intervals[i][1]

 */

// A：先排序
// 然后将第一个放入mergerd数组，
// 如果 	intervals的第一个可以和merged数组合并,合并，
// 否则将 interval的第一个放入数组

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

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Sort(data(intervals))
	merged := [][]int{}
	merged = append(merged,intervals[0])
	for i := 1;i<len(intervals);i++ {
		// 只是有交集的一种情况
		if intervals[i][0] <= merged[len(merged)-1][1] {
			merged[len(merged)-1] = []int{merged[len(merged)-1][0],
				Max(merged[len(merged)-1][1],intervals[i][1]),
				}
		} else {
			merged = append(merged,intervals[i])
		}
	}
	return merged
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}