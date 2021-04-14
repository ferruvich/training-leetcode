package move_zeros_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/arrays_strings/move_zeros"
)

func TestMergeIntervals(t *testing.T) {
	for _, tt := range []struct {
		input          []int
		expectedResult []int
	}{
		{input: []int{1, 2, 0, 3, 4, 0}, expectedResult: []int{1, 2, 3, 4, 0, 0}},
		{input: []int{1}, expectedResult: []int{1}},
		{input: []int{0}, expectedResult: []int{0}},
		{input: []int{0, 0, 1}, expectedResult: []int{1, 0, 0}},
		{input: []int{0, 1, 0, 1, -2, 3}, expectedResult: []int{1, 1, -2, 3, 0, 0}},
		{input: []int{}, expectedResult: []int{}},
	} {
		t.Run(
			fmt.Sprintf("should return %v when input is %v", tt.expectedResult, tt.input),
			func(t *testing.T) {
				move_zeros.MoveZeroes(tt.input)
				assert.Equal(t, tt.expectedResult, tt.input)
			},
		)
	}
}
