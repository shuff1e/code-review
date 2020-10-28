package main

import (
	"fmt"
)

/*

437. 路径总和 III
给定一个二叉树，它的每个结点都存放着一个整数值。

找出路径和等于给定数值的路径总数。

路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。

示例：

root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

返回 3。和等于 8 的路径有:

1.  5 -> 3
2.  5 -> 2 -> 1
3.  -3 -> 11

 */

func main() {
	node1 := &TreeNode{Val: 1}
	//node2 := &TreeNode{Val: 1}
	//node3 := &TreeNode{Val: 1}
	//node4 := &TreeNode{Val: 3}
	//node5 := &TreeNode{Val: 2}
	//node7 := &TreeNode{Val: 11}
	//
	//node8 := &TreeNode{Val: 3}
	//node9 := &TreeNode{Val: -2}
	//node10 := &TreeNode{Val: 1}
	connectNodes(node1,nil,nil)
	//connectNodes(node2,node4,node5)
	//connectNodes(node3,nil,node7)
	//connectNodes(node4,node8,node9)
	//connectNodes(node5,nil,node10)

	result := pathSum(node1,0)
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

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// sum 没有出现，dict[sum] = level，记录和第一次出现的位置
// if dict.contains(sum-k) 说明找到了，然后 count ++

// 如果sum是在当前层第一次出现的，删除sum

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	dict := map[int]int{}
	// index,count
	dict[0] = 1
	count := help(root,sum,0,dict)
	return count
}

func help(root *TreeNode,k ,sum int,dict map[int]int) int {
	if root == nil {
		return 0
	}
	sum += root.Val

	res := 0
	res += dict[sum-k]

	dict[sum] = dict[sum] + 1

	res += help(root.Left,k,sum,dict)
	res += help(root.Right,k,sum,dict)

	dict[sum] = dict[sum] - 1

	return res
}

/*

class Solution {
    public int pathSum(TreeNode root, int sum) {
        // key是前缀和, value是大小为key的前缀和出现的次数
        Map<Integer, Integer> prefixSumCount = new HashMap<>();
        // 前缀和为0的一条路径
        prefixSumCount.put(0, 1);
        // 前缀和的递归回溯思路
        return recursionPathSum(root, prefixSumCount, sum, 0);
    }

    private int recursionPathSum(TreeNode node, Map<Integer, Integer> prefixSumCount, int target, int currSum) {
        // 1.递归终止条件
        if (node == null) {
            return 0;
        }
        // 2.本层要做的事情
        int res = 0;
        // 当前路径上的和
        currSum += node.val;

        //---核心代码
        // 看看root到当前节点这条路上是否存在节点前缀和加target为currSum的路径
        // 当前节点->root节点反推，有且仅有一条路径，如果此前有和为currSum-target,而当前的和又为currSum,两者的差就肯定为target了
        // currSum-target相当于找路径的起点，起点的sum+target=currSum，当前点到起点的距离就是target
        res += prefixSumCount.getOrDefault(currSum - target, 0);
        // 更新路径上当前节点前缀和的个数
        prefixSumCount.put(currSum, prefixSumCount.getOrDefault(currSum, 0) + 1);
        //---核心代码

        // 3.进入下一层
        res += recursionPathSum(node.left, prefixSumCount, target, currSum);
        res += recursionPathSum(node.right, prefixSumCount, target, currSum);

        // 4.回到本层，恢复状态，去除当前节点的前缀和数量
        prefixSumCount.put(currSum, prefixSumCount.get(currSum) - 1);
        return res;
    }
}

 */