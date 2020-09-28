package main

import "fmt"

/*
30. 串联所有单词的子串
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。



示例 1：

输入：
s = "barfoothefoobarman",
words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
s = "wordgoodgoodgoodbestword",
words = ["word","good","best","word"]
输出：[]

 */

// A：类似字符串匹配算法
// 但是这里可以是无序的，可以使用map

// 另外使用滑动窗口，temp_map
// 窗口右端right右移，当right的元素是valid的时候，left不动，窗口是有效的

// 否则，窗口左端left左移，并从temp_map中减去left位置的元素，窗口是无效的，
// 再检查right的元素是否valid，如果是valid的话，left不动，窗口是有效的

// 计算是否达到了要求

// 每次有新的invalid的right进来时，要移动left，来消除这个right带来的影响

// the the foo bar yoo woo yoo  7*3=21
// [foo bar yoo woo]  4*3 = 12
// 21-12=9 (b的位置）

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo","bar"}
	fmt.Printf("%#v\n",findSubstring(s,words))
}

func findSubstring(s string, words []string) []int {
	result := []int{}
	words_length := len(words)
	one_word_length := len(words[0])

	if len(s) == 0 || words_length == 0 || one_word_length == 0 {
		return result
	}

	dict := map[string]int{}
	for _,word := range words {
		dict[word] = dict[word] + 1
	}

	// 这个地方i设置为i<= len(s) - words_length * one_word_length 会有重复
	// 假设one_word_length为3的话，
	// i=3和i=0的情况重复了
	// 如果 i == 0
	// 那么i=1和i=2的情况就没办法覆盖到
	for i := 0;i< one_word_length ;i++ {
		temp_dict := map[string]int{}
		count := 0
		left,right := i,i
		for ;right <= len(s) - one_word_length;right += one_word_length {
			count ++
			r_word := s[right:right+one_word_length]
			temp_dict[r_word] = temp_dict[r_word] + 1
			for temp_dict[r_word] > dict[r_word] {
				l_word := s[left:left+one_word_length]
				temp_dict[l_word] = temp_dict[l_word] - 1
				left += one_word_length
				count --
			}
			if count == words_length {
				result = append(result,left)
			}
		}
	}
	return result
}