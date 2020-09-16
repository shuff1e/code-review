package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

// 32（二）：分行从上到下打印二叉树
// 题目：从上到下按层打印二叉树，同一层的结点按从左到右的顺序打印，每一层
// 打印到一行。

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func printTreeInLine(root *TreeNode) {
	if root == nil {
		return
	}
	queue1 := arraylist.New()
	queue2 := arraylist.New()
	fmt.Println(root.value," ")
	queue1.Add(root)
	for !queue1.Empty() {
		for !queue1.Empty() {
			temp,_ := queue1.Get(0)
			node := temp.(*TreeNode)
			queue1.Remove(0)

			if node.left != nil {
				fmt.Print(node.left.value," ")
				queue2.Add(node.left)
			}
			if node.right != nil {
				fmt.Print(node.right.value," ")
				queue2.Add(node.right)
			}
		}
		fmt.Println()
		temp := queue1
		queue1 = queue2
		queue2 = temp
	}
}

func main() {
	Test1()
	Test2()
	Test3()
}

// 测试二叉树：出叶子结点之外，左右的结点都有且只有一个右子结点
//            8
//             7
//              6
//               5
//                4
func Test1() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode7 := CreateBinaryTreeNode(7);
	pNode6 := CreateBinaryTreeNode(6);
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);

	ConnectTreeNodes(pNode8, nil, pNode7);
	ConnectTreeNodes(pNode7, nil, pNode6);
	ConnectTreeNodes(pNode6, nil, pNode5);
	ConnectTreeNodes(pNode5, nil, pNode4);

	printTreeInLine(pNode8)
}

// 测试完全二叉树：除了叶子节点，其他节点都有两个子节点
//            8
//        6      10
//       5 7    9  11
func Test2() {
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

	printTreeInLine(pNode8)
}

//        100
//        /
//       50-1
//         \
//         150
func Test3() {
	pNode100 := CreateBinaryTreeNode(100);
	pNode50 := CreateBinaryTreeNode(50);
	pNode150 := CreateBinaryTreeNode(150);

	ConnectTreeNodes(pNode100, pNode50, nil);
	ConnectTreeNodes(pNode50, nil, pNode150);
	printTreeInLine(pNode100)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}