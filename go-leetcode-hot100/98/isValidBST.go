package main

/*
98. 验证二叉搜索树
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
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

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 左子树向上报告自己的最大值和最小值
// 并返回左子树是否是搜索树
func isValidBST(root *TreeNode) bool {
	ok,_,_,_ := help(root)
	return ok
}

// 二叉搜索树，如果对其进行中序遍历，得到的值序列是递增有序的

func isValidBST2(root *TreeNode) bool {
	stack := []*TreeNode{}
	pred := (*TreeNode)(nil)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if pred != nil && pred.Val >= root.Val {
			return false
		}
		pred = root
		root = root.Right
	}
	return true
}


func help(root *TreeNode) (ok bool,isNil bool,min,max int) {
	if root == nil {
		return true,true,0,0
	}

	lok,ok1,lmin,lmax := help(root.Left)
	rok,ok2,rmin,rmax := help(root.Right)

	if !lok || !rok {
		return false,false,0,0
	}

	// 左子树不是nil，比较lmax
	if (!ok1 && lmax >= root.Val) || (!ok2 && rmin <= root.Val) {
		return false,false,0,0
	}
	// 左子树是nil
	if ok1 {
		lmin = root.Val
	}
	if ok2 {
		rmax = root.Val
	}
	return true,false,lmin,rmax
}
