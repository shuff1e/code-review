package main

import "fmt"

/*
430. 扁平化多级双向链表
多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。

给你位于列表第一级的头节点，请你扁平化列表，使所有结点出现在单级双链表中。



示例 1：

输入：head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
输出：[1,2,3,7,8,11,12,9,10,4,5,6]
解释：

输入的多级列表如下图所示：



扁平化后的链表如下图：


示例 2：

输入：head = [1,2,null,3]
输出：[1,3,2]
解释：

输入的多级列表如下图所示：

1---2---NULL
|
3---NULL
示例 3：

输入：head = []
输出：[]


如何表示测试用例中的多级链表？

以 示例 1 为例：

1---2---3---4---5---6--NULL
|
7---8---9---10--NULL
|
11--12--NULL
序列化其中的每一级之后：

[1,2,3,4,5,6,null]
[7,8,9,10,null]
[11,12,null]
为了将每一级都序列化到一起，我们需要每一级中添加值为 null 的元素，以表示没有节点连接到上一级的上级节点。

[1,2,3,4,5,6,null]
[null,null,7,8,9,10,null]
[null,11,12,null]
合并所有序列化结果，并去除末尾的 null 。

[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]


提示：

节点数目不超过 1000
1 <= Node.val <= 10^5
 */

type Node struct {
	    Val int
	    Prev *Node
	    Next *Node
	    Child *Node
}


// dfs

//     1---2---3---4---5---6--NULL
//     |
// 7---8---9---10--NULL
//     |
//    11--12--NULL

func main() {
	root := &Node{Val: 1}

	root.Next = &Node{Val: 2}
	root.Next.Prev = root

	root.Next.Next = &Node{Val: 3}
	root.Next.Next.Prev = root.Next

	root.Next.Next.Next = &Node{Val: 4}
	root.Next.Next.Next.Prev = root.Next.Next

	root.Next.Next.Next.Next = &Node{Val: 5}
	root.Next.Next.Next.Next.Prev = root.Next.Next.Next

	root.Next.Next.Next.Next.Next = &Node{Val: 6}
	root.Next.Next.Next.Next.Next.Prev = root.Next.Next.Next.Next

	root1 := &Node{Val: 7}
	//
	root1.Next = &Node{Val: 8}
	root1.Next.Prev = root1

	root1.Next.Next = &Node{Val: 9}
	root1.Next.Next.Prev = root1.Next

	root1.Next.Next.Next = &Node{Val: 10}
	root1.Next.Next.Next.Prev = root1.Next.Next

	root2 := &Node{Val: 11}
	//
	root2.Next = &Node{Val: 12}
	root2.Next.Prev = root2

	root.Child = root1.Next
	root1.Next.Child = root2


	//root.Child = root1
	//root1.Child = root2

	head := flatten2(root)

	for head != nil {
		fmt.Print(head.Val,"->")
		head = head.Next
	}
}

func flatten(root *Node) *Node {
	h,_ := help(root)
	return h
}

func help(root *Node) (*Node,*Node) {
	if root == nil {
		return nil,nil
	}
	temp := root
	prev := &Node{}
	prev.Next = temp

	for temp != nil {
		next := temp.Next
		if temp.Child != nil {

			temp.Next = nil
			if next != nil {
				next.Prev = nil
			}

			child := temp.Child
			temp.Child = nil

			h,t := help(findHead(child))

			temp.Next = h
			h.Prev = temp

			t.Next = next
			if next != nil {
				next.Prev = t
			}

			temp = next
			prev = t
		} else {
			temp = next
			prev = prev.Next
		}
	}

	return root,prev
}

func findHead(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node.Prev != nil {
		node = node.Prev
	}
	return node
}

func flatten2(root *Node) *Node {
	if root == nil {
		return root
	}
	dummy := &Node{Next: root}
	prev := dummy
	stack := []*Node{}
	stack = append(stack,root)

	for len(stack) > 0 {
		// 出栈的时候设置prev，next
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Prev = prev
		prev.Next = cur

		if cur.Next != nil {
			stack = append(stack,cur.Next)
		}
		if cur.Child != nil {
			stack = append(stack,cur.Child)
			cur.Child = nil
		}
		prev = cur
	}
	dummy.Next.Prev = nil
	return dummy.Next
}

/*

// Definition for a Node.
class Node {
    public int val;
    public Node prev;
    public Node next;
    public Node child;

    public Node() {}

    public Node(int _val,Node _prev,Node _next,Node _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
  public Node flatten(Node head) {
    if (head == null) return head;
    // pseudo head to ensure the `prev` pointer is never none
    Node pseudoHead = new Node(0, null, head, null);

    flattenDFS(pseudoHead, head);

    // detach the pseudo head from the real head
    pseudoHead.next.prev = null;
    return pseudoHead.next;
  }

  // return the tail of the flatten list
  public Node flattenDFS(Node prev, Node curr) {
    if (curr == null) return prev;
    curr.prev = prev;
    prev.next = curr;

    // the curr.next would be tempered in the recursive function
    Node tempNext = curr.next;

    Node tail = flattenDFS(curr, curr.child);
    curr.child = null;

    return flattenDFS(tail, tempNext);
  }
}

 */

/*

// Definition for a Node.
class Node {
    public int val;
    public Node prev;
    public Node next;
    public Node child;

    public Node() {}

    public Node(int _val,Node _prev,Node _next,Node _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
  public Node flatten(Node head) {
    if (head == null) return head;

    Node pseudoHead = new Node(0, null, head, null);
    Node curr, prev = pseudoHead;

    Deque<Node> stack = new ArrayDeque<>();
    stack.push(head);

    while (!stack.isEmpty()) {
      curr = stack.pop();
      prev.next = curr;
      curr.prev = prev;

      if (curr.next != null) stack.push(curr.next);
      if (curr.child != null) {
        stack.push(curr.child);
        // don't forget to remove all child pointers.
        curr.child = null;
      }
      prev = curr;
    }
    // detach the pseudo node from the result
    pseudoHead.next.prev = null;
    return pseudoHead.next;
  }
}

 */