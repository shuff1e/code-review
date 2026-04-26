package main

import (
	"fmt"
)

/*
148. 排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

 */

/*
解答一：归并排序（递归法）
题目要求时间空间复杂度分别为O(nlogn)O(nlogn)和O(1)O(1)，根据时间复杂度我们自然想到二分法，从而联想到归并排序；

对数组做归并排序的空间复杂度为 O(n)O(n)，分别由新开辟数组O(n)O(n)和递归函数调用O(logn)O(logn)组成，而根据链表特性：

数组额外空间：链表可以通过修改引用来更改节点顺序，无需像数组一样开辟额外空间；
递归额外空间：递归调用函数将带来O(logn)O(logn)的空间复杂度，因此若希望达到O(1)O(1)空间复杂度，则不能使用递归。
通过递归实现链表归并排序，有以下两个环节：

分割 cut 环节： 找到当前链表中点，并从中点将链表断开（以便在下次递归 cut 时，链表片段拥有正确边界）；
我们使用 fast,slow 快慢双指针法，奇数个节点找到中点，偶数个节点找到中心左边的节点。
找到中点 slow 后，执行 slow.next = None 将链表切断。
递归分割时，输入当前链表左端点 head 和中心节点 slow 的下一个节点 tmp(因为链表是从 slow 切断的)。
cut 递归终止条件： 当head.next == None时，说明只有一个节点了，直接返回此节点。
合并 merge 环节： 将两个排序链表合并，转化为一个排序链表。
双指针法合并，建立辅助ListNode h 作为头部。
设置两指针 left, right 分别指向两链表头部，比较两指针处节点值大小，由小到大加入合并链表头部，指针交替前进，直至添加完两个链表。
返回辅助ListNode h 作为头部的下个节点 h.next。
时间复杂度 O(l + r)，l, r 分别代表两个链表长度。
当题目输入的 head == None 时，直接返回None。
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next.Next = &ListNode{Val: 3}
	//head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	printList(head)
	head = sortListMerge(head)
	printList(head)
}

func printList(root *ListNode) {
	for root != nil {
		fmt.Print(root.Val,"->")
		root = root.Next
	}
	fmt.Println("nil")
}

func sortListMerge(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// split
	slow,fast := head,head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil

	left := sortListMerge(head)
	right := sortListMerge(next)
	result := &ListNode{
		Val: -1,
	}
	temp := result
	for left != nil && right != nil {
		if left.Val < right.Val {
			result.Next = left

			left = left.Next
			result = result.Next
		} else {
			result.Next = right

			right = right.Next
			result = result.Next
		}
	}
	if left != nil {
		result.Next = left
	}
	if right != nil {
		result.Next = right
	}
	return temp.Next
}


func sortListQuick(head *ListNode) *ListNode {
	result,_ := help(head)
	return result
}

// partition
func partition(head *ListNode) (*ListNode,*ListNode) {
	tempPre := &ListNode{
		Val: -1,
	}
	tempPre.Next = head
	temp := head

	headPre := &ListNode{
		Val: -1,
	}
	headPre.Next = head

	markPre := &ListNode{
		Val: -1,
	}
	markPre.Next = head
	mark := head

	for temp != nil {
		if temp.Val < head.Val {
			mark = mark.Next
			markPre = markPre.Next
			// swap
			a,b := swap(markPre,tempPre)
			markPre = a
			tempPre = b
			mark = markPre.Next
			temp = tempPre.Next
		}
		temp = temp.Next
		tempPre = tempPre.Next
	}
	// swap
	a,b := swap(headPre,markPre)
	headPre = a
	markPre = b
	head = headPre.Next
	mark = markPre.Next

	return headPre.Next,mark
}

// 链表中，已知前驱，交换后面的节点
func swap(a,b *ListNode) (*ListNode,*ListNode){
	if a == b {
		return a,b
	}
	if a.Next == b {
		bNextNext := b.Next.Next
		a.Next = b.Next
		a.Next.Next = b
		a.Next.Next.Next = bNextNext
		return a,a.Next
	} else {
		aNext := a.Next
		bNext := b.Next
		aNextNext := a.Next.Next
		bNextNext := b.Next.Next

		a.Next = bNext
		a.Next.Next = aNextNext
		b.Next = aNext
		b.Next.Next = bNextNext
		return a,b
	}
}

func help(head *ListNode) (*ListNode,*ListNode) {
	if head == nil || head.Next == nil {
		return head,head
	}

	head,tail := partition(head)

	next := tail.Next
	tail.Next = nil

	head1,tail1 := help(head)
	head2,tail2 := help(next)
	tail1.Next = head2
	return head1,tail2
}

/*

循环的归并排序

方法二：自底向上归并排序
使用自底向上的方法实现归并排序，则可以达到 O(1) 的空间复杂度。

首先求得链表的长度 length，然后将链表拆分成子链表进行合并。

具体做法如下。

用 subLength 表示每次需要排序的子链表的长度，初始时 subLength=1。

每次将链表拆分成若干个长度为 subLength 的子链表（最后一个子链表的长度可以小于 subLength），
按照每两个子链表一组进行合并，合并后即可得到若干个长度为 subLength×2 的有序子链表（最后一个子链表的长度可以小于 subLength×2）。
合并两个子链表仍然使用「21. 合并两个有序链表」的做法。

将 subLength 的值加倍，重复第 2 步，对更长的有序子链表进行合并操作，直到有序子链表的长度大于或等于 length，整个链表排序完毕。

class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null) {
            return head;
        }
        int length = 0;
        ListNode node = head;
        while (node != null) {
            length++;
            node = node.next;
        }
        ListNode dummyHead = new ListNode(0, head);
        for (int subLength = 1; subLength < length; subLength <<= 1) {
            ListNode prev = dummyHead, curr = dummyHead.next;
            while (curr != null) {
                ListNode head1 = curr;
                for (int i = 1; i < subLength && curr.next != null; i++) {
                    curr = curr.next;
                }
                ListNode head2 = curr.next;
                curr.next = null;
                curr = head2;
                for (int i = 1; i < subLength && curr != null && curr.next != null; i++) {
                    curr = curr.next;
                }
                ListNode next = null;
                if (curr != null) {
                    next = curr.next;
                    curr.next = null;
                }
                ListNode merged = merge(head1, head2);
                prev.next = merged;
                while (prev.next != null) {
                    prev = prev.next;
                }
                curr = next;
            }
        }
        return dummyHead.next;
    }

    public ListNode merge(ListNode head1, ListNode head2) {
        ListNode dummyHead = new ListNode(0);
        ListNode temp = dummyHead, temp1 = head1, temp2 = head2;
        while (temp1 != null && temp2 != null) {
            if (temp1.val <= temp2.val) {
                temp.next = temp1;
                temp1 = temp1.next;
            } else {
                temp.next = temp2;
                temp2 = temp2.next;
            }
            temp = temp.next;
        }
        if (temp1 != null) {
            temp.next = temp1;
        } else if (temp2 != null) {
            temp.next = temp2;
        }
        return dummyHead.next;
    }
}

 */