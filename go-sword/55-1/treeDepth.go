package main

import "fmt"

// 55（一）：二叉树的深度
// 题目：输入一棵二叉树的根结点，求该树的深度。从根结点到叶结点依次经过的
// 结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度。

// A：递归。将root的问题分解为left的问题和right的问题

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getMaxDepth(root.left)
	right := getMaxDepth(root.right)
	return 1 + Max(left,right)
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
}

func Test(name string,node *TreeNode,expected int) {
	fmt.Println(name,getMaxDepth(node) == expected)
}

//            1
//         /      \
//        2        3
//       /\         \
//      4  5         6
//        /
//       7
func Test1() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);
	pNode6 := CreateBinaryTreeNode(6);
	pNode7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNode1, pNode2, pNode3);
	ConnectTreeNodes(pNode2, pNode4, pNode5);
	ConnectTreeNodes(pNode3, nil, pNode6);
	ConnectTreeNodes(pNode5, pNode7, nil);

	Test("Test1", pNode1, 4);
}

//               1
//              /
//             2
//            /
//           3
//          /
//         4
//        /
//       5
func Test2() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, pNode2, nil);
	ConnectTreeNodes(pNode2, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode5, nil);

	Test("Test2", pNode1, 5);

}

func Test3() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, nil, pNode2);
	ConnectTreeNodes(pNode2, nil, pNode3);
	ConnectTreeNodes(pNode3, nil, pNode4);
	ConnectTreeNodes(pNode4, nil, pNode5);

	Test("Test3", pNode1, 5);
}

func Test4() {
	pNode1 := CreateBinaryTreeNode(1);
	Test("Test4", pNode1, 1);
}

func Test5() {
	Test("Test5", nil, 0);
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}