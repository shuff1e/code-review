package main

import "fmt"

// 23：链表中环的入口结点
// 题目：一个链表中包含环，如何找出环的入口结点？例如，在图3.8的链表中，
// 环的入口结点是结点3。
//
//		 |--------|
//       |		  |
// 1->2->3->4->5->6

// A：快慢指针。
// 快指针一次走两步，慢指针一次走一步，如果有环，肯定会碰撞，两者第一次碰撞的位置为？
// 假设无环部分长度为k，慢指针走k步时，快指针走了2k步，快指针领先慢指针k步，或者慢指针领先快指针loop_size-k步
// 然后快指针追赶慢指针，每次都会追赶一步，这样快指针经过loop_size-k步之后，会和慢指针碰撞
// 因此第一次碰撞的位置距离环的入口处，慢指针已经走了loop_size-k步，距离环的入口处k步
// 快指针回到链表头部，然后快指针和慢指针都一次走一步，则第二次碰撞的位置为环的入口处


// 1.快指针一次走两步，慢指针一次走一步，两者第一次碰撞
// 快指针 1，3，5，3，5，3，5，5，5
// 慢指针 1，2，3，4，5，6，3，4，5

// 2.快指针回到头部，一次走一步，慢指针继续一次走一步
// 慢指针 5，6，3
// 快指针 1，2，3
// 这样第二次碰撞的位置就是入口处

type Node struct {
	value int
	next *Node
}

func getEntryNodeInListLoop(root *Node) *Node {
	if root == nil || root.next == nil {
		return nil
	}
	fast := root
	slow := root
	// do while
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			break
		}
	}
	// 如果是环状，则第一次碰撞的位置为root
	// 1,3,5,1,3,5,1
	// 1,2,3,4,5,6,1
	if fast == nil || fast != slow {
		return nil
	}
	// first collision
	fast = root
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	// second collision
	return fast
}

func main() {
	Test1()
	Test2()
	Test3()
	Test4()
	Test5()
	Test6()
	Test7()
}

func Test1() {
	root := &Node{value: 1}
	fmt.Println(getEntryNodeInListLoop(root) == nil)
}

func Test2() {
	root := &Node{value: 1}
	root.next = root
	fmt.Println(getEntryNodeInListLoop(root).value == 1)
}

func Test3() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	root.next.next.next.next.next = root.next.next
	fmt.Println(getEntryNodeInListLoop(root).value == 3)
}

func Test4() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	root.next.next.next.next.next = root
	fmt.Println(getEntryNodeInListLoop(root).value == 1)
}

func Test5() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	root.next.next.next.next.next = root.next.next.next.next
	fmt.Println(getEntryNodeInListLoop(root).value == 5)
}

func Test6() {
	root := &Node{value: 1}
	root.next = &Node{value: 2}
	root.next.next = &Node{value: 3}
	root.next.next.next = &Node{value: 4}
	root.next.next.next.next = &Node{value: 5}
	fmt.Println(getEntryNodeInListLoop(root) == nil)
}

func Test7() {
	root := (*Node)(nil)
	fmt.Println(getEntryNodeInListLoop(root) == nil)
}