package main

import (
	"algorithm/common/help"
	"fmt"
)

func maxLength(arr []int,k int) int {
	if len(arr) == 0{
		return 0
	}
	m := map[int]int{}
	result := 0
	sum := 0
	// 初始化，sum为0的第一次出现的位置为-1
	m[0] = -1
	for i := 0;i<len(arr);i++ {
		sum += arr[i]
		if v,ok := m[sum-k];ok {
			result = help.Max(result,i-v)
		}
		if _,ok := m[sum];!ok {
			m[sum] = i
		}
	}
	return result
}

func main() {
	arr := []int{1,2,3,4,5,3,7,2,3,8}
	fmt.Println(maxLength(arr,8))
	var preArr = []int{1,2,4,7,3,5,6,8}
	var midArr = []int{4,7,2,1,5,3,8,6}
	node := build(preArr,midArr)
	fmt.Println(maxTreeLength(node,18))
}

func build(pre []int,mid []int) *help.TreeNode{
	if len(pre) == 0 {
		return nil
	}
	value := pre[0]
	index := findIndex(value,mid) //0
	leftLen := index //0
	rightLen := len(mid)-1-index//0

	left := build(pre[1:1+leftLen],mid[0:leftLen])
	right := build(pre[len(pre)-rightLen:len(pre)],mid[index+1:])
	root := &help.TreeNode{Value: value,Left: left,Right: right}
	return root
}

func findIndex(value int,mid []int) int {
	for index,v := range mid {
		if value == v {
			return index
		}
	}
	return -1
}

func maxTreeLength(node *help.TreeNode,k int) int {
	m := map[int]int{}
	m[0] = -1
	sum := 0
	maxLength := 0
	maxLength = digui(node,sum,maxLength,k,0,m)
	return maxLength
}

func digui(node *help.TreeNode, sum,maxLength ,k ,level int,m map[int]int) int {
	if node == nil {
		return maxLength
	}
	sum += node.Value
	if v,ok := m[sum - k];ok {
		maxLength = help.Max(maxLength,level-v)
	}
	if _,ok:=m[sum];!ok {
		m[sum] = level
	}

	maxLength = digui(node.Left,sum,maxLength,k,level+1,m)
	maxLength = digui(node.Right,sum,maxLength,k,level+1,m)
	if m[sum] == level {
		delete(m,sum)
	}

	return maxLength
}