package main

/*

863. 二叉树中所有距离为 K 的结点
给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 K 。

返回到目标结点 target 距离为 K 的所有结点的值的列表。 答案可以以任何顺序返回。

示例 1：

输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
输出：[7,4,1]
解释：
所求结点为与目标结点（值为 5）距离为 2 的结点，
值分别为 7，4，以及 1



注意，输入的 "root" 和 "target" 实际上是树上的结点。
上面的输入仅仅是对这些对象进行了序列化描述。


提示：

给定的树是非空的。
树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
目标结点 target 是树上的结点。
0 <= K <= 1000.

 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	parent := map[*TreeNode]*TreeNode{}
	dfs(root,nil,parent)

	queue := []*TreeNode{}
	queue = append(queue,nil)
	queue = append(queue,target)

	seen := map[*TreeNode]struct{}{}
	seen[nil] = struct{}{}
	seen[target] = struct{}{}


	distance := 0
	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		if temp == nil {
			if distance == K {
				ans := []int{}
				for i := 0;i<len(queue);i++ {
					ans = append(ans,queue[i].Val)
				}
				return ans
			}
			distance ++
			queue = append(queue,nil)
		} else {
			if _,ok := seen[temp.Left];!ok {
				queue = append(queue,temp.Left)
				seen[temp.Left] = struct{}{}
			}
			if _,ok := seen[temp.Right];!ok {
				queue = append(queue,temp.Right)
				seen[temp.Right] = struct{}{}
			}
			if _,ok := seen[parent[temp]];!ok {
				queue = append(queue,parent[temp])
				seen[parent[temp]] = struct{}{}
			}
		}
	}
	return nil
}

func dfs(node *TreeNode,parent *TreeNode,dict map[*TreeNode]*TreeNode) {
	if node != nil {
		dict[node] = parent
		dfs(node.Left,node,dict)
		dfs(node.Right,node,dict)
	}
}

func distanceK2(root *TreeNode, target *TreeNode, K int) []int {
	ans := make([]int,0)
	dfs2(root,target,K,&ans)
	return ans
}

// 返回上层节点距离target的距离

func dfs2(root *TreeNode,target *TreeNode,K int,ans *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		subtreeAdd(root,0,K,ans)
		return 1
	}

	L := dfs2(root.Left,target,K,ans)
	R := dfs2(root.Right,target,K,ans)

	if L != -1 {
		if L == K {
			*ans = append(*ans,root.Val)
		}
		subtreeAdd(root.Right,L + 1,K,ans)
		return L + 1
	} else if R != -1 {
		if R == K {
			*ans = append(*ans,root.Val)
		}
		subtreeAdd(root.Left,R+1,K,ans)
		return R + 1
	} else {
		return -1
	}
}

func subtreeAdd(node *TreeNode,dist ,K int,ans *[]int) {
	if node == nil {
		return
	}
	if dist == K {
		*ans = append(*ans,node.Val)
	} else {
		subtreeAdd(node.Left,dist+1,K,ans)
		subtreeAdd(node.Right,dist+1,K,ans)
	}
}