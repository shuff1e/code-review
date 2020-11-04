package main

import "fmt"

/*

433. 最小基因变化
一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。

假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。

例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。

与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，
请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

注意:

起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
所有的目标基因序列必须是合法的。
假定起始基因序列与目标基因序列是不一样的。
示例 1:

start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]

返回值: 1
示例 2:

start: "AACCGGTT"
end:   "AAACGGTA"
bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

返回值: 2
示例 3:

start: "AAAAACCC"
end:   "AACCCCCC"
bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

返回值: 3

 */

/*

115   不同的子序列
126   单词接龙 II
127. 单词接龙

207  课程表
210. 课程表 II

 */

func main() {
	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	fmt.Println(minMutation(start,end,bank))
}

type pair struct {
	word string
	level int
}

func minMutation(start string, end string, bank []string) int {
	allDict := map[string][]string{}
	for _,v := range bank {
		for i := 0;i<len(v);i++ {
			key := v[0:i] + "*" + v[i+1:]
			if _,ok := allDict[key];!ok {
				allDict[key] = []string{v}
			} else {
				allDict[key] = append(allDict[key],v)
			}
		}
	}

	visited := map[string]struct{}{}
	queue := []pair{{start,0}}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		poll := queue[0]
		queue = queue[1:]
		if poll.word == end {
			return poll.level
		}

		word := poll.word
		level := poll.level
		for i := 0;i<len(poll.word);i++ {
			key := word[0:i] + "*" + word[i+1:]
			for _,tempWord := range allDict[key] {
				if _,ok := visited[tempWord];!ok {
					queue = append(queue,pair{tempWord,level+1})
					visited[tempWord] = struct{}{}
				}
			}
		}
	}
	return -1
}