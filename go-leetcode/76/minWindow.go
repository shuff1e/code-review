package main

import "fmt"

/*
76. 最小覆盖子串
给你一个字符串 S、一个字符串 T 。
请你设计一种算法，可以在 O(n) 的时间复杂度内，从字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：

输入：S = "ADOBECODEBANC", T = "ABC"
输出："BANC"


提示：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

 */

// A：滑动窗口
// l,r
// r一直向右延伸，当l,r是一个符合要求的窗口之后，更新length，并看能否向右移动l，减少窗口的长度
//

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s,t))
}

func minWindow(s string, t string) string {
	dictT := map[byte]int{}
	dictS := map[byte]int{}
	for i := 0;i<len(t);i++ {
		dictT[t[i]] = dictT[t[i]] + 1
	}

	l,r := 0,0
	ansL,ansR := -1,-1
	length := 0x7fffffff
	for ;r<len(s);r++ {
		if _,ok := dictT[s[r]];ok {
			dictS[s[r]] = dictS[s[r]] + 1
			for checkValid(dictS,dictT) && l<=r {
				if r-l+1 <length {
					ansL = l
					ansR = r
					length = r-l+1
				}
				if _,ok := dictT[s[l]];!ok {
					l++
				} else if dictS[s[l]] > dictT[s[l]] {
					dictS[s[l]] = dictS[s[l]] - 1
					l ++
				} else {
					break
				}
			}
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR+1]
}

func checkValid(dictS,dictT map[byte]int) bool {
	for k,v := range dictT {
		if dictS[k] < v {
			return false
		}
	}
	return true
}