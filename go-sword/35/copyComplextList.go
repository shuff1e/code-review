package main

import (
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
)

// 35：复杂链表的复制
// 题目：请实现函数ComplexListNode* Clone(ComplexListNode* pHead)，复
// 制一个复杂链表。在复杂链表中，每个结点除了有一个m_pNext指针指向下一个
// 结点外，还有一个m_pSibling 指向链表中的任意结点或者nullptr。

// A：下图是一个含有5个节点的复杂链表
// 第一步就是复制原始链表上的每个节点，并用next指针链接🔗起来
// 假设原始链表中的某个节点N的sibling指向节点S，由于S在链表中可能在N的前面，
// 也可能在N的后面，所以要定位S的位置需要从原始链表的头节点开始找
// 总体的时间复杂度为O(n^2)

//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//  |       |      /|\             /|\
//  --------+--------               |
//          -------------------------

// 由于上述方法的时间主要花费在定位节点的sibling上，
// 因此可以将原始链表节点N的复制节点N'，将N和N'的配对信息放到一个哈希表中
// 这样可以在O(1)时间内根据S找到S'

// 主要的时间花费在定位节点的sibling上
// 第一步，根据每个节点N创建对应的N'，但是把N'链接在N的后面
// 第二步，设置复制出来的节点的sibling，N'的sibling就是N的sibling的next节点
// 第三步，将长链表拆分成2个链表，把奇数位置的节点用next链接起来就是原始链表
// 将偶数位置的节点用next链接起来就是复制出来的链表

func copyListInPlace(root *Node) *Node {
	if root == nil {
		return nil
	}
	node := root
	//result := &Node{value: -1}
	//temp := result
	for node != nil {
		next := node.next
		node.next = &Node{value: node.value}
		node.next.next = next
		node = next
	}
	node = root
	for node != nil {
		if node.sibling != nil {
			node.next.sibling = node.sibling.next
		}
		node = node.next.next
	}

	result1 := &Node{value: -1}
	result2 := &Node{value: -1}
	temp1 := result1
	temp2 := result2

	node = root
	for node != nil {
		// 1 ->1 ->2 ->2
		// 处理完1 ->1，将问题递归到2 ->2
		temp1.next = node
		next1 := node.next
		next2 := node.next.next
		temp2.next = next1

		node = next2
		temp1 = temp1.next
		temp2 = temp2.next
	}
	return result2.next
}

type Node struct {
	value int
	next *Node
	sibling *Node
}

func copyListWithMap(root *Node) *Node {
	if root == nil {
		return nil
	}
	mmp := hashmap.New()
	result := &Node{
		value: root.value,
	}
	mmp.Put(root,result)

	node := root.next
	temp := result
	for node != nil {
		temp.next = &Node{
			value: node.value,
		}
		temp = temp.next
		mmp.Put(node,temp)

		node = node.next
	}

	node = root
	temp = result
	for node != nil {
		if node.sibling != nil {
			tt,_ := mmp.Get(node.sibling)
			temp.sibling = tt.(*Node)
		}
		node = node.next
		temp = temp.next
	}

	return result
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
}

//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//  |       |      /|\             /|\
//  --------+--------               |
//          -------------------------
func Test1() {
	pNode1 := CreateNode(1);
	pNode2 := CreateNode(2);
	pNode3 := CreateNode(3);
	pNode4 := CreateNode(4);
	pNode5 := CreateNode(5);

	BuildNodes(pNode1, pNode2, pNode3);
	BuildNodes(pNode2, pNode3, pNode5);
	BuildNodes(pNode3, pNode4, nil);
	BuildNodes(pNode4, pNode5, pNode2);

	Test("Test1", pNode1);
}

// m_pSibling指向结点自身
//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//         |       | /|\           /|\
//         |       | --             |
//         |------------------------|
func Test2() {
	pNode1 := CreateNode(1);
	pNode2 := CreateNode(2);
	pNode3 := CreateNode(3);
	pNode4 := CreateNode(4);
	pNode5 := CreateNode(5);

	BuildNodes(pNode1, pNode2, nil);
	BuildNodes(pNode2, pNode3, pNode5);
	BuildNodes(pNode3, pNode4, pNode3);
	BuildNodes(pNode4, pNode5, pNode2);
	Test("Test2",pNode1)
}

// m_pSibling形成环
//          -----------------
//         \|/              |
//  1-------2-------3-------4-------5
//          |              /|\
//          |               |
//          |---------------|
func Test3() {
	pNode1 := CreateNode(1);
	pNode2 := CreateNode(2);
	pNode3 := CreateNode(3);
	pNode4 := CreateNode(4);
	pNode5 := CreateNode(5);

	BuildNodes(pNode1, pNode2, nil);
	BuildNodes(pNode2, pNode3, pNode4);
	BuildNodes(pNode3, pNode4, nil);
	BuildNodes(pNode4, pNode5, pNode2);

	Test("Test3", pNode1);
}

// 只有一个结点
func Test4() {
	pNode1 := CreateNode(1);
	BuildNodes(pNode1, nil, pNode1);

	Test("Test4", pNode1);
}

// 鲁棒性测试
func Test5() {
	Test("Test5", nil);
}

func Test(name string,node *Node) {
	fmt.Println(name)
	printList(node)
	//temp := copyListWithMap(node)
	temp := copyListInPlace(node)
	printList(temp)
}

func CreateNode(v int) *Node {
	return &Node{
		value: v,
	}
}

func BuildNodes(node1,next,sibling *Node) {
	if node1 != nil {
		node1.next = next
		node1.sibling = sibling
	}
}

func printList(node *Node) {
	for node != nil {
		if node.sibling != nil {
			fmt.Print(node.value," sibling ",node.sibling.value," ->")
		} else {
			fmt.Print(node.value," ->")
		}
		node = node.next
	}
	fmt.Println()
}