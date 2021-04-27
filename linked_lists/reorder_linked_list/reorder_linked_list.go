package reorder_linked_list

// ListNode is the definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	i := 0
	listLen := 0

	el := head
	for el != nil {
		el = el.Next
		listLen++
	}

	el = head
	for el != nil && i < listLen {
		// We're getting the Previous element of the one we should have
		// as we need to set its next to nil
		previousLast := getListNode(el, listLen-i-2)
		elNext := el.Next
		if previousLast != nil && previousLast.Next != nil && previousLast != el {
			el.Next = previousLast.Next
			el.Next.Next = elNext
			previousLast.Next = nil
		}

		el = elNext
		i += 2
	}
}

func getListNode(root *ListNode, index int) *ListNode {
	if root == nil || index < 0 {
		return nil
	}

	i := 0
	for root != nil {
		i++
		if i > index {
			return root
		}
		root = root.Next
	}

	return root
}
