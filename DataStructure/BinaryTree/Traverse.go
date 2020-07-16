package BinaryTree

import "fmt"

func PreOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func MidOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	MidOrder(root.Left)
	fmt.Println(root.Data)
	MidOrder(root.Right)
}