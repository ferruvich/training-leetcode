package substring_no_rep_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/substring_no_rep"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tt := range []struct {
		inputString       string
		expectedMaxLength int
	}{
		{inputString: "", expectedMaxLength: 0},
		{inputString: "pwwkew", expectedMaxLength: 3},
		{inputString: "bbbbb", expectedMaxLength: 1},
		{inputString: "abcabcbb", expectedMaxLength: 3},
	} {
		t.Run(
			fmt.Sprintf("should return correctly return length %d as the string is %s", tt.expectedMaxLength, tt.inputString),
			func(t *testing.T) {
				res := substring_no_rep.LengthOfLongestSubstring(tt.inputString)
				assert.Equal(t, tt.expectedMaxLength, res)
			},
		)
	}
}
