package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
)

// Q：给定一颗二叉树的头节点head，和一个整数sum
// 二叉树上的每个节点都有数字
// 我们规定路径必须是从上往下的
// 求二叉树上累加和为sum的最长路径长度

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func getMaxLength(root *TreeNode,k int) int {
	mmp1 := hashmap.New()
	mmp2 := hashmap.New()
	mmp1.Put(0,-1)
	mmp2.Put(-1,0)
	return doit(root,0,0,k,mmp1,mmp2)
}

func doit(root *TreeNode,sum ,level ,k int,mmp1,mmp2 *hashmap.Map) int {
	if root == nil {
		return 0
	}
	curSum := root.value + sum
	if _,ok := mmp1.Get(curSum);!ok {
		mmp1.Put(curSum,level)
		mmp2.Put(level,curSum)
	}
	max := 0
	if tempLevel,ok := mmp1.Get(curSum - k);ok {
		if level - tempLevel.(int) > max {
			max = level - tempLevel.(int)
		}
	}
	leftMax := doit(root.left,curSum,level + 1,k,mmp1,mmp2)
	rightMax := doit(root.right,curSum,level +1,k,mmp1,mmp2)
	if leftMax > max {
		max = leftMax
	}
	if rightMax > max {
		max = rightMax
	}
	if temp,ok := mmp2.Get(level);ok {
		mmp2.Remove(level)
		mmp1.Remove(temp)
	}
	return max
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
	fmt.Println(getMaxLength(pNode10,22))
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
	fmt.Println(getMaxLength(pNode10,15))
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
	fmt.Println(getMaxLength(pNode5,15))
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
	fmt.Println(getMaxLength(pNode1,16))
}

// 树中只有1个结点
func Test5() {
	pNode1 := CreateBinaryTreeNode(1);
	fmt.Println(getMaxLength(pNode1,1))
}

func Test6() {
	fmt.Println(getMaxLength((*TreeNode)(nil),0))
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}
