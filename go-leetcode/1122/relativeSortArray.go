package main

import (
	"fmt"
	"sort"
)

/*

1122. 数组的相对排序
给你两个数组，arr1 和 arr2，

arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。



示例：

输入：arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
输出：[2,2,2,1,4,3,3,9,6,7,19]


提示：

arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中

 */

func main() {
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	fmt.Printf("%#v\n",relativeSortArray2(arr1,arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	rank := map[int]int{}
	for i := 0;i<len(arr2);i++ {
		rank[arr2[i]] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		index1,ok1 := rank[arr1[i]]
		index2,ok2 := rank[arr1[j]]
		if ok1 && ok2 {
			return index1 < index2
		}
		if ok1 {
			return true
		}
		if ok2 {
			return false
		}
		return arr1[i] < arr1[j]
	})

	return arr1
}

/*

方法一：自定义排序
一种容易想到的方法是使用排序并自定义比较函数。

由于数组arr2规定了比较顺序，因此我们可以使用哈希表对该顺序进行映射：即对于数组 arr2 中的第 i 个元素，我们将 (arr2[i],i) 这一键值对放入哈希表 rank 中，就可以很方便地对数组 arr1中的元素进行比较。

比较函数的写法有很多种，例如我们可以使用最基础的比较方法，对于元素 x 和 y：

如果 x 和 y 都出现在哈希表中，那么比较它们对应的值 \textit{rank}[x]rank[x] 和 \textit{rank}[y]rank[y]；

如果 x 和 y 都没有出现在哈希表中，那么比较它们本身；

对于剩余的情况，出现在哈希表中的那个元素较小。

 */

/*

方法二：计数排序
思路与算法

注意到本题中元素的范围为 [0,1000]，这个范围不是很大，我们也可以考虑不基于比较的排序，例如「计数排序」。

具体地，我们使用一个长度为 10011001（下标从 00 到 10001000）的数组 frequency，记录每一个元素在数组 arr1中出现的次数。随后我们遍历数组 arr2
​
，当遍历到元素 x 时，我们将 frequency[x] 个 x 加入答案中，并将 frequency[x] 清零。当遍历结束后，所有在 arr2
​
中出现过的元素就已经有序了。

此时还剩下没有在 arr2中出现过的元素，因此我们还需要对整个数组 frequency 进行一次遍历。当遍历到元素 x 时，如果 frequency[x] 不为 0，我们就将frequency[x] 个 x 加入答案中。

细节

我们可以对空间复杂度进行一个小优化。实际上，我们不需要使用长度为 10011001 的数组，而是可以找出数组 arr1 中的最大值 upper，使用长度为 upper+1 的数组即可。

 */

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	max := getMax(arr1)

	frequency := make([]int,max+1)
	for i := 0;i<len(arr1);i++ {
		frequency[arr1[i]] ++
	}

	result := []int{}
	for i := 0;i<len(arr2);i++ {
		for frequency[arr2[i]] > 0 {
			result = append(result,arr2[i])
			frequency[arr2[i]] --
		}
	}

	for i := 0;i<len(frequency);i++ {
		for frequency[i] > 0 {
			result = append(result,i)
			frequency[i] --
		}
	}

	return result
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