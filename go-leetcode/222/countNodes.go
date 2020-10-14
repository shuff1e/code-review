package main

import (
	"fmt"
)

/*
222. 完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，
并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6
 */

func main() {
	//            1
	//        2      5
	//       3 4   6   nil
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 4}
	//node5 := (*TreeNode)(nil)
	//node6 := &TreeNode{Val: 9}
	//node6 := &TreeNode{Val: 6}
	node6 := (*TreeNode)(nil)
	node7 := (*TreeNode)(nil)
	connectNodes(node1, node2, node3)
	connectNodes(node2, node4, node5)
	connectNodes(node3, node6, node7)
	fmt.Println(countNodes(node1))
	fmt.Println(countNodes2(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

//1 + 2 +4
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	k := getLevel(root)
	count := 0
	help2(root,1,k,&count)
	result := 0
	temp := 1
	for i :=0;i<k;i++ {
		result += temp
		temp = temp*2
	}
	return result+count
}

func getLevel(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + getLevel(node.Left)
}

func help2(node *TreeNode,level int,k int,count *int) bool {
	if level == k {
		if node != nil {
			return true
		} else {
			*count --
			return false
		}
	}
	ok := help2(node.Right,level+1,k,count)
	if ok {
		return true
	}
	ok = help2(node.Left,level+1,k,count)
	return ok
}

func Pow(x,y int) int {
	result := 1
	for y > 0 {
		if y & 1 > 0 {
			result = result*x
		}
		x = x*x
		y = y >> 1
	}
	return result
}

func getDepth(node *TreeNode) int {
	result := 0
	for node.Left != nil {
		result ++
		node = node.Left
	}
	return result
}

// 假设depth为3，如果只有一个节点，depth为0

// 这样最后一层，如果全满的话一共8个节点，
// 最后一层的节点为0,1,2,3,4,5,6,7

//让我们来用二分搜索来构造从根节点到 idx 的移动序列。如，idx = 4。
//idx 位于 0,1,2,3,4,5,6,7 的后半部分，
//因此第一步是 node 节点从根节点开始，向右移动；然后 idx 位于 4,5,6,7 的前半部分，
//因此第二部是 node 节点向左移动；idx 位于 4,5 的前半部分，因此下一步是node节点向左移动。
// 这时候节点到了4的位置，判断该位置的node是不是nil

func exists(index int,depth int,node *TreeNode) bool {
	left := 0
	right := Pow(2,depth)-1
	for i :=0;i<depth;i++ {
		mid := (left+right)/2
		if index <= mid {
			right = mid
			node = node.Left
		} else {
			left = mid + 1
			node = node.Right
		}
	}
	return node != nil
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := getDepth(root)
	if depth == 0 {
		return 1
	}
	// 0的位置肯定有元素
	left := 1
	right := Pow(2,depth)-1
	for left <= right {
		mid := (left + right)/2
		if exists(mid,depth,root) {
			left = mid + 1
		} else {
			right  = mid -1
		}
	}
	return Pow(2,depth) - 1 + left
}
