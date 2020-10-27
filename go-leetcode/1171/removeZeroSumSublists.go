package main

/*

1171. 从链表中删去总和值为零的连续节点
给你一个链表的头节点 head，请你编写代码，反复删去链表中由 总和 值为 0 的连续节点组成的序列，直到不存在这样的序列为止。

删除完毕后，请你返回最终结果链表的头节点。

你可以返回任何满足题目要求的答案。

（注意，下面示例中的所有序列，都是对 ListNode 对象序列化的表示。）

示例 1：

输入：head = [1,2,-3,3,1]
输出：[3,1]
提示：答案 [1,2,1] 也是正确的。
示例 2：

输入：head = [1,2,3,-3,4]
输出：[1,2,4]
示例 3：

输入：head = [1,2,3,-3,-2]
输出：[1]


提示：

给你的链表中可能有 1 到 1000 个节点。
对于链表中的每个节点，节点的值：-1000 <= node.val <= 1000.

 */

func main() {
	arr := []int{1,2,3,-3,4}
	//arr = []int{1,2,3,-3,-2}
	head := generateList(arr)
	removeZeroSumSublists(head)
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

func removeZeroSumSublists(head *ListNode) *ListNode {
	dict := map[int]*ListNode{}

	dummy := &ListNode{Val: 0}
	dummy.Next = head

	sum := 0
	temp := dummy
	for ;temp != nil; temp = temp.Next {
		sum += temp.Val
		dict[sum] = temp
	}

	sum = 0
	temp = dummy
	for ;temp != nil; temp = temp.Next {
		sum += temp.Val
		temp.Next = dict[sum].Next
	}

	return dummy.Next
}

/*
func removeZeroSumSublists(head *ListNode) *ListNode {
	dict := map[int]int{}

	dict[-1] = 0

	// start -> end
	dict2 := map[int]int{}

	sum := 0

	index := -1
	for ;head != nil;head = head.Next {
		index ++
		sum += head.Val
		if i,ok := dict[sum - 0];ok {
			dict2[i] = index
		}

		// 记录该和第一次出现的位置
		if _,ok := dict[sum];!ok {
			dict[sum] = index
		}
	}
	fmt.Printf("%#v\n",dict2)
	return nil
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

 */
