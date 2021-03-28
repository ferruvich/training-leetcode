package closest_value

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	bigFloat = float64(1 << 64)
)

func ClosestValue(root *TreeNode, target float64) int {
	closest := -1
	minDistance := bigFloat

	el := root
	for el != nil {
		distance := math.Abs(float64(el.Val) - target)
		if distance < minDistance {
			closest = el.Val
			minDistance = distance
		}

		if float64(el.Val) < target {
			el = el.Right
		} else {
			el = el.Left
		}
	}

	return closest
}
