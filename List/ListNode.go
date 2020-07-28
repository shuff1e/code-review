package List

// 123 456 789 10
// 翻转成 321 654 987 10

type ListNode struct {
	Data int
	Next *ListNode
}

func ReverseKList(node *ListNode,k int) *ListNode {
	if node == nil {
		return nil
	}
	head := node
	temp := k-1
	for temp > 0 && node !=nil {
		temp--
		node = node.Next
	}
	if temp > 0 || node == nil {
		return head
	}
	next := node.Next
	node.Next = nil
	left,right := reverseList(head)
	right.Next = ReverseKList(next,k)
	return left
}

func reverseList(node *ListNode) (*ListNode,*ListNode) {
	tail := node
	cur := node
	var prev *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev,tail
}
