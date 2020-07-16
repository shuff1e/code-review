package Greddy

import (
	"github.com/shuff1e/code-review/DataStructure"
	"github.com/shuff1e/code-review/DataStructure/BinaryTree"
)

type MyInt int

func (m MyInt) CompareTo(other interface{}) int {
	b := other.(MyInt)
	return int(b)-int(m)
}

func HuffmanCoding(arr []int) *BinaryTree.BinaryTreeNode {
	heap := DataStructure.NewPriorityQueue(len(arr))
	for _,v := range arr {
		temp := &BinaryTree.BinaryTreeNode{
			Data: MyInt(v),
		}
		heap.Offer(temp)
	}

	for heap.Length() > 1 {
		left := heap.Poll().(*BinaryTree.BinaryTreeNode)
		right := heap.Poll().(*BinaryTree.BinaryTreeNode)
		temp := &BinaryTree.BinaryTreeNode{
			Left: left,
			Right: right,
			Data: MyInt(int(left.Data.(MyInt)) + int(right.Data.(MyInt))),
		}
		heap.Offer(temp)
	}
	return heap.Poll().(*BinaryTree.BinaryTreeNode)
}
