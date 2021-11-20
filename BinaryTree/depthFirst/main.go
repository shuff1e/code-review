package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func depthFirstSearch(root *TreeNode) {
    if root == nil {
        return
    }
    stack := NewStack()
    stack.Push(root)

    for len(stack) > 0 {
        temp := stack.Pop()
        fmt.Println(temp.Val)
        if temp.Right != nil {
            stack.Push(temp.Right)
        }
        if temp.Left != nil {
            stack.Push(temp.Left)
        }
    }
}

type Stack []*TreeNode

func NewStack() Stack {
    return make([]*TreeNode,0)
}

func (s *Stack) Push(node *TreeNode) {
    *s = append(*s,node)
}

func (s *Stack) Pop() *TreeNode {
    result := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return result
}


func main() {
    root := &TreeNode{
        Val: 1,
    }
    root.Right = &TreeNode{
        Val: 2,
    }
    root.Right.Left = &TreeNode{
        Val: 3,
    }

    depthFirstSearch(root)
}
