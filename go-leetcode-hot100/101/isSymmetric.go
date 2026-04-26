package main

/*
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

   1
  / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
    3    3

进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
// 因此，该问题可以转化为：两个树在什么情况下互为镜像？
//
// 如果同时满足下面的条件，两个树互为镜像：
// 它们的两个根结点具有相同的值
// 每个树的右子树都与另一个树的左子树镜像对称

// 将根节点的问题 转换为左子树和右子树的问题
// left.left和right.right是否镜像对称

// 我们可以实现这样一个递归函数，通过「同步移动」两个指针的方法来遍历这棵树，
// p 指针和 qq 指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，q 右移。
// 每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称。

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left,root.Right)
}

func mirror(left,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil) != (right == nil) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return mirror(left.Left,right.Right) && mirror(left.Right,right.Left)
}

// 首先我们引入一个队列，
// 这是把递归程序改写成迭代程序的常用方法。初始化时我们把根节点入队两次。
// 每次提取两个结点并比较它们的值（队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像），
// 然后将两个结点的左右子结点按相反的顺序插入队列中。
// 当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。
//

func isSymmetric2(root *TreeNode) bool {
	u,v := root,root
	queue := []*TreeNode{}
	queue = append(queue,u,v)
	for len(queue) > 0 {
		u,v = queue[0],queue[1]
		queue = queue[2:]
		if u == nil && v == nil {
			continue
		}
		if (u == nil) != (v== nil) {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		queue = append(queue,u.Left,v.Right)
		queue = append(queue,u.Right,v.Left)
	}
	return true
}
