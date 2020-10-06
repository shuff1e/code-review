package main

import "fmt"

/*
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	node := buildTree([]int{3,9,20,15,7},[]int{9,3,15,20,7})
	fmt.Println(node.Left.Val)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return help(preorder,inorder)
}

func help(pre,in []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}

	index := findIndex(in,pre[0])

	leftPre := pre[1:1 + index]
	leftIn := in[0:index]

	rightPre := pre[1+index:]
	rightIn := in[index+1:]

	node := &TreeNode{Val: pre[0]}
	node.Left = help(leftPre,leftIn)
	node.Right = help(rightPre,rightIn)
	return node
}

func findIndex(arr []int,target int) int {
	for i := 0;i<len(arr);i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

/*
以树
        3
       / \
      9  20
     /  /  \
    8  15   7
   / \
  5  10
 /
4
为例
preorder = [3, 9, 8, 5, 4, 10, 20, 15, 7]
inorder = [4, 5, 8, 10, 9, 3, 15, 20, 7]

我们用一个栈 stack 来维护「当前节点的所有还没有考虑过右儿子的祖先节点」，
栈顶就是当前节点。也就是说，只有在栈中的节点才可能连接一个新的右儿子。
同时，我们用一个指针 index 指向中序遍历的某个位置，初始值为 0。
index 对应的节点是「当前节点不断往左走达到的最终节点」，这也是符合中序遍历的，它的作用在下面的过程中会有所体现。

我们遍历 10，这时情况就不一样了。
我们发现 index 恰好指向当前的栈顶节点 4，也就是说 4 没有左儿子，那么 10 必须为栈中某个节点的右儿子。
那么如何找到这个节点呢？栈中的节点的顺序和它们在前序遍历中出现的顺序是一致的，
而且每一个节点的右儿子都还没有被遍历过，那么这些节点的顺序和它们在中序遍历中出现的顺序一定是相反的。

因此我们可以把 index 不断向右移动，并与栈顶节点进行比较。
如果 index 对应的元素恰好等于栈顶节点，那么说明我们在中序遍历中找到了栈顶节点，
所以将 index 增加 1 并弹出栈顶节点，直到 index 对应的元素不等于栈顶节点。
按照这样的过程，我们弹出的最后一个节点 x 就是 10 的双亲节点，
这是因为 10 出现在了 x 与 x 在栈中的下一个节点的中序遍历之间，因此 10 就是 x 的右儿子。

回到我们的例子，我们会依次从栈顶弹出 4，5 和 8，
并且将 index 向右移动了三次。我们将 10 作为最后弹出的节点 8 的右儿子，并将 10 入栈。

stack = [3, 9, 10]
index -> inorder[3] = 10

我们归纳出上述例子中的算法流程：

我们用一个栈和一个指针辅助进行二叉树的构造。初始时栈中存放了根节点（前序遍历的第一个节点），指针指向中序遍历的第一个节点；

我们依次枚举前序遍历中除了第一个节点以外的每个节点。
如果 index 恰好指向栈顶节点，那么我们不断地弹出栈顶节点并向右移动 index，
并将当前节点作为最后一个弹出的节点的右儿子；如果 index 和栈顶节点不同，我们将当前节点作为栈顶节点的左儿子；

无论是哪一种情况，我们最后都将当前的节点入栈。
 */

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{}
	stack = append(stack,root)

	inorderIndex := 0

	for i := 1;i<len(preorder);i++ {
		node :=stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorder[i]}
			stack = append(stack,node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex ++
			}
			node.Right = &TreeNode{Val: preorder[i]}
			stack = append(stack,node.Right)
		}
	}
	return root
}