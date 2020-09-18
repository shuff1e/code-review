package main

import "algorithm/common/help"

func reverse(head *help.Node) *help.Node {
	var next *help.Node = nil
	var prev *help.Node = nil

	for head != nil {
		// 闹住next
		// 然后修改next
		next = head.Next
		head.Next = prev
		prev= head
		head = next
	}
	return prev
}

func main() {
	node := help.GenerateNList(10)
	help.PrintList(node)

	result := reverse(node)
	help.PrintList(result)
}

func reverseDoubleNode(head *help.DoubleNode) *help.DoubleNode {
	var next *help.DoubleNode = nil
	var prev *help.DoubleNode = nil
	for head != nil {
		next = head.Next
		head.Next = prev
		head.Prev = next
		prev = head
		head = next
	}
	return prev
}
