package main

import (
	"fmt"
	"strconv"
)

/*

1028. 从先序遍历还原二叉树
我们从二叉树的根节点 root 开始进行深度优先搜索。

在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。
（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。

如果节点只有一个子节点，那么保证该子节点为左子节点。

给出遍历输出 S，还原树并返回其根节点 root。

示例 1：

        1
       / \
      2   5
     / \  / \
    3   4 6  7


输入："1-2--3--4-5--6--7"
输出：[1,2,5,3,4,6,7]
示例 2：

         1
        / \
       2   5
      /    /
     3    6
    /    /
   4    7

输入："1-2--3---4-5--6---7"
输出：[1,2,5,3,null,6,null,4,null,7]
示例 3：

         1
        /
       401
       /  \
      349  88
     /
    90


输入："1-401--349---90--88"
输出：[1,401,null,349,88,90]


提示：

原始树中的节点数介于 1 和 1000 之间。
每个节点的值介于 1 和 10 ^ 9 之间。

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	str := "1-2--3---4-5--6---7"
	str = "1-2--3--4-5--6--7"
	str = "1-401--349---90--88"
	node := recoverFromPreorder(str)
	fmt.Println(node)
}

func recoverFromPreorder(S string) *TreeNode {
	i := 0
	for ;i<len(S);i++ {
		if S[i] == '-' {
			break
		}
	}
	val,_ := strconv.Atoi(S[:i])
	node := &TreeNode{Val: val}
	help(node,S,i,1)
	return node
}

func help(node *TreeNode,str string, index int,level int) int {
	if index >= len(str) {
		return index
	}

	// 看是在第几层
	i := index
	for ;i<len(str) && str[i] == '-';i++ {
	}

	if i - index < level {
		return index
	}

	// 这个层的value
	ii := i
	for ;ii <len(str)&& str[ii] != '-';ii++ {
	}

	val,_ := strconv.Atoi(str[i:ii])
	node.Left = &TreeNode{Val: val}

	// 左子节点
	leftIndex := help(node.Left,str,ii,level + 1)

	//
	i = leftIndex
	for ;i<len(str) && str[i] == '-';i++ {
	}

	if i - leftIndex < level {
		return leftIndex
	}

	//
	ii = i
	for ;ii <len(str)&& str[ii] != '-';ii++ {
	}

	//
	val,_ = strconv.Atoi(str[i:ii])
	node.Right = &TreeNode{Val: val}

	rightIndex := help(node.Right,str,ii,level+1)
	return rightIndex
}

/*

class Solution {
    public TreeNode recoverFromPreorder(String S) {
        Deque<TreeNode> path = new LinkedList<TreeNode>();
        int pos = 0;
        while (pos < S.length()) {
            int level = 0;
            while (S.charAt(pos) == '-') {
                ++level;
                ++pos;
            }
            int value = 0;
            while (pos < S.length() && Character.isDigit(S.charAt(pos))) {
                value = value * 10 + (S.charAt(pos) - '0');
                ++pos;
            }
            TreeNode node = new TreeNode(value);
            if (level == path.size()) {
                if (!path.isEmpty()) {
                    path.peek().left = node;
                }
            }
            else {
                while (level != path.size()) {
                    path.pop();
                }
                path.peek().right = node;
            }
            path.push(node);
        }
        while (path.size() > 1) {
            path.pop();
        }
        return path.peek();
    }
}

 */