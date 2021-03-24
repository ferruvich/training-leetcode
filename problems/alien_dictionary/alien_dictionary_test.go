package alien_dictionary_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/problems/alien_dictionary"
)

func TestIsAlienSorted(t *testing.T) {
	for _, tt := range []struct {
		input  []string
		order  string
		sorted bool
	}{
		{input: []string{"app", "apple"}, order: "abcdefghijklmnopqrstuvwxyz", sorted: true},
		{input: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz", sorted: true},
		{input: []string{"hello", "leetcode"}, order: "lhabcdefgijkmnopqrstuvwxyz", sorted: false},
		{input: []string{"word","world","row"}, order: "worldabcefghijkmnpqstuvxyz", sorted: false},
	} {
		t.Run(
			fmt.Sprintf("should return %t as the strings are %v and order is %s", tt.sorted, tt.input, tt.order),
			func(t *testing.T) {
				res := alien_dictionary.IsAlienSorted(tt.input, tt.order)
				assert.Equal(t, tt.sorted, res)
			},
		)
	}
}
