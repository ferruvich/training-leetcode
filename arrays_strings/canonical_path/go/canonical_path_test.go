package canonical_path_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	canonical_path "github.com/ferruvich/training-leetcode/arrays_strings/canonical_path/go"
)

func TestSimplifyPath(t *testing.T) {
	for _, tt := range []struct {
		path           string
		simplifiedPath string
	}{
		{path: "/home/", simplifiedPath: "/home"},
		{path: "/../", simplifiedPath: "/"},
		{path: "/home//foo", simplifiedPath: "/home/foo"},
		{path: "/a/./b/../../c/", simplifiedPath: "/c"},
	} {
		t.Run(
			fmt.Sprintf("should correctly return %s as the path is %s", tt.simplifiedPath, tt.path),
			func(t *testing.T) {
				res := canonical_path.SimplifyPath(tt.path)
				assert.Equal(t, tt.simplifiedPath, res)
			},
		)
	}
}
