package add_two_numbers_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ferruvich/training-leetcode/linked_lists/add_two_numbers/go"
)

func TestAddTwoNumbers(t *testing.T) {
	for _, tt := range []struct {
		inputListOne   []int
		inputListTwo   []int
		expectedResult []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9,9,9,9,9,9,9}, []int{9,9,9,9}, []int{8,9,9,9,0,0,0,1}},
	} {
		t.Run(
			fmt.Sprintf("should return correct sum"),
			func(t *testing.T) {
				res := add_two_numbers.AddTwoNumbers(
					buildLinkedList(t, tt.inputListOne),
					buildLinkedList(t, tt.inputListTwo),
				)
				testResult(t, tt.expectedResult, res)
			},
		)
	}
}

func buildLinkedList(t *testing.T, l []int) *add_two_numbers.ListNode {
	if len(l) == 0 {
		return nil
	}

	return &add_two_numbers.ListNode{
		Val:  l[0],
		Next: buildLinkedList(t, l[1:]),
	}
}

func testResult(t *testing.T, res []int, l *add_two_numbers.ListNode) {
	t.Helper()

	if len(res) == 0 && l == nil {
		return
	}

	require.NotNil(t, l)
	require.NotEmpty(t, res)

	assert.Equal(t, res[0], l.Val)

	testResult(t, res[1:], l.Next)
}
