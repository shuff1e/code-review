package main

/*

589. N叉树的前序遍历
给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :

         1
     /   |  \
     3   2   4
   /  \
   5   6

返回其前序遍历: [1,3,5,6,2,4]。

说明: 递归法很简单，你可以使用迭代法完成此题吗?

 */

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{}
	result := []int{}
	stack = append(stack,root)

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result,temp.Val)

		for i := len(temp.Children)-1;i>=0;i-- {
			stack = append(stack,temp.Children[i])
		}
	}
	return result
}
