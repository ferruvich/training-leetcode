package max_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	maxSum = int(math.Inf(-1) - 1)
)

func MaxPathSum(root *TreeNode) int {
	maxSum = int(math.Inf(-1) - 1)
	maxGain(root)

	return maxSum
}

func maxGain(node *TreeNode) int {
	if node == nil {
		return 0
	}

	var leftGain, rightGain int
	lg := maxGain(node.Left)
	if lg > 0 {
		leftGain = lg
	}

	rg := maxGain(node.Right)
	if rg > 0 {
		rightGain = rg
	}

	priceNewPath := node.Val + leftGain + rightGain

	if priceNewPath > maxSum {
		maxSum = priceNewPath
	}

	maxPath := leftGain
	if rightGain > maxPath {
		maxPath = rightGain
	}

	return node.Val + maxPath
}
