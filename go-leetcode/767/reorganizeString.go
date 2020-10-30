package main

import (
	"container/heap"
	"fmt"
	"sort"
	"unsafe"
)

/*

767. 重构字符串
给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

若可行，输出任意可行的结果。若不可行，返回空字符串。

示例 1:

输入: S = "aab"
输出: "aba"
示例 2:

输入: S = "aaab"
输出: ""
注意:

S 只包含小写字母并且长度在[1, 500]区间内。

 */

func main() {
	str := "aab"
	fmt.Println(reorganizeString2(str))
}

func reorganizeString(S string) string {
	counts := make([]int,26)

	for i := 0;i<len(S);i++ {
		counts[S[i]-'a'] += 100
	}

	for i := 0;i<26;i++ {
		counts[i] += i
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	result := []byte(S)
	for i := 0;i<len(result);i ++ {
		result[i] = 0
	}

	index := 0
	for i := 0;i<len(counts);i++ {
		ct := counts[i]/100
		char := 'a' + byte(counts[i]%100)
		if ct > (len(S)+1)/2 {
			return ""
		}
		for j := 0;j<ct;j++ {
			index = index %len(S)
			for result[index] != 0 {
				index = (index + 1)%len(S)
			}
			result[index] = char
			index += 2
		}
	}
	return String(result)
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

// 类似霍夫曼编码
// 每次从大顶堆中pop两个元素
// 然后append

func reorganizeString2(S string) string {
	result := []byte{}

	counts := make([]int,26)
	for i := 0;i<len(S);i++ {
		counts[S[i] - 'a'] ++
	}
	h := &maxHeap{}

	for i := 0;i<len(counts);i++ {
		if counts[i] == 0 {
			continue
		}
		if counts[i] > (len(S)+1)/2 {
			 return ""
		}
		heap.Push(h,&pair{letter: byte(i),count: counts[i]})
	}

	for h.Len() >=2 {
		p1 := heap.Pop(h).(*pair)
		p2 := heap.Pop(h).(*pair)
		result = append(result,'a' + p1.letter)
		result = append(result,'a' + p2.letter)

		p1.count --
		p2.count --
		if p1.count > 0 {
			heap.Push(h,p1)
		}
		if p2.count > 0 {
			heap.Push(h,p2)
		}
	}

	if h.Len() > 0 {
		result = append(result,h.Pop().(*pair).letter + 'a')
	}
	return String(result)
}

type pair struct {
	letter byte
	count int
}

type maxHeap []*pair

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i,j int) bool {
	if h[i].count == h[j].count {
		return h[i].letter > h[j].letter
	}
	return h[i].count > h[j].count
}

func (h maxHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h,x.(*pair))
}

func (h *maxHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return result
}

func (h *maxHeap) Peek() *pair {
	return (*h)[0]
}
