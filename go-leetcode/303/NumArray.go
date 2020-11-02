package main

import "fmt"

/*

303. 区域和检索 - 数组不可变
给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。

实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）


示例：

输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]

解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))


提示：

0 <= nums.length <= 104
-105 <= nums[i] <= 105
0 <= i <= j < nums.length
最多调用 104 次 sumRange 方法

 */

func main() {
	arr := Constructor([]int{-2,0,3,-5,2,-1})
	fmt.Println(arr.SumRange(0,2))
	fmt.Println(arr.SumRange(2,5))
	fmt.Println(arr.SumRange(0,5))
}

type NumArray struct {
	sumArr []int
}

// sum[i]的位置存放了0到i的和
// range(i,j) = sum[j] - sum[i-1]

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	sumArr := make([]int,len(nums))
	sumArr[0] = nums[0]
	for i := 1 ;i<len(nums);i++ {
		sumArr[i] = sumArr[i-1] + nums[i]
	}
	return NumArray{
		sumArr: sumArr,
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sumArr[j]
	}
	return this.sumArr[j] - this.sumArr[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

/*

private int[] sum;

public NumArray(int[] nums) {
    sum = new int[nums.length + 1];
    for (int i = 0; i < nums.length; i++) {
        sum[i + 1] = sum[i] + nums[i];
    }
}

public int sumRange(int i, int j) {
    return sum[j + 1] - sum[i];
}

注意，在上面的代码中，我们插入了一个虚拟 0 作为 sum 数组中的第一个元素。
这个技巧可以避免在 sumrange 函数中进行额外的条件检查。

另外，例如 lc 437
 */