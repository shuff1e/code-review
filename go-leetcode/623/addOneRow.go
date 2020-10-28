package main

import "fmt"

/*

623. 在二叉树中增加一行
给定一个二叉树，根节点为第1层，深度为 1。在其第 d 层追加一行值为 v 的节点。

添加规则：给定一个深度值 d （正整数），针对深度为 d-1 层的每一非空节点 N，为 N 创建两个值为 v 的左子树和右子树。

将 N 原先的左子树，连接为新节点 v 的左子树；将 N 原先的右子树，连接为新节点 v 的右子树。

如果 d 的值为 1，深度 d - 1 不存在，则创建一个新的根节点 v，原先的整棵树将作为 v 的左子树。

示例 1:

输入:
二叉树如下所示:
       4
     /   \
    2     6
   / \   /
  3   1 5

v = 1

d = 2

输出:
       4
      / \
     1   1
    /     \
   2       6
  / \     /
 3   1   5

示例 2:

输入:
二叉树如下所示:
      4
     /
    2
   / \
  3   1

v = 1

d = 3

输出:
      4
     /
    2
   / \
  1   1
 /     \
3       1
注意:

输入的深度值 d 的范围是：[1，二叉树最大深度 + 1]。
输入的二叉树至少有一个节点。

 */

func main() {
	//            8
	//        6      10
	//       5 7    9  11
	node1 := &TreeNode{Val: 8}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 10}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 9}
	node7 := &TreeNode{Val: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	node := addOneRow(node1,1,5)
	fmt.Println(node.Val)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		temp := &TreeNode{Val: v}
		temp.Left = root
		return temp
	}
	help(root,1,d - 1,v)
	return root
}

func help(root *TreeNode,level int,k int,v int) {
	if root == nil {
		return
	}
	if level == k {
		temp := root.Left
		root.Left = &TreeNode{Val: v}
		root.Left.Left = temp

		temp = root.Right
		root.Right = &TreeNode{Val: v}
		root.Right.Right = temp
		return
	}
	help(root.Left,level+1,k,v)
	help(root.Right,level+1,k,v)
}