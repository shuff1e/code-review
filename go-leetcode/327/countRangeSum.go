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