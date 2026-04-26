package main

import "fmt"

/*
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
二叉树：[3,9,20,null,null,15,7],

3
/ \
9  20
   /  \
  15   7
返回其层次遍历结果：

[
[3],
[9,20],
[15,7]
]
 */

func main() {
	//            8
	//        6      10
	//       5 7    9  11
	node1 := &TreeNode{Val: 8}
	node2 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 10}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 9}
	node7 := &TreeNode{Val: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	node1 = (*TreeNode)(nil)
	fmt.Printf("%#v\n",levelOrder(node1))
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}
	q1 := NewQueue()
	q2 := NewQueue()

	q1.AddLast(root)

	for !q1.Empty() || !q2.Empty() {
		tempList := []int{}
		for !q1.Empty() {
			temp := q1.RemoveFirst()
			tempNode := temp.(*TreeNode)
			tempList = append(tempList, tempNode.Val)
			if tempNode.Left != nil {
				q2.AddLast(tempNode.Left)
			}
			if tempNode.Right != nil {
				q2.AddLast(tempNode.Right)
			}
		}
		temp := q1
		q1 = q2
		q2 = temp
		if len(tempList) > 0 {
			result = append(result,tempList)
		}
	}
	return result
}

type queue struct {
	data []interface{}
}

func NewQueue() *queue {
	return &queue{
		data: make([]interface{},0),
	}
}

func (q *queue) AddLast(v interface{}) {
	q.data = append(q.data,v)
}

func (q *queue) RemoveFirst() interface{} {
	result := q.data[0]
	q.data = q.data[1:]
	return result
}

func (q *queue) Empty() bool {
	return len(q.data) == 0
}