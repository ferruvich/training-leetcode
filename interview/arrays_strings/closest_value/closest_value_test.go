package closest_value_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/interview/arrays_strings/closest_value"
)

func TestClosestValue(t *testing.T) {
	for _, tt := range []struct {
		tree          *closest_value.TreeNode
		target        float64
		expectedValue int
	}{
		{
			tree:          buildTreeFromSlice(t, []int{4, 2, 5, 1, 3}),
			target:        3.714286,
			expectedValue: 4,
		},
		{
			tree:          buildTreeFromSlice(t, []int{4, 2, 5, 1, 3}),
			target:        3.214286,
			expectedValue: 3,
		},
		{
			tree:          buildTreeFromSlice(t, []int{1}),
			target:        3.214286,
			expectedValue: 1,
		},
	} {
		t.Run(
			fmt.Sprintf("should correctly return value %d as the target is %f", tt.expectedValue, tt.target),
			func(t *testing.T) {
				res := closest_value.ClosestValue(tt.tree, tt.target)
				assert.Equal(t, tt.expectedValue, res)
			},
		)
	}
}

func buildTreeFromSlice(t *testing.T, treeSlice []int) *closest_value.TreeNode {
	t.Helper()

	if len(treeSlice) == 0 {
		return nil
	}

	root := &closest_value.TreeNode{Val: treeSlice[0]}
	for _, el := range treeSlice {
		root = insertElement(t, root, el)
	}

	return root
}

func insertElement(t *testing.T, node *closest_value.TreeNode, val int) *closest_value.TreeNode {
	if node == nil {
		return &closest_value.TreeNode{Val: val}
	}

	if node.Val < val {
		node.Right = insertElement(t, node.Right, val)
	} else {
		node.Left = insertElement(t, node.Left, val)
	}

	return node
}
