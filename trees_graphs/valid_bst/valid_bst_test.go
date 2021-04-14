package valid_bst_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/trees_graphs/valid_bst"
)

func TestIsValidBST(t *testing.T) {
	for _, tt := range []struct {
		tree  []int
		valid bool
	}{
		{
			tree:  []int{4, 2, 5, 1, 3},
			valid: true,
		},
		{
			tree: []int{4, 2, 5, 1, 3},
		},
		{
			tree:  []int{1},
			valid: true,
		},
		{
			tree: []int{1, 2, 3, 4, 5, 6},
		},
	} {
		t.Run(
			fmt.Sprintf("should correctly flatten the tree"),
			func(t *testing.T) {
				valid := valid_bst.IsValidBST(buildTreeFromSlice(t, tt.tree, tt.valid))
				assert.Equal(t, tt.valid, valid)
			},
		)
	}
}

func buildTreeFromSlice(t *testing.T, treeSlice []int, shouldBeValid bool) *valid_bst.TreeNode {
	t.Helper()

	if len(treeSlice) == 0 {
		return nil
	}

	root := &valid_bst.TreeNode{Val: treeSlice[0]}
	for _, el := range treeSlice[1:] {
		root = insertElement(t, root, el, shouldBeValid)
	}

	return root
}

func insertElement(t *testing.T, node *valid_bst.TreeNode, val int, shouldBeValid bool) *valid_bst.TreeNode {
	if node == nil {
		return &valid_bst.TreeNode{Val: val}
	}

	if node.Val < val && shouldBeValid {
		node.Right = insertElement(t, node.Right, val, shouldBeValid)
	} else {
		node.Left = insertElement(t, node.Left, val, shouldBeValid)
	}

	return node
}
