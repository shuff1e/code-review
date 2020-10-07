package main

/*
111. 二叉树的最小深度
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

   3
  / \
 9  20
   /  \
  15   7
返回它的最小深度  2.
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

const MAX = 0x7fffffff

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right ==nil {
		return 1
	}

	min1,min2 := MAX,MAX
	if root.Left != nil {
		min1 = minDepth(root.Left)
	}
	if root.Right != nil {
		min2 = minDepth(root.Right)
	}
	return Min(min1,min2) + 1
}

func Min(x,y int) int {
	if x < y {
		return x
	}
	return y
}

/*
同样，我们可以想到使用广度优先搜索的方法，遍历整棵树。

当我们找到一个叶子节点时，直接返回这个叶子节点的深度。广度优先搜索的性质保证了最先搜索到的叶子节点的深度一定最小。

class Solution {
    class QueueNode {
        TreeNode node;
        int depth;

        public QueueNode(TreeNode node, int depth) {
            this.node = node;
            this.depth = depth;
        }
    }

    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }

        Queue<QueueNode> queue = new LinkedList<QueueNode>();
        queue.offer(new QueueNode(root, 1));
        while (!queue.isEmpty()) {
            QueueNode nodeDepth = queue.poll();
            TreeNode node = nodeDepth.node;
            int depth = nodeDepth.depth;
            if (node.left == null && node.right == null) {
                return depth;
            }
            if (node.left != null) {
                queue.offer(new QueueNode(node.left, depth + 1));
            }
            if (node.right != null) {
                queue.offer(new QueueNode(node.right, depth + 1));
            }
        }

        return 0;
    }
}

 */