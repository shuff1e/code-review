package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 32（三）：之字形打印二叉树
// 题目：请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺
// 序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，
// 其他行以此类推。

//            10
//         /      \
//        6        14
//       /\        /\
//      4  8     12  16

// 打印顺序为 10,14,6,4,8,12,16

// A：入队之前先打印
// 如果从左到右，先左子树入队，然后下一层从后pop，从右到左
// 如果从右到左，先右子树入队，然后下一层从后pop，从左到右
// 那看来其实是用了一个stack

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func travelTree2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack1 := linkedliststack.New()
	stack1.Push(root)

	stack2 := linkedliststack.New()

	result := make([]int,1)
	result[0] = root.value

	leftToRight := false

	for !stack1.Empty() || !stack2.Empty() {
		for !stack1.Empty() {
			temp,_ := stack1.Pop()
			node := temp.(*TreeNode)
			if !leftToRight {
				if node.right != nil {
					result = append(result,node.right.value)
					stack2.Push(node.right)
				}
				if node.left != nil {
					result = append(result,node.left.value)
					stack2.Push(node.left)
				}
			} else {
				if node.left != nil {
					result = append(result,node.left.value)
					stack2.Push(node.left)
				}
				if node.right != nil {
					result = append(result,node.right.value)
					stack2.Push(node.right)
				}
			}
		}
		// swap stack1 and stack2
		temp := stack1
		stack1 = stack2
		stack2 = temp
		leftToRight = !leftToRight
	}
	return result
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
func Test3() {
	pNode8 := CreateBinaryTreeNode(8);
	//pNode7 := CreateBinaryTreeNode(7);
	//pNode6 := CreateBinaryTreeNode(6);
	//pNode5 := CreateBinaryTreeNode(5);
	//pNode4 := CreateBinaryTreeNode(4);

	//ConnectTreeNodes(pNode8, nil, pNode7);
	//ConnectTreeNodes(pNode7, nil, pNode6);
	//ConnectTreeNodes(pNode6, nil, pNode5);
	//ConnectTreeNodes(pNode5, nil, pNode4);

	result := travelTree2(pNode8)
	fmt.Println(result)
}

// 测试二叉树：出叶子结点之外，左右的结点都有且只有一个左子结点
//            8
//          7
//        6
//      5
//    4
func Test2() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode7 := CreateBinaryTreeNode(7);
	pNode6 := CreateBinaryTreeNode(6);
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);

	ConnectTreeNodes(pNode8, pNode7, nil);
	ConnectTreeNodes(pNode7, pNode6, nil);
	ConnectTreeNodes(pNode6, pNode5, nil);
	ConnectTreeNodes(pNode5, pNode4, nil);

	result := travelTree2(pNode8)
	fmt.Println(result)
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

	result := travelTree2(pNode8)
	fmt.Println(result)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}