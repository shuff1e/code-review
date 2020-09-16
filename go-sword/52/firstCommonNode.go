package main

import "fmt"

// 52：两个链表的第一个公共结点
// 题目：输入两个链表，找出它们的第一个公共结点。

// A：因为是单链表，只有next指针
// 因此共同节点之后的部分，不会分叉，因此是Y形

// 遍历链表，用栈保存节点
// 然后后进先出，第一个p1.next != p2.next
// 这样需要O(m+n)的空间复杂度，时间复杂度也是遍历两次链表


// 也可以第一次遍历链表，得到长度
// 然后m>n，先走m-n步
// 然后两者会一同走到共同的节点























type Node struct {
	value int
	next *Node
}

func getCommon(node1,node2 *Node) *Node {
	if node1 == nil || node2 == nil {
		return nil
	}
	m := getLength(node1)
	n := getLength(node2)
	if m < n {
		temp := node1
		node1 = node2
		node2 = temp

		tempNum := m
		m = n
		n = tempNum
	}

	for i:=0;i<m-n;i++ {
		node1 = node1.next
	}

	for node1 != node2 {
		node1 = node1.next
		node2 = node2.next
	}

	return node1
}

func getLength(node *Node) int {
	count := 0
	for node != nil {
		node = node.next
		count ++
	}
	return count
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}

func Test(name string,node1,node2,expected *Node) {
	fmt.Println(name)
	temp := getCommon(node1,node2)
	if temp != expected {
		panic("fuck")
	}
}

// 第一个公共结点在链表中间
// 1 - 2 - 3 \
//            6 - 7
//     4 - 5 /
func Test1() {
	pNode1 := CreateListNode(1);
	pNode2 := CreateListNode(2);
	pNode3 := CreateListNode(3);
	pNode4 := CreateListNode(4);
	pNode5 := CreateListNode(5);
	pNode6 := CreateListNode(6);
	pNode7 := CreateListNode(7);

	ConnectListNodes(pNode1, pNode2);
	ConnectListNodes(pNode2, pNode3);
	ConnectListNodes(pNode3, pNode6);
	ConnectListNodes(pNode4, pNode5);
	ConnectListNodes(pNode5, pNode6);
	ConnectListNodes(pNode6, pNode7);

	Test("Test1", pNode1, pNode4, pNode6);
}

// 没有公共结点
// 1 - 2 - 3 - 4
//
// 5 - 6 - 7
func Test2() {
	pNode1 := CreateListNode(1);
	pNode2 := CreateListNode(2);
	pNode3 := CreateListNode(3);
	pNode4 := CreateListNode(4);
	pNode5 := CreateListNode(5);
	pNode6 := CreateListNode(6);
	pNode7 := CreateListNode(7);

	ConnectListNodes(pNode1, pNode2);
	ConnectListNodes(pNode2, pNode3);
	ConnectListNodes(pNode3, pNode4);
	ConnectListNodes(pNode5, pNode6);
	ConnectListNodes(pNode6, pNode7);

	Test("Test2", pNode1, pNode5, nil);
}

// 公共结点是最后一个结点
// 1 - 2 - 3 - 4 \
//                7
//         5 - 6 /
func Test3() {
	pNode1 := CreateListNode(1);
	pNode2 := CreateListNode(2);
	pNode3 := CreateListNode(3);
	pNode4 := CreateListNode(4);
	pNode5 := CreateListNode(5);
	pNode6 := CreateListNode(6);
	pNode7 := CreateListNode(7);

	ConnectListNodes(pNode1, pNode2);
	ConnectListNodes(pNode2, pNode3);
	ConnectListNodes(pNode3, pNode4);
	ConnectListNodes(pNode4, pNode7);
	ConnectListNodes(pNode5, pNode6);
	ConnectListNodes(pNode6, pNode7);

	Test("Test3", pNode1, pNode5, pNode7);
}

// 公共结点是第一个结点
// 1 - 2 - 3 - 4 - 5
// 两个链表完全重合
func Test4() {
	pNode1 := CreateListNode(1);
	pNode2 := CreateListNode(2);
	pNode3 := CreateListNode(3);
	pNode4 := CreateListNode(4);
	pNode5 := CreateListNode(5);

	ConnectListNodes(pNode1, pNode2);
	ConnectListNodes(pNode2, pNode3);
	ConnectListNodes(pNode3, pNode4);
	ConnectListNodes(pNode4, pNode5);

	Test("Test4", pNode1, pNode1, pNode1);
}

// 输入的两个链表有一个空链表
func Test5() {
	pNode1 := CreateListNode(1);
	pNode2 := CreateListNode(2);
	pNode3 := CreateListNode(3);
	pNode4 := CreateListNode(4);
	pNode5 := CreateListNode(5);

	ConnectListNodes(pNode1, pNode2);
	ConnectListNodes(pNode2, pNode3);
	ConnectListNodes(pNode3, pNode4);
	ConnectListNodes(pNode4, pNode5);

	Test("Test5", nil, pNode1, nil);
}

func Test6() {
	Test("Test6", nil, nil, nil);
}

func CreateListNode(v int) *Node {
	return &Node{
		value: v,
	}
}

func ConnectListNodes(node1,node2 *Node) {
	node1.next = node2
}