package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersRec(l1, l2, 0)
}

func addTwoNumbersRec(l1 *ListNode, l2 *ListNode, rest int) *ListNode {
	fv, sv := 0,0
	var l1Next, l2Next *ListNode

	switch l := []*ListNode{l1, l2}; {
	case l[0] == nil && l[1] == nil:
		if rest == 0 {
			return nil
		}
		return &ListNode{Val: rest}
	case l[1] == nil:
		fv = l1.Val
		l1Next = l1.Next
	case l[0] == nil:
		sv = l2.Val
		l2Next = l2.Next
	default:
		fv = l1.Val
		sv = l2.Val
		l1Next = l1.Next
		l2Next = l2.Next
	}

	sum := fv + sv + rest
	unit := sum
	r := 0
	if sum >= 10 {
		unit = sum % 10
		r = sum/10
	}

	return &ListNode{
		Val: unit,
		Next: addTwoNumbersRec(l1Next, l2Next, r),
	}
}
