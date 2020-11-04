package main

import "fmt"

/*

450. 删除二叉搜索树中的节点
给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，
并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。

示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 找到对应节点，如果是叶节点，直接删除
// 如果只有左子节点，该节点的值，替换为左子节点的值，删除左子节点
// 如果只有右子节点，该节点的值，替换为右子节点的值，删除右子节点

// 如果左右子节点都有，该节点的值，替换为右子节点的值，删除右子节点

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

	result := deleteNode(node1,8)
	fmt.Println(result)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right != nil {
			// 替换为右子节点中的最小值
			parent,delete := help(root,root.Right,false)
			root.Val = delete.Val
			if parent == root {
				root.Right = deleteNode(delete,delete.Val)
			} else {
				parent.Left = deleteNode(delete,delete.Val)
			}
			return root
		} else {
			// 替换为左子节点中的最大值
			parent,delete := help(root,root.Left,true)
			root.Val = delete.Val
			if parent == root {
				root.Left = deleteNode(delete,delete.Val)
			} else {
				parent.Right = deleteNode(delete,delete.Val)
			}
			return root
		}
	} else if root.Val < key {
		root.Right = deleteNode(root.Right,key)
		return root
	} else {
		root.Left = deleteNode(root.Left,key)
		return root
	}
	return nil
}

func help(parent,root *TreeNode,max bool) (child *TreeNode,toDeleted *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return parent,root
	}
	if max {
		if root.Right != nil {
			return help(root,root.Right,max)
		} else {
			return parent,root
		}
	}

	if !max {
		if root.Left != nil {
			return help(root,root.Left,max)
		} else {
			return parent,root
		}
	}

	return nil,nil
}
