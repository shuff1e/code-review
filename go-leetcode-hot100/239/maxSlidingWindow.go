package main

import (
	"fmt"
)

/*
239. 滑动窗口最大值
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。
滑动窗口每次只向右移动一位。

返回滑动窗口中的最大值。

进阶：

你能在线性时间复杂度内解决此题吗？

示例:

输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:

滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
1 [3  -1  -3] 5  3  6  7       3
1  3 [-1  -3  5] 3  6  7       5
1  3  -1 [-3  5  3] 6  7       5
1  3  -1  -3 [5  3  6] 7       6
1  3  -1  -3  5 [3  6  7]      7


提示：

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
1 <= k <= nums.length
 */

func main() {
	arr := []int{1,3,-1,-3,5,3,6,7}
	arr = []int{1,1}
	k := 3
	//arr = []int{}
	fmt.Println(maxSlidingWindow(arr,k))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	result := []int{}
	queue := maxQueue{}
	for i := 0;i<k && i < len(nums);i++ {
		queue.Add(nums[i])
	}
	result = append(result,queue.Max())
	for i := 1;i<=len(nums)-k;i++ {
		if len(queue) > 0 && nums[i-1] == queue.Max() {
			queue.RemoveFirst()
		}
		queue.Add(nums[i+k-1])
		result = append(result,queue.Max())
	}
	return result
}

type maxQueue []int

func (queue *maxQueue) Add(x int) {
	for len(*queue) > 0 && (*queue)[len(*queue)-1] < x {
		*queue = (*queue)[:len(*queue)-1]
	}
	*queue = append(*queue,x)
}

func (queue *maxQueue) Max() int {
	return (*queue)[0]
}

func (queue *maxQueue) RemoveFirst() int {
	result := (*queue)[0]
	*queue = (*queue)[1:]
	return result
}

/*

方法二：双向队列
直觉

如何优化时间复杂度呢？首先想到的是使用堆，因为在最大堆中 heap[0] 永远是最大的元素。在大小为 k 的堆中插入一个元素消耗 log(k) 时间，
因此算法的时间复杂度为 O(Nlog(k))。

能否得到只要 O(N) 的算法？

我们可以使用双向队列，该数据结构可以从两端以常数时间压入/弹出元素。

存储双向队列的索引比存储元素更方便，因为两者都能在数组解析中使用。

算法

算法非常直截了当：

处理前 k 个元素，初始化双向队列。

遍历整个数组。在每一步 :

清理双向队列 :

  - 只保留当前滑动窗口中有的元素的索引。

  - 移除比当前元素小的所有元素，它们不可能是最大的。
将当前元素添加到双向队列中。
将 deque[0] 添加到输出中。
返回输出数组。

class Solution {
  ArrayDeque<Integer> deq = new ArrayDeque<Integer>();
  int [] nums;

  public void clean_deque(int i, int k) {
    // remove indexes of elements not from sliding window
    if (!deq.isEmpty() && deq.getFirst() == i - k)
      deq.removeFirst();

    // remove from deq indexes of all elements
    // which are smaller than current element nums[i]
    while (!deq.isEmpty() && nums[i] > nums[deq.getLast()])
      deq.removeLast();
  }

  public int[] maxSlidingWindow(int[] nums, int k) {
    int n = nums.length;
    if (n * k == 0) return new int[0];
    if (k == 1) return nums;

    // init deque and output
    this.nums = nums;
    int max_idx = 0;
    for (int i = 0; i < k; i++) {
      clean_deque(i, k);
      deq.addLast(i);
      // compute max in nums[:k]
      if (nums[i] > nums[max_idx]) max_idx = i;
    }
    int [] output = new int[n - k + 1];
    output[0] = nums[max_idx];

    // build output
    for (int i  = k; i < n; i++) {
      clean_deque(i, k);
      deq.addLast(i);
      output[i - k + 1] = nums[deq.getFirst()];
    }
    return output;
  }
}

 */

//index: 0 1  2   3 4 5 6 7
//
//arr:  [1,3,-1,-3,5,3,6,7]
//
//left:  1,3, 3,-3,5,5,6,7
//
//right: 3 3  -1 5 5 3,7,7
//
//max:   3 3  5  5 6 7

// offset by one

/*

方法三: 动态规划
直觉

这是另一个 O(N) 的算法。本算法的优点是不需要使用 数组 / 列表 之外的任何数据结构。

算法的思想是将输入数组分割成有 k 个元素的块。
若 n % k != 0，则最后一块的元素个数可能更少。


开头元素为 i ，结尾元素为 j 的当前滑动窗口可能在一个块内，也可能在两个块中。


情况 1 比较简单。 建立数组 left， 其中 left[j] 是从块的开始到下标 j 最大的元素，方向 左->右。


为了处理更复杂的情况 2，我们需要数组 right，其中 right[j] 是从块的结尾到下标 j 最大的元素，方向 右->左。right 数组和 left 除了方向不同以外基本一致。


两数组一起可以提供两个块内元素的全部信息。考虑从下标 i 到下标 j的滑动窗口。 根据定义，right[i] 是左侧块内的最大元素， left[j] 是右侧块内的最大元素。因此滑动窗口中的最大元素为 max(right[i], left[j])。



算法

算法十分直截了当：

从左到右遍历数组，建立数组 left。

从右到左遍历数组，建立数组 right。

建立输出数组 max(right[i], left[i + k - 1])，其中 i 取值范围为 (0, n - k + 1)。

 */

/*

class Solution {
  public int[] maxSlidingWindow(int[] nums, int k) {
    int n = nums.length;
    if (n * k == 0) return new int[0];
    if (k == 1) return nums;

    int [] left = new int[n];
    left[0] = nums[0];
    int [] right = new int[n];
    right[n - 1] = nums[n - 1];
    for (int i = 1; i < n; i++) {
      // from left to right
      if (i % k == 0) left[i] = nums[i];  // block_start
      else left[i] = Math.max(left[i - 1], nums[i]);

      // from right to left
      int j = n - i - 1;
      if ((j + 1) % k == 0) right[j] = nums[j];  // block_end
      else right[j] = Math.max(right[j + 1], nums[j]);
    }

    int [] output = new int[n - k + 1];
    for (int i = 0; i < n - k + 1; i++)
      output[i] = Math.max(left[i + k - 1], right[i]);

    return output;
  }
}

 */