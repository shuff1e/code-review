package main

import "algorithm/common/help"

func getBiggestSubArea(root *help.TreeNode) int {

}

func dealLeft(root *help.TreeNode2,pivot int) int {
	if root == nil {
		return 0
	}
	temp := root
	for temp != nil {
		if temp.Value > pivot {
			root.Num -= temp.Num
			break
		}
		temp = temp.Right
	}
	return root.Num
}

func dealRight(root *help.TreeNode2,pivot int) int {
	if root == nil {
		return 0
	}
	temp := root
	for temp != nil {
		if temp.Value < pivot {
			root.Num -= temp.Num
			break
		}
		temp = temp.Left
	}
	return root.Num
}

func give(root *help.TreeNode2) int {
	if root == nil {
		return 0
	}

	leftGive := give(root.Left)
	leftGive = dealLeft(root.Left,root.Value)

	rightGive := give(root.Right)
	rightGive = dealRight(root.Right,root.Value)

	root.Num = leftGive + rightGive + 1

	return root.Num
}

//                      2
//                3
//            4
//        1       1
//