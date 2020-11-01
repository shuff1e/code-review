package main

import (
	"fmt"
	"sort"
	"strings"
)

/*

761. 特殊的二进制序列
特殊的二进制序列是具有以下两个性质的二进制序列：

0 的数量与 1 的数量相等。
二进制序列的每一个前缀码中 1 的数量要大于等于 0 的数量。
给定一个特殊的二进制序列 S，以字符串形式表示。定义一个操作 为首先选择 S 的两个连续且非空的特殊的子串，然后将它们交换。（两个子串为连续的当且仅当第一个子串的最后一个字符恰好为第二个子串的第一个字符的前一个字符。)

在任意次数的操作之后，交换后的字符串按照字典序排列的最大的结果是什么？

示例 1:

输入: S = "11011000"
输出: "11100100"
解释:
将子串 "10" （在S[1]出现） 和 "1100" （在S[3]出现）进行交换。
这是在进行若干次操作后按字典序排列最大的结果。
说明:

S 的长度不超过 50。
S 保证为一个满足上述定义的特殊 的二进制序列。

 */

func main() {
	str := "1010101100"
	fmt.Println(makeLargestSpecial(str))
	//slice := []string{"01","10","11"}
	//sort.Slice(slice, func(i, j int) bool {
	//	return strings.Compare(slice[i],slice[j]) > 0
	//})
	//fmt.Println(slice)
}


// 0123456789
// 1010101100
// pair是1和0的配对关系
// pair是[1,?,3,?,5,?,9,8,?,?]

func makeLargestSpecial(S string) string {
	pair := make([]int,len(S))
	stack := []int{}
	for i := 0;i<len(S);i++ {
		if S[i] == '1' {
			stack = append(stack,i)
		} else {
			temp := stack[len(stack)-1]
			pair[temp] = i
			stack = stack[:len(stack)-1]
		}
	}
	result := arrange(S,0,len(S)-1,pair)
	return result
}

func arrange(str string ,left ,right int,pair []int) string {
	if left > right {
		return ""
	}
	slice := []string{}
	for i := left;i<right;i = pair[i] + 1 {
		// 比如( () (()) )
		// 去掉两边的括号，里面的递归，返回最大的string
		result := arrange(str,i+1,pair[i]-1,pair)
		// 然后需要加上左边的1和右边的0
		slice = append(slice,"1" + result + "0")
		// 如果是 () () () (())
		// 同样
	}

	// 比如当前为 () () () (())
	// 在这层递归中， (())应该放在前面
	sort.Slice(slice, func(i, j int) bool {
		return strings.Compare(slice[i],slice[j]) > 0
	})

	return strings.Join(slice,"")
}