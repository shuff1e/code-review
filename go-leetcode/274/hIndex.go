package main

import (
	"fmt"
	"sort"
)

/*
274. H 指数
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。

h 指数的定义：h 代表“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）总共有 h 篇论文分别被引用了至少 h 次。
（其余的 N - h 篇论文每篇被引用次数 不超过 h 次。）

例如：某人的 h 指数是 20，这表示他已发表的论文中，每篇被引用了至少 20 次的论文总共有 20 篇。

示例：

输入：citations = [3,0,6,1,5]
输出：3
解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。

提示：如果 h 有多种可能的值，h 指数是其中最大的那个。
 */

// 0 1 2 3 4
// 6 5 2 1 0
// arr[i] > i

//我们想象一个直方图，其中 x 轴表示文章，y 轴表示每篇文章的引用次数。
//如果将这些文章按照引用次数降序排序并在直方图上进行表示，那么直方图上的最大的正方形的边长 h 就是我们所要求的 h。
//
//首先我们将引用次数降序排序，在排完序的数组 citations 中，如果 citations[i]>i，那么说明第 0 到 i 篇论文都有至少 i+1 次引用。因此我们只要找到最大的 i 满足 citations[i]>i，那么 h 指数即为 i+1。例如：
//
//其中最大的满足 citations[i]>i 的 i 值为 2，因此 h=i+1=3。
//
//找到最大的 i 的方法有很多，可以对数组进行线性扫描，也可以使用二分查找。由于排序的时间复杂度已经为 O(nlogn)，因此无论是线性扫描 O(n) 还是二分查找 O(logn)，都不会改变算法的总复杂度。

func main() {
	arr := []int{1,1}
	fmt.Println(hIndex2(arr))
}

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	if len(citations) == 1 {
		if citations[0] >= 1 {
			return 1
		}
		return 0
	}

	sort.Slice(citations,func(i, j int) bool{
		return citations[i] > citations[j]
	})
	index := 0
	for i := 0;i < len(citations);i++ {
		if citations[i] <= index {
			break
		}
		index ++
	}
	return index
}

// 方法一中，我们通过降序排序得到了 h 指数，
//然而，所有基于比较的排序算法，例如堆排序，合并排序和快速排序，都存在时间复杂度下界 O(nlogn)。要得到时间复杂度更低的算法. 可以考虑最常用的不基于比较的排序，计数排序。
//
//然而，论文的引用次数可能会非常多，这个数值很可能会超过论文的总数n，因此使用计数排序是非常不合算的（会超出空间限制）。
//在这道题中，我们可以通过一个不难发现的结论来让计数排序变得有用，即：
//
//如果一篇文章的引用次数超过论文的总数 n，那么将它的引用次数降低为 n 也不会改变 h 指数的值。
//
//由于 h 指数一定小于等于 n，因此这样做是正确的。在直方图中，将所有超过 y 轴值大于 n 的变为 n 等价于去掉 y>n 的整个区域。
//
//
//我们用一个例子来说明如何使用计数排序得到 hh 指数。首先，引用次数如下所示：
//
//citations=[1,3,2,3,100]
//
//将所有大于 n=5 的引用次数变为 n，得到：
//
//citations=[1,3,2,3,5]
//
//计数排序得到的结果如下：
//
//     k  0   1   2   3   4   5
//count   0   1   1   2   0   1
//sk      5   5   4   3   1   1
//
//其中 sk表示至少有 k 次引用的论文数量，在表中即为在它之后的列（包括本身）的 count 一行的和。根据定义，最大的满足 k≤sk
//​的 k 即为所求的 h。在表中，这个 k 为 3，因此 h 指数为 3。

func hIndex2(citations []int) int {
	papers := make([]int,len(citations) + 1)
	for i := 0;i<len(citations);i++ {
		papers[Min(citations[i],len(citations))] ++
	}
	k := len(citations)
	for s := papers[k];k>s;s+=papers[k]{
		k--
	}
	return k
}

func Min(x,y int) int {
	if x > y {
		return y
	}
	return x
}