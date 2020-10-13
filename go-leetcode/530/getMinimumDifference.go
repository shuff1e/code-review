package main

import "fmt"

/*
530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例：

输入：

1
 \
  3
 /
2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

提示：

树中至少有 2 个节点。
 */

func main() {
//            8
//        6      10
//       5 7    9  11
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 2}
	connectNodes(node1,nil,node2)
	connectNodes(node2,node3,nil)
	fmt.Println(getMinimumDifference(node1))
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

// 中序遍历
func getMinimumDifference(root *TreeNode) int {
	diff := 0x7fffffff
	prev := -1
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == -1 {
			prev = root.Val
		} else {
			diff = Min(diff,root.Val-prev)
			prev = root.Val
		}
		root = root.Right
	}
	return diff
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}