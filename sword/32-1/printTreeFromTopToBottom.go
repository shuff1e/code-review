package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

// 层次遍历，先入先出，用queue

func travelInLevel(root *TreeNode) {
	if root == nil {
		return
	}
	queue := arraylist.New()
	queue.Add(root)
	for !queue.Empty() {
		temp,_ := queue.Get(0)
		node := temp.(*TreeNode)
		fmt.Print(node.value," ")
		queue.Remove(0)
		if node.left != nil {
			queue.Add(node.left)
		}
		if node.right != nil {
			queue.Add(node.right)
		}
	}
	fmt.Println()
}

func main() {
	Test1()
}

// 测试完全二叉树：除了叶子节点，其他节点都有两个子节点
//            8
//        6      10
//       5 7    9  11
func Test1() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode6 := CreateBinaryTreeNode(6);
	pNode10 := CreateBinaryTreeNode(10);
	pNode5 := CreateBinaryTreeNode(5);
	pNode7 := CreateBinaryTreeNode(7);
	pNode9 := CreateBinaryTreeNode(9);
	pNode11 := CreateBinaryTreeNode(11);
	ConnectTreeNodes(pNode8, pNode6, pNode10);
	ConnectTreeNodes(pNode6, pNode5, pNode7);
	ConnectTreeNodes(pNode10, pNode9, pNode11);

	travelInLevel(pNode8)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}
