package merge_sorted_array_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/arrays_strings/merge_sorted_array"
)

func TestMerge(t *testing.T) {
	for _, tt := range []struct {
		nums1          []int
		m              int
		nums2          []int
		n              int
		expectedResult []int
	}{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, expectedResult: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, expectedResult: []int{1}},
		{nums1: []int{0, 0, 0, 0}, m: 0, nums2: []int{1, 2, 3, 4}, n: 4, expectedResult: []int{1, 2, 3, 4}},
	} {
		t.Run(
			fmt.Sprintf("should return merged array"),
			func(t *testing.T) {
				merge_sorted_array.Merge(tt.nums1, tt.m, tt.nums2, tt.n)
				assert.Equal(t, tt.expectedResult, tt.nums1)
			},
		)
	}
}
