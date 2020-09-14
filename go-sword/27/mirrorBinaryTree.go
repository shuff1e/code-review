package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

// 27：二叉树的镜像
// 题目：请完成一个函数，输入一个二叉树，该函数输出它的镜像。

type TreeNode struct {
	left *TreeNode
	right *TreeNode
	value int
}

// 类似前序遍历
func mirror(node *TreeNode) {
	if node == nil {
		return
	}
	temp := node.left
	node.left = node.right
	node.right = temp
	mirror(node.left)
	mirror(node.right)
}

func mirrorIteratively(root *TreeNode) {
	stack := linkedliststack.New()
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			// swap
			temp := node.left
			node.left = node.right
			node.right = temp

			node = node.left
		}
		temp,_ := stack.Pop()
		node = temp.(*TreeNode)
		node = node.right
	}
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
}

// 层次遍历，先入先出，用queue

func travelInLevel(root *TreeNode) {
	if root == nil {
		return
	}
	queue := arraylist.New()
	queue.Add(root)
	for !queue.Empty() {
		temp,_ := queue.Get(0)
		node := temp.(*TreeNode)
		fmt.Print(node.value," ")
		queue.Remove(0)
		if node.left != nil {
			queue.Add(node.left)
		}
		if node.right != nil {
			queue.Add(node.right)
		}
	}
	fmt.Println()
}

// 前序遍历
func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value," ")
	preOrder(root.left)
	preOrder(root.right)
}

// 类似递归和回溯的过程
// 尽量用画图的方式思考
// 入栈之前打印
func preOrderIteratively(root *TreeNode) {
	stack := linkedliststack.New()
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			fmt.Print(node.value," ")
			stack.Push(node)
			node = node.left
		}
		temp,_ := stack.Pop()
		node = temp.(*TreeNode)
		node = node.right
	}
	fmt.Println()
}

// 中序遍历
func midOrder(root *TreeNode) {
	if root == nil {
		return
	}
	midOrder(root.left)
	fmt.Println(root.value)
	midOrder(root.right)
}

// 出栈时打印
func midOrderIteratively(root *TreeNode) {
	stack := linkedliststack.New()
	node := root
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.left
		}
		temp,_ := stack.Pop()
		node = temp.(*TreeNode)
		fmt.Println(node.value)
		node = node.right
	}
	fmt.Println()
}

func postOrder(root *TreeNode) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.value," ")
}

// 出栈之后，如果右子树不为空，或者右子树已经访问过
// 再次入栈
func postOrderIteratively(root *TreeNode) {
	stack := linkedliststack.New()
	node := root
	lastVisited := (*TreeNode)(nil)
	for node != nil || !stack.Empty() {
		for node != nil {
			stack.Push(node)
			node = node.left
		}
		temp,_ := stack.Pop()
		node = temp.(*TreeNode)
		if node.right == nil || lastVisited == node.right {
			fmt.Print(node.value," ")
			lastVisited = node
			node = nil
		} else {
			stack.Push(node)
			node = node.right
		}
	}
	fmt.Println()
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

	postOrderIteratively(pNode8)
	travelInLevel(pNode8)
	mirrorIteratively(pNode8)
	travelInLevel(pNode8)
	mirror(pNode8)
	travelInLevel(pNode8)
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

	postOrderIteratively(pNode8)
	travelInLevel(pNode8)
	mirrorIteratively(pNode8)
	travelInLevel(pNode8)
	mirror(pNode8)
	travelInLevel(pNode8)
}

// 测试二叉树：出叶子结点之外，左右的结点都有且只有一个右子结点
//            8
//             7
//              6
//               5
//                4
func Test3() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode7 := CreateBinaryTreeNode(7);
	pNode6 := CreateBinaryTreeNode(6);
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);

	ConnectTreeNodes(pNode8, nil, pNode7);
	ConnectTreeNodes(pNode7, nil, pNode6);
	ConnectTreeNodes(pNode6, nil, pNode5);
	ConnectTreeNodes(pNode5, nil, pNode4);

	postOrderIteratively(pNode8)
	travelInLevel(pNode8)
	mirrorIteratively(pNode8)
	travelInLevel(pNode8)
	mirror(pNode8)
	travelInLevel(pNode8)
}

func Test4() {
	pNode8 := (*TreeNode)(nil)
	postOrderIteratively(pNode8)
	travelInLevel(pNode8)
	mirrorIteratively(pNode8)
	travelInLevel(pNode8)
	mirror(pNode8)
	travelInLevel(pNode8)
}

func Test5() {
	pNode8 := CreateBinaryTreeNode(8);
	postOrderIteratively(pNode8)
	travelInLevel(pNode8)
	mirrorIteratively(pNode8)
	travelInLevel(pNode8)
	mirror(pNode8)
	travelInLevel(pNode8)
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}
