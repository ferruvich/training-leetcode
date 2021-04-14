package copy_random_list_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ferruvich/training-leetcode/linked_lists/copy_random_list"
)

func TestCopyRandomList(t *testing.T) {
	for _, tt := range []struct {
		input *copy_random_list.Node
	}{
		{input: fromSliceToLinkedList(t, [][]int{})},
		{input: fromSliceToLinkedList(t, [][]int{{1}})},
		{input: fromSliceToLinkedList(t, [][]int{{1}, {2}, {3}})},
		{input: fromSliceToLinkedList(t, [][]int{{1, 1}, {2, 0}, {3}})},
		{input: fromSliceToLinkedList(t, [][]int{{1, 0}, {2, 2}, {3}})},
		{input: fromSliceToLinkedList(t, [][]int{{1, 2}, {2, 2}, {3}})},
	} {
		t.Run("should return copied list", func(t *testing.T) {
			out := copy_random_list.CopyRandomList(tt.input)
			for out != nil && tt.input != nil {
				assert.NotSame(t, tt.input, out)
				assert.Equal(t, tt.input, out)

				if out.Random != nil {
					assert.NotSame(t, tt.input.Random, out.Random)
					assert.EqualValues(t, tt.input.Random, out.Random)
				}

				out = out.Next
				tt.input = tt.input.Next
			}
		})
	}
}

func fromSliceToLinkedList(t *testing.T, s [][]int) *copy_random_list.Node {
	var elements []*copy_random_list.Node

	// Nodes
	var prev *copy_random_list.Node
	for _, entry := range s {
		el := &copy_random_list.Node{Val: entry[0]}
		elements = append(elements, el)

		if prev != nil {
			prev.Next = el
		}

		prev = el
	}

	// Random
	for i, el := range elements {
		if len(s[i]) == 2 {
			el.Random = elements[s[i][1]]
		}
	}

	if len(elements) > 0 {
		return elements[0]
	}
	return nil
}
