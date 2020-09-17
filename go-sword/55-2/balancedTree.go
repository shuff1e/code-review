package main

import "fmt"

// 55（二）：平衡二叉树
// 题目：输入一棵二叉树的根结点，判断该树是不是平衡二叉树。如果某二叉树中
// 任意结点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

// A：该节点为平衡二叉树
// 左子树是平衡的，而且右子树是平衡的
// 并且左右子树的高度差小于1

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func isBalancedTree(node *TreeNode) bool {
	_,ok := balanceTree(node)
	return ok
}

func balanceTree(node *TreeNode) (height int,balanced bool) {
	if node == nil {
		return 0,true
	}
	leftHeight,ok := balanceTree(node.left)
	if !ok {
		return leftHeight,ok
	}
	rightHeight,ok := balanceTree(node.right)
	if !ok {
		return rightHeight,ok
	}
	if Abs(leftHeight-rightHeight) > 1 {
		return 0,false
	}
	return 1 + Max(leftHeight,rightHeight),true
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()
}

func Test(name string,node *TreeNode,expected bool) {
	fmt.Println(name,isBalancedTree(node) == expected)
}

// 完全二叉树
//             1
//         /      \
//        2        3
//       /\       / \
//      4  5     6   7
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
	ConnectTreeNodes(pNode3, pNode6, pNode7);

	Test("Test1", pNode1, true);
}

// 不是完全二叉树，但是平衡二叉树
//             1
//         /      \
//        2        3
//       /\         \
//      4  5         6
//        /
//       7
func Test2() {
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

	Test("Test2", pNode1, true);
}

// 不是平衡二叉树
//             1
//         /      \
//        2        3
//       /\
//      4  5
//        /
//       6
func Test3() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);
	pNode6 := CreateBinaryTreeNode(6);

	ConnectTreeNodes(pNode1, pNode2, pNode3);
	ConnectTreeNodes(pNode2, pNode4, pNode5);
	ConnectTreeNodes(pNode5, pNode6, nil);

	Test("Test3", pNode1, false);
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
func Test4() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, pNode2, nil);
	ConnectTreeNodes(pNode2, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode5, nil);

	Test("Test4", pNode1, false);
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
func Test5() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, nil, pNode2);
	ConnectTreeNodes(pNode2, nil, pNode3);
	ConnectTreeNodes(pNode3, nil, pNode4);
	ConnectTreeNodes(pNode4, nil, pNode5);

	Test("Test5", pNode1, false);
}

// 树中只有1个结点
func Test6() {
	pNode1 := CreateBinaryTreeNode(1);
	Test("Test6", pNode1, true);
}

// 树中没有结点
func Test7() {
	Test("Test7", nil, true);
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}