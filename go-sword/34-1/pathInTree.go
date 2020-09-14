package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

// 34-1：二叉树中和为某一值的路径
// 题目：输入一棵二叉树和一个整数，打印出二叉树中结点值的和为输入整数的所
// 有路径。从树的根结点开始往下一直到叶结点所经过的结点形成一条路径。

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func getMaxLength(root *TreeNode,k int) {
	result := arraylist.New()
	doIt(root,k,result)
}

func doIt(root *TreeNode,k int,result *arraylist.List) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		result.Add(root.value)
		if caculateSum(result) == k {
			printSlice(result)
		}
		result.Remove(result.Size()-1)
		return
	} else {
		result.Add(root.value)
		doIt(root.left,k,result)
		doIt(root.right,k,result)
		result.Remove(result.Size()-1)
		return
	}
}

func caculateSum(result *arraylist.List) int {
	sum := 0
	result.Each(func(_ int,value interface{}){
		sum += value.(int)
	})
	return sum
}

func printSlice(result *arraylist.List) {
	result.Each(func(_ int,value interface{}){
		fmt.Print(value," ")
	})
	fmt.Println()
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}

//            10
//         /      \
//        5        12
//       /\
//      4  7
// 有两条路径上的结点和为22
func Test1() {
	pNode10 := CreateBinaryTreeNode(10);
	pNode5 := CreateBinaryTreeNode(5);
	pNode12 := CreateBinaryTreeNode(12);
	pNode4 := CreateBinaryTreeNode(4);
	pNode7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNode10, pNode5, pNode12);
	ConnectTreeNodes(pNode5, pNode4, pNode7);
	getMaxLength(pNode10,22)
}

//            10
//         /      \
//        5        12
//       /\
//      4  7
// 没有路径上的结点和为15

func Test2() {
	pNode10 := CreateBinaryTreeNode(10);
	pNode5 := CreateBinaryTreeNode(5);
	pNode12 := CreateBinaryTreeNode(12);
	pNode4 := CreateBinaryTreeNode(4);
	pNode7 := CreateBinaryTreeNode(7);

	ConnectTreeNodes(pNode10, pNode5, pNode12);
	ConnectTreeNodes(pNode5, pNode4, pNode7);
	getMaxLength(pNode10,15)
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
// 有一条路径上面的结点和为15
func Test3() {
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);
	pNode3 := CreateBinaryTreeNode(3);
	pNode2 := CreateBinaryTreeNode(2);
	pNode1 := CreateBinaryTreeNode(1);

	ConnectTreeNodes(pNode5, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode2, nil);
	ConnectTreeNodes(pNode2, pNode1, nil);
	getMaxLength(pNode5,15)
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
// 没有路径上面的结点和为16
func Test4() {
	pNode1 := CreateBinaryTreeNode(1);
	pNode2 := CreateBinaryTreeNode(2);
	pNode3 := CreateBinaryTreeNode(3);
	pNode4 := CreateBinaryTreeNode(4);
	pNode5 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, nil, pNode2);
	ConnectTreeNodes(pNode2, nil, pNode3);
	ConnectTreeNodes(pNode3, nil, pNode4);
	ConnectTreeNodes(pNode4, nil, pNode5);
	getMaxLength(pNode1,16)
}

// 树中只有1个结点
func Test5() {
	pNode1 := CreateBinaryTreeNode(1);
	getMaxLength(pNode1,1)
}

func Test6() {
	getMaxLength((*TreeNode)(nil),0)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}