package valid_bst

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	return isValidRec(root, int(math.Inf(-1)), int(math.Inf(1))-1)
}

func isValidRec(bst *TreeNode, min, max int) bool {
	if bst == nil {
		return true
	}

	if bst.Val <= min || bst.Val >= max {
		return false
	}

	return isValidRec(bst.Left, min, bst.Val) && isValidRec(bst.Right, bst.Val, max)
}
