package main

import (
	"sort"
	"strconv"
)

/*
179. 最大数
给定一组非负整数 nums，重新排列它们每位数字的顺序使之组成一个最大的整数。

注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。



示例 1：

输入：nums = [10,2]
输出："210"
示例 2：

输入：nums = [3,30,34,5,9]
输出："9534330"
示例 3：

输入：nums = [1]
输出："1"
示例 4：

输入：nums = [10]
输出："10"


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 109
 */

type data []int

func (d data) Len() int {
	return len(d)
}

func (d data) Less(i,j int) bool {
	return (strconv.Itoa(d[i]) + strconv.Itoa(d[j])) > (strconv.Itoa(d[j]) + strconv.Itoa(d[i]))
}

func (d data) Swap(i,j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func largestNumber(nums []int) string {
	sort.Sort(data(nums))
	if nums[0] == 0 {
		return "0"
	}
	result := ""
	for i := 0;i<len(nums);i++ {
		result += strconv.Itoa(nums[i])
	}
	return result
}

/*

class Solution {
    private class LargerNumberComparator implements Comparator<String> {
        @Override
        public int compare(String a, String b) {
            String order1 = a + b;
            String order2 = b + a;
           return order2.compareTo(order1);
        }
    }

    public String largestNumber(int[] nums) {
        // Get input integers as strings.
        String[] asStrs = new String[nums.length];
        for (int i = 0; i < nums.length; i++) {
            asStrs[i] = String.valueOf(nums[i]);
        }

        // Sort strings according to custom comparator.
        Arrays.sort(asStrs, new LargerNumberComparator());

        // If, after being sorted, the largest number is `0`, the entire number
        // is zero.
        if (asStrs[0].equals("0")) {
            return "0";
        }

        // Build largest number from sorted array.
        String largestNumberStr = new String();
        for (String numAsStr : asStrs) {
            largestNumberStr += numAsStr;
        }

        return largestNumberStr;
    }
}

 */
