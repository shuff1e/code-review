package main

import (
	"fmt"
)

/*
103. 二叉树的锯齿形层次遍历
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

3
/ \
9  20
   /  \
  15   7
返回锯齿形层次遍历如下：

[
[3],
[20,9],
[15,7]
]

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
	//node1 = (*TreeNode)(nil)
	fmt.Printf("%#v\n",zigzagLevelOrder(node1))
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

// A：最先访问到的元素的子孙节点
// 下一次最后访问到
// 因此用栈

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int,0)
	stack1 := make([]interface{},0)
	stack2 := make([]interface{},0)
	leftToRight := true
	//push
	stack1 = append(stack1,root)
	for len(stack1) > 0 {
		tempList := []int{}
		for len(stack1) > 0 {
			// pop
			temp := stack1[len(stack1)-1]
			stack1 = stack1[0:len(stack1)-1]
			tempList = append(tempList,temp.(*TreeNode).Val)
			if leftToRight {
				if temp.(*TreeNode).Left != nil {
					stack2 = append(stack2,temp.(*TreeNode).Left)
				}
				if temp.(*TreeNode).Right != nil {
					stack2 = append(stack2,temp.(*TreeNode).Right)
				}
			} else {
				if temp.(*TreeNode).Right != nil {
					stack2 = append(stack2,temp.(*TreeNode).Right)
				}
				if temp.(*TreeNode).Left != nil {
					stack2 = append(stack2,temp.(*TreeNode).Left)
				}
			}
		}
		if len(tempList) > 0 {
			result = append(result,tempList)
		}
		leftToRight = !leftToRight
		temp := stack1
		stack1 = stack2
		stack2 = temp
	}
	return result
}

