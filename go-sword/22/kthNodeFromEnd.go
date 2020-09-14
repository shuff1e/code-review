package main

import "fmt"

// 22：链表中倒数第k个结点
// 题目：输入一个链表，输出该链表中倒数第k个结点。为了符合大多数人的习惯，
// 本题从1开始计数，即链表的尾结点是倒数第1个结点。例如一个链表有6个结点，
// 从头结点开始它们的值依次是1、2、3、4、5、6。这个链表的倒数第3个结点是
// 值为4的结点。

// A：倒数第3个，就是正数第4个，一共6个节点的话
// 倒数第k个，就是正数n-k+1

// 一次遍历的话，使用快慢指针
// 快指针先走k步，到达正数第k个的位置，然后慢指针和快指针同时走n-k步，快指针到达倒数第一个节点的位置
// 则慢指针所在的位置为n-(n-k)=k

type Node struct {
	value int
	next *Node
}
func getLength(root *Node) int {
	count := 0
	for root != nil {
		count ++
		root = root.next
	}
	return count
}

func getKthNode(root *Node,k int) *Node {
	length := getLength(root)
	if length - k +1 <=0 {
		return nil
	}
	for i:=1;i<length-k+1;i++ {
		root = root.next
	}
	return root
}

// 1 2 3 4 5 6 k=3
// fast 先到3
func getKthNode2(root *Node,k int) *Node {
	if root == nil || k == 0{
		return nil
	}
	fast := root
	for i := 1;i<k;i++ {
		if fast.next == nil {
			return nil
		}
		fast = fast.next
	}
	slow := root
	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
}

func Test1() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(root,2).value == 4)
}

func Test2() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(root,1).value == 5)
}

func Test3() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(root,5).value == 1)
}

func Test4() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(nil,100))
}

func Test5() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(root,6))
}

func Test6() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getKthNode2(root,0))
}
