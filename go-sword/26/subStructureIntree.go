package main

import "fmt"

// 26：树的子结构
// 题目：输入两棵二叉树A和B，判断B是不是A的子结构。

// A结构              B结构
//       8             8
//      /  \          / \
//     8   7         9   2
//    / \
//   9   2
//      / \
//     4   7
// 如图，B是A的子结构

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

// 左子树或者右子树包含，或者以当前节点能匹配到，就返回true
func isSub(p,c *TreeNode) bool {
	if c == nil {
		return true
	}
	if p == nil {
		return false
	}
	return isSubStructure(p,c) || isSub(p.left,c) || isSub(p.right,c)
}

// c是否是以p开头的子结构
// c.value == p.value && isSubStructure(c.left,p.left) && isSubStructure(c.right,p.right)
func isSubStructure(p,c *TreeNode) bool {
	if c == nil {
		return true
	}
	if p == nil {
		return false
	}
	if p.value != c.value {
		return false
	}
	return isSubStructure(p.left,c.left) && isSubStructure(p.right,c.right)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}

// 树中结点含有分叉，树B是树A的子结构
//                  8                8
//              /       \           / \
//             8         7         9   2
//           /   \
//          9     2
//               / \
//              4   7
func Test1() {
	pNodeA1 := CreateBinaryTreeNode(8);
	pNodeA2 := CreateBinaryTreeNode(8);
	pNodeA3 := CreateBinaryTreeNode(7);
	pNodeA4 := CreateBinaryTreeNode(9);
	pNodeA5 := CreateBinaryTreeNode(2);
	pNodeA6 := CreateBinaryTreeNode(4);
	pNodeA7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNodeA1, pNodeA2, pNodeA3);
	ConnectTreeNodes(pNodeA2, pNodeA4, pNodeA5);
	ConnectTreeNodes(pNodeA5, pNodeA6, pNodeA7);

	pNodeB1 := CreateBinaryTreeNode(8);
	pNodeB2 := CreateBinaryTreeNode(9);
	pNodeB3 := CreateBinaryTreeNode(2);

	ConnectTreeNodes(pNodeB1, pNodeB2, pNodeB3);

	Test("Test1", pNodeA1, pNodeB1, true);
}

// 树中结点含有分叉，树B不是树A的子结构
//                  8                8
//              /       \           / \
//             8         7         9   2
//           /   \
//          9     3
//               / \
//              4   7
func Test2() {
	pNodeA1 := CreateBinaryTreeNode(8);
	pNodeA2 := CreateBinaryTreeNode(8);
	pNodeA3 := CreateBinaryTreeNode(7);
	pNodeA4 := CreateBinaryTreeNode(9);
	pNodeA5 := CreateBinaryTreeNode(3);
	pNodeA6 := CreateBinaryTreeNode(4);
	pNodeA7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNodeA1, pNodeA2, pNodeA3);
	ConnectTreeNodes(pNodeA2, pNodeA4, pNodeA5);
	ConnectTreeNodes(pNodeA5, pNodeA6, pNodeA7);

	pNodeB1 := CreateBinaryTreeNode(8);
	pNodeB2 := CreateBinaryTreeNode(9);
	pNodeB3 := CreateBinaryTreeNode(2);

	ConnectTreeNodes(pNodeB1, pNodeB2, pNodeB3);

	Test("Test2", pNodeA1, pNodeB1, false);
}

// 树中结点只有左子结点，树B是树A的子结构
//                8                  8
//              /                   /
//             8                   9
//           /                    /
//          9                    2
//         /
//        2
//       /
//      5

func Test3() {
	pNodeA1 := CreateBinaryTreeNode(8);
	pNodeA2 := CreateBinaryTreeNode(8);
	pNodeA3 := CreateBinaryTreeNode(9);
	pNodeA4 := CreateBinaryTreeNode(2);
	pNodeA5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNodeA1, pNodeA2, nil);
	ConnectTreeNodes(pNodeA2, pNodeA3, nil);
	ConnectTreeNodes(pNodeA3, pNodeA4, nil);
	ConnectTreeNodes(pNodeA4, pNodeA5, nil);

	pNodeB1 := CreateBinaryTreeNode(8);
	pNodeB2 := CreateBinaryTreeNode(9);
	pNodeB3 := CreateBinaryTreeNode(2);

	ConnectTreeNodes(pNodeB1, pNodeB2, nil);
	ConnectTreeNodes(pNodeB2, pNodeB3, nil);

	Test("Test3", pNodeA1, pNodeB1, true);
}

func Test(name string,node1,node2 *TreeNode,match bool) {
	fmt.Println(name,isSub(node1,node2) == match)
}

func main() {
	Test1()
	Test2()
	Test3()
}