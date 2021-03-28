package remove_duplicates_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/interview/arrays_strings/remove_duplicates/go"
)

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range []struct {
		input          []int
		expectedResult []int
	}{
		{input: []int{}, expectedResult: []int{}},
		{input: []int{1}, expectedResult: []int{1}},
		{input: []int{1, 1}, expectedResult: []int{1}},
		{input: []int{1, 1, 2}, expectedResult: []int{1, 2}},
		{input: []int{1, 1, 1, 2, 3, 4, 4}, expectedResult: []int{1, 2, 3, 4}},
		{input: []int{1, 2, 3, 3, 4, 5, 5, 6}, expectedResult: []int{1, 2, 3, 4, 5, 6}},
	} {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			resLen := remove_duplicates.RemoveDuplicates(tt.input)
			assert.Len(t, tt.expectedResult, resLen)
		})
	}
}
