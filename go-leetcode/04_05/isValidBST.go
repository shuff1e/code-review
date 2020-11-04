package main

import "fmt"

/*

面试题 04.05. 合法二叉搜索树
实现一个函数，检查一棵二叉树是否为二叉搜索树。

示例 1:
输入:
    2
   / \
  1   3
输出: true
示例 2:
输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

 */

func main() {
	//            1
	//        2      5
	//       3 4   nil 6
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 3}
	connectNodes(node1,node2,node3)
	result := isValidBST(node1)
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

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return help(root,-0x800000000000,0x7fffffffffff)
}

func help(root *TreeNode,low,high int) bool {
	if root == nil {
		return true
	}
	tempLow := low
	if root.Left != nil {
		tempLow = Max(root.Left.Val,low)
	}
	tempHigh := high
	if root.Right != nil {
		tempHigh = Min(high,root.Right.Val)
	}
	if root.Val <= tempLow || root.Val >= tempHigh {
		return false
	}
	return help(root.Left,low,Min(high,root.Val)) &&
		help(root.Right,Max(low,root.Val),high)
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*

原理：二叉搜索树的中序遍历是递增的

定义一个pre节点保存当前节点的前一个节点，比较大小，不是严格递增即返回false


class Solution {
    TreeNode pre ;
    public boolean isValidBST(TreeNode root) {
        if(root == null) return true ;
        Boolean L = isValidBST(root.left);
        if(pre != null && pre.val >= root.val) return false ;
        pre = root ;
        Boolean R = isValidBST(root.right) ;
        return L && R ;
    }
}

 */