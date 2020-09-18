package main

import (
	"algorithm/common/help"
	"fmt"
)

type trieNode struct {
	value byte
	count int
	children []*trieNode
}

func buildTrie(arrary []string) *trieNode {
	root := &trieNode{}
	rootTemp := root
	for _,str := range arrary {
		for i:=0;i<len(str);i++ {
			if findNode,ok := find(rootTemp.children,str[i]);!ok {
				tempNode := &trieNode{
					str[i],
					1,
					make([]*trieNode,0),
				}
				rootTemp.children = append(rootTemp.children,tempNode)
				rootTemp = tempNode
			} else {
				findNode.count ++
				rootTemp = findNode
			}
		}
		rootTemp = root
	}
	return root
}

func getResult(array []string,root *trieNode) (result []string) {
	result = []string{}
	temp := root
	i := 0
	for _,str := range array {
		match := false
		for i = 0;i<len(str);i++ {
			if tt,ok := find(temp.children,str[i]);ok && tt.count > 1{
				temp = tt
			} else {
				match = true
				result = append(result,str[:i+1])
				break
			}
		}
		if !match {
			result = append(result,str)
		}
		temp = root
	}
	return
}
func find(children []*trieNode,b byte) (*trieNode,bool) {
	for _,v := range children {
		if v.value == b {
			return v,true
		}
	}
	return nil,false
}

func main() {
	array := []string{"abc","abcd","abcdefg","bcd","bcdefg"}
	root := buildTrie(array)
	temp := root
	stack := *help.NewMyQueue()
	stack.Add(temp)

	for stack.Length() > 0 {
		tempTemp,_ := stack.Poll()
		tempTrie := tempTemp.(*trieNode)
		temp = tempTrie
		fmt.Println(string(temp.value),temp.count)

		for _,v := range temp.children {
			fmt.Println(v.value)
			stack.Add(v)
		}
	}
	result := getResult(array,root)
	fmt.Printf("%#v\n",result)
}