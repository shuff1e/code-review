package main

/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{}
	lastVisited := (*TreeNode)(nil)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		temp := stack[len(stack)-1]
		if temp.Right == nil || temp.Right == lastVisited {
			// 出栈的
			lastVisited = temp
			result = append(result,temp.Val)
			stack = stack[0:len(stack)-1]
		} else {
			root = temp.Right
		}
	}
	return result
}