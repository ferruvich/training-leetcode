package is_palindrome_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/arrays_strings/is_palindrome"
)

func TestIsPalindrome(t *testing.T) {
	for _, tt := range []struct {
		input          string
		expectedResult bool
	}{
		{input: "abba", expectedResult: true},
		{input: "ab!ba!", expectedResult: true},
		{input: "123rerer321", expectedResult: true},
		{input: "a", expectedResult: true},
		{input: "not a palindrome string at all", expectedResult: false},
		{input: "a b b a", expectedResult: true},
		{input: "!?!?!!!", expectedResult: true},
		{input: "A man, a plan, a canal: Panama", expectedResult: true},
		{input: "race a car", expectedResult: false},
	} {
		t.Run(
			fmt.Sprintf("should return isPalindrome=%t as the input is %s", tt.expectedResult, tt.input),
			func(t *testing.T) {
				profit := is_palindrome.IsPalindrome(tt.input)
				assert.Equal(t, tt.expectedResult, profit)
			},
		)
	}
}
