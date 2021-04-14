package flatten_bt_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ferruvich/training-leetcode/trees_graphs/flatten_bt"
)

func TestFlatten(t *testing.T) {
	for _, tt := range []struct {
		tree []int
	}{
		{
			tree: []int{4, 2, 5, 1, 3},
		},
		{
			tree: []int{4, 2, 5, 1, 3},
		},
		{
			tree: []int{1},
		},
		{
			tree: []int{1, 2, 3, 4, 5, 6},
		},
	} {
		t.Run(
			fmt.Sprintf("should correctly flatten the tree"),
			func(t *testing.T) {
				tr := buildTreeFromSlice(t, tt.tree)
				flatten_bt.Flatten(tr)

				var i int
				for tr != nil {
					assert.Contains(t, tt.tree, tr.Val)
					i++
					tr = tr.Right
				}
				assert.Equal(t, len(tt.tree), i)
			},
		)
	}
}

func buildTreeFromSlice(t *testing.T, treeSlice []int) *flatten_bt.TreeNode {
	t.Helper()

	if len(treeSlice) == 0 {
		return nil
	}

	root := &flatten_bt.TreeNode{Val: treeSlice[0]}
	for _, el := range treeSlice[1:] {
		root = insertElement(t, root, el)
	}

	return root
}

func insertElement(t *testing.T, node *flatten_bt.TreeNode, val int) *flatten_bt.TreeNode {
	if node == nil {
		return &flatten_bt.TreeNode{Val: val}
	}

	if node.Val < val {
		node.Right = insertElement(t, node.Right, val)
	} else {
		node.Left = insertElement(t, node.Left, val)
	}

	return node
}
