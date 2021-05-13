package bst_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	traversal []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var traversal []*TreeNode
	for root != nil {
		traversal = append(traversal, root)
		root = root.Left
	}

	return BSTIterator{traversal: traversal}
}

func (bst *BSTIterator) Next() int {
	next := bst.traversal[len(bst.traversal)-1]

	bst.traversal = bst.traversal[:len(bst.traversal)-1]
	el := next.Right
	for el != nil {
		bst.traversal = append(bst.traversal, el)
		el = el.Left
	}

	return next.Val
}

func (bst *BSTIterator) HasNext() bool {
	return len(bst.traversal) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
