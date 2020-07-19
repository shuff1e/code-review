package Greddy

import (
	"github.com/shuff1e/code-review/DataStructure"
	"github.com/shuff1e/code-review/DataStructure/BinaryTree"
)

// 给定来自字母表A的n个字符的集合，已知每个字符出现的频率freq(c)
// 为每一个字符找到一个二进制编码，使得 freq(c)*Length(binarycode(c))的值最小
// 其中Length(binarycode(c))表示字符c的二进制编码的长度
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
