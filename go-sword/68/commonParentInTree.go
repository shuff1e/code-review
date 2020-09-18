package main

import (
	"fmt"
)

// 68：树中N个结点的最低公共祖先
// 题目：输入N个树结点，求它们的最低公共祖先。

// A：一个节点是祖先节点，
// 需要左子树包含全部节点，或者右子树包含全部节点
// 或者左右各奉献一个节点
// 节点需要子节点返回上述值，则自己也要向上返回这些值

// 返回1表示找到了一个节点，返回2表示全部找到了，返回0表示一个都没找到

type TreeNode struct {
	value int
	children []*TreeNode
}


func GetParent(root *TreeNode,target []*TreeNode) *TreeNode {
	_,result := getCommonParent(root,target,0)
	return result
}

// 只向上返回自己和自己的子孙节点的贡献
// 如果加上祖上节点的贡献，满足要求，直接返回
// 如果当前层不是要求的node，返回nil
func getCommonParent(root *TreeNode,target []*TreeNode,parentCount int) (int,*TreeNode) {
	if root == nil {
		return 0,nil
	}
	// 当前节点的贡献
	cur := 0
	for _,v := range target {
		if root == v {
			cur ++
		}
	}
	count := cur
	// 如果到当前节点为止，已经达到了要求，直接返回
	// 如果 parentCount == len(target)，根本不会到这一步
	// 所以count == 1
	if count + parentCount == len(target) {
		return count,nil
	}

	// 在当前节点和父亲节点的贡献的基础上，左子节点计算
	// 如果左子节点和左子节点的子孙节点的贡献达到了要求
	for _,child := range root.children {
		countLeft,nodeLeft := getCommonParent(child,target,count+parentCount)
		if countLeft == len(target) {
			if nodeLeft != nil {
				return countLeft,nodeLeft
			} else {
				return countLeft,root
			}
		}

		count += countLeft
		// parentCount 为0，cur为0的话，直接返回cur
		// 否则不是要求的节点，返回nil
		// parentCount 为0，cur为1的话，返回nil
		// parentCount 不为0,不是要求的节点，返回nil
		if count + parentCount == len(target) {
			if parentCount == 0 && cur == 0 {
				return count,root
			} else {
				return count,nil
			}
		}
	}

	// 返回当前节点和孩子节点的贡献
	return count,nil
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
}

// 形状普通的树
//              1
//            /   \
//           2     3
//       /       \
//      4         5
//     / \      / |  \
//    6   7    8  9  10
func Test1() {
	pNode1 := CreateTreeNode(1);
	pNode2 := CreateTreeNode(2);
	pNode3 := CreateTreeNode(3);
	pNode4 := CreateTreeNode(4);
	pNode5 := CreateTreeNode(5);
	pNode6 := CreateTreeNode(6);
	pNode7 := CreateTreeNode(7);
	pNode8 := CreateTreeNode(8);
	pNode9 := CreateTreeNode(9);
	pNode10 := CreateTreeNode(10);

	ConnectTreeNodes(pNode1, pNode2);
	ConnectTreeNodes(pNode1, pNode3);

	ConnectTreeNodes(pNode2, pNode4);
	ConnectTreeNodes(pNode2, pNode5);

	ConnectTreeNodes(pNode4, pNode6);
	ConnectTreeNodes(pNode4, pNode7);

	ConnectTreeNodes(pNode5, pNode8);
	ConnectTreeNodes(pNode5, pNode9);
	ConnectTreeNodes(pNode5, pNode10);

	Test("Test1", pNode1, []*TreeNode{pNode6, pNode8,pNode3}, pNode1);
}

// 树退化成一个链表
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
	pNode1 := CreateTreeNode(1);
	pNode2 := CreateTreeNode(2);
	pNode3 := CreateTreeNode(3);
	pNode4 := CreateTreeNode(4);
	pNode5 := CreateTreeNode(5);

	ConnectTreeNodes(pNode1, pNode2);
	ConnectTreeNodes(pNode2, pNode3);
	ConnectTreeNodes(pNode3, pNode4);
	ConnectTreeNodes(pNode4, pNode5);

	Test("Test2", pNode1, []*TreeNode{pNode5, pNode4}, pNode3);
}

// 树退化成一个链表，一个结点不在树中
//               1
//              /
//             2
//            /
//           3
//          /
//         4
//        /
//       5
func Test3() {
	pNode1 := CreateTreeNode(1);
	pNode2 := CreateTreeNode(2);
	pNode3 := CreateTreeNode(3);
	pNode4 := CreateTreeNode(4);
	pNode5 := CreateTreeNode(5);

	ConnectTreeNodes(pNode1, pNode2);
	ConnectTreeNodes(pNode2, pNode3);
	ConnectTreeNodes(pNode3, pNode4);
	ConnectTreeNodes(pNode4, pNode5);

	pNode6 := CreateTreeNode(6);

	Test("Test3", pNode1, []*TreeNode{pNode5, pNode6}, nil);
}

// 输入nullptr
func Test4() {
	Test("Test4", nil, nil, nil );
}

func Test5() {
	Test("Test5", nil, []*TreeNode{}, nil );
}

func Test(name string,root *TreeNode,target []*TreeNode,expected *TreeNode) {
	fmt.Println(name)
	fmt.Println(GetParent(root,target) == expected)
}

func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{
		value: value,
	}
}

func ConnectTreeNodes(root,child *TreeNode) {
	if root.children == nil {
		root.children = []*TreeNode{child}
	} else {
		root.children = append(root.children,child)
	}
}