package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*

218. 天际线问题
城市的天际线是从远处观看该城市中所有建筑物形成的轮廓的外部轮廓。
现在，假设您获得了城市风光照片（图A）上显示的所有建筑物的位置和高度，请编写一个程序以输出由这些建筑物形成的天际线（图B）。

Buildings Skyline Contour

每个建筑物的几何信息用三元组 [Li，Ri，Hi] 表示，其中 Li 和 Ri 分别是第 i 座建筑物左右边缘的 x 坐标，Hi 是其高度。可以保证 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX 和 Ri - Li > 0。您可以假设所有建筑物都是在绝对平坦且高度为 0 的表面上的完美矩形。

例如，图A中所有建筑物的尺寸记录为：[ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] 。

输出是以 [ [x1,y1], [x2, y2], [x3, y3], ... ] 格式的“关键点”（图B中的红点）的列表，它们唯一地定义了天际线。
关键点是水平线段的左端点。请注意，最右侧建筑物的最后一个关键点仅用于标记天际线的终点，并始终为零高度。
此外，任何两个相邻建筑物之间的地面都应被视为天际线轮廓的一部分。

例如，图B中的天际线应该表示为：[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ]。

说明:

任何输入列表中的建筑物数量保证在 [0, 10000] 范围内。
输入列表已经按左 x 坐标 Li  进行升序排列。
输出列表必须按 x 位排序。
输出天际线中不得有连续的相同高度的水平线。例如 [...[2 3], [4 5], [7 5], [11 5], [12 7]...] 是不正确的答案；
三条高度为 5 的线应该在最终输出中合并为一个：[...[2 3], [4 5], [12 7], ...]

 */

//type Interface interface {
//	sort.Interface
//	Push(x interface{}) // add x as element Len()
//	Pop() interface{}   // remove and return element Len() - 1.
//}

//type Interface interface {
	//Len() int
	//Less(i, j int) bool
	//Swap(i, j int)
//}

func main() {
	matrix := [][]int{ {2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8} }
	result := getSkyline(matrix)
	for i := 0;i<len(result);i++ {
		fmt.Printf("%#v\n",result[i])
	}
}

// 线扫描法
func getSkyline(buildings [][]int) [][]int {
	result := [][]int{}
	pairs := make([][2]int,len(buildings)*2)
	index := 0
	for _,v := range buildings {
		// 左边的点
		pairs[index] = [2]int{v[0],-v[2]}
		index ++
		// 右边的点
		pairs[index] = [2]int{v[1],v[2]}
		index ++
	}
	// 升序排列
	sort.Slice(pairs,func(i,j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] < pairs[j][0]
		}
		// 这样[2,-10]会排在[2,5]前面
		return pairs[i][1] < pairs[j][1]
	})
	maxHeap := &intHeap{}
	prev := 0
	for _,pair := range pairs {
		// 例如[0,-5] [2,5] [2,-10] [6,10]
		//  pairs中的顺序为 [0,-5][2,-10][2,5][6,10]
		// 左边的入heap
		if pair[1] < 0 {
			heap.Push(maxHeap,-pair[1])
		} else {
			for i := 0;i<maxHeap.Len();i++ {
				if maxHeap.Get(i) == pair[1] {
					heap.Remove(maxHeap,i)
					break
				}
			}
		}
		top := maxHeap.Peek()
		if prev != top {
			result = append(result,[]int{pair[0],top})
			prev = top
		}
	}
	return result
}

type intHeap []int

func (h intHeap) Len() int {return len(h)}
func (h intHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}
func (h intHeap) Less(i,j int) bool {
	return h[i] > h [j] // > 表示最大堆，< 表示最小堆
}



func (h *intHeap) Push(x interface{}) {
	*h = append(*h,x.(int))
}



func (h *intHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
	return result
}



func (h intHeap) Peek() int {
	if len(h) > 0{
		return h[0]
	}
	return 0
}

func (h intHeap) Get(index int) int {return h[index]}

