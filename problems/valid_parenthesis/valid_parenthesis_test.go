package valid_parenthesis_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/problems/valid_parenthesis"
)

func TestValidParenthesis(t *testing.T) {
	for _, tt := range []struct {
		input string
		valid bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	} {
		t.Run(
			fmt.Sprintf("should return valid=%t as string is %s", tt.valid, tt.input),
			func(t *testing.T) {
				res := valid_parenthesis.ValidParenthesis(tt.input)
				assert.Equal(t, tt.valid, res)
			},
		)
	}
}
