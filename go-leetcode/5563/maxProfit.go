package main

import "fmt"

/*

5563. 销售价值减少的颜色球
你有一些球的库存 inventory ，里面包含着不同颜色的球。一个顾客想要 任意颜色 总数为 orders 的球。

这位顾客有一种特殊的方式衡量球的价值：每个球的价值是目前剩下的 同色球 的数目。比方说还剩下 6 个黄球，那么顾客买第一个黄球的时候该黄球的价值为 6 。这笔交易以后，只剩下 5 个黄球了，所以下一个黄球的价值为 5 （也就是球的价值随着顾客购买同色球是递减的）

给你整数数组 inventory ，其中 inventory[i] 表示第 i 种颜色球一开始的数目。同时给你整数 orders ，表示顾客总共想买的球数目。你可以按照 任意顺序 卖球。

请你返回卖了 orders 个球以后 最大 总价值之和。由于答案可能会很大，请你返回答案对 109 + 7 取余数 的结果。



示例 1：


输入：inventory = [2,5], orders = 4
输出：14
解释：卖 1 个第一种颜色的球（价值为 2 )，卖 3 个第二种颜色的球（价值为 5 + 4 + 3）。
最大总和为 2 + 5 + 4 + 3 = 14 。
示例 2：

输入：inventory = [3,5], orders = 6
输出：19
解释：卖 2 个第一种颜色的球（价值为 3 + 2），卖 4 个第二种颜色的球（价值为 5 + 4 + 3 + 2）。
最大总和为 3 + 2 + 5 + 4 + 3 + 2 = 19 。
示例 3：

输入：inventory = [2,8,4,10,6], orders = 20
输出：110
示例 4：

输入：inventory = [1000000000], orders = 1000000000
输出：21
解释：卖 1000000000 次第一种颜色的球，总价值为 500000000500000000 。 500000000500000000 对 109 + 7 取余为 21 。


提示：

1 <= inventory.length <= 105
1 <= inventory[i] <= 109
1 <= orders <= min(sum(inventory[i]), 109)

 */

/*

方法一：贪心 + 二分查找
思路与算法

首先，贪心的思路很容易想到：我们每次会找到当前剩余最多的那一类球（如果有多个类剩余的球数相同，那么任意选择一个类即可），然后将一个这类的球卖给顾客。我们连续这样操作 orders 次，就可以卖出最大的价值。

既然我们每一次操作都要「找最大值」，那么我们可以想到使用「优先队列（大根堆）」这一数据结构，它可以使得我们：

在初始时把每一类球的数量全部放入优先队列中；

每一次操作时，取出堆顶的元素并累加入答案，再将其减去 11 放回堆中。

这样做的时间复杂度为 O(orders⋅logn)，而本题中 orders 可以到 10^9，会导致其超出时间限制。那么有什么可以优化的地方呢？

我们可以这样想：由于每次我们都是将当前的最大值减去 1，那么可以看成我们维护了一个「最大值集合」：如果其中有 x 个元素，那么我们需要 x 次操作把它们都减去 1。在这之后，可能会有不在「最大值集合」中的元素现在也变成最大值了，我们就将这些元素也加入集合中，并且继续轮流减去 1，直到进行了 orders 次操作。

orders 次操作看成一个整体，那么一定存在一个「阈值」T（也就是最后「最大值集合」对应的那个值），使得：

初始时所有小于等于 T 的元素都保持不变；

初始时所有大于等于 T 的元素要么变成了 T−1（在「最大值集合」中，并且减去了 1），要么变成了 T（在最大值集合中，但是没来得及被减去 1）。

那么如何求出这个 T 呢？对于每一个元素 ai
​
 ，如果它大于等于 T，那么它被减去 1 的次数要么是 ai−T，要么是 ai−T+1，所以满足题目要求的 T 即为满足

ai ≥ T ∑(ai−T)≤orders< ai ≥ T ∑(ai−T+1)

的 T 值。由于随着 T 的减小，∑ai≥T(ai−T) 是单调递增的，所以满足上述不等式要求的 T 值是唯一的，并且我们可以使用二分查找的方法找出这个 T，即为最小的满足

ai≥T∑(ai−T)≤orders

的 T 值。二分查找的下界为 0，上界为所有 ai中的最大值。

在求出了 T 值之后，我们也可以很方便地算出答案了：令 rest=orders−∑ai≥T(ai−T)，即表示有 rest 个大于等于 T 的元素 ai最后变成了 T−1，
其余的变成了 T。随后我们遍历每一个元素：

如果 ai<T，那么它不会有任何变化，对答案也没有贡献；

如果 ai≥T，那么根据 rest 的值考虑将 [T,ai] 或者 [T+1,ai] 计入答案。

 */

func main() {
	arr := []int{2,8,4,10,6}
	orders := 20

	arr = []int{3,5}
	orders = 6

	arr = []int{1000000000}
	orders = 1000000000
	fmt.Println(maxProfit(arr,orders))
}

const mod = 1000000007

func maxProfit(inventory []int, orders int) int {
	l := 0
	r := getMax(inventory)
	T := -1
	for l <= r {
		mid := (l+r)/2
		sum := getSum(mid,inventory)
		if sum <= orders {
			T = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	rest := orders - getSum(T,inventory)

	result := 0
	for i := 0;i<len(inventory);i++ {
		if inventory[i] < T {
			continue
		}
		if rest > 0 {
			result += rangeSum(T,inventory[i])
			rest --
		} else {
			result += rangeSum(T+1,inventory[i])
		}
	}
	return result % mod
}

func getSum(pivot int,arr []int) int {
	sum := 0
	for i := 0;i<len(arr);i++ {
		if arr[i] > pivot {
			sum += arr[i] - pivot
		}
	}
	return sum
}

func rangeSum(pivot int,x int) int {
	return (pivot + x) * (x - pivot + 1) / 2
}

func getMax(arr []int) int {
	max := arr[0]
	for i := 1;i<len(arr);i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}