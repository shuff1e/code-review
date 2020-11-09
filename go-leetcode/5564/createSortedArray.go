package main

import "fmt"

/*

5564. 通过指令创建有序数组
给你一个整数数组 instructions ，你需要根据 instructions 中的元素创建一个有序数组。
一开始你有一个空的数组 nums ，你需要 从左到右 遍历 instructions 中的元素，
将它们依次插入 nums 数组中。每一次插入操作的 代价 是以下两者的 较小值 ：

nums 中 严格小于  instructions[i] 的数字数目。
nums 中 严格大于  instructions[i] 的数字数目。
比方说，如果要将 3 插入到 nums = [1,2,3,5] ，
那么插入操作的 代价 为 min(2, 1) (元素 1 和  2 小于 3 ，元素 5 大于 3 ），插入后 nums 变成 [1,2,3,3,5] 。

请你返回将 instructions 中所有元素依次插入 nums 后的 总最小代价 。
由于答案会很大，请将它对 10^9 + 7 取余 后返回。



示例 1：

输入：instructions = [1,5,6,2]
输出：1
解释：一开始 nums = [] 。
插入 1 ，代价为 min(0, 0) = 0 ，现在 nums = [1] 。
插入 5 ，代价为 min(1, 0) = 0 ，现在 nums = [1,5] 。
插入 6 ，代价为 min(2, 0) = 0 ，现在 nums = [1,5,6] 。
插入 2 ，代价为 min(1, 2) = 1 ，现在 nums = [1,2,5,6] 。
总代价为 0 + 0 + 0 + 1 = 1 。
示例 2:

输入：instructions = [1,2,3,6,5,4]
输出：3
解释：一开始 nums = [] 。
插入 1 ，代价为 min(0, 0) = 0 ，现在 nums = [1] 。
插入 2 ，代价为 min(1, 0) = 0 ，现在 nums = [1,2] 。
插入 3 ，代价为 min(2, 0) = 0 ，现在 nums = [1,2,3] 。
插入 6 ，代价为 min(3, 0) = 0 ，现在 nums = [1,2,3,6] 。
插入 5 ，代价为 min(3, 1) = 1 ，现在 nums = [1,2,3,5,6] 。
插入 4 ，代价为 min(3, 2) = 2 ，现在 nums = [1,2,3,4,5,6] 。
总代价为 0 + 0 + 0 + 0 + 1 + 2 = 3 。
示例 3：

输入：instructions = [1,3,3,3,2,4,2,1,2]
输出：4
解释：一开始 nums = [] 。
插入 1 ，代价为 min(0, 0) = 0 ，现在 nums = [1] 。
插入 3 ，代价为 min(1, 0) = 0 ，现在 nums = [1,3] 。
插入 3 ，代价为 min(1, 0) = 0 ，现在 nums = [1,3,3] 。
插入 3 ，代价为 min(1, 0) = 0 ，现在 nums = [1,3,3,3] 。
插入 2 ，代价为 min(1, 3) = 1 ，现在 nums = [1,2,3,3,3] 。
插入 4 ，代价为 min(5, 0) = 0 ，现在 nums = [1,2,3,3,3,4] 。
​​​​​插入 2 ，代价为 min(1, 4) = 1 ，现在 nums = [1,2,2,3,3,3,4] 。
插入 1 ，代价为 min(0, 6) = 0 ，现在 nums = [1,1,2,2,3,3,3,4] 。
插入 2 ，代价为 min(2, 4) = 2 ，现在 nums = [1,1,2,2,2,3,3,3,4] 。
总代价为 0 + 0 + 0 + 0 + 1 + 0 + 1 + 0 + 2 = 4 。


提示：

1 <= instructions.length <= 105
1 <= instructions[i] <= 105

 */

func main() {
	arr := []int{1,3,3,3,2,4,2,1,2}
	fmt.Println(createSortedArray(arr))
}

type segTree struct {
	n int
	nodes []int
}

func New(n int) *segTree {
	return &segTree{
		n:n,
		nodes: make([]int,4*n),
	}
}

func (s *segTree) Update(x int) {
	s.update(0,1,s.n,x)
}

func (s *segTree) Query(left,right int) int {
	return s.query(0,1,s.n,left,right)
}

func (s *segTree) query(index int,l,r int,ql,qr int) int {
	if l > qr || r < ql {
		return 0
	}
	if ql <= l && r <= qr {
		return s.nodes[index]
	}

	mid := (l + r) / 2
	return s.query(2*index+1,l,mid,ql,qr) +
		s.query(2*index+2,mid+1,r,ql,qr)
}

func (s *segTree) update( index int,l,r int,x int) {
	if l > x || r < x {
		return
	}
	s.nodes[index] ++
	if l == r {
		return
	}
	mid := (l+r)/2
	s.update(2*index+1,l,mid,x)
	s.update(2*index+2,mid+1,r,x)
}

const mod = 1000000007

func createSortedArray(instructions []int) int {
	max := getMax(instructions)
	s := New(max)
	result := 0
	for i := 0;i<len(instructions);i++ {
		x := instructions[i]
		smaller := s.Query(1,x-1)
		larger := s.Query(x+1,s.n)
		result += Min(smaller,larger)
		s.Update(x)
	}
	return result % mod
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

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

/*

https://juejin.im/post/6858619792157638670

public class SegmentTree<E> {
    private E[] tree; //线段树
    private E[] data; //数据
    private Merger<E> merger;//融合器

    public SegmentTree(E[] arr, Merger<E> merger) {
        this.merger = merger;
        data = (E[]) new Object[arr.length];
        tree = (E[]) new Object[arr.length * 4]; //大小为 4 * n
        for (int i = 0; i < arr.length; i++) {
            data[i] = arr[i];
        }
        //构建线段树
        buildSegmentTree(0, 0, data.length - 1);
    }

    // 返回数组元素个数
    public int getSize() {
        return data.length;
    }

    // 根据索引获取数据
    public E get(int index) {
        if (index < 0 || index > data.length)
            throw new IllegalArgumentException("Index is illegal");
        return data[index];
    }

    //根据一个节点的索引 index，返回这个节点的左孩子的索引
    private int leftChild(int index) {
        return 2 * index + 1;
    }

    //根据一个节点的索引 index，返回这个节点的右孩子的索引
    private int rightChild(int index) {
        return 2 * index + 2;
    }

    // 在 treeIndex 的位置创建表示区间 [l,r] 的线段树
    private void buildSegmentTree(int treeIndex, int l, int r) {
        // base case：递归到叶子节点了
        if (l == r) {
            tree[treeIndex] = data[l];
            return;
        }

        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        //划分区间
        int mid = l + (r - l) / 2;
        // 求（左孩子）左区间的统计值
        buildSegmentTree(leftTreeIndex, l, mid);
        // 求（右孩子）右区间的统计值
        buildSegmentTree(rightTreeIndex, mid + 1, r);
        //求当前节点 [左区间+右区间] 的统计值
        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    //查询区间，返回区间 [queryL, queryR] 的统计值
    public E query(int queryL, int queryR) {
        //首先进行边界检查
        if (queryL < 0 || queryL > data.length || queryR < 0 || queryR > data.length || queryL > queryR) {
            throw new IllegalArgumentException("Index is illegal");
        }
        return query(0, 0, data.length - 1, queryL, queryR);
    }

    //在以 treeIndex 为根的线段树中 [l,r] 的范围里，搜索区间 [queryL, queryR]
    private E query(int treeIndex, int l, int r, int queryL, int queryR) {
        if (l == queryL && r == queryR) {
            return tree[treeIndex];
        }
        int mid = l + (r - l) / 2;
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        // 如果左边界大于中间节点，则查询右区间
        if (queryL > mid)
            return query(rightTreeIndex, mid + 1, r, queryL, queryR);
        // 如果右边界小于等于中间节点，则查询左区间
        if (queryR <= mid)
            return query(leftTreeIndex, l, mid, queryL, queryR);
        // 如果上述两种情况都不是，则根据中间节点，拆分为两个查询区间
        E leftResult = query(leftTreeIndex, l, mid, queryL, mid);
        E rightResult = query(rightTreeIndex, mid + 1, r, mid + 1, queryR);
        //合并左右区间的查询结果
        return merger.merge(leftResult, rightResult);
    }
    //将 index 位置的值，更新为 e
    public void update(int index, E e) {
        if (index < 0 || index >= data.length)
            throw new IllegalArgumentException("Index is illegal");
        data[index] = e;
        //更新线段树相应的节点
        updateTree(0, 0, data.length - 1, index, e);
    }

    // 在以 treeIndex 为根的线段树中，更新 index 的值为 e
    private void updateTree(int treeIndex, int l, int r, int index, E e) {
        //递归终止条件
        if (l == r) {
            tree[treeIndex] = e;
            return;
        }
        int mid = l + (r - l) / 2;
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);
        if (index > mid)
            updateTree(rightTreeIndex, mid + 1, r, index, e);
        else //index <= mid
            updateTree(leftTreeIndex, l, mid, index, e);
        //更新当前节点
        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    public String toString() {
        StringBuffer res = new StringBuffer();
        res.append('[');
        for (int i = 0; i < tree.length; i++) {
            if (tree[i] != null)
                res.append(tree[i]);
            else res.append("null");
            if (i != tree.length - 1)
                res.append(", ");
        }
        res.append(']');
        return res.toString();
    }
}

 */

