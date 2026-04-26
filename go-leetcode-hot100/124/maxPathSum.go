package main

import "fmt"

/*
124. 二叉树中的最大路径和
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1：

输入：[1,2,3]

  1
 / \
2   3

输出：6
示例 2：

输入：[-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出：42
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	//           -10
	//        6      -10
	//       5 7    9  -11
	node1 := &TreeNode{Val: -10}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: -10}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 9}
	node7 := &TreeNode{Val: -11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	//node1 = (*TreeNode)(nil)
	fmt.Printf("%#v\n",maxPathSum(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func maxPathSum(root *TreeNode) int {
	global := -0x80000000
	help(root,&global)
	return global
}

func help(root *TreeNode,global *int) int {
	if root == nil {
		return 0
	}

	lv := help(root.Left,global)
	rv := help(root.Right,global)

	if lv > 0 && rv > 0 {
		*global = Max(*global,lv + rv + root.Val)
	} else if lv > 0 {
		*global = Max(*global,lv + root.Val)
	} else if rv > 0 {
		*global = Max(*global,rv + root.Val)
	} else {
		*global = Max(*global,root.Val)
	}

	max := Max(lv,rv)
	result := 0
	if max > 0 {
		result = max + root.Val
	} else {
		result = root.Val
	}
	return result
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}