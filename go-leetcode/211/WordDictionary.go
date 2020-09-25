package main

/*
211. 添加与搜索单词 - 数据结构设计
如果数据结构中有任何与word匹配的字符串，则bool search（word）返回true，否则返回false。
单词可能包含点“。” 点可以与任何字母匹配的地方。

请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。
word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。


示例：

输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True


提示：

1 <= word.length <= 500
addWord 中的 word 由小写英文字母组成
search 中的 word 由 '.' 或小写英文字母组成
最调用多 50000 次 addWord 和 search


 */

type WordDictionary struct {
	children map[byte]*WordDictionary
	isEnd bool
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		children: make(map[byte]*WordDictionary,0),
	}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	temp := this
	for i := 0;i<len(word);i++ {
		if node,ok := temp.children[word[i]];!ok {
			node = &WordDictionary{
				children: map[byte]*WordDictionary{},
			}
			temp.children[word[i]] = node
			temp = node
		} else {
			temp = node
		}
	}
	temp.isEnd = true
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(word) == 1 {
		if word == "." {
			for _,v := range this.children {
				if v.isEnd {
					return true
				}
			}
			return false
		}
		if node,ok := this.children[word[0]];ok {
			if node.isEnd {
				return true
			}
		}
		if node,ok := this.children['.'];ok {
			if node.isEnd {
				return true
			}
		}
		return false
	}

	if word[0] == '.' {
		for _,node := range this.children {
			if node.Search(word[1:]) {
				return true
			}
		}
		return false
	}
	if node,ok := this.children[word[0]];ok {
		if node.Search(word[1:]) {
			return true
		}
	}
	if node,ok := this.children['.'];ok {
		if node.Search(word[1:]) {
			return true
		}
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

