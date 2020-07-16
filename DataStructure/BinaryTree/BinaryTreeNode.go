package BinaryTree

import "github.com/shuff1e/code-review/DataStructure"

type BinaryTreeNode struct {
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
	Data  DataStructure.Element
}

func (a *BinaryTreeNode) CompareTo(other interface{}) int {
	b := other.(*BinaryTreeNode)
	return a.Data.CompareTo(b.Data)
}