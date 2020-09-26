package main

/*
113. 路径总和 II
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

     5
    / \
   4   8
  /   / \
 11  13  4
/  \    / \
7    2  5   1
返回:

[
[5,4,11,2],
[5,8,4,5]
]

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	temp := []int{}
	help(root,sum,&result,&temp)
	return result
}

func help(root *TreeNode,target int,result *[][]int,temp *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if target == root.Val {
			*temp = append(*temp,root.Val)

			temp2 := make([]int,len(*temp))
			copy(temp2,*temp)
			*result = append(*result,temp2)

			*temp = (*temp)[:len(*temp)-1]
		}
		return
	}

	if root == nil {
	}

	*temp = append(*temp,root.Val)

	help(root.Left,target-root.Val,result,temp)
	help(root.Right,target-root.Val,result,temp)

	*temp = (*temp)[:len(*temp)-1]
}