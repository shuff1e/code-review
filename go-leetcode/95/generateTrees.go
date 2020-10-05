package main

import "fmt"

/*
95. 不同的二叉搜索树 II
给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。

示例：

输入：3
输出：
[
[1,null,3,2],
[3,2,null,1],
[3,1,null,null,2],
[2,1,3],
[1,null,2,null,3]
]
解释：
以上的输出对应以下 5 种不同结构的二叉搜索树：

1         3     3      2      1
 \       /     /      / \      \
  3     2     1      1   3      2
 /     /       \                 \
2     1         2                 3
 */

// 1234 5 6789
// 1 234 5

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	result := generateTrees(4)
	fmt.Println(len(result))
	for _,v := range result {
		levelOrder(v)
		fmt.Println()
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return help(1,n)
}

// 左子节点返回自己生成的树
// 当前节点也要向父节点返回自己生成的树
func help(start,end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}

	for index := start;index<=end;index++ {
		leftNodes := help(start,index-1)
		rightNodes := help(index+1,end)

		for i := 0;i<len(leftNodes);i++ {
			for j :=0;j<len(rightNodes);j++ {
				node := &TreeNode{Val: index}
				node.Left = leftNodes[i]
				node.Right = rightNodes[j]
				result = append(result,node)
			}
		}
	}
	return result
}

func levelOrder(node *TreeNode) {
	queue1 := []*TreeNode{}
	queue2 := []*TreeNode{}
	queue1 = append(queue1,node)
	for len(queue1) > 0 {
		for len(queue1) > 0 {
			temp := queue1[0]
			fmt.Print(temp.Val," ")
			if temp.Left!= nil {
				queue2 = append(queue2,temp.Left)
			}
			if temp.Right != nil {
				queue2 = append(queue2,temp.Right)
			}
			queue1 = queue1[1:]
		}
		fmt.Println()
		temp := queue1
		queue1 = queue2
		queue2 = temp
	}
}