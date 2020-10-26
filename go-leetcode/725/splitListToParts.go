package main

import "fmt"

/*

725. 分隔链表
给定一个头结点为 root 的链表, 编写一个函数以将链表分隔为 k 个连续的部分。

每部分的长度应该尽可能的相等: 任意两部分的长度差距不能超过 1，也就是说可能有些部分为 null。

这k个部分应该按照在链表中出现的顺序进行输出，并且排在前面的部分的长度应该大于或等于后面的长度。

返回一个符合上述规则的链表的列表。

举例： 1->2->3->4, k = 5 // 5 结果 [ [1], [2], [3], [4], null ]

示例 1：

输入:
root = [1, 2, 3], k = 5
输出: [[1],[2],[3],[],[]]
解释:
输入输出各部分都应该是链表，而不是数组。
例如, 输入的结点 root 的 val= 1, root.next.val = 2, root.next.next.val = 3, 且 root.next.next.next = null。
第一个输出 output[0] 是 output[0].val = 1, output[0].next = null。
最后一个元素 output[4] 为 null, 它代表了最后一个部分为空链表。
示例 2：

输入:
root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
输出: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
解释:
输入被分成了几个连续的部分，并且每部分的长度相差不超过1.前面部分的长度大于等于后面部分的长度。


提示:

root 的长度范围： [0, 1000].
输入的每个节点的大小范围：[0, 999].
k 的取值范围： [1, 50].

 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	root := &ListNode{Val: 1}
	//root.Next = &ListNode{Val: 2}
	//root.Next.Next = &ListNode{Val: 3}
	//root.Next.Next.Next = &ListNode{Val: 4}
	//root.Next.Next.Next.Next = &ListNode{Val: 5}
	//root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	//root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	//root.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	//root.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	//root.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 10}
	//root.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 11}

	result := splitListToParts(root,3)
	fmt.Println(len(result))
	for _,v := range result {
		for v != nil {
			fmt.Print(v.Val,"->")
			v = v.Next
		}
		fmt.Println()
	}
}

// 1 2 3 4 5 6 7 8 9 10
//
func splitListToParts(root *ListNode, k int) []*ListNode {
	length := getLength(root)
	if length == 0 {
		return make([]*ListNode,k)
	}
	n := length/k
	rest := 0
	if n == 0 {
		n = 1
	} else {
		rest = length % k
	}

	result := make([]*ListNode,0)
	split(root,n,rest,&result)

	tempLength := len(result)
	if tempLength < k {
		index := 0
		for index < k - tempLength {
			result = append(result,(*ListNode)(nil))
			index ++
		}
	}
	return result
}

func split(root *ListNode,k int,rest int,result *[]*ListNode) {
	if root == nil {
		return
	}
	// 1 2
	// 1 2
	// 1 2
	temp := root
	index := 1

	tempLength := k
	if rest > 0 {
		tempLength = k + 1
		rest --
	}

	for temp != nil && index < tempLength {
		temp = temp.Next
		index ++
	}

	if index < k {
		*result = append(*result,root)
		return
	}
	if temp == nil {
		*result = append(*result,root)
		return
	}

	next := temp.Next
	temp.Next = nil

	*result = append(*result,root)

	split(next,k,rest,result)
}

func getLength(root *ListNode) int {
	index := 0
	for root != nil {
		index ++
		root = root.Next
	}
	return index
}

/*

class Solution {
    public ListNode[] splitListToParts(ListNode root, int k) {
        ListNode cur = root;
        int N = 0;
        while (cur != null) {
            cur = cur.next;
            N++;
        }

        int width = N / k, rem = N % k;

        ListNode[] ans = new ListNode[k];

        cur = root;
        for (int i = 0; i < k; ++i) {
            ListNode head = new ListNode(0), write = head;
            for (int j = 0; j < width + (i < rem ? 1 : 0); ++j) {
                write = write.next = new ListNode(cur.val);
                if (cur != null) cur = cur.next;
            }
            ans[i] = head.next;
        }
        return ans;
    }
}


class Solution {
    public ListNode[] splitListToParts(ListNode root, int k) {
        ListNode cur = root;
        int N = 0;
        while (cur != null) {
            cur = cur.next;
            N++;
        }

        int width = N / k, rem = N % k;

        ListNode[] ans = new ListNode[k];
        cur = root;
        for (int i = 0; i < k; ++i) {
            ListNode head = cur;
            for (int j = 0; j < width + (i < rem ? 1 : 0) - 1; ++j) {
                if (cur != null) cur = cur.next;
            }
            if (cur != null) {
                ListNode prev = cur;
                cur = cur.next;
                prev.next = null;
            }
            ans[i] = head;
        }
        return ans;
    }
}

 */