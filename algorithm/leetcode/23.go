package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	const INTEGER_MAX = 1<<31-1
	minValue := INTEGER_MAX
	var minNode *ListNode
	minIndex := -1

	head := &ListNode{}
	result := head

	for index, list := range lists {
		if list == nil {
			continue
		}
		if list.Val < minValue {
			minNode = list
			minValue = list.Val
			minIndex = index
		}
	}
	if minValue == INTEGER_MAX {
		return nil
	}

	head.Next = minNode
	if lists[minIndex].Next != nil {
		lists[minIndex] = lists[minIndex].Next
	} else {
		lists = append(lists[:minIndex],lists[minIndex+1:]...)
	}

	minNode.Next = mergeKLists(lists)

	return result.Next
}

func main() {

}
