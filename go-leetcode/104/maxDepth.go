package main

import "fmt"

/*
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node1 := &TreeNode{Val: 8}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 10}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	fmt.Println(maxDepth(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1
	if root.Left != nil {
		help(root.Left,2,&max)
	}
	if root.Right != nil {
		help(root.Right,2,&max)
	}
	return max
}

func help(root *TreeNode,level int,max *int) {
	*max = Max(*max,level)
	if root.Left != nil {
		help(root.Left,level+1,max)
	}
	if root.Right != nil {
		help(root.Right,level+1,max)
	}
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*
class Solution {
    public int maxDepth(TreeNode root) {
        if (root == null) {
            return 0;
        } else {
            int leftHeight = maxDepth(root.left);
            int rightHeight = maxDepth(root.right);
            return Math.max(leftHeight, rightHeight) + 1;
        }
    }
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    maxDepth := rightDepth

    if leftDepth > rightDepth {
        maxDepth = leftDepth
    }

    return maxDepth + 1

}
 */
