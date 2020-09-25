package main

import (
	"fmt"
	"unsafe"
)

/*
212. 单词搜索 II
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。
同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ['oath','pea','eat','rain'] and board =
[
['o','a','a','n'],
['e','t','a','e'],
['i','h','k','r'],
['i','f','l','v']
]

输出: ['eat','oath']
说明:
你可以假设所有输入都由小写字母 a-z 组成。

提示:

你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？
散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

 */


// A：主体是回溯，从bord中的每个节点开始找
// 关键是如何判断当前回溯到的单词，没有以该单词开始的word，而不用一直回溯到最后
// 相比于set，trie是判断前缀的更好的结构

// 如果发现了匹配的，就将该word从trie中删除

func main() {
	board := [][]byte{
		{'o','a','a','n'},
		{'e','t','a','e'},
		{'i','h','k','r'},
		{'i','f','l','v'},
	}
	words := []string{"oath","pea","eat","rain"}
	board = [][]byte{
		{'a','a'},
	}
	words = []string{"a"}

	fmt.Println(findWords(board,words))
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _,v := range words {
		trie.Insert(v)
	}

	visited := make([][]bool,len(board))
	for i := 0;i<len(visited);i++ {
		visited[i] = make([]bool,len(board[0]))
	}
	list := []byte{}
	result := []string{}

	for i := 0;i<len(board);i++ {
		for j := 0;j<len(board[0]);j++ {
			doIt(board,visited,i,j,&trie,&list,&result)
		}
	}

	return result
}

func doIt(board [][]byte,visited [][]bool,row,col int,trie *Trie,list *[]byte,result *[]string) {

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return
	}

	if visited[row][col] {
		return
	}


	visited[row][col] = true
	*list = append(*list,board[row][col])

	defer func() {
		*list = (*list)[0:len(*list)-1]
		visited[row][col] = false
	}()

	node := trie.searchPrefix(String(*list))

	if node == nil {
		return
	}

	if node.isEnd {
		trie.Delete(String(*list))
		temp2 := make([]byte,len(*list))
		copy(temp2,*list)
		*result = append(*result,String(temp2))
	}

	doIt(board,visited,row+1,col,trie,list,result)
	doIt(board,visited,row-1,col,trie,list,result)
	doIt(board,visited,row,col+1,trie,list,result)
	doIt(board,visited,row,col-1,trie,list,result)
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func String(arr []byte) string {
	return *(*string)(unsafe.Pointer(&arr))
}

type Trie struct {
	children map[rune]*Trie
	isEnd bool
	Val rune
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: map[rune]*Trie{},
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	temp := this
	for _,char := range word {
		if node,ok := temp.children[char];!ok {
			node = &Trie{
				children: make(map[rune]*Trie,0),
				Val: char,
			}
			temp.children[char] = node
			temp = node
		} else {
			temp = node
		}
	}
	temp.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchPrefix(prefix)
	return node != nil

}

func (this *Trie) searchPrefix(word string) *Trie {
	temp := this
	for _,char := range word {
		if node,ok := temp.children[char];ok {
			temp = node
		} else {
			return nil
		}
	}
	return temp
}

func (this *Trie) Delete(word string) {
	// 从 this开始的所有的父节点
	parents := []*Trie{}

	temp := this
	for _,char := range word {
		parents = append(parents,temp)
		if node,ok := temp.children[char];ok {
			temp = node
		} else {
			// 如果根本就search不到该前缀
			return
		}
	}

	// 发现是前缀，不是单词
	if !temp.isEnd {
		return
	}

	// 如果是前缀
	temp.isEnd = false
	if len(temp.children) > 0 {
		return
	}

	// 只是单词
	for len(parents) > 0 {
		parent := parents[len(parents)-1]
		parents = parents[:len(parents)-1]
		delete(parent.children,temp.Val)
		if len(parent.children) > 0 || parent.isEnd {
			return
		}
		temp = parent
	}

}

func (this *Trie) Walk() {
	var walk func(string,*Trie)
	walk = func(str string,this *Trie) {
		if this == nil {
			return
		}
		str += string(this.Val)
		if this.isEnd {
			fmt.Println(str)
		}
		for _,v := range this.children {
			walk(str,v)
		}
	}
	walk("",this)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
