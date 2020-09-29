package main

/*
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	result := []int{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			result = append(result,root.Val)
			// push
			stack = append(stack,root)
			root = root.Left
		}
		temp := stack[len(stack)-1]
		if temp.Right != nil {
			root = temp.Right
			stack = stack[0:len(stack)-1]
		} else {
			// pop
			stack = stack[0:len(stack)-1]
		}
	}
	return result
}