package main

/*
100. 相同的树
给定两个二叉树，编写一个函数来检验它们是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

示例 1:

输入:       1         1
           / \       / \
          2   3     2   3

          [1,2,3],   [1,2,3]

输出: true
示例 2:

输入:      1          1
          /           \
         2             2

        [1,2],     [1,null,2]

输出: false
示例 3:

输入:       1         1
           / \       / \
          2   1     1   2

         [1,2,1],   [1,1,2]

输出: false
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	queue1,queue2 := []*TreeNode{},[]*TreeNode{}
	queue1 = append(queue1,p)
	queue2 = append(queue2,q)
	for len(queue1) > 0 && len(queue2) > 0 {
		// poll
		temp1,temp2 := queue1[0],queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]

		if temp1.Val != temp2.Val {
			return false
		}
		if (temp1.Left == nil) != (temp2.Left == nil) {
			return false
		}
		if (temp1.Right == nil) != (temp2.Right == nil) {
			return false
		}
		if temp1.Left != nil {
			queue1 = append(queue1,temp1.Left)
			queue2 = append(queue2,temp2.Left)
		}
		if temp1.Right != nil {
			queue1 = append(queue1,temp1.Right)
			queue2 = append(queue2,temp2.Right)
		}
	}
	return len(queue1) == 0 && len(queue2) == 0
}