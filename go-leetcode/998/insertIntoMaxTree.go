package main

import "fmt"

/*

998. 最大二叉树 II
最大树定义：一个树，其中每个节点的值都大于其子树中的任何其他值。

给出最大树的根节点 root。

就像之前的问题那样，给定的树是从表 A（root = Construct(A)）递归地使用下述 Construct(A) 例程构造的：

如果 A 为空，返回 null
否则，令 A[i] 作为 A 的最大元素。创建一个值为 A[i] 的根节点 root
root 的左子树将被构建为 Construct([A[0], A[1], ..., A[i-1]])
root 的右子树将被构建为 Construct([A[i+1], A[i+2], ..., A[A.length - 1]])
返回 root
请注意，我们没有直接给定 A，只有一个根节点 root = Construct(A).

假设 B 是 A 的副本，并附加值 val。保证 B 中的值是不同的。

返回 Construct(B)。



示例 1：

                  5
                 /
  4             4
 / \           / \
1   3         1   3
   /             /
  2             2

输入：root = [4,1,3,null,null,2], val = 5
输出：[5,4,null,1,3,null,null,2]
解释：A = [1,4,2,3], B = [1,4,2,3,5]
示例 2：

  5             5
 / \           / \
2   4         2   4
 \             \   \
  1             1   3

输入：root = [5,2,4,null,1], val = 3
输出：[5,2,4,null,1,null,3]
解释：A = [2,1,5,4], B = [2,1,5,4,3]
示例 3：

  5              5
 / \           /   \
2   3         2     4
 \             \   /
  1             1 3

输入：root = [5,2,3,null,1], val = 4
输出：[5,2,4,null,1,3]
解释：A = [2,1,5,3], B = [2,1,5,3,4]


提示：

1 <= B.length <= 100

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	//            5
	//        2      3
	//         1
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 1}
	connectNodes(node1,node2,node3)
	connectNodes(node2,nil,node4)
	result := insertIntoMaxTree(node1,4)
	fmt.Println(result)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	dup := &TreeNode{Val: root.Val}
	dup = help(root,dup,val)
	return dup
}

func help(root *TreeNode,dup *TreeNode,val int) *TreeNode {
	if val >= dup.Val {
		temp := &TreeNode{Val: val}
		temp.Left = dup
		dfs(root,dup)
		return temp
	}

	if root.Right == nil {

		dup.Right = &TreeNode{Val: val}

		if root.Left != nil {
			dup.Left = &TreeNode{Val: root.Left.Val}
			dfs(root.Left,dup.Left)
		}

	} else {
		if root.Left != nil {
			dup.Left = &TreeNode{Val: root.Left.Val}
			dfs(root.Left,dup.Left)
		}
		dup.Right = &TreeNode{Val: root.Right.Val}

		dup.Right = help(root.Right,dup.Right,val)
	}
	return dup
}

func dfs(root,dup *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		dup.Left = &TreeNode{Val: root.Left.Val}
		dfs(root.Left,dup.Left)
	}
	if root.Right != nil {
		dup.Right = &TreeNode{Val: root.Right.Val}
		dfs(root.Right,dup.Right)
	}
}