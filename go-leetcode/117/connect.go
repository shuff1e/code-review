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

/*
方法二：使用已建立的 next 指针
思路与算法

因为必须处理树上的所有节点，所以无法降低时间复杂度，但是可以尝试降低空间复杂度。

在方法一中，因为对树的结构一无所知，所以使用队列保证有序访问同一层的所有节点，
并建立它们之间的连接。然而不难发现：一旦在某层的节点之间建立了 next 指针，
那这层节点实际上形成了一个链表。因此，如果先去建立某一层的 next 指针，再去遍历这一层，就无需再使用队列了。

基于该想法，提出降低空间复杂度的思路：如果第 ii 层节点之间已经建立 next 指针，
就可以通过 next 指针访问该层的所有节点，同时对于每个第 i 层的节点，
我们又可以通过它的 left 和 right 指针知道其第 i+1 层的孩子节点是什么，
所以遍历过程中就能够按顺序为第 i+1 层节点建立 next 指针。

具体来说：

从根节点开始。因为第 0 层只有一个节点，不需要处理。可以在上一层为下一层建立 next 指针。
该方法最重要的一点是：位于第 x 层时为第 x+1 层建立 next 指针。
一旦完成这些连接操作，移至第 x+1 层为第 x+2 层建立 next 指针。
 */

func connect2(root *Node) *Node {
	last,nextStart := (*Node)(nil),(*Node)(nil)
	start := root
	for start != nil {
		last,nextStart = (*Node)(nil),(*Node)(nil)
		// 沿着链表走
		for p := start;p!=nil;p = p.Next {
			if p.Left != nil {
				handle(p.Left,&last,&nextStart)
			}
			if p.Right != nil {
				handle(p.Right,&last,&nextStart)
			}
		}
		start = nextStart
	}
	return root
}

func handle(node *Node,last,nextStart **Node) {
	if *last != nil {
		(*last).Next = node
	}
	if *nextStart == nil {
		*nextStart = node
	}
	*last = node
}