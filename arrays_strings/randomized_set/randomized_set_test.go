package randomized_set_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/training-leetcode/arrays_strings/randomized_set"
)

func TestRandomizedSet_Behaviour(t *testing.T) {
	const (
		insert    = "insert"
		remove    = "remove"
		getRandom = "getRandom"
	)
	for _, tt := range []struct {
		commands []string
		params   []interface{}
		expected []interface{}
	}{
		{
			commands: []string{"insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
			params:   []interface{}{1, 2, 2, nil, 1, 2, nil},
			expected: []interface{}{true, false, true, 2, true, false, 2},
		},
		{
			commands: []string{"remove", "remove", "insert", "getRandom", "remove", "insert"},
			params:   []interface{}{0, 0, 0, nil, 0, 0},
			expected: []interface{}{false, false, true, 0, true, true},
		},
	} {
		t.Run(
			fmt.Sprintf("should return correctly handle a Randomized set with input long %d", len(tt.commands)),
			func(t *testing.T) {
				set := randomized_set.Constructor()
				for i, c := range tt.commands {
					switch c {
					case insert:
						inserted := set.Insert(tt.params[i].(int))
						require.Equal(t, tt.expected[i].(bool), inserted)
					case remove:
						removed := set.Remove(tt.params[i].(int))
						require.Equal(t, tt.expected[i].(bool), removed)
					case getRandom:
						val := set.GetRandom()
						require.Equal(t, tt.expected[i].(int), val)
					}
				}
			},
		)
	}
}
