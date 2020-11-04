package main

import "fmt"

/*
127. 单词接龙
给定两个单词（beginWord 和 endWord）和一个字典，
找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回 0。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
返回它的长度 5。
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。
 */

// 给wordList做预处理,key是通用状态，value是所有具有通用状态的单词

// 队列，将word和level的元组放入队列
// 将所有可以转换的入队

// 用visited数组标志是否访问过该单词
//

func main() {
	w1 := "a"
	w2 := "c" // dot
	wlist := []string{"a","b","c"}
	fmt.Println(ladderLength2(w1,w2,wlist))
}

type pair struct {
	word string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	allDict := map[string][]string{}
	for _,v := range wordList {
		for i := 0;i<len(v);i++ {
			key := v[0:i] + "*" + v[i+1:]
			value := allDict[key]
			value = append(value,v)
			allDict[key] = value
		}
	}
	visited := map[string]struct{}{}
	stack := []pair{}
	stack = append(stack,pair{beginWord,0})
	for len(stack) > 0 {
		// 这个地方不能用stack，要用queue
		// stack是深度优先
		// queue是广度优先
		// 这里是求最短转换序列的长度
		// 要广度优先
		curr := stack[0]
		stack = stack[1:]
		word := curr.word
		level := curr.level
		if word == endWord {
			return level + 1
		}

		for i := 0;i<len(word);i++ {
			key := word[0:i] + "*" + word[i+1:]
			for _,tempWord := range allDict[key] {
				if _,ok := visited[tempWord];!ok {
					visited[word] = struct{}{}
					stack = append(stack,pair{tempWord,level+1})
				}
			}
		}
	}
	return 0
}

/*
如果使用两个同时进行的广搜可以有效地减少搜索空间。
一边从 beginWord 开始，另一边从 endWord 开始。我们每次从两边各扩展一个节点，当发现某一时刻两边都访问了某一顶点时就停止搜索。
这就是双向广度优先搜索，它可以可观地减少搜索空间大小，从而降低时间和空间复杂度。

算法与之前描述的标准广搜方法相类似。

唯一的不同是我们从两个节点同时开始搜索，同时搜索的结束条件也有所变化。

我们现在有两个访问数组，分别记录从对应的起点是否已经访问了该节点。

如果我们发现一个节点被两个搜索同时访问，就结束搜索过程。因为我们找到了双向搜索的交点。
过程如同从中间相遇而不是沿着搜索路径一直走。

双向搜索的结束条件是找到一个单词被两边搜索都访问过了。

最短变换序列的长度就是中间节点在两边的层次之和。因此，我们可以在访问数组中记录节点的层次。
 */

func ladderLength2(beginWord string, endWord string, wordList []string) int {
	found := false
	// endWord "cog" 不在字典中，所以无法进行转换。
	for _,v := range wordList {
		if v == endWord {
			found = true
			break
		}
	}
	if !found {
		return 0
	}
	allDict := map[string][]string{}
	for _,v := range wordList {
		for i := 0;i<len(v);i++ {
			key := v[0:i] + "*" + v[i+1:]
			value := allDict[key]
			value = append(value,v)
			allDict[key] = value
		}
	}

	visitedBegin := map[string]int{}
	visitedEnd := map[string]int{}
	visitedBegin[beginWord] = 0
	visitedEnd[endWord] = 0

	queueBegin := []pair{}
	queueEnd := []pair{}
	queueBegin = append(queueBegin,pair{beginWord,0})
	queueEnd = append(queueEnd,pair{endWord,0})

	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		ans := visit(&queueBegin,visitedBegin,visitedEnd,allDict)
		if ans != -1 {
			return ans
		}
		ans = visit(&queueEnd,visitedEnd,visitedBegin,allDict)
		if ans != -1 {
			return ans
		}
	}
	return 0
}

func visit(queue *[]pair,visited,otherVisited map[string]int,allDict map[string][]string) int {
	curr := (*queue)[0]
	*queue = (*queue)[1:]
	word := curr.word
	level := curr.level
	for i := 0;i<len(word);i++ {
		key := word[0:i] + "*" + word[i+1:]
		for _,v := range allDict[key] {
			if oLevel,ok := otherVisited[v];ok {
				//fmt.Println(word,v)
				return level + 1 + oLevel + 1
			}
			if _,ok := visited[v];!ok {
				*queue = append(*queue,pair{v,level+1})
				//fmt.Printf("%#v\n",*queue)
				visited[v] = level + 1
			}
		}
	}
	return -1
}