package main

import (
	"bytes"
	"fmt"
)

/*

395. 至少有K个重复字符的最长子串
找到给定字符串（由小写字符组成）中的最长子串 T ， 要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。

示例 1:

输入:
s = "aaabb", k = 3

输出:
3

最长子串为 "aaa" ，其中 'a' 重复了 3 次。
示例 2:

输入:
s = "ababbc", k = 2

输出:
5

最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

 */

// 遍历字符串
// 找到出现次数小于k的字符，那这个字符串中包含这些字符的，肯定不行
// 以这些字符，将字符串再分割成各个子串

// 递归调用

func main() {
	s := "ababbc"
	k := 2

	s = ""
	k = 3

	fmt.Println(longestSubstring(s,k))
}

func longestSubstring(s string, k int) int {
	return help([]byte(s),k)
}

func help(arr []byte,k int) int {
	if len(arr) == 0 {
		return 0
	}

	dict := make(map[byte]int,0)
	for i := 0;i<len(arr);i++ {
		dict[arr[i]] = dict[arr[i]] + 1
	}

	for i := 0;i<len(arr);i++ {
		if dict[arr[i]] < k {
			arr[i] = ','
		}
	}

	slice := bytes.Split(arr,[]byte{','})
	if len(slice) ==  1 {
		return len(slice[0])
	}

	result := 0

	for i := 0;i<len(slice);i++ {
		result = Max(result,help(slice[i],k))
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}