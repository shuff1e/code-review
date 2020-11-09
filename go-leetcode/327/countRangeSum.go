package main

import (
	"fmt"
	"sort"
)

/*

327. 区间和的个数
给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。
区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

说明:
最直观的算法复杂度是 O(n2) ，请在此基础上优化你的算法。

示例:

输入: nums = [-2,5,-1], lower = -2, upper = 2,
输出: 3
解释: 3个区间分别是: [0,0], [2,2], [0,2]，它们表示的和分别为: -2, -1, 2。

 */

func main() {
	arr := []int{-2,5,-1}
	lower := -2
	upper := 2
	fmt.Println(countRangeSum(arr,lower,upper))
}

func countRangeSum(nums []int, lower int, upper int) int {
	preSum := make([]int,len(nums) + 1)
	sum := 0

	for i := 0;i<len(nums);i++ {
		sum += nums[i]
		preSum[i+1] = sum
	}

	// 离散化
	allNumbersHelp := []int{}
	for i := 0;i<len(preSum);i++ {
		allNumbersHelp = append(allNumbersHelp,
			preSum[i],
			preSum[i] - lower,
			preSum[i] - upper)
	}

	// 去重
	dict := map[int]struct{}{}
	for i := 0;i<len(allNumbersHelp);i++ {
		dict[allNumbersHelp[i]] = struct{}{}
	}
	// 排序
	allNumbers := make([]int,0)
	for k,_ := range dict {
		allNumbers = append(allNumbers,k)
	}
	sort.Slice(allNumbers, func(i, j int) bool {
		return allNumbers[i] < allNumbers[j]
	})

	result := 0
	st := NewSegTree(len(allNumbers))
	for i := 0;i<len(preSum);i++ {
		// 查询也是查询离散化之后的
		left := lowBound(allNumbers,preSum[i] - upper)
		right := lowBound(allNumbers,preSum[i] - lower)
		result += st.Query(left,right)
		// 插入也是插入离散化之后的
		st.Update(lowBound(allNumbers,preSum[i]))
	}
	return result
}

