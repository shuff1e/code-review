package main

import "fmt"

/*
114. 二叉树展开为链表
给定一个二叉树，原地将它展开为一个单链表。

例如，给定二叉树

    1
   / \
  2   5
 / \   \
3   4   6
将其展开为：

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// A：前序遍历

func main() {
	//            1
	//        2      5
	//       3 4   nil 6
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 4}
	//node6 := &TreeNode{Val: 9}
	node6 := (*TreeNode)(nil)
	node7 := &TreeNode{Val: 6}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node7,node6)
	//node1 := (*TreeNode)(nil)
	flatten(node1)
	for node1 != nil {
		fmt.Print(node1.Val,"->")
		if node1.Left != nil {
			panic("fuck")
		}
		node1 = node1.Right
	}
	//node1 = (*TreeNode)(nil)
	//fmt.Printf("%#v\n",levelOrder(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	pre := &TreeNode{Val: -1}
	help(root,pre)
}

func help(root *TreeNode, pre *TreeNode) *TreeNode {
	left := root.Left
	right := root.Right

	// 记得把left设置为nil
	root.Left = nil
	root.Right = nil

	pre.Right = root
	pre = pre.Right

	if left != nil {
		pre = help(left,pre)
	}
	if right != nil {
		pre = help(right,pre)
	}
	return pre
}

	//stack := []*TreeNode{}
	//for len(stack) > 0 || root != nil {
	//	for root != nil {
	//		pre.Right = root
	//		pre = pre.Right
	//		stack = append(stack,root)
	//		root = root.Left
	//	}
	//	root = stack[len(stack)-1]
	//	stack = stack[0:len(stack)-1]
	//	root = root.Right
	//}

// 迭代
func flatten2(root *TreeNode)  {
	if root == nil {
		return
	}
	stack := []*TreeNode{}
	stack = append(stack,root)
	prev := (*TreeNode)(nil)
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left = nil
			prev.Right = temp
		}
		if temp.Right != nil {
			stack = append(stack,temp.Right)
		}
		if temp.Left != nil {
			stack = append(stack,temp.Left)
		}
		prev = temp
	}
}

func flatten3(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left = nil
			curr.Right = next
		}
		curr = curr.Right
	}
}
