package main

import (
	"algorithm/common/help"
)

// 左子树返回最深的深度，最大的距离
func getPath(root *help.TreeNode) (depth int ,path int) {
	if root == nil {
		return 0,0
	}
	leftDepth,leftPath := getPath(root.Left)
	rightDepth,rightPath := getPath(root.Right)

	depth = help.Max(leftDepth,rightDepth) + 1
	path = help.Max(leftPath,rightPath)
	path = help.Max(path,leftDepth + rightDepth + 1)
	return
}