func lowBound(arr []int,x int ) int {
	left := 0
	right := len(arr)-1
	for left <= right {
		mid := (left + right)/2
		if x < arr[mid] {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return right
}

type segTree struct {
	n int
	nodes []int
}

func NewSegTree(n int) *segTree {
	return &segTree{
		n: n,
		nodes: make([]int,4*n),
	}
}

func (s *segTree) Update(x int) {
	s.update(0,0,s.n,x)
}

func (s *segTree) Query(left,right int) int {
	return s.query(0,0,s.n,left,right)
}

func (s *segTree) query(index int,l,r int,ql,qr int) int {
	if qr < l || r < ql {
		return 0
	}
	if ql <= l && r <= qr {
		return s.nodes[index]
	}
	mid := (l + r) / 2
	return s.query(2*index+1,l,mid,ql,qr) + 
		s.query(2*index+2,mid+1,r,ql,qr)
}

func (s *segTree) update(index int,l,r int,x int) {
	if l > x || r < x {
		return
	}
	s.nodes[index] ++
	if l == r {
		return
	}
	mid := (l + r) / 2
	s.update(2*index+1,l,mid,x)
	s.update(2*index+2,mid+1,r,x)
}

/*

https://leetcode-cn.com/problems/count-of-range-sum/solution/xian-ren-zhi-lu-ru-he-xue-xi-ke-yi-jie-jue-ben-ti-/

https://juejin.im/post/6858619792157638670

nums = [-2,5,-1], lower = -2, upper = 2,
preSum = [0,-2,3,2]
allNumbers = [0,2,-2,   -2,0,-4,   3,5,1,   2,4,0]


class Solution {
    public int countRangeSum(int[] nums, int lower, int upper) {
        long sum = 0;
        long[] preSum = new long[nums.length + 1];
        for (int i = 0; i < nums.length; ++i) {
            sum += nums[i];
            preSum[i + 1] = sum;
        }


        // 对于preSum 0，放入三个数字，0，2，-2
        // 当你寻找2的时候，可以找到这个地方，寻找-2的时候，也可以找到这个地方

        // nums = [-2,5,-1]
        // preSum = [0,-2,3,2]
        // 5这个数字，会出现在preSum[1] 和 preSum[2]对应的范围内

        Set<Long> allNumbers = new TreeSet<Long>();
        for (long x : preSum) {
            allNumbers.add(x);
            allNumbers.add(x - lower);
            allNumbers.add(x - upper);
        }

        // 利用哈希表进行离散化
        Map<Long, Integer> values = new HashMap<Long, Integer>();
        int idx = 0;
        for (long x : allNumbers) {
            values.put(x, idx);
            idx++;
        }

        SegNode root = build(0, values.size() - 1);
        int ret = 0;
        for (long x : preSum) {
            int left = values.get(x - upper), right = values.get(x - lower);
            ret += count(root, left, right);
            insert(root, values.get(x));
        }
        return ret;
    }

    public SegNode build(int left, int right) {
        SegNode node = new SegNode(left, right);
        if (left == right) {
            return node;
        }
        int mid = (left + right) / 2;
        node.lchild = build(left, mid);
        node.rchild = build(mid + 1, right);
        return node;
    }

    public int count(SegNode root, int left, int right) {
        if (left > root.hi || right < root.lo) {
            return 0;
        }
        if (left <= root.lo && root.hi <= right) {
            return root.add;
        }
        return count(root.lchild, left, right) + count(root.rchild, left, right);
    }

    public void insert(SegNode root, int val) {
        root.add++;
        if (root.lo == root.hi) {
            return;
        }
        int mid = (root.lo + root.hi) / 2;
        if (val <= mid) {
            insert(root.lchild, val);
        } else {
            insert(root.rchild, val);
        }
    }
}

class SegNode {
    int lo, hi, add;
    SegNode lchild, rchild;

    public SegNode(int left, int right) {
        lo = left;
        hi = right;
        add = 0;
        lchild = null;
        rchild = null;
    }
}

 */

/*

频数数组
很多数据结构都是基于「频数数组」。

给定数组 t 以及它的下标范围 [L,R]，t[x] 就表示元素 x 在数据结构中的出现次数。基于此，上述的两个操作可以变为：

操作 1「查询」：给定一个范围 [left,right]，查询 t[left] 到 t[right] 的和；

操作 2「更新」：给定一个元素 x，将 t[x] 增加 1。

这也是线段树和树状数组的基础，它们需要的空间都与数组 t 的下标范围 [L,R] 正相关。在本题数据规模较大的情况下（例如测试数据中，出现了元素值达到 32 位有符号整数的上下界），线段树和树状数组都会超出空间限制，因此需要借助「离散化」操作，将这些元素映射到一个较小规模的区间内。

离散化
给定数组元素 [1,22,333,4444,55555]，如果我们只关注它们之间的大小关系，那么该数组其实和 [1,2,3,4,5] 是等价的。

这就是离散化的技巧。我们将所有会涉及到比较操作的数全部放入一个列表中，进行排序，再从小到大依次给它们赋予一个新的值。在离散化完成后，任何两个数之间的相对大小都不会改变，但是元素的上下界范围被限制住，这使得我们可以方便地使用一些数据结构。

在本题中，我们可以将所有的 P[j],P[j]−upper,P[j]−lower 一起进行离散化，并将它们从小到大，依次赋予一个从 1 开始的整数值。这样一来，我们所有涉及到的元素值都会在 [1,3(n+1)] 的范围内。

线段树
当我们将元素离散化后，就可以直接使用线段树了。最基础的线段树恰好就支持这两种操作：

操作 1「查询」：给定一个范围 [left,right]，查询 t[left] 到 t[right] 的和；

操作 2「更新」：给定一个元素 x，将 t[x] 增加 δ。

我们只需要时刻令 δ=1 即可。两种操作的时间复杂度均为 O(logn)。

 */