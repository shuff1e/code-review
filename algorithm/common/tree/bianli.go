package main

import (
	"algorithm/common/help"
	"fmt"
)

func preOrder111(node *help.TreeNode) {
	root := node
	stack := help.NewMyStack()
	for root != nil || stack.Length() > 0 {
		// 入栈的时候，转移为左子节点
		for root != nil {
			fmt.Println(root.Value)
			stack.Push(root)
			root = root.Left
		}
		if stack.Length() > 0 {
			root,_ := stack.Pop()
			//出栈的时候转为右子节点
			root = root.(*help.TreeNode).Right
		}
	}
}

func midOrder111(node *help.TreeNode) {
	root := node
	stack := help.NewMyStack()
	for root != nil || stack.Length() > 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if stack.Length() > 0 {
			temp,_ := stack.Pop()
			root := temp.(*help.TreeNode)
			fmt.Println(root.Value)
			root = root.Right
		}
	}
}

func doubleQueueOrder111(node *help.TreeNode) {
	root := node
	stackData := help.NewMyStack()
	stackPrint := help.NewMyStack()
	for root != nil || stackData.Length() > 0 {
		for root != nil {
			stackData.Push(root)
			stackPrint.Push(root)
			root = root.Right
		}
		if stackData.Length() > 0 {
			temp,_ := stackData.Pop()
			root = temp.(*help.TreeNode)
			root = root.Left
		}
	}

	for stackPrint.Length() > 0 {
		temp,_ := stackPrint.Pop()
		fmt.Println(temp.(*help.TreeNode).Value)
	}
}

func postOrder111(node *help.TreeNode) {
	root := node
	prev := node
	stack := help.NewMyStack()
	for root != nil || stack.Length() > 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		if stack.Length() > 029 {
			temp,_ := stack.Peek()
			root := temp.(*help.TreeNode)
			right := root.Right

			if right == nil || prev == right{
				temp,_ = stack.Pop()
				fmt.Println(temp.(*help.TreeNode).Value)
				root = nil
				prev = right
				//出栈
				// 打印输出
			} else {
				root = right
			}
		}
	}
}

func cengcibianli2(node *help.TreeNode) {
	queue := help.NewMyQueue()
	queue.Add(node)
	for queue.Length() > 0 {
		temp,_ := queue.Poll()
		node := temp.(*help.TreeNode)
		fmt.Print(node.Value)
		fmt.Println()

		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}
}