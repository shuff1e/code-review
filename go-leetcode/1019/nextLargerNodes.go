package main

import "fmt"

/*

1019. 链表中的下一个更大节点
给出一个以头节点 head 作为第一个节点的链表。链表中的节点分别编号为：node_1, node_2, node_3, ... 。

每个节点都可能有下一个更大值（next larger value）：对于 node_i，如果其 next_larger(node_i) 是 node_j.val，那么就有 j > i 且  node_j.val > node_i.val，而 j 是可能的选项中最小的那个。如果不存在这样的 j，那么下一个更大值为 0 。

返回整数答案数组 answer，其中 answer[i] = next_larger(node_{i+1}) 。

注意：在下面的示例中，诸如 [2,1,5] 这样的输入（不是输出）是链表的序列化表示，其头节点的值为 2，第二个节点值为 1，第三个节点值为 5 。



示例 1：

输入：[2,1,5]
输出：[5,5,0]
示例 2：

输入：[2,7,4,3,5]
输出：[7,0,5,5,0]
示例 3：

输入：[1,7,5,1,9,2,5,1]
输出：[7,9,9,9,0,5,0,0]


提示：

对于链表中的每个节点，1 <= node.val <= 10^9
给定列表的长度在 [0, 10000] 范围内

 */

func main() {
	arr := []int{1,7,5,1,9,2,5,1}
	head := generateList(arr)
	fmt.Println(head.Val,head.Next.Val)
	fmt.Println(nextLargerNodes(head))
}

func generateList(array []int) *ListNode {
	head := &ListNode{}
	temp := head
	for _,v := range array {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	return head.Next
}

type ListNode struct {
	Val int
	Next *ListNode
}

// 单调栈

// [1,7,5,1,9,2,5,1]


func nextLargerNodes(head *ListNode) []int {
	list := []int{}
	for head != nil {
		list = append(list,head.Val)
		head = head.Next
	}

	result := make([]int,len(list))

	stack := []int{}
	for i := 0;i<len(list);i++ {
		for len(stack) > 0 && list[stack[len(stack)-1]] < list[i] {
			tempIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[tempIndex] = list[i]
		}
		stack = append(stack,i)
	}

	return result
}

// [1,7,5,1,9,2,5,1]

// 0
// 1

// 1 2 3
// 7 7 5 1

// 4
// 7 9 9 9 9

// 7
//
// 7 9 9 9 9 5 5 1

func nextLargerNodes2(head *ListNode) []int {
	list := []int{}
	stack := []int{}
	for ;head != nil;head = head.Next {
		for len(stack) > 0 && list[stack[len(stack)-1]] < head.Val {
			tempIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			list[tempIndex] = head.Val
		}
		stack = append(stack,len(list))
		list = append(list,head.Val)
	}

	for i := 0;i<len(stack);i++ {
		list[stack[i]] = 0
	}
	return list
}

// 归并求逆序对，标记每次归并最大的那一对
