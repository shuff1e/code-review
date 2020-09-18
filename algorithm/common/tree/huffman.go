package main

import (
	"algorithm/common/help"
	"fmt"
)

func getWPL(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	if len(array) == 2 {
		return array[0] + array[1]
	}
	help.MergeSort(array)
	result := 0
	for len(array) > 2 {
		temp := array[0] + array[1]
		result += temp
		array = array[2:]
		array = append(array,temp)
		help.MergeSort(array)
	}
	return result + array[0] + array[1]
}

func main() {
	array := []int{1,2,2,5,9}
	fmt.Println(getWPL(array))
}

func buildHuffman(array []int) *help.TreeNode {
	node := &help.TreeNode{Value: array[0]}
	for i:=1;i<len(array);i++ {
		temp := &help.TreeNode{Value: array[i]}
		node.Left = temp
		node.Right = node
		node.Value = node.Left.Value + node.Right.Value
	}
	return node
}
