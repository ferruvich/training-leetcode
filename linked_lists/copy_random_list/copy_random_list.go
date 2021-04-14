package copy_random_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	newHead := &Node{Val: head.Val}
	oldNew := map[*Node]*Node{head: newHead}

	el := newHead
	for head != nil {
		if head.Next != nil {
			newNext, ok := oldNew[head.Next]
			if !ok {
				newNext = &Node{Val: head.Next.Val}
				oldNew[head.Next] = newNext
			}

			el.Next = newNext
		}

		if head.Random != nil {
			newRandom, ok := oldNew[head.Random]
			if !ok {
				newRandom = &Node{Val: head.Random.Val}
				oldNew[head.Random] = newRandom
			}

			el.Random = newRandom
		}

		oldNew[head] = el

		el = el.Next
		head = head.Next
	}

	return newHead
}
