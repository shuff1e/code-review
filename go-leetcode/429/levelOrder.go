package main

/*

429. N叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。

例如，给定一个 3叉树 :

		1
	/   |   \
   3    2    4
 /  \
5    6

返回其层序遍历:

[
     [1],
     [3,2,4],
     [5,6]
]


说明:

树的深度不会超过 1000。
树的节点总数不会超过 5000。

 */

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	result := [][]int{}

	stack1 := []*Node{}
	stack1 = append(stack1,root)

	result = append(result,[]int{root.Val})

	for len(stack1) > 0 {
		help := []*Node{}
		helpArr := []int{}

		for len(stack1) > 0 {
			// queue !!!
			temp := stack1[0]
			stack1 = stack1[1:]

			for _,v := range temp.Children {
				help = append(help,v)
				helpArr = append(helpArr,v.Val)
			}

		}
		if len(helpArr) > 0 {
			result = append(result,helpArr)
		}
		stack1 = help
	}
	return result
}