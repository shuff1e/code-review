package main

import "fmt"

// 36：二叉搜索树与双向链表
// 题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求
// 不能创建任何新的结点，只能调整树中结点指针的指向。

// A：树的话，就是分为左子树和右子树，这样递归去解决

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func convert(node *TreeNode) *TreeNode {
	head,_ := convertHelper(node)
	return head
}

func convertHelper(node *TreeNode) (head , tail *TreeNode) {
	if node == nil {
		return nil,nil
	}

	leftHead,leftEnd := convertHelper(node.left)
	node.left = leftEnd
	if leftEnd != nil {
		leftEnd.right = node
		head = leftHead
	} else {
		head = node
	}

	rightHead,rightEnd := convertHelper(node.right)
	node.right = rightHead
	if rightHead != nil {
		rightHead.left = node
		tail = rightEnd
	} else {
		tail = node
	}

	return

}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
}

//            10
//         /      \
//        6        14
//       /\        /\
//      4  8     12  16

func Test1() {
	pNode10 := CreateBinaryTreeNode(10);
	pNode6 := CreateBinaryTreeNode(6);
	pNode14 := CreateBinaryTreeNode(14);
	pNode4 := CreateBinaryTreeNode(4);
	pNode8 := CreateBinaryTreeNode(8);
	pNode12 := CreateBinaryTreeNode(12);
	pNode16 := CreateBinaryTreeNode(16);

	ConnectTreeNodes(pNode10, pNode6, pNode14);
	ConnectTreeNodes(pNode6, pNode4, pNode8);
	ConnectTreeNodes(pNode14, pNode12, pNode16);

	Test("Test1", pNode10);
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
func Test2() {
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);
	pNode3 := CreateBinaryTreeNode(3);
	pNode2 := CreateBinaryTreeNode(2);
	pNode1 := CreateBinaryTreeNode(1);

	ConnectTreeNodes(pNode5, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode2, nil);
	ConnectTreeNodes(pNode2, pNode1, nil);

	Test("Test2", pNode5);
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

	Test("Test3", pNode1);
}

func Test4() {
	pNode1 := CreateBinaryTreeNode(1);
	Test("Test4", pNode1);
}

func Test5() {
	Test("Test5", nil)
}

func Test(name string,node *TreeNode) {
	fmt.Println(name)
	printTree(node)
	head := convert(node)
	printDoubleLinkedList(head)
}

func printTree(node *TreeNode) {
	printTreeHelper(node)
	fmt.Println()
}

func printTreeHelper(node *TreeNode) {
	if node == nil {
		return
	}
	printTreeHelper(node.left)
	fmt.Print(node.value," ")
	printTreeHelper(node.right)
}

func printDoubleLinkedList(node *TreeNode) {
	if node == nil {
		return
	}
	for {
		fmt.Print(node.value," ")
		if node.right == nil {
			break
		}
		node = node.right
	}
	fmt.Println()
	for {
		fmt.Print(node.value," ")
		if node.left == nil {
			break
		}
		node = node.left
	}
	fmt.Println()
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}