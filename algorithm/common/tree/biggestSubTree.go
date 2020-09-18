package main

import "algorithm/common/help"

// 左子树是否是合格的
// 左子树最大值
// 左子树最大长度
func getMax(root *help.TreeNode) (max int,result int,valid bool) {

	if root == nil {
		return 0,0,true
	}
	leftMax,leftResult,leftValid := getMax(root.Left)
	rightMax,rightResult,rightValid := getMax(root.Right)

	if leftValid && rightValid && leftMax < root.Value && root.Value < rightMax {
		result = leftResult + rightResult + 1
		valid = true
		max = help.Max(leftMax,rightMax)
		max = help.Max(max,root.Value)
	} else {
		result = help.Max(leftResult,rightResult)
	}
	return
}
