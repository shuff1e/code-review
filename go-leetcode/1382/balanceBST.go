package main

import "fmt"

/*

1382. 将二叉搜索树变平衡
给你一棵二叉搜索树，请你返回一棵 平衡后 的二叉搜索树，新生成的树应该与原来的树有着相同的节点值。

如果一棵二叉搜索树中，每个节点的两棵子树高度差不超过 1 ，我们就称这棵二叉搜索树是 平衡的 。

如果有多种构造方法，请你返回任意一种。



示例：



输入：root = [1,null,2,null,3,null,4,null,null]
输出：[2,1,3,null,null,null,4]
解释：这不是唯一的正确答案，[3,1,4,null,2,null,null] 也是一个可行的构造方案。


提示：

树节点的数目在 1 到 10^4 之间。
树节点的值互不相同，且在 1 到 10^5 之间。

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	//            1
	//        2      5
	//       3 4   nil 6
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	connectNodes(node1,nil,node2)
	connectNodes(node2,nil,node3)
	result := balanceBST(node1)
	fmt.Println(result)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

// 二叉搜索树的中序遍历，是递增的
// 从中序遍历的结果再次构造二叉树

func balanceBST(root *TreeNode) *TreeNode {
	slice := midOrder(root)
	return build(slice,0,len(slice)-1)
}

func midOrder(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result,root.Val)
		root = root.Right
	}
	return result
}

func build(arr []int,l,r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (l+r)/2
	result := &TreeNode{
		Val: arr[mid],
	}
	result.Left = build(arr,l,mid-1)
	result.Right = build(arr,mid+1,r)
	return result
}