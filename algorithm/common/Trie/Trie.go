package main

import "fmt"

type TrieNode struct {
	links []*TrieNode
	R int
	isEnd bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		links: make([]*TrieNode,26),
		R:26,
		isEnd: false,
	}
}

func (n *TrieNode) containsKey(key byte) bool {
	return n.links[key-97] != nil
}

func (n *TrieNode) get(key byte) *TrieNode {
	return n.links[key-97]
}

func (n *TrieNode) put(key byte,node *TrieNode) {
	n.links[key-97] = node
}

func (n *TrieNode) setEnd() {
	n.isEnd = true
}

func (n *TrieNode) End() bool {
	return n.isEnd
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) insert(word string) {
	temp := t.root
	for i :=0;i<len(word);i++ {
		if !temp.containsKey(word[i]) {
			temp.put(word[i],NewTrieNode())
		}
		temp = temp.get(word[i])
	}
	temp.setEnd()
}

func (t *Trie) searchPrefix(word string) *TrieNode {
	temp := t.root
	for i :=0;i<len(word);i++ {
		if temp.get(word[i]) != nil {
			temp = temp.get(word[i])
		} else {
			return nil
		}
	}
	return temp
}

func (t *Trie) search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.End()
}

func (t *Trie) satrtsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}

func main() {
	t := &Trie{
		root: NewTrieNode(),
	}
	t.insert("helloworld")
	fmt.Println(t.satrtsWith("helloo"))
}