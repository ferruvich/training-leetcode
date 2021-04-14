package flatten_bt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	_ = flattenRec(root)
}

func flattenRec(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	flattenedLeft := flattenRec(root.Left)
	flattenedRight := flattenRec(root.Right)

	root.Right = flattenedLeft

	if root.Right == nil {
		root.Right = flattenedRight
	} else {
		el := root.Right
		for el.Right != nil {
			el = el.Right
		}
		el.Right = flattenedRight
	}
	root.Left = nil

	return root
}
