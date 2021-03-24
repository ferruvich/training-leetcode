package lru_cache_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ferruvich/training-leetcode/problems/lru_cache"
)

func TestLRUCache_Results(t *testing.T) {
	const (
		getCommand = "get"
		putCommand = "put"
	)
	for _, tt := range []struct {
		capacity int
		commands []string
		params   [][]int
		expected []interface{}
	}{
		{
			capacity: 10,
			commands: []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"},
			params:   [][]int{{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}},
			expected: []interface{}{nil, nil, nil, nil, nil, -1, nil, 19, 17, nil, -1, nil, nil, nil, -1, nil, -1, 5, -1, 12, nil, nil, 3, 5, 5, nil, nil, 1, nil, -1, nil, 30, 5, 30, nil, nil, nil, -1, nil, -1, 24, nil, nil, 18, nil, nil, nil, nil, -1, nil, nil, 18, nil, nil, -1, nil, nil, nil, nil, nil, 18, nil, nil, -1, nil, 4, 29, 30, nil, 12, -1, nil, nil, nil, nil, 29, nil, nil, nil, nil, 17, 22, 18, nil, nil, nil, -1, nil, nil, nil, 20, nil, nil, nil, -1, 18, 18, nil, nil, nil, nil, 20, nil, nil, nil, nil, nil, nil, nil},
		},
	} {
		t.Run(
			fmt.Sprintf("should return correctly handle a LRU with capacity %d with %d commands", tt.capacity, len(tt.commands)),
			func(t *testing.T) {
				cache := lru_cache.Constructor(tt.capacity)
				for i, c := range tt.commands {
					switch c {
					case getCommand:
						val := cache.Get(tt.params[i][0])
						require.Equal(t, tt.expected[i].(int), val)
					case putCommand:
						cache.Put(tt.params[i][0], tt.params[i][1])
					}
				}
			},
		)
	}
}
