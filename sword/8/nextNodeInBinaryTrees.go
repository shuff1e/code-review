package main

import "fmt"

// 8：二叉树的下一个结点
// 题目：给定一棵二叉树和其中的一个结点，如何找出中序遍历顺序的下一个结点？
// 树中的结点除了有两个分别指向左右子结点的指针以外，还有一个指向父结点的指针。

// A：中序遍历
// do(root.left)
// print(root.value)
// do(root.right)

// 如果当前节点有右子节点，则下一个就是右子节点中最左边的节点

// 如果回溯之后，当前节点是父节点的左子节点，则下一个就是父节点

// 如果回溯之后，当前节点是父节点的右子节点，
// 继续回溯，如果父节点是祖父节点的左子节点，则下一个就是祖父节点

// 如果父节点是祖父节点的右子节点，继续回溯
// 直到根节点

type Node struct {
	left *Node
	right *Node
	value int
	parent *Node
}

func FindNextNode(n *Node) *Node {
	if n.right != nil {
		return findMostLeft(n.right)
	} else if n.parent.left == n {
		return n.parent
	} else {
		return findHelper(n)
	}
}

func findHelper(n *Node) *Node {
	if n == nil || n.parent == nil {
		return nil
	}
	if n == n.parent.left {
		return n.parent
	} else {
		return findHelper(n.parent)
	}
}

func findMostLeft(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return findMostLeft(n.left)
}

func main() {
	//            8
	//        6      10
	//       5 7    9  11
	node1 := &Node{value: 8}
	node2 := &Node{value: 6}
	node3 := &Node{value: 10}
	node4 := &Node{value: 5}
	node5 := &Node{value: 7}
	node6 := &Node{value: 9}
	node7 := &Node{value: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	result := FindNextNode(node5)
	if result == nil {
		fmt.Println(result)
	} else {
		fmt.Println(result.value)
	}
	//            5
	//          4
	//        3
	//      2
	node1 = &Node{value: 5}
	node2 = &Node{value: 4}
	node3 = &Node{value: 3}
	node4 = &Node{value: 2}
	connectNodes(node1,node2,nil)
	connectNodes(node2,node3,nil)
	connectNodes(node3,node4,nil)
	result = FindNextNode(node3)
	if result == nil {
		fmt.Println(result)
	} else {
		fmt.Println(result.value)
	}
	//        2
	//         3
	//          4
	//           5
	node1 = &Node{value: 2}
	node2 = &Node{value: 3}
	node3 = &Node{value: 4}
	node4 = &Node{value: 5}
	connectNodes(node1,nil,node2)
	connectNodes(node2,nil,node3)
	connectNodes(node3,nil,node4)
	result = FindNextNode(node4)
	if result == nil {
		fmt.Println(result)
	} else {
		fmt.Println(result.value)
	}
}

func connectNodes(p,left,right *Node) {
	if left != nil {
		p.left = left
		left.parent = p
	}
	if right != nil {
		p.right = right
		right.parent = p
	}
}