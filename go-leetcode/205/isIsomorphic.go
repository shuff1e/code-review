package main

import "fmt"

/*

205. 同构字符串
给定两个字符串 s 和 t，判断它们是否是同构的。

如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。

所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:

输入: s = "egg", t = "add"
输出: true
示例 2:

输入: s = "foo", t = "bar"
输出: false
示例 3:

输入: s = "paper", t = "title"
输出: true
说明:
你可以假设 s 和 t 具有相同的长度。

 */
func main() {
	fmt.Println(isIsomorphic("paper","title"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict1 := map[byte]byte{}
	dict2 := map[byte]byte{}

	for i := 0;i<len(s);i++ {
		if temp,ok := dict1[s[i]];ok {
			if temp != t[i] {
				return false
			}
			continue
		}

		if temp,ok := dict2[t[i]];ok {
			if temp != s[i] {
				return false
			}
			continue
		}

		dict1[s[i]] = t[i]
		dict2[t[i]] = s[i]
	}
	return true
}
