package main

import (
	"fmt"
	"github.com/shuff1e/code-review/DataStructure/BinaryTree"
	"github.com/shuff1e/code-review/Greddy"
)

func main() {
	arr := []int{12,2,7,13,14,85}
	node := Greddy.HuffmanCoding(arr)
	BinaryTree.PreOrder(node)
	fmt.Println()
	BinaryTree.MidOrder(node)
}
