package main

import "fmt"

/*
68. 文本左右对齐
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

说明:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。
示例:

输入:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
输出:
[
"This    is    an",
"example  of text",
"justification.  "
]
示例 2:

输入:
words = ["What","must","be","acknowledgment","shall","be"]
maxWidth = 16
输出:
[
"What   must   be",
"acknowledgment  ",
"shall be        "
]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
因为最后一行应为左对齐，而不是左右两端对齐。
第二行同样为左对齐，这是因为这行只包含一个单词。
示例 3:

输入:
words = ["Science","is","what","we","understand","well","enough","to","explain",
"to","a","computer.","Art","is","everything","else","we","do"]
maxWidth = 20
输出:
[
"Science  is  what we",
"understand      well",
"enough to explain to",
"a  computer.  Art is",
"everything  else  we",
"do                  "
]
 */

func main() {
	words := []string{"Science","is","what","we","understand","well","enough","to","explain",
		"to","a","computer.","Art","is","everything","else","we","do"}
	words = []string{"What","must","be","acknowledgment","shall","be"}
	words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	result := fullJustify(words,16)
	for i := 0;i<len(result);i++ {
		fmt.Println(result[i])
	}
}

func fullJustify(words []string, maxWidth int) []string {
	sum := 0
	start := 0
	result := []string{}
	for i := 0;i<len(words);i++ {
		sum += len(words[i])
		if sum > maxWidth {
			result = append(result,getStr(words,start,i-1,maxWidth))
			start = i
			sum = len(words[i]) + 1
		} else {
			sum += 1
		}
	}
	result = append(result,getStr(words,start,len(words)-1,maxWidth))
	return result
}

func getStr(words []string,start,end , maxWidth int) string {
	if end - start == 0{
		result := words[start]
		for i := len(words[start]);i<maxWidth;i++ {
			result += " "
		}
		return result
	}

	result := ""
	length := 0
	for i := start;i<=end;i++ {
		length += len(words[i])
	}
	rest := maxWidth - length

	if end == len(words) - 1 {
		result := ""
		for i := start;i<end;i++ {
			result += words[i] + " "
			rest --
		}
		result += words[end]
		for j := 0;j<rest;j++ {
			result += " "
		}
		return result
	}

	for i := start;i<end;i++ {
		gap := rest/(end-i)
		if rest%(end-i) != 0 {
			gap = gap + 1
		}
		rest -= gap
		result += words[i]
		for j:= 0;j<gap;j++ {
			result = result + " "
		}
	}
	result += words[end]
	return result
}