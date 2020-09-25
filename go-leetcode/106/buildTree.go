package main

/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

 3
/ \
9  20
  /  \
 15   7

 */

// A：后续遍历，最后一个就是头节点

type TreeNode struct {
	   Val int
	   Left *TreeNode
	   Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	node := &TreeNode{
		Val: postorder[len(postorder)-1],
	}

	index := -1
	for i,v := range inorder {
		if postorder[len(postorder)-1] == v {
			index = i
			break
		}
	}

	// 左子树在后序遍历中的长度
	length2 := index

	node.Left = buildTree(inorder[0:index],postorder[0:length2])
	node.Right = buildTree(inorder[index+1:],postorder[length2:len(postorder)-1])

	return node
}