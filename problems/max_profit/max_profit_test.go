package max_profit_test

import (
	"fmt"
	"testing"

	"github.com/ferruvich/training-leetcode/max_profit"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	for _, tt := range []struct {
		input          []int
		expectedProfit int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, expectedProfit: 5},
		{input: []int{7, 6, 4, 3, 1}, expectedProfit: 0},
	} {
		t.Run(
			fmt.Sprintf("should correctly return a profit equal to %d as the list of values is %v", tt.expectedProfit, tt.input),
			func(t *testing.T) {
				profit := max_profit.MaxProfit(tt.input)
				assert.Equal(t, tt.expectedProfit, profit)
			},
		)
	}
}
