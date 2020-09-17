package main

import "fmt"

// 54：二叉搜索树的第k个结点
// 题目：给定一棵二叉搜索树，请找出其中的第k大的结点。



























//            8
//        6      10
//       5 7    9  11

// A：中序遍历，就是从小到大的
//

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func GetKthNode(node *TreeNode,k int) int {
	cur := 0
	result := getKthNode(node,k,&cur)
	if result == nil {
		return -1
	}
	return result.value
}

func getKthNode(node *TreeNode,k int,cur *int) *TreeNode {
	if node == nil {
		return nil
	}
	temp := getKthNode(node.left,k,cur)
	if temp != nil {
		return temp
	}
	*cur = *cur + 1
	if *cur == k {
		return node
	}
	temp = getKthNode(node.right,k,cur)
	if temp != nil {
		return temp
	}
	return nil
}

func main() {
	TestA()
	TestB()
	TestC()
	TestD()
	TestE()
}

func TestA() {
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

	Test("TestA0", pNode8, 0, -1);
	Test("TestA1", pNode8, 1,  5);
	Test("TestA2", pNode8, 2,  6);
	Test("TestA3", pNode8, 3,  7);
	Test("TestA4", pNode8, 4,  8);
	Test("TestA5", pNode8, 5,  9);
	Test("TestA6", pNode8, 6,  10);
	Test("TestA7", pNode8, 7,  11);
	Test("TestA8", pNode8, 8, -1);
}

//               5
//              /
//             4
//            /
//           3
//          /
//         2
//        /
//       1
func TestB() {
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);
	pNode3 := CreateBinaryTreeNode(3);
	pNode2 := CreateBinaryTreeNode(2);
	pNode1 := CreateBinaryTreeNode(1);

	ConnectTreeNodes(pNode5, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode2, nil);
	ConnectTreeNodes(pNode2, pNode1, nil);

	Test("TestB0", pNode5, 0, -1);
	Test("TestB1", pNode5, 1,  1);
	Test("TestB2", pNode5, 2,  2);
	Test("TestB3", pNode5, 3,  3);
	Test("TestB4", pNode5, 4,  4);
	Test("TestB5", pNode5, 5,  5);
	Test("TestB6", pNode5, 6, -1);
}

// 1
//  \
//   2
//    \
//     3
//      \
//       4
//        \
//         5
func TestC() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, nil, pNode2);
	ConnectTreeNodes(pNode2, nil, pNode3);
	ConnectTreeNodes(pNode3, nil, pNode4);
	ConnectTreeNodes(pNode4, nil, pNode5);

	Test("TestC0", pNode1, 0, -1);
	Test("TestC1", pNode1, 1,  1);
	Test("TestC2", pNode1, 2,  2);
	Test("TestC3", pNode1, 3,  3);
	Test("TestC4", pNode1, 4,  4);
	Test("TestC5", pNode1, 5,  5);
	Test("TestC6", pNode1, 6, -1);
}

// There is only one node in a tree
func TestD() {
	pNode1 := CreateBinaryTreeNode(1);

	Test("TestD0", pNode1, 0, -1);
	Test("TestD1", pNode1, 1,  1);
	Test("TestD2", pNode1, 2, -1);
}

// empty tree
func TestE() {
	Test("TestE0", nil, 0,  -1);
	Test("TestE1", nil, 1,  -1);
}

func Test(name string,node *TreeNode,k,expected int) {
	fmt.Println(name,GetKthNode(node,k),expected)
	if GetKthNode(node,k) != expected {
		panic("fuck")
	}
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}