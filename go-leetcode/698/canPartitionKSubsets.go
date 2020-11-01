package main

import (
	"fmt"
	"sort"
)

/*

698. 划分为k个相等的子集
给定一个整数数组  nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。

示例 1：

输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。


提示：

1 <= k <= len(nums) <= 16
0 < nums[i] < 10000

 */

// 得到sum，以及max
// 如果max>sum/k ,false

// 1，2，2，3，3，4，5
// 1, 1,1,1,1,1,1 ,3,4

// 1 ,2 ,3 ,4 ,5

func main() {
	arr := []int{10,10,10,7,7,7,7,7,7,6,6,6}
	k := 3
	fmt.Println(canPartitionKSubsets2(arr,k))
}

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) < k {
		return false
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum := 0
	for i := 0;i<len(nums);i++ {
		sum += nums[i]
	}
	if sum % k != 0 {
		return false
	}

	target := sum / k
	if nums[len(nums)-1] > target {
		return false
	}

	end := len(nums) - 1
	for end >= 0 {
		if nums[end] == target {
			end --
			k --
		} else {
			break
		}
	}
	if k  == 0 {
		return true
	}

	result := make([]int,k)
	return help(result,nums,end,target)
}

func help(result []int,arr []int,index int,target int) bool {
	if index < 0 {
		return true
	}

	// 这个v可以放在result的任何一个位置

	// arr[index]放在0的位置
	// 下一层调用的时候，arr[index-1]也可以放在0的位置
	// 如果arr[index-1]也可以放在0的位置不行，就减去arr[index-1]
	// 然后尝试放在1的位置

	v := arr[index]
	index --
	// 这个v放在
	for i := 0;i<len(result);i++ {
		if result[i] + v <= target {
			result[i] += v
			if help(result,arr,index,target) {
				return true
			}
			result[i] -= v
		}
		if result[i] == 0 {
			break
		}
	}
	return false
}

/*

class Solution {
    public boolean search(int[] groups, int row, int[] nums, int target) {
        if (row < 0) return true;
        int v = nums[row--];
        for (int i = 0; i < groups.length; i++) {
            if (groups[i] + v <= target) {
                groups[i] += v;
                if (search(groups, row, nums, target)) return true;
                groups[i] -= v;
            }
            if (groups[i] == 0) break;
        }
        return false;
    }

    public boolean canPartitionKSubsets(int[] nums, int k) {
        int sum = Arrays.stream(nums).sum();
        if (sum % k > 0) return false;
        int target = sum / k;

        Arrays.sort(nums);
        int row = nums.length - 1;
        if (nums[row] > target) return false;
        while (row >= 0 && nums[row] == target) {
            row--;
            k--;
        }
        return search(new int[k], row, nums, target);
    }
}

以数组为[10,10,10,7,7,7,7,7,7,6,6,6]，分成3部分为例，
即分成3个桶，每个桶的和为30

排序后数组为[6,6,6,7,7,7,7,7,7,10,10,10]，需要放在3个桶中，每个桶的和为30

第一层递归，末尾的10放在第一个桶中
|   |   |   |   |   |
|   |	|   |   |   |
|   |	|   |   |   |
|10 |	|   |   |   |
 ---	 ---	 ---
桶1      桶2      桶3

第二层和第三层递归，倒数第二个和倒数第三个10都放在桶1中
|   |   |   |   |   |
|10 |	|   |   |   |
|10 |	|   |   |   |
|10 |	|   |   |   |
 ---	 ---	 ---
桶1      桶2      桶3

第四层递归，倒数第一个7就不能放在桶1中了，因为30+7>30
所以放在桶2中
|   |   |   |   |   |
|10 |	|   |   |   |
|10 |	|   |   |   |
|10 |	|7  |   |   |
 ---	 ---	 ---
桶1      桶2      桶3

后续直到桶2中放了4个7，

|   |   |7  |   |   |
|10 |	|7  |   |   |
|10 |	|7  |   |   |
|10 |	|7  |   |   |
 ---	 ---	 ---
桶1      桶2      桶3

再有7就不能放到桶2中了，因为5*7>30
后续的2个7和2个6放到了桶3中

|   |   |7  |   |6  |
|10 |	|7  |   |6  |
|10 |	|7  |   |7  |
|10 |	|7  |   |7  |
 ---	 ---	 ---
桶1      桶2      桶3

正数第一个6这时没地方放了，因为放到任何一个桶中，都大于30
这时遍历3个桶，都没法放进去之后，返回false

然后递归返回，

递归返回到正数第二个6，正数第二个6从桶3中出栈，但是没有桶4可以让正数第二个6放进去了，
for循环直接结束了，同时返回false
同理，桶3中的元素会依次从桶3中出栈

然后桶2中的栈顶的7，会尝试放到桶3中，再递归下去，（剩余的数组元素为[6,6,6,7,7])
当然我们知道这种情况也是无解的，

|   |   |   |   |   |
|10 |	|7  |   |   |
|10 |	|7  |   |   |
|10 |	|7  |   |7  |
 ---	 ---	 ---
桶1      桶2      桶3

最终桶2中的元素，会全部依次出栈，此时数组中剩余的元素为[6,6,6,7,7,7,7,7,7]

|   |   |   |   |   |
|10 |	|   |   |   |
|10 |	|   |   |   |
|10 |	|   |   |   |
 ---	 ---	 ---
桶1      桶2      桶3

如果没有 if (groups[i] == 0) break; 这行代码

我们知道了桶2已经是空了，但是仍然会运行for循环，把倒数第一个位置的7，放到桶3中再次尝试，并继续递归下去
但是这样其实没有意义，因为桶2和桶3的地位是一样的，
这种情况也是无解的，所以可以剪枝

同样，桶1中的栈顶的2个10都会出栈，最终会平均分配到3个桶中，剩余的元素，也都会平均分配到每个桶中

|6  |   |6  |   |6  |
|7  |	|7  |   |7  |
|7  |	|7  |   |7  |
|10 |	|10 |   |10 |
 ---	 ---	 ---
桶1      桶2      桶3

最后一次递归的时候 row就是-1了，这时直接返回true

*/

func canPartitionKSubsets2(nums []int, k int) bool {
	sum := 0
	for i := 0;i<len(nums);i++ {
		sum += nums[i]
	}
	if sum % k != 0 {
		return false
	}
	// 每个nums[i]都有选择和不选择两种选项，这样一共有2^(len(nums)-1)种选择
	result := make([]int,1<<len(nums))
	for i := 0;i<len(result)-1;i++ {
		result[i] = -1
	}
	// 如果所有的nums[i]都已经被选择了，那么是一种有效的解
	result[len(result)-1] = 1
	return search(0,sum,result,nums,sum / k )
}

// 记忆化dp和状态压缩
func search(used int,todo int,memo []int,arr []int,target int) bool {
	if memo[used] == -1 {
		memo[used] = 0
		// 防止todo % target == 0
		targ := (todo-1)% target + 1
		for i := 0;i<len(arr);i++ {
			if (used & (1 << i)) == 0 && arr[i] <= targ {
				if search(used | (1 << i),todo-arr[i],memo,arr,target) {
					memo[used] = 1
					break
				}
			}
		}
	}
	return memo[used] == 1
}