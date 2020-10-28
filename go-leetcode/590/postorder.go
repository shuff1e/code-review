package main

/*

590. N叉树的后序遍历
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :

         1
     /   |  \
     3   2   4
   /  \
   5   6

返回其后序遍历: [5,6,3,2,4,1].

说明: 递归法很简单，你可以使用迭代法完成此题吗?

 */

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{}
	stack = append(stack,root)

	result := []int{}
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{temp.Val},result...)
		for i := 0;i<len(temp.Children);i++ {
			stack = append(stack,temp.Children[i])
		}
	}

	return result
}