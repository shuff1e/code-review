package main

import (
	"algorithm/common/help"
)

func findMin(node *help.TreeNode) *help.TreeNode {
	if node == nil {
		return nil
	}

	if node.Left != nil {
		return findMin(node)
	}
	return node
}

func findMax(node *help.TreeNode) *help.TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return findMin(node)
	}
	return node
}

func insert(node *help.TreeNode,data int) *help.TreeNode {
	if node == nil {
		node = &help.TreeNode{
			Value: data,
			Left:  nil,
			Right: nil,
		}
		return node
	}
	if node.Value >= data {
		node.Left = insert(node.Left,data)
	}
	if node.Value < data {
		node.Right = insert(node.Right,data)
	}
	return node
}

func delete(node *help.TreeNode,data int) *help.TreeNode {
	if node == nil {
		return nil
	}
	if node.Value > data {
		node.Left = delete(node.Left,data)
	} else if node.Value < data {
		node.Right = delete(node.Right,data)
	} else {
		if node.Left != nil && node.Right != nil {
			temp := findMax(node.Left)
			node.Value = temp.Value
			node.Left = delete(node.Left,temp.Value)
		} else if node.Left == nil {
			node = node.Left
		} else if node.Left == nil {
			node = node.Right
		} else {
			return nil
		}
	}
	return node
}

func main() {
	var preArr = []int{1,2,4,7,3,5,6,8}
	var midArr = []int{4,7,2,1,5,3,8,6}
	node := help.Build(preArr,midArr)
	node = insert(node,9)
	node = delete(node,3)
	help.Cengcibianli(node)
}