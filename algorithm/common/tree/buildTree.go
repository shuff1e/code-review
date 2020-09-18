package main

import (
	"algorithm/common/help"
	"fmt"
)

var preArr = []int{1,2,4,11,3,5,6,8}
var midArr = []int{4,2,11,1,5,3,8,6}

func build(pre []int,mid []int) *help.TreeNode{
	if len(pre) == 0 {
		return nil
	}
	value := pre[0]
	index := findIndex(value,mid) //0
	leftLen := index //0
	rightLen := len(mid)-1-index//0

	left := build(pre[1:1+leftLen],mid[0:leftLen])
	right := build(pre[len(pre)-rightLen:len(pre)],mid[index+1:])
	root := &help.TreeNode{Value: value,Left: left,Right: right}
	return root
}

var  myleft *help.TreeNode = nil
var myright *help.TreeNode = nil

type element struct {
	value int
	index int
	leftLen int
	rightLen int
	pre []int
	mid []int
	node *help.TreeNode
}

var count = 0
var lastVisited = make([]int,0)
var P = 1

func isEqual(a []int,b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i,v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func contains(a int,slice []int) bool {
	for _,v := range slice {
		if a == v {
			return true
		}
	}
	return false
}

type nihao struct {
	slice []int
	node *help.TreeNode
}

var nihaoslice = make([]nihao,0)

func get(slice []int) *help.TreeNode {
	for _,v := range nihaoslice {
		if isEqual(slice,v.slice) {
			return v.node
		}
	}
	return nil
}

func set(slice []int,node *help.TreeNode) {
	nihaoslice = append(nihaoslice,nihao{
		slice: slice,
		node:  node,
	})
}

var node *help.TreeNode

func build2(preArr []int,midArr []int) *help.TreeNode {
	stack := &help.MyStack{}
	pre := preArr
	mid := midArr
	for len(pre) > 0 || stack.Length() > 0{
		for len(pre) != 0 {
			value := pre[0]
			index := findIndex(value,mid)
			leftLen := index //0
			rightLen := len(mid)-1-index//0
			stack.Push(element{
				value:    value,
				index:    index,
				leftLen:  leftLen,
				rightLen: rightLen,
				pre:      pre,
				mid:      mid,
				node: nil,
			})
			pre = pre[1:1+leftLen]
			mid = mid[0:leftLen]
		}
		if stack.Length() > 0 {
			temp,_ := stack.Peek()
			ele := temp.(element)
			//fmt.Printf("%#v\n",ele)
			pre = ele.pre[len(ele.pre)-ele.rightLen:len(ele.pre)]
			if len(pre) == 0 || isEqual(pre,lastVisited) {
				temp,_ := stack.Pop()
				ele := temp.(element)

				node = &help.TreeNode{}
				node.Value = ele.value
				node.Left = get(ele.pre[1:1+ele.leftLen])
				node.Right = get(ele.pre[len(ele.pre)-ele.rightLen:len(ele.pre)])
				set(ele.pre,node)

				ele.node = node

				lastVisited = ele.pre
				pre = []int{}
			} else {
				pre = ele.pre[len(ele.pre)-ele.rightLen:len(ele.pre)]
				mid = ele.mid[ele.index+1:]
			}
		}
	}
	return node
}

func findIndex(value int,mid []int) int {
	for index,v := range mid {
		if value == v {
			return index
		}
	}
	return -1
}

func cengcibianli(node *help.TreeNode) {
	queue := help.NewMyQueue()
	queue.Add(node)
	for queue.Length() > 0 {
		temp,_ := queue.Poll()
		node := temp.(*help.TreeNode)
		fmt.Print(node.Value)
		fmt.Print(" ")

		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}
}

func main() {
	node := build2(preArr,midArr)
	isBalanced := isBalanced(node)
	fmt.Println(isBalanced)
	preOrder(node)
	//fmt.Println()
	//midOrder(node)
	//fmt.Println()
	//doubleQueueOrder(node)
	//fmt.Println()
	//cengcibianli(node)
	//fmt.Println()
	//fmt.Println("--------")
	//
	//node = build(preArr,midArr)
	//preOrder(node)
	//fmt.Println()
	//midOrder(node)
	//fmt.Println()
	//doubleQueueOrder(node)
}

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
	if !rightBalanced {
		return rightHeight,rightBalanced
	}

	fmt.Println(leftHeight,rightHeight)
	if help.Abs(leftHeight-rightHeight) > 1 {
		return help.Max(leftHeight,rightHeight),false
	}
	return help.Max(leftHeight,rightHeight),true
}

func preOrder2(node *help.TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value)
	fmt.Print(" ")

	preOrder2(node.Left)
	preOrder2(node.Right)
}

func midOrder2(node *help.TreeNode) {
	if node == nil {
		return
	}
	midOrder2(node.Left)
	fmt.Print(node.Value)
	fmt.Print(" ")
	midOrder2(node.Right)
}

func preOrder(node *help.TreeNode) {
	root := node
	stack := help.NewMyStack()
	for root != nil || stack.Length() > 0 {
		// 入栈的时候，转移为左子节点
		for root != nil {
			fmt.Print(root.Value)
			fmt.Print(" ")
			stack.Push(root)
			root = root.Left
		}
		if stack.Length() > 0 {
			temp,_ := stack.Pop()
			root = temp.(*help.TreeNode)
			//出栈的时候转为右子节点
			root = root.Right
		}
	}
}

func midOrder(node *help.TreeNode) {
	root := node
	stack := help.NewMyStack()
	for root != nil || stack.Length() > 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}

		temp,_ := stack.Pop()
		root = temp.(*help.TreeNode)
		fmt.Print(root.Value)
		fmt.Print(" ")
		root = root.Right

	}
}

func doubleQueueOrder(node *help.TreeNode) {
	root := node
	stackData := help.NewMyStack()
	stackPrint := help.NewMyStack()
	for root != nil || stackData.Length() > 0 {
		for root != nil {
			stackData.Push(root)
			stackPrint.Push(root)
			root = root.Right
		}
		if stackData.Length() > 0 {
			temp,_ := stackData.Pop()
			root = temp.(*help.TreeNode)
			root = root.Left
		}
	}

	for stackPrint.Length() > 0 {
		temp,_ := stackPrint.Pop()
		fmt.Print(temp.(*help.TreeNode).Value)
		fmt.Print(" ")
	}
}