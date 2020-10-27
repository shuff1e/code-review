package main

/*

1367. 二叉树中的列表
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，
那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。



示例 1：



输入：head = [4,2,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
解释：树中蓝色的节点构成了与链表对应的子路径。
示例 2：



输入：head = [1,4,2,6], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：true
示例 3：

输入：head = [1,4,2,6,8], root = [1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]
输出：false
解释：二叉树中不存在一一对应链表的路径。


提示：

二叉树和链表中的每个节点的值都满足 1 <= node.val <= 100 。
链表包含的节点数目在 1 到 100 之间。
二叉树包含的节点数目在 1 到 2500 之间。

 */

type ListNode struct {
 	Val int
 	Next *ListNode
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil && head == nil {
		return true
	}
	if root == nil && head != nil {
		return false
	}
	if root != nil && head == nil {
		return true
	}

	return dfs(head,root) || isSubPath(head,root.Left) || isSubPath(head,root.Right)
}

func dfs(head *ListNode,root *TreeNode) bool {
	if root == nil && head == nil {
		return true
	}
	if root == nil && head != nil {
		return false
	}
	if root != nil && head == nil {
		return true
	}

	if head.Val != root.Val {
		return false
	}
	return dfs(head.Next,root.Left) || dfs(head.Next,root.Right)
}

func isSubPath2(head *ListNode, root *TreeNode) bool {
	arr := getArrFromList(head)
	lps := computeTempArray(arr)
	return treeKMP(root,arr,lps,0)
}

func computeTempArray(arr []int) []int {
	result := make([]int,len(arr))

	index := 0
	for i := 1;i<len(arr); {
		if arr[i] == arr[index] {
			result[i] = index + 1
			i ++
			index ++
		} else {
			if index != 0 {
				index = result[index-1]
			} else {
				result[i] = 0
				i ++
			}
		}
	}
	return result
}

func getArrFromList(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result,head.Val)
		head = head.Next
	}
	return result
}

func treeKMP(root *TreeNode, arr,lps []int, j int) bool {
	if root == nil && j < len(arr) {
		return false
	}
	if root == nil && j == len(arr) {
		return true
	}
	if root != nil && j == len(arr) {
		return true
	}

	if root.Val == arr[j] {
		j++
		return treeKMP(root.Left,arr,lps,j) ||
			treeKMP(root.Right,arr,lps,j)
	} else if j != 0 {
		j = lps[j-1]
		return treeKMP(root,arr,lps,j)
	} else {
		return treeKMP(root.Left,arr,lps,j) || treeKMP(root.Right,arr,lps,j)
	}
}

/*

实现的代码非常的人性化。
算法步骤：

将链表转为数组，复杂度O(m)，m为链表长度。
求链表的nexts数组，O(m)。
KMP匹配，O(n)

public class Solution {

    public int lengthOf(ListNode head) {
        int len = 0;
        while (head != null) {
            head = head.next;
            len++;
        }
        return len;
    }

    public int[] valsOf(ListNode head, int len) {
        int[] vals = new int[len];
        int i = 0;
        while (head != null) {
            vals[i++] = head.val;
            head = head.next;
        }
        return vals;
    }

    public int[] nextsOf(int[] p) {
        int[] nexts = new int[p.length];
        nexts[0] = -1;
        int j = 0;
        for (int i = 1; i < p.length - 1; ) {
            if (p[i] == p[j]) {
                if (p[++i] != p[++j]) nexts[i] = j;
                else nexts[i] = nexts[j];
            } else if (j > 0) {
                j = nexts[j];
            } else {
                nexts[++i] = 0;
            }
        }
        return nexts;
    }

    public boolean findSingle(int headVal, TreeNode root) {
        if (root == null) return false;
        if (root.val == headVal) return true;
        return findSingle(headVal, root.left) || findSingle(headVal, root.right);
    }

    public boolean treeKMP(TreeNode root, int[] p, int[] nexts, int j) {
        if (root == null) return false;
        if (root.val == p[j]) {
            j++;
            if (j == p.length) return true;
            return treeKMP(root.left, p, nexts, j) || treeKMP(root.right, p, nexts, j);
        } else if (nexts[j] != -1) {
            return treeKMP(root, p, nexts, nexts[j]);
        } else {
            return treeKMP(root.left, p, nexts, j) || treeKMP(root.right, p, nexts, j);
        }
    }

    public boolean isSubPath(ListNode head, TreeNode root) {
        int listLen = lengthOf(head);
        if (listLen == 0) return true;
        else if (listLen == 1) return findSingle(head.val, root);
        int[] p = valsOf(head, listLen);
        return treeKMP(root, p, nextsOf(p), 0);
    }
}

 */