package main

// 一棵二叉树
// 返回任意两节点之间的距离

import "fmt"

func main() {
	//           8
	//        9    10
	//      5  7  6  11
	node1 := &TreeNode{Val: 8}
	node2 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 10}
	node4 := &TreeNode{Val: 5}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	findOne,findAll,distance := getDistance(node1,node1,node2)
	fmt.Println(findOne,findAll,distance)
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
	Left *TreeNode
	Right *TreeNode
	Val int
}

// 类似最低的共同的父节点
// 在寻找最低的共同父节点的同时，返回距离

// distance,报告上层节点，上层节点距离target1或者target2的距离
func getDistance(root *TreeNode,target1,target2 *TreeNode) (findOne bool,findAll bool,distance int) {
	if root == nil {
		return false,false,0
	}
	leftFindOne,leftFindAll, leftDistance := getDistance(root.Left,target1,target2)
	rightFindOne,rightFindAll, rightDistance := getDistance(root.Right,target1,target2)
	if leftFindAll {
		return true,true,leftDistance
	}
	if rightFindAll {
		return true,true,rightDistance
	}

	if leftFindOne && rightFindOne {
		return true,true,leftDistance + rightDistance
	}

	if (leftFindOne || rightFindOne) && (root == target1 || root == target2) {
		if leftFindOne {
			return true,true,leftDistance
		}
		if rightFindOne {
			return true,true,rightDistance
		}
	}

	if leftFindOne {
		return true,false,leftDistance + 1
	}

	if rightFindOne {
		return true,false,rightDistance + 1
	}

	if root == target1 || root == target2 {
		return true,false,1
	}
	return false,false,0
}
