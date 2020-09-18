package main

import (
	"algorithm/common/help"
)

func isBalanced(node *help.TreeNode) bool {
	_,isBalanced := getHeight(node,1)
	return isBalanced
}

func getHeight(node *help.TreeNode,level int) (height int,isBalanced bool) {
	if node == nil {
		return level,true
	}
	leftHeight,leftBalanced := getHeight(node.Left,level+1)
	if !leftBalanced {
		return leftHeight,leftBalanced
	}

	rightHeight,rightBalanced := getHeight(node.Right,level+1)
	if !isBalanced {
		return rightHeight,rightBalanced
	}

	if help.Abs(leftHeight-rightHeight) > 1 {
		return help.Max(leftHeight,rightHeight),false
	}
	return help.Max(leftHeight,rightHeight),true
}

func main() {

}