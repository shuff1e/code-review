package main

import "fmt"

/*
208. 实现 Trie (前缀树)
实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母 a-z 构成的。
保证所有输入均为非空字符串。

 */

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("fuck")
	trie.Insert("fuck you")
	trie.Insert("你好")
	trie.Insert("你不好")

	fmt.Println(trie.Search("fuck"))
	trie.Walk()
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
	if len(temp.children) > 0 {
		temp.isEnd = false
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
