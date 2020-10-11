package main

import "fmt"

/*
173. 二叉搜索树迭代器
实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。

调用 next() 将返回二叉搜索树中的下一个最小的数。

示例：

BSTIterator iterator = new BSTIterator(root);
iterator.next();    // 返回 3
iterator.next();    // 返回 7
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 9
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 15
iterator.hasNext(); // 返回 true
iterator.next();    // 返回 20
iterator.hasNext(); // 返回 false

提示：

next() 和 hasNext() 操作的时间复杂度是 O(1)，并使用 O(h) 内存，其中 h 是树的高度。
你可以假设 next() 调用总是有效的，也就是说，当调用 next() 时，BST 中至少存在一个下一个最小的数。
 */

// 中序遍历

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
    //node7 := &TreeNode{Val: 11}
    connectNodes(node1,node2,node3)
    connectNodes(node2,node4,node5)
    connectNodes(node3,node6,nil)
    it := Constructor(node1)
    for it.HasNext() {
        fmt.Println(it.Next())
    }
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

type BSTIterator struct {
    stack []*TreeNode
    root *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        stack: make([]*TreeNode,0),
        root: root,
    }
}


/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	stack := this.stack
	root := this.root

    if len(stack) > 0  || root != nil {
        for root != nil {
            stack = append(stack,root)
            root = root.Left
        }
        // 出栈
        temp := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        root = temp.Right
        this.stack = stack
        this.root = root
        return temp.Val
    }
    return -1
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0 || this.root != nil
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */