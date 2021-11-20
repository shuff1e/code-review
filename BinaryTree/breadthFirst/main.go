package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func breadthFirstSearch(root *TreeNode) {
    if root == nil {
        return
    }
    queue := NewQueue()
    queue.Add(root)

    for len(queue) > 0 {
        temp := queue.Front()
        fmt.Println(temp.Val)
        if temp.Left != nil {
            queue.Add(temp.Left)
        }
        if temp.Right != nil {
            queue.Add(temp.Right)
        }
    }
}

type Queue []*TreeNode

func NewQueue() Queue {
    return make([]*TreeNode,0)
}

func (q *Queue) Add(node *TreeNode) {
    *q = append([]*TreeNode{node},(*q)...)
}

func (q *Queue) Front() *TreeNode {
    result := (*q)[0]
    *q = (*q)[1:]
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

    breadthFirstSearch(root)
}
