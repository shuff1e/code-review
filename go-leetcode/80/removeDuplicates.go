package main

import "fmt"

/*
80. 删除排序数组中的重复项 II
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定 nums = [1,1,1,2,2,3],

函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,1,2,3,3],

函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

你不需要考虑数组中超出新长度后面的元素。
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
print(nums[i]);
}
 */

// index指针指向下一个要填充元素的位置

func main() {
	arr := []int{1,1,1,2,2,3}
	fmt.Println(removeDuplicates(arr))
	arr = []int{0,0,1,1,1,1,2,3,3}
	fmt.Println(removeDuplicates(arr))
	arr = []int{1,2,3,4,4,5,5,5,5,6,6,7,7,7}
	fmt.Println(removeDuplicates(arr))
	fmt.Printf("%#v\n",arr)
}

func removeDuplicates(nums []int) int {
	index := -1
	curCount := 0
	for i := 0;i<len(nums);i++ {
		if i == 0 {
			curCount = 1
		} else if nums[i] == nums[i-1] {
			curCount += 1
		} else {
			curCount = 1
		}
		if curCount <= 2 {
			index ++
			nums[index] = nums[i]
		}
	}
	return index + 1
}
