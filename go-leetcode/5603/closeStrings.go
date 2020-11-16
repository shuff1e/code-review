package main

/*

5603. 确定两个字符串是否接近
如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。



示例 1：

输入：word1 = "abc", word2 = "bca"
输出：true
解释：2 次操作从 word1 获得 word2 。
执行操作 1："abc" -> "acb"
执行操作 1："acb" -> "bca"
示例 2：

输入：word1 = "a", word2 = "aa"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
示例 3：

输入：word1 = "cabbba", word2 = "abbccc"
输出：true
解释：3 次操作从 word1 获得 word2 。
执行操作 1："cabbba" -> "caabbb"
执行操作 2："caabbb" -> "baaccc"
执行操作 2："baaccc" -> "abbccc"
示例 4：

输入：word1 = "cabbba", word2 = "aabbss"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。


提示：

1 <= word1.length, word2.length <= 105
word1 和 word2 仅包含小写英文字母

 */

import (
	"fmt"
	"sort"
)

func main() {
	s1 := "uau"
	s2 := "ssx"
	fmt.Println(closeStrings(s1,s2))
}

func closeStrings(word1 string, word2 string) bool {
	dict1,slice1 := getSlice(word1)
	dict2,slice2 := getSlice(word2)
	if len(dict1) != len(dict2) {
		return false
	}

	for k,_ := range dict1 {
		if _,ok := dict2[k];!ok {
			return false
		}
	}

	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0;i<len(slice1);i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func getSlice( word1 string) (map[byte]int,[]int) {
	dict1 := map[byte]int{}
	for i := 0;i<len(word1);i++ {
		dict1[word1[i]] ++
	}

	slice1 := []int{}
	for _,v := range dict1 {
		slice1 = append(slice1,v)
	}
	sort.Ints(slice1)
	return dict1,slice1
}

/*

把题目要求翻译成人话就是，
如果两个字符串：

包含的字符种类完全一样；
把各个字符的重复次数放在一个数组里，数组在排序后完全一样；
那么这两个字符串接近。

所以：

如果两个字符串长度不一样，那么直接返回false；
遍历两个字符串，用两个长度 2626 的数组存放次数；
同时遍历这两个数组，如果在某下标i处出现一个是0一个不是0（即异或结果是1）的情况，那么直接返回false；
排序后如果数组不相同，也返回false；
否则返回true。

class Solution {
public:
    bool closeStrings(string word1, string word2)
    {
        int m = word1.size();
        int n = word2.size();
        if (m != n)
            return false;
        vector<int> repeat1(26), repeat2(26);
        for (int i = 0; i < m; ++i)
        {
            ++repeat1[word1[i] - 'a'];
            ++repeat2[word2[i] - 'a'];
        }
        for (int i = 0; i < 26; ++i)
            if ((repeat1[i] == 0) ^ (repeat2[i] == 0))
                return false;
        sort(repeat1.begin(), repeat1.end());
        sort(repeat2.begin(), repeat2.end());
        for (int i = 0; i < 26; ++i)
            if (repeat1[i] != repeat2[i])
                return false;
        return true;
    }
};

 */