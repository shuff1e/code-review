package main

import "fmt"

/*
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	node := &ListNode{Val: 1}
	//node.Next = &ListNode{Val: 2}
	//node.Next.Next = &ListNode{Val: 4}
	//node.Next.Next.Next = &ListNode{Val: 3}
	//node.Next.Next.Next.Next = &ListNode{Val: 2}
	//node.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	fmt.Println(isPalindrome(node))
}

//1 2 3 4 5 6 7
//slow 1 2 3 4
//fast 1 3 5 7
//
//1 2 3 4 5 6
//slow 1 2 3
//fast 1 3 5
//
//1 2 3 4 5
//slow 1 2 3
//fast 1 3 5
//
//1 2 3 4
//slow 1 2
//fast 1 3
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow,fast := head,head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	otherHalf := slow.Next
	slow.Next = nil
	otherHalf = reverseList(otherHalf)

	result := true
	fast = head
	temp := otherHalf
	for temp != nil {
		if temp.Val != fast.Val {
			result = false
			break
		}
		temp = temp.Next
		fast = fast.Next
	}
	slow.Next = reverseList(otherHalf)

	return result
}

func reverseList(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

/*

class Solution {
    public boolean isPalindrome(ListNode head) {
        List<Integer> vals = new ArrayList<>();

        // Convert LinkedList into ArrayList.
        ListNode currentNode = head;
        while (currentNode != null) {
            vals.add(currentNode.val);
            currentNode = currentNode.next;
        }

        // Use two-pointer technique to check for palindrome.
        int front = 0;
        int back = vals.size() - 1;
        while (front < back) {
            // Note that we must use ! .equals instead of !=
            // because we are comparing Integer, not int.
            if (!vals.get(front).equals(vals.get(back))) {
                return false;
            }
            front++;
            back--;
        }
        return true;
    }
}

因为在许多语言中，堆栈帧很大（如 Python），
并且最大的运行时堆栈深度为 1000（可以增加，但是有可能导致底层解释程序内存出错）。
为每个节点创建堆栈帧极大的限制了算法能够处理的最大链表大小。


我们可以将链表的后半部分反转（修改链表结构），然后将前半部分和后半部分进行比较。
比较完成后我们应该将链表恢复原样。虽然不需要恢复也能通过测试用例，因为使用该函数的人不希望链表结构被更改。


找到前半部分链表的尾节点。
反转后半部分链表。
判断是否为回文。
恢复链表。
返回结果。
 */