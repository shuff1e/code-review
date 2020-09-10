package main

// 27：二叉树的镜像
// 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。




// TODO
// 迭代，而不是递归
// 树的遍历，前序，中序，后序，层次遍历
// 广度优先，深度优先












type TreeNode struct {
	left *TreeNode
	right *TreeNode
	value int
}

func mirror(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.left
	node.left = node.right
	node.right = temp
	mirror(node.left)
	mirror(node.right)
}

func main() {

}

func Test1() {
	pNodeA1 := CreateBinaryTreeNode(8);
	pNodeA2 := CreateBinaryTreeNode(8);
	pNodeA3 := CreateBinaryTreeNode(7);
	pNodeA4 := CreateBinaryTreeNode(9);
	pNodeA5 := CreateBinaryTreeNode(3);
	pNodeA6 := CreateBinaryTreeNode(4);
	pNodeA7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNodeA1, pNodeA2, pNodeA3);
	ConnectTreeNodes(pNodeA2, pNodeA4, pNodeA5);
	ConnectTreeNodes(pNodeA5, pNodeA6, pNodeA7);
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}
