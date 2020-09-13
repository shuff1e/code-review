package main

import (
	"bufio"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"io"
	"os"
	"strconv"
)

// 37：序列化二叉树
// 题目：请实现两个函数，分别用来序列化和反序列化二叉树。

type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}

func serialize(root *TreeNode,fileName string) {
	f,err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	serializeHelper(root,w)
	// do not forget
	w.Flush()
}

func serializeHelper(root *TreeNode, w *bufio.Writer) {
	if root == nil {
		_,err := w.WriteString("$,")
		if err != nil {
			panic(err)
		}
		return
	}
	_,err := w.WriteString(strconv.Itoa(root.value) + ",")
	if err != nil {
		panic(err)
	}
	serializeHelper(root.left,w)
	serializeHelper(root.right,w)
}

func deSerialize(fileName string) *TreeNode {
	f,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	return deSerializeHelper(r)
}

func deSerializeHelper(r *bufio.Reader) *TreeNode {
	str,err := r.ReadString(',')
	if err != nil && err != io.EOF {
		panic(err)
	}
	if str == "$," {
		return nil
	}
	value,err := strconv.Atoi(str[0:len(str)-1])
	if err != nil {
		panic(err)
	}
	node := &TreeNode{
		value: value,
	}
	node.left = deSerializeHelper(r)
	node.right = deSerializeHelper(r)
	return node
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}

//            8
//        6      10
//       5 7    9  11
func Test1() {
	pNode8 := CreateBinaryTreeNode(8);
	pNode6 := CreateBinaryTreeNode(6);
	pNode10 := CreateBinaryTreeNode(10);
	pNode5 := CreateBinaryTreeNode(5);
	pNode7 := CreateBinaryTreeNode(7);
	pNode9 := CreateBinaryTreeNode(9);
	pNode11 := CreateBinaryTreeNode(11);

	ConnectTreeNodes(pNode8, pNode6, pNode10);
	ConnectTreeNodes(pNode6, pNode5, pNode7);
	ConnectTreeNodes(pNode10, pNode9, pNode11);

	Test("Test1", pNode8);
}

//            5
//          4
//        3
//      2
func Test2() {
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);
	pNode3 := CreateBinaryTreeNode(3);
	pNode2 := CreateBinaryTreeNode(2);

	ConnectTreeNodes(pNode5, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode3, nil);
	ConnectTreeNodes(pNode3, pNode2, nil);

	Test("Test2", pNode5);
}

//        5
//         4
//          3
//           2
func Test3() {
	pNode5 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(4);
	pNode3 := CreateBinaryTreeNode(3);
	pNode2 := CreateBinaryTreeNode(2);

	ConnectTreeNodes(pNode5, nil, pNode4);
	ConnectTreeNodes(pNode4, nil, pNode3);
	ConnectTreeNodes(pNode3, nil, pNode2);

	Test("Test3", pNode5);
}

func Test4() {
	pNode5 := CreateBinaryTreeNode(5);
	Test("Test4", pNode5);
}

func Test5() {
	Test("Test5", nil);
}

//        5
//         5
//          5
//         5
//        5
//       5 5
//      5   5
func Test6() {
	pNode1 := CreateBinaryTreeNode(5);
	pNode2 := CreateBinaryTreeNode(5);
	pNode3 := CreateBinaryTreeNode(5);
	pNode4 := CreateBinaryTreeNode(5);
	pNode5 := CreateBinaryTreeNode(5);
	pNode61 := CreateBinaryTreeNode(5);
	pNode62 := CreateBinaryTreeNode(5);
	pNode71 := CreateBinaryTreeNode(5);
	pNode72 := CreateBinaryTreeNode(5);

	ConnectTreeNodes(pNode1, nil, pNode2);
	ConnectTreeNodes(pNode2, nil, pNode3);
	ConnectTreeNodes(pNode3, pNode4, nil);
	ConnectTreeNodes(pNode4, pNode5, nil);
	ConnectTreeNodes(pNode5, pNode61, pNode62);
	ConnectTreeNodes(pNode61, pNode71, nil);
	ConnectTreeNodes(pNode62, nil, pNode72);

	Test("Test6", pNode1);
}

func Test(name string,root *TreeNode) {
	fmt.Println(name)
	printTreeInLine(root)
	fileName := "a.txt"
	serialize(root,fileName)
	node := deSerialize(fileName)
	printTreeInLine(node)
	if !isSameTree(root,node) {
		panic("fuck")
	}
}

func CreateBinaryTreeNode(x int) *TreeNode {
	return &TreeNode{value: x}
}

func ConnectTreeNodes(p,l,r *TreeNode) {
	p.left = l
	p.right = r
}

func printTreeInLine(root *TreeNode) {
	if root == nil {
		return
	}
	queue1 := arraylist.New()
	queue2 := arraylist.New()
	fmt.Println(root.value," ")
	queue1.Add(root)
	for !queue1.Empty() {
		for !queue1.Empty() {
			temp,_ := queue1.Get(0)
			node := temp.(*TreeNode)
			queue1.Remove(0)

			if node.left != nil {
				fmt.Print(node.left.value," ")
				queue2.Add(node.left)
			}
			if node.right != nil {
				fmt.Print(node.right.value," ")
				queue2.Add(node.right)
			}
		}
		fmt.Println()
		temp := queue1
		queue1 = queue2
		queue2 = temp
	}
}

func isSameTree(node1,node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil {
		if node1.value != node2.value {
			return false
		}
		return isSameTree(node1.left,node2.left) &&
			isSameTree(node1.right,node2.right)
	}
	return false
}