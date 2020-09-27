package main

import "fmt"

/*
235. 二叉搜索树的最近公共祖先
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
				 6
                / \
	           2   8
		      / \  /\
             0   4 7 9
                / \
               3   5
示例 1:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
输出: 6
解释: 节点 2 和节点 8 的最近公共祖先是 6。
示例 2:

输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
输出: 2
解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉搜索树中。

 */


func main() {
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2}
	node8 := &TreeNode{Val: 8}
	node0 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}
	node7 := &TreeNode{Val: 7}
	node9 := &TreeNode{Val: 9}
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	ConnectTreeNodes(node6,node2,node8)
	ConnectTreeNodes(node2,node0,node4)
	ConnectTreeNodes(node8,node7,node9)
	ConnectTreeNodes(node4,node3,node5)
	fmt.Println(lowestCommonAncestor(node6,node2,node8))
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.Left = l
	p.Right = r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestorBest(root, p, q *TreeNode) (ancestor *TreeNode) {
	ancestor = root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return
		}
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//找到后再入栈，这样栈顶就是root
	stack1 := make([]*TreeNode,0)
	stack2 := make([]*TreeNode,0)
	findWithStack(root,p,&stack1)
	findWithStack(root,q,&stack2)

	parent := (*TreeNode)(nil)
	for len(stack1) > 0 && len(stack2) > 0 && stack1[len(stack1)-1] == stack2[len(stack2)-1] {
		parent = stack1[len(stack1)-1]
		stack1 = stack1[0:len(stack1)-1]
		stack2 = stack2[0:len(stack2)-1]
	}
	return parent
}

func findWithStack(root,node *TreeNode,stack *[]*TreeNode) {
	if node.Val < root.Val {
		findWithStack(root.Left,node,stack)
		*stack = append(*stack,root)
	} else if node.Val > root.Val {
		findWithStack(root.Right,node,stack)
		*stack = append(*stack,root)
	} else {
		*stack = append(*stack,node)
		return
	}
}

func lowestCommonAncestorBad(root, p, q *TreeNode) *TreeNode {
	result,_ := help(root,p,q)
	return result
}

func help(root,p,q *TreeNode) (*TreeNode,int){
	if root == nil {
		return nil,0
	}
	// 需要左子树和右子树传给root
	// 在自己这里发现了多少个节点
	// 以及是否发现了公共祖先
	result1,result2 := (*TreeNode)(nil),(*TreeNode)(nil)
	num1,num2 := 0,0

	if p.Val < root.Val || q.Val < root.Val {
		result1,num1 = help(root.Left,p,q)
	}
	if p.Val > root.Val || q.Val > root.Val {
		result2,num2 = help(root.Right,p,q)
	}

	if num1 == 2 {
		return result1,2
		//if result1 != nil {
		//	return result1,2
		//} else {
		//	return root,2
		//}
	}

	if num2 == 2 {
		return result2,2
		//if result2 != nil {
		//	return result2,2
		//} else {
		//	return root,2
		//}
	}

	if num1 == 1 && num2 == 1 {
		return root,2
	}

	if (num1 == 1 || num2 == 1) && (root == p || root == q) {
		return root,2
	}
	if num1 == 1 {
		return nil,num1
	}
	if num2 == 1 {
		return nil,num2
	}
	if root == p {
		return nil,1
	}
	if root == q {
		return nil,1
	}
	return nil,0
}