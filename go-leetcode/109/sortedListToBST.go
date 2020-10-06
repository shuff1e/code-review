package main

/*
109. 有序链表转换二叉搜索树
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

     0
    / \
  -3   9
  /   /
-10  5
 */

type ListNode struct {
	Val int
	Next *ListNode
	}

type TreeNode struct {
	 Val int
 	Left *TreeNode
 	Right *TreeNode
}

// 1 2 3 4 5
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	pre,fast,slow := head,head,head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}

	next := slow.Next
	slow.Next = nil
	pre.Next = nil

	node := &TreeNode{Val: slow.Val}

	if pre != slow {
		node.Left = sortedListToBST(head)
	}

	node.Right = sortedListToBST(next)
	return node
}

/*
方法一的时间复杂度的瓶颈在于寻找中位数节点。
由于构造出的二叉搜索树的中序遍历结果就是链表本身，因此我们可以将分治和中序遍历结合起来，减少时间复杂度。

链表节点的编号为 [0,n)。中序遍历的顺序是「左子树 - 根节点 - 右子树」，那么在分治的过程中，
我们不用急着找出链表的中位数节点，而是使用一个占位节点，等到中序遍历到该节点时，再填充它的值。

这样一来，我们其实已经知道了这棵二叉搜索树的结构，并且题目给定了它的中序遍历结果，
那么我们只要对其进行中序遍历，就可以还原出整棵二叉搜索树了。
*/

func sortedListToBST2(head *ListNode) *TreeNode {
	globalHead := head
	length := getLength(head)
	return buildTree(0, length - 1,&globalHead)
}

func getLength(head *ListNode) int {
	count := 0
	for head != nil {
		head = head.Next
		count ++
	}
	return count
}

func buildTree(left,right int,global **ListNode) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1)/2
	node := &TreeNode{}
	node.Left = buildTree(left,mid-1,global)
	node.Val = (*global).Val
	*global = (*global).Next
	node.Right = buildTree(mid+1,right,global)
	return node
}