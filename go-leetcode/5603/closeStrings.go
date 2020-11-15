package main

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