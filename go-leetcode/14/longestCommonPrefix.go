package main

/*
14. 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
 */


// A：用不到字典树
//

func main() {
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1;i<len(strs);i++ {
		prefix = getPrefix(prefix,strs[i])
		if prefix == "" {
			return prefix
		}
	}
	return prefix
}

func getPrefix(str1,str2 string) string {
	index := 0
	for index < len(str1) && index < len(str2) && str1[index] == str2[index] {
		index ++
	}
	return str1[0:index]
}