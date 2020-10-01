package main

/*
49. 字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
["ate","eat","tea"],
["nat","tan"],
["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。

 */

// A：用a b c 的个数作为key

func groupAnagrams(strs []string) [][]string {
	dict := map[string][]string{}
	for _,s := range strs{
		key := [26]int{}
		for i := 0;i<len(s);i++ {
			key[s[i]-'a'] = key[s[i]-'a'] + 1
		}

		keyStr := ""
		for i := 0;i<len(key);i++ {
			keyStr = keyStr + "#" + string(key[i])
		}

		if _,ok := dict[keyStr];!ok {
			dict[keyStr] = make([]string,0)
		}

		dict[keyStr] = append(dict[keyStr],s)
	}
	result := [][]string{}
	for _,v := range dict {
		result = append(result,v)
	}
	return result
}