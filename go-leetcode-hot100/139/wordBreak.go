package main

import (
	"fmt"
	"strings"
)

/*
139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
 */

func main() {
	str := "applepenapple"
	arr := []string{"apple", "pen"}
	//str = "catsandog"
	//arr = []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(str,arr))
}

func wordBreak(s string, wordDict []string) bool {
	memo := make([]int,len(s)+1)
	return help(s,0,wordDict,memo)
}

func help(s string,start int,wordDict []string,memo []int) bool {
	if len(s) == start {
		return true
	}
	if memo[start] != 0 {
		if memo[start] == 1 {
			return true
		} else {
			return false
		}
	}
	for _,v := range wordDict {
		if strings.HasPrefix(s[start:],v) {
			ok := help(s,start+len(v),wordDict,memo)
			if ok {
				memo[start+len(v)] = 1
				return true
			} else {
				memo[start+len(v)] = -1
			}
		}
	}
	memo[start] = -1
	return false
}

/*
改为dp

public class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        Set<String> wordDictSet = new HashSet(wordDict);
        boolean[] dp = new boolean[s.length() + 1];
        dp[0] = true;
        for (int i = 1; i <= s.length(); i++) {
            for (int j = 0; j < i; j++) {
                if (dp[j] && wordDictSet.contains(s.substring(j, i))) {
                    dp[i] = true;
                    break;
                }
            }
        }
        return dp[s.length()];
    }
}

 */
