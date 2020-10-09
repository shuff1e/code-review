package main

import (
	"fmt"
)

/*
129. 求根到叶子节点数字之和
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字。

例如，从根到叶子节点路径 1->2->3 代表数字 123。

计算从根到叶子节点生成的所有数字之和。

说明: 叶子节点是指没有子节点的节点。

示例 1:

输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.
示例 2:

输入: [4,9,0,5,1]
     4
    / \
   9   0
  / \
 5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	//            8
	//        6      1
	//       5 7    9  1
	node1 := &TreeNode{Val: 8}
	//node2 := &TreeNode{Val: 6}
	//node3 := &TreeNode{Val: 1}
	//node4 := &TreeNode{Val: 5}
	//node5 := &TreeNode{Val: 7}
	//node6 := &TreeNode{Val: 9}
	//node7 := &TreeNode{Val: 1}
	//connectNodes(node1,node2,node3)
	//connectNodes(node2,node4,node5)
	//connectNodes(node3,node6,node7)
	fmt.Println(sumNumbers(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	help(root,0*10,&sum)
	return sum
}

func help(root *TreeNode,curr int,sum *int) {
	if root.Left == nil && root.Right == nil {
		*sum += curr + root.Val
		return
	}
	curr += root.Val
	if root.Left != nil {
		help(root.Left,curr*10,sum)
	}
	if root.Right != nil {
		help(root.Right,curr*10,sum)
	}
}