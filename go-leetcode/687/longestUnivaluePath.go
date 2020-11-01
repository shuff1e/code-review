package main

import (
	"fmt"
)

/*

687. 最长同值路径
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。

 */

func main() {
	//            1
	//        2      5
	//       3 4   nil 6
	node1 := &TreeNode{Val: 5}
	node2 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 1}
	node5 := &TreeNode{Val: 1}
	//node6 := &TreeNode{Val: 9}
	node6 := (*TreeNode)(nil)
	node7 := &TreeNode{Val: 5}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node7,node6)
	fmt.Println(longestUnivaluePath(node1))
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

func longestUnivaluePath(root *TreeNode) int {
	count := 0
	help(root,&count)
	return count
}

// 返回值：以该节点向左或者向右能延伸的最大长度
// 每层递归，结算count的时候，为向左和向右的长度的和

func help(root *TreeNode,count *int) int {
	if root == nil {
		return 0
	}
	leftDistance := help(root.Left,count)
	rightDistance := help(root.Right,count)

	currLeftDistance := 0
	// 向左延伸的最大长度
	if root.Left != nil && root.Left.Val == root.Val {
		currLeftDistance = leftDistance + 1
	}

	currRightDistance := 0
	// 向右延伸的最大长度
	if root.Right != nil && root.Right.Val == root.Val {
		currRightDistance = rightDistance + 1
	}

	*count = Max(*count,currLeftDistance + currRightDistance)
	return Max(currLeftDistance,currRightDistance)
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

/*

也不算特别难的题，可能由于之前写过类似的题目。 其实二叉树递归的难点就在于怎么构思：子节点向父节点返回的是什么?或者说，当前节点向其父节点返回的是什么?
这题中，当前节点返回给父节点的值就是： 从当前节点出发，向下延伸与其值相同的最大深度 于是返回值分两种情况：
1.if( 如果当前节点与其左右节点都不相等)，那么深度为0，则返回0 2. else，这个最大深度就取其 左右子树返回值中的较大者 + 1

然后，在上面这个dfs的遍历过程中，还可以做一些其他的事情，这题做的就是 计算路径长度。由于子树的返回值已经确定了，所以需要额外的一个全局变量。
如何计算路径长度呢？其实知道了和自己数值相等的左右子树的最大高度了，那么 把左右子树返回的值相加 就是贯穿自己的最长路径了。

 */

// 下面是自己写的错误答案

/*
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	help(root,&count,0)
	return count
}

func help(root *TreeNode,count *int,curr int) int{
	if curr > *count {
		*count = curr
	}

	if root.Left != nil && root.Right != nil {
		if root.Left.Val == root.Val && root.Right.Val == root.Val {
			leftLen := help(root.Left,count,curr + 1)
			rightLen := help(root.Right,count,curr + 1)
			if leftLen + rightLen > *count {
				*count = leftLen + rightLen
			}
			return Max(leftLen,rightLen)
		}
	}

	if root.Left != nil {
		if root.Left.Val == root.Val {
			return help(root.Left,count,curr + 1)
		} else {
			help(root.Left,count,0)
		}
	}

	if root.Right != nil {
		if root.Right.Val == root.Val {
			return help(root.Right,count,curr + 1)
		} else {
			help(root.Right,count,0)
		}
	}
	return curr
}

func Max(x,y int) int {
	if x > y {
		return x
	}
	return y
}

func longestUnivaluePath2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_,mine,notMine,_ := help2(root)
	return Max(mine,notMine)
}

// 返回包含我在内的最大长度，
// 以及不包含我在内的最大长度

func help2(root *TreeNode) (value int,pathLenOfMine int,pathLenNotMine int,valid bool) {
	if root == nil {
		return 0,0,0,false
	}

	leftValue,leftLen,notLeftLen,ok1 := help2(root.Left)
	rightValue,rightLen,notRightLen,ok2 := help2(root.Right)

	if !ok1 && !ok2 {
		return root.Val,0,0,true
	}

	if ok1 && ok2 {
		if leftValue == root.Val && rightValue == root.Val {
			return root.Val,leftLen + 1 + rightLen + 1,Max(notLeftLen,notRightLen),true
		}
		if leftValue == root.Val {
			return root.Val,leftLen + 1,
			Max(Max(notLeftLen,notRightLen),rightLen),true
		}
		if rightValue == root.Val {
			return root.Val,rightLen + 1,
			Max(Max(notLeftLen,notRightLen),leftLen),true
		}
		return root.Val,0,
		Max(Max(notLeftLen,notRightLen),
		Max(leftLen,rightLen)),true
	}

	if ok1 {
		if leftValue == root.Val {
			return root.Val,leftLen + 1,notLeftLen,true
		} else {
			return root.Val,0,Max(leftLen,notLeftLen),true
		}
	}

	if ok2 {
		if rightValue == root.Val {
			return root.Val,rightLen + 1,notRightLen,true
		} else {
			return root.Val,0,Max(rightLen,notRightLen),true
		}
	}
	return 0,0,0,false
}
 */
