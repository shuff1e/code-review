package main

/*
117. 填充每个节点的下一个右侧节点指针 II
给定一个二叉树

struct Node {
int val;
Node *left;
Node *right;
Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

示例：

    1                1 ->
   / \              / \
   2  3            2-> 3->
  / \  \          / \   \
 4   5  7        4 ->5 ->7
输入：root = [1,2,3,4,5,null,7]
输出：[1,#,2,3,#,4,5,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。

提示：

树中的节点数小于 6000
-100 <= node.val <= 100

 */

// A：层次遍历

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue1 := []*Node{}
	queue2 := []*Node{}
	queue1 = append(queue1,root)

	for len(queue1) > 0 {
		for len(queue1) > 0 {
			// 出队列的时候设置next指针
			temp := queue1[0]
			queue1 = queue1[1:]
			if len(queue1) > 0 {
				temp.Next = queue1[0]
			}
			if temp.Left != nil {
				queue2= append(queue2,temp.Left)
			}
			if temp.Right != nil {
				queue2 = append(queue2,temp.Right)
			}
		}
		temp := queue2
		queue2 = queue1
		queue1 = temp
	}
	return root
}