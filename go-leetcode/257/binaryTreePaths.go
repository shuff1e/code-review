package main

import (
	"strconv"
	"strings"
)

/*

257. 二叉树的所有路径
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	temp := make([]string,0)
	result := make([]string,0)
	help(root,temp,&result)
	return result
}

func help(root *TreeNode,temp []string,result *[]string) {
	if root.Left == nil && root.Right == nil {
		length := len(temp)
		temp = append(temp,strconv.Itoa(root.Val))
		*result = append(*result,strings.Join(temp,"->"))
		temp = temp[:length]
		return
	}

	if root.Left != nil {
		length := len(temp)
		help(root.Left,append(temp,strconv.Itoa(root.Val)),result)
		temp = temp[:length]
	}
	if root.Right != nil {
		length := len(temp)
		help(root.Right,append(temp,strconv.Itoa(root.Val)),result)
		temp = temp[:length]
	}
}