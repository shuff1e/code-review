package main

import "fmt"

/*
763. 划分字母区间

给定一个字符串 s，将它划分为尽可能多的片段，使得同一个字母最多只出现在一个片段中。
返回每个片段的长度。

划分后的片段按原顺序连接后，应当仍然等于原字符串。

示例 1：
输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：划分结果为 "ababcbaca"、"defegde"、"hijhklij"。

示例 2：
输入：s = "eccbbbbdec"
输出：[10]

提示：
1 <= s.length <= 500
s 仅包含小写英文字母。
*/

func main() {
	str := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(str))

	str = "eccbbbbdec"
	fmt.Println(partitionLabels(str))

	str = "abab"
	fmt.Println(partitionLabels(str))
}

func partitionLabels(s string) []int {
	ret := []int{}

	dict := map[byte]int{}
	for i := 0; i < len(s); i++ {
		dict[s[i]] = i
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if end < dict[s[i]] {
			end = dict[s[i]]
		}
		if end == i {
			ret = append(ret, end-start+1)
			start = end + 1
		}
	}

	return ret
}
