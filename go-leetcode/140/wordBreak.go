package main

import (
	"fmt"
	"strings"
)

/*
140. 单词拆分 II
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，
使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：

分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
"cats and dog",
"cat sand dog"
]
示例 2：

输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
"pine apple pen apple",
"pineapple pen apple",
"pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
 */

func main() {
	str := "applepenapple"
	arr := []string{"apple", "pen"}
	str = "catsandog"
	arr = []string{"cats", "dog", "sand", "and", "cat"}
	str = "catsanddog"
	arr = []string{"cat", "cats", "and", "sand", "dog"}
	result := wordBreak2(str,arr)
	for _,v := range result {
		fmt.Println(v)
	}
}

func wordBreak(s string, wordDict []string) []string {
	result,temp := []string{},[]string{}
	help(s,0,wordDict,&result,&temp)
	return result
}

func help(s string,start int,wordDict []string,result,temp *[]string) {
	if len(s) == start {
		*result = append(*result,strings.Join(*temp," "))
		return
	}
	for _,v := range wordDict {
		if strings.HasPrefix(s[start:],v) {
			*temp = append(*temp,v)
			help(s,start+len(v),wordDict,result,temp)
			*temp = (*temp)[0:len(*temp)-1]
		}
	}
}

func wordBreak2(s string, wordDict []string) []string {
	memo := map[int][]string{}
	result := helpBetter(s,0,wordDict,memo)
	return result
}
func helpBetter(s string,start int,wordDict []string,memo map[int][]string) []string {
	if start == len(s) {
		return []string{""}
	}
	if memo[start] != nil {
		return memo[start]
	}

	curr := []string{}
	for _,v := range wordDict {
		if strings.HasPrefix(s[start:],v) {
			list := helpBetter(s,start+len(v),wordDict,memo)
			for _,word := range list {
				if word == "" {
					curr = append(curr,v)
				} else {
					curr = append(curr,v + " " + word)
				}
			}
		}
	}
	memo[start] = curr
	return curr
}

/*

public class Solution {
   public List<String> wordBreak(String s, Set<String> wordDict) {
       LinkedList<String>[] dp = new LinkedList[s.length() + 1];
       LinkedList<String> initial = new LinkedList<>();
       initial.add("");
       dp[0] = initial;
       for (int i = 1; i <= s.length(); i++) {
           LinkedList<String> list = new LinkedList<>();
           for (int j = 0; j < i; j++) {
               if (dp[j].size() > 0 && wordDict.contains(s.substring(j, i))) {
                   for (String l : dp[j]) {
                       list.add(l + (l.equals("") ? "" : " ") + s.substring(j, i));
                   }
               }
           }
           dp[i] = list;
       }
       return dp[s.length()];
   }
}

 */
