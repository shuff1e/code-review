package main

/*
107. 二叉树的层次遍历 II
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

     3
    / \
   9  20
   /  \
  15   7
返回其自底向上的层次遍历为：

[
[15,7],
[9,20],
[3]
]
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
	}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := [][]int{}
	queue1,queue2 := []*TreeNode{},[]*TreeNode{}
	queue1 = append(queue1,root)
	for len(queue1) > 0 {
		arr := []int{}
		for len(queue1) > 0 {
			temp := queue1[0]
			queue1 = queue1[1:]

			arr = append(arr,temp.Val)

			if temp.Left != nil {
				queue2 = append(queue2,temp.Left)
			}
			if temp.Right != nil {
				queue2 = append(queue2,temp.Right)
			}
		}
		result = append([][]int{arr},result...)
		temp := queue1
		queue1 = queue2
		queue2 = temp
	}
	return result
}

