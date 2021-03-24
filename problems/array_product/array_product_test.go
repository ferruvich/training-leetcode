package array_product_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/array_product"
)

func TestProductExceptSelf(t *testing.T) {
	for _, tt := range []struct {
		input          []int
		expectedResult []int
	}{
		{input: []int{1, 2, 3, 4}, expectedResult: []int{24, 12, 8, 6}},
		{input: []int{-1, 1, 0, -3, 3}, expectedResult: []int{0, 0, 9, 0, 0}},
	} {
		t.Run(
			fmt.Sprintf("should return %v as the input is %v", tt.expectedResult, tt.input),
			func(t *testing.T) {
				res := array_product.ProductExceptSelf(tt.input)
				assert.Equal(t, tt.expectedResult, res)
			},
		)
	}
}
