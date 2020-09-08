package main

import "fmt"

// 7：重建二叉树
// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输
// 入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列{1,
// 2, 4, 7, 3, 5, 6, 8}和中序遍历序列{4, 7, 2, 1, 5, 3, 8, 6}，则重建出
// 图2.6所示的二叉树并输出它的头结点。

// A：前序遍历的第一个节点为1
// 则1为根节点，1的左子树的 前序遍历为2，4，7；中序遍历为4，7，2

type Node struct {
	left *Node
	right *Node
	value int
}

func construct(preOrder []int,midOrder []int) *Node {
	// 为空
	if len(preOrder) == 0 {
		return nil
	}
	// 只有一个节点
	root := &Node{value: preOrder[0]}
	if len(preOrder) == 1 {
		return root
	}
	// 2个或者多个节点
	index := findIndex(midOrder,preOrder[0])
	// 只有左子树
	if index == len(midOrder) - 1 {
		root.left = construct(preOrder[1:1+index],midOrder[0:index])
		// 只有右子树
	} else if index == 0 {
		root.right = construct(preOrder[1+index:],midOrder[index+1:])
	} else {
		// 左右都有
		root.left = construct(preOrder[1:1+index],midOrder[0:index])
		root.right = construct(preOrder[1+index:],midOrder[index+1:])
	}
	return root
}

func findIndex(arr []int,n int) int {
	for i := 0;i<len(arr);i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}

func main() {
	// 普通二叉树
	//              1
	//           /     \
	//          2       3
	//         /       / \
	//        4       5   6
	//         \         /
	//          7       8
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	node := construct(preOrder, midOrder)
	fmt.Println(node.value)

	// 所有结点都没有右子结点
	//            1
	//           /
	//          2
	//         /
	//        3
	//       /
	//      4
	//     /
	//    5
	preOrder = []int{1, 2, 3, 4, 5}
	midOrder = []int{5, 4, 3, 2, 1}
	node = construct(preOrder, midOrder)
	fmt.Println(node.value)
	// 所有结点都没有左子结点
	//            1
	//             \
	//              2
	//               \
	//                3
	//                 \
	//                  4
	//                   \
	//                    5
	preOrder = []int{1, 2, 3, 4, 5}
	midOrder = []int{1, 2, 3, 4, 5}
	node = construct(preOrder, midOrder)
	fmt.Println(node.value)
	// 完全二叉树
	//              1
	//           /     \
	//          2       3
	//         / \     / \
	//        4   5   6   7
	preOrder = []int{1, 2, 4, 5, 3, 6, 7}
	midOrder = []int{4, 2, 5, 1, 6, 3, 7}
	node = construct(preOrder,midOrder)
	fmt.Println(node.value)
}