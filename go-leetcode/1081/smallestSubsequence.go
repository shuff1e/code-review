package main

import (
	"fmt"
	"unsafe"
)

/*

1081. 不同字符的最小子序列
返回字符串 text 中按字典序排列最小的子序列，该子序列包含 text 中所有不同字符一次。

示例 1：

输入："cdadabcc"
输出："adbc"
示例 2：

输入："abcd"
输出："abcd"
示例 3：

输入："ecbacba"
输出："eacb"
示例 4：

输入："leetcode"
输出："letcod"


提示：

1 <= text.length <= 1000
text 由小写英文字母组成


注意：本题目与 316 https://leetcode-cn.com/problems/remove-duplicate-letters/ 相同

 */

func main() {
	str := "cdadabcc"
	str = "leetcode"
	str = "ecbacba"
	str = "cbaacabcaaccaacababa"
	fmt.Println(smallestSubsequence(str))
}

func smallestSubsequence(s string) string {
	last_appear := map[byte]int{}
	for i := 0;i<len(s);i++ {
		last_appear[s[i]] = i
	}

	seen := map[byte]struct{}{}

	stack := []byte{}
	for i := 0;i<len(s);i++ {
		if _,ok := seen[s[i]];!ok {
			for len(stack) > 0 &&
				stack[len(stack)-1] > s[i] &&
				last_appear[stack[len(stack)-1]] > i {

				delete(seen,stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			seen[s[i]] = struct{}{}
			stack = append(stack,s[i])
		}
	}
	return String(stack)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}