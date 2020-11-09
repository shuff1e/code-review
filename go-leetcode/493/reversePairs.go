package main

import (
	"fmt"
	"sort"
)

/*

493. 翻转对
给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。

你需要返回给定数组中的重要翻转对的数量。

示例 1:

输入: [1,3,2,3,1]
输出: 2
示例 2:

输入: [2,4,3,5,1]
输出: 3
注意:

给定数组的长度不会超过50000。
输入数组中的所有数字都在32位整数的表示范围内。

 */

/*

https://mingshan.fun/2019/11/29/binary-indexed-tree/

https://oi-wiki.org/ds/fenwick/

 */

func main() {
	arr := []int{2,4,3,5,1}
	arr = []int{1,3,2,3,1}
	fmt.Println(reversePairs(arr))
}

type Bit struct {
	n int
	nodes []int
}

func (b *Bit) update(x int) {
	for x <= b.n {
		b.nodes[x] ++
		x += lowBit(x)
	}
}

func (b *Bit) query(x int) int {
	result := 0
	for x > 0 {
		result += b.nodes[x]
		x -= lowBit(x)
	}
	return result
}

func NewBit(n int) *Bit {
	return &Bit{
		n: n,
		nodes: make([]int,n+1),
	}
}

func lowBit(x int) int {
	return x & (-x)
}

func reversePairs(nums []int) int {
	dict := map[int]struct{}{}
	for i := 0;i<len(nums);i++ {
		dict[nums[i]] = struct{}{}
	}

	temp := []int{}

	for k,_ := range dict {
		temp = append(temp,k)
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})

	// 离散化
	// 插入的数字，变为1，2，... 直到 len(temp)
	// 因为有 for x > 0
	// 2,4,3,5,1
	b := NewBit(len(temp))
	result := 0
	for i := 0;i<len(nums);i++ {
		x := lowBound(temp,2*nums[i]) + 1
		result += b.query(b.n) - b.query(x)
		x = lowBound(temp,nums[i]) + 1
		b.update(x)
	}
	return result
}

func lowBound(arr []int,x int) int {
	left,right := 0,len(arr)-1
	for left <= right {
		mid := (left + right)/2
		if arr[mid] > x {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return right
}

/*

解题思路
题目求的是i < j && nums[i] > 2 * nums[j]的数目对，
换个位置2 * nums[j] < nums[i] ， 也就是求，此时走到下标j， 统计 j 前面 大于2 * nums[j]的数的个数
区间查询， 单点修改，这个是树状数组的经典例题了。

但是由于本题数据范围比较大， 使用树状数组前， 需要先离散化一下。

1. 离散化
把nums[j]和nums[j] * 2， 所有数组加入离散化数组ys，接下来是离散化常规三连。
a. 排序 Collections.sort(ys);
b. 去重 unique(ys);
c. 二分查找 binaryFind(ys, target);

2. 树状数组
离散化后， 开一个跟ys.size() + 1大小的数组用于存放树状数组。
接下来也是常规三连。
a. 求数字x对应二进制最右边为1的数字 lowBit(x)
b. 树状数组单点增加 add 操作
c. 树状数组前缀和查询 query(x);

时间复杂度： O(NlogN)

class Solution {
    int N;
    long[] tr;
    public int reversePairs(int[] nums) {
        List<Long> ys = new ArrayList();
        for(int i: nums) {
            ys.add((long)i);
            ys.add((long)i * 2);
        }
        Collections.sort(ys);
        ys = unique(ys);
        N = ys.size();
        tr = new long[N + 1];
        int ans = 0;
        for(int i = 0; i < nums.length; i++){
            long target = (long)nums[i] * 2;
            int left = binaryFind(ys, target) + 1;
            int right = N;
            ans += query(right) - query(left);
            add(binaryFind(ys, nums[i]) + 1, 1);
        }
        return ans;
    }

    public void add(int x, int c){
        for(int i = x; i <= N; i += lowBit(i)) tr[i] += c;
    }

    public int query(int x){
        int res = 0;
        for(int i = x; i > 0; i -= lowBit(i)) res += tr[i];
        return res;
    }

    public int lowBit(int x){return x & -x;}



    public List<Long> unique(List<Long> list){
        List<Long> res = new ArrayList(list.size());
        for(int i = 0; i < list.size(); i++){
            if(res.isEmpty() || res.get(res.size() - 1) - list.get(i) != 0){
                res.add(list.get(i));
            }
        }
        return res;
    }

    public int binaryFind(List<Long> list, long target){
        int l = 0, r = list.size() - 1;
        while(l < r){
            int mid = l + r >> 1;
            if(list.get(mid) >= target) r = mid;
            else l = mid + 1;
        }
        return l;
    }
}

 */

/*

方法三：归并排序MergeSort
分治
count：统计符合题意 arr[idx] > 2arr[j] 的个数
i：用于合并数组时，左侧数组的指针
j：用于合并数组时，右侧数组的指针；且用于查找符合 arr[idx] > 2arr[j] 时，右侧数组的指针
k：合并后的数组的指针
idx：用于查找符合 arr[idx] > 2*arr[j] 时，左侧数组的指针
归并，合并两个有序数组：左侧数组 l - mid；右侧数组 mid+1 - r
2.1.合并左侧数组
2.2.合并右侧数组
2.3.将左侧数组剩余的未合并的“大数”移至 temp 的末尾（由于temp最终要考到到原数组中，所以可以直接移至原数组 l+k: 的位置）
查询左侧数组中符合 arr[idx] > 2arr[j] 的数
3.1.找出当前左侧数组中，符合 arr[idx] > 2arr[j] 的idx位置；arr[idx+1],arr[idx+2],...,arr[mid]，都符合 > 2arr[j]
arr[idx]+1)>>1 <= arr[j]：+1 降低溢出风险（偷个懒）
3.2.累加 arr[idx] > 2arr[j] 个数：idx-mid 之间的数，都符合 > 2*arr[j]
将合并后的数组拷贝回原数组

func reversePairs(nums []int) int {
	// 归并排序
	return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(arr []int, l int, r int) int {
	if l >= r {
		return 0
	}
	count, mid := 0, (l+r)>>1
	count = mergeSort_(arr, l, mid) + mergeSort_(arr, mid+1, r) // 1
	temp, i, k := make([]int, r-l+1), l, 0
	// 对于右边的一个位置j
	// 左边小于该位置的放到temp
	// 从最左边的位置，开始计算，直到 (arr[idx] + 1) / 2 > arr[j]
	for j, idx := mid+1, l; j <= r; j++ { // 2
		for ; i <= mid && arr[i] < arr[j]; i++ {
			temp[k], k = arr[i], k+1 // 2.1
		}
		temp[k], k = arr[j], k+1                      // 2.2
		for idx <= mid && (arr[idx]+1)>>1 <= arr[j] { // 3
			idx++ // 3.1
		}
		count += mid - idx + 1 // 3.2
	}
	copy(arr[l+k:], arr[i:mid+1]) // 2.3
	copy(arr[l:], temp[:k])       // 4
	return count
}

 */

/*

方法二：树状数组BIT
构建原数组 + 原数组*2 的集合（去重），并排序
1.1.构建（集合 cache）
1.2.排序（数组 a）
构建BIT
2.1.构建树状数组 c，长度为数组a长度 +1
2.2.query：求 a[1 - i] 的和
2.3.update：在 a[i] 位置加上 k，此处 k=1
遍历原数组，尾先入（后入的元素就可以判断，有多少个先入的元素是小于该后入元素）
3.1.二分查找元素排序后的位置（首次出现）
idx1：用于查询树状数组
idx2：用于更新树状数组
3.2.查询树状数组
3.3.更新树状数组，为什么 idx+1？
lowbit：取最后一位 1，所以需要 idx > 0

func reversePairs(nums []int) int {
	// 树状数组：BIT
	cache := make(map[int]struct{}) // 1
	for _, v := range nums {        // 1.1
		cache[v] = struct{}{}
		cache[v<<1] = struct{}{}
	}
	n := len(cache)
	a, c, count, i := make([]int, n), make([]int, n+1), 0, 0 // 2.1
	for k, _ := range cache {                                // 1.2
		a[i], i = k, i+1
	}
	sort.Ints(a)
	for i := len(nums) - 1; i >= 0; i-- { // 3
		idx1, idx2 := binarySearch(a, 0, n, nums[i]), binarySearch(a, 0, n, nums[i]<<1) // 3.1
		count += query(c, idx1)                                                         // 3.2
		update(c, idx2+1)                                                               // 2 & 3.3
	}
	return count
}
func binarySearch(arr []int, l, r int, val int) int {
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
func update(bit []int, i int) { // 2.3
	for i < len(bit) {
		bit[i] ++
		i += i & -i
	}
}
func query(bit []int, i int) int { // 2.2
	count := 0
	for i > 0 {
		count += bit[i]
		i -= i & -i
	}
	return count
}

 */

/*

方法一：二叉搜索树BST
构建struct并创建root，默认count=1
查询BST
2.1.到叶子节点
2.2.节点 val <= 目标，说明要往 left 节点查询
2.3.节点 val > 目标，累加左子树的个数，并往 right 节点查询
BST插入元素
3.1.要插入的节点不存在，新建节点并返回该节点
3.2.存在，返回该节点
3.3.插入 val > 节点val，往右插入
3.4.插入 val < 节点val，节点count++（左子树节点个数+1），往左插入

func reversePairs(nums []int) int {
	// BST：超时（数组如果单调递增，就退化为链表）
	ans, n := 0, len(nums)
	if n > 0 {
		root := &Node{nil, nil, nums[n-1], 1} // 1
		for i := n - 2; i >= 0; i-- {
			ans += queryBST(root, (nums[i]+1)>>1) // 2
			insertBST(root, nums[i])              // 3
		}
	}
	return ans
}
func insertBST(root *Node, val int) *Node {
	if root == nil { // 3.1
		return &Node{nil, nil, val, 1}
	}
	switch {
	case root.val == val: // 3.2
		root.count++
	case root.val < val: // 3.3
		root.r = insertBST(root.r, val)
	case root.val > val: // 3.4
		root.count++
		root.l = insertBST(root.l, val)
	}
	return root
}
func queryBST(root *Node, half int) int {
	if root == nil { // 2.1
		return 0
	}
	if root.val >= half { // 2.2
		return queryBST(root.l, half)
	} else { // 2.3
		return root.count + queryBST(root.r, half)
	}
}
type Node struct {
	l, r       *Node
	val, count int
}

 */