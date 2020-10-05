package main

import "fmt"

/*
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
 3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？
 */

// 二叉搜索树，如果对其进行中序遍历，得到的值序列是递增有序的

//1 6 3 4 5 2 7 8
//x       y

// 1 2 4 3 5 6 7 8
//     x y

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree1(root *TreeNode) {
	result := midOrder(root)
	a,b := findTwoNumber(result)
	insert(root,a.Val,b.Val)
}

func insert(root *TreeNode,a,b int)  {
	if root != nil {
		if root.Val == a {
			root.Val = b
		} else if root.Val == b {
			root.Val = a
		}
		insert(root.Left,a,b)
		insert(root.Right,a,b)
	}
}

func findTwoNumber(arr []*TreeNode) (a ,b *TreeNode) {
	pred := (*TreeNode)(nil)
	for _,v := range arr {
		if pred != nil && pred.Val > v.Val {
			b = v
			if a == nil {
				a=pred
			} else {
				break
			}
		}
		pred = v
	}
	return
}

func midOrder(root *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		result = append(result,root)
		root = root.Right
	}
	return result
}

func recoverTree2(root *TreeNode) {
	x,y,pred := (*TreeNode)(nil),(*TreeNode)(nil),(*TreeNode)(nil)
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if pred != nil && pred.Val > root.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	swap(x,y)
}

func recoverTree3(root *TreeNode)  {
	x,y,pred,predecessor := (*TreeNode)(nil),(*TreeNode)(nil),(*TreeNode)(nil),(*TreeNode)(nil)
	for root != nil {
		if root.Left != nil {
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			// first travel
			if predecessor.Right != root {
				predecessor.Right = root
				root = root.Left
				// already travel
			} else {
				predecessor.Right = nil
				if pred != nil && pred.Val > root.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				root = root.Right
			}

		} else  {
			//1 6 3 4 5 2 7 8
			//x       y
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	swap(x,y)
}

func swap(x,y *TreeNode) {
	//fmt.Println(x,y)
	temp := x.Val
	x.Val = y.Val
	y.Val = temp
}

//          1
//         /
//        2
//      /  \
//     3    4
// 首先root = 1，设置4.right=1

// root = 2,设置3.right = 2
// root = 3,root.left == nil ,root = root.right,root = 2
// 3.right = 2，说明左边已经遍历玩了，砍断指针，root = 4
// root.left == nil，root = root.right root = 1

func mirrorTravel(root *TreeNode) {
	predecessor := (*TreeNode)(nil)
	for root != nil {
		if root.Left != nil {
			// 设置predecessor
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}
			// 第一次遍历
			if predecessor.Right != root {
				predecessor.Right = root
				root = root.Left
				// 已经遍历完了
			} else {
				predecessor.Right = nil
				fmt.Print(root.Val," ")
				root = root.Right
			}
		} else {
			fmt.Print(root.Val," ")
			root = root.Right
		}
	}
	fmt.Println()
}

func main() {
	//            8
	//        6      10
	//       5 7    9  11
	//node1 := &TreeNode{Val: 8}
	//node2 := &TreeNode{Val: 9}
	//node3 := &TreeNode{Val: 10}
	//node4 := &TreeNode{Val: 5}
	//node5 := &TreeNode{Val: 7}
	//node6 := &TreeNode{Val: 6}
	//node7 := &TreeNode{Val: 11}
	//connectNodes(node1,node2,node3)
	//connectNodes(node2,node4,node5)
	//connectNodes(node3,node6,node7)
	//mirrorTravel(node1)
	//recoverTree2(node1)
	//mirrorTravel(node1)
	//3
//	 / \
//	1   4
//	   /
//	  2
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 1}
	node3 := &TreeNode{Val: 4}
	node4 := &TreeNode{Val: 2}
	connectNodes(node1,node2,node3)
	connectNodes(node3,node4,nil)
	result := midOrder(node1)
	for _,v := range result {
		fmt.Println(v.Val)
	}
	fmt.Println(findTwoNumber(result))
	mirrorTravel(node1)
	recoverTree1(node1)
	mirrorTravel(node1)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}
