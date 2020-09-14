package main

import "fmt"

// 28：对称的二叉树
// 题目：请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和
// 它的镜像一样，那么它是对称的。

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	value int
}

// 而且 左子树和右子树 互为镜像
func mirror(left,right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if !((left != nil) && (right != nil)) {
		return false
	}
	if left.value != right.value {
		return false
	}
	return mirror(left.left,right.right) && mirror(left.right,right.left)
}

func check(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.left,root.right)
}
func main() {
	Test1()
	Test2()
}

//            8
//        6      6
//       5 7    7 5
func Test1() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode61 := CreateBinaryTreeNode(6);
	pNode62 := CreateBinaryTreeNode(6);
	pNode51 := CreateBinaryTreeNode(5);
	pNode71 := CreateBinaryTreeNode(7);
	pNode72 := CreateBinaryTreeNode(7);
	pNode52 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode8, pNode61, pNode62);
	ConnectTreeNodes(pNode61, pNode51, pNode71);
	ConnectTreeNodes(pNode62, pNode72, pNode52);

	fmt.Println(check(pNode8))
}

//            8
//        6      9
//       5 7    7 5
func Test2() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode61 := CreateBinaryTreeNode(6);
	pNode9 := CreateBinaryTreeNode(9);
	pNode51 := CreateBinaryTreeNode(5);
	pNode71 := CreateBinaryTreeNode(7);
	pNode72 := CreateBinaryTreeNode(7);
	pNode52 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode8, pNode61, pNode9);
	ConnectTreeNodes(pNode61, pNode51, pNode71);
	ConnectTreeNodes(pNode9, pNode72, pNode52);

	fmt.Println(check(pNode8))
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}