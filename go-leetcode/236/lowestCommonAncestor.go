package main

import "fmt"

/*
236. 二叉树的最近公共祖先
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]

         3
       /   \
      5     1
     / \   / \
     6  2  0  8
       / \
       7 4

示例 1:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
示例 2:

输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。


说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

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
	node7 := &TreeNode{Val: 11}
	connectNodes(node1,node2,node3)
	connectNodes(node2,node4,node5)
	connectNodes(node3,node6,node7)
	result := lowestCommonAncestor(node1,node4,node6)
	fmt.Println(result.Val)
}

func connectNodes(p,left,right *TreeNode) {
	if left != nil {
		p.Left = left
	}
	if right != nil {
		p.Right = right
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_,result := help(root,p,q)
	return result
}

func help(root ,p,q *TreeNode) (count int,result *TreeNode) {
	if root == nil {
		return 0,nil
	}
	leftCount,leftResult := help(root.Left,p,q)
	if leftCount == 2 {
		return 2,leftResult
	}

	rightCount,rightResult := help(root.Right,p,q)
	if rightCount == 2 {
		return 2,rightResult
	}

	if leftCount == 1 && rightCount == 1 {
		return 2,root
	}

	tempCount := 0
	if root == p || root == q {
		tempCount = 1
	}

	if tempCount == 1 && (leftCount == 1 || rightCount == 1) {
		return 2,root
	}


	if tempCount == 1 || leftCount == 1 || rightCount == 1 {
		return 1,nil
	}

	return 0,nil

}

/*

1.从根节点开始遍历整棵二叉树，用哈希表记录每个节点的父节点指针。
2.从 p 节点开始不断往它的祖先移动，并用数据结构记录已经访问过的祖先节点。
3.同样，我们再从 q 节点开始不断往它的祖先移动，如果有祖先已经被访问过，即意味着这是 p 和 q 的深度最深的公共祖先，即 LCA 节点。

class Solution {
    Map<Integer, TreeNode> parent = new HashMap<Integer, TreeNode>();
    Set<Integer> visited = new HashSet<Integer>();

    public void dfs(TreeNode root) {
        if (root.left != null) {
            parent.put(root.left.val, root);
            dfs(root.left);
        }
        if (root.right != null) {
            parent.put(root.right.val, root);
            dfs(root.right);
        }
    }

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        dfs(root);
        while (p != null) {
            visited.add(p.val);
            p = parent.get(p.val);
        }
        while (q != null) {
            if (visited.contains(q.val)) {
                return q;
            }
            q = parent.get(q.val);
        }
        return null;
    }
}

 */