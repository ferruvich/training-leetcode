package robot_bounded_test

import (
	"fmt"
	"testing"

	"github.com/ferruvich/training-leetcode/robot_bounded"
	"github.com/stretchr/testify/assert"
)

func TestIsRobotBounded(t *testing.T) {
	for _, tt := range []struct {
		command string
		bounded bool
	}{
		{command: "GGLLGG", bounded: true},
		{command: "GG", bounded: false},
		{command: "GL", bounded: true},
	} {
		t.Run(
			fmt.Sprintf("should return correctly return bounded=%t when command is %s", tt.bounded, tt.command),
			func(t *testing.T) {
				res := robot_bounded.IsRobotBounded(tt.command)
				assert.Equal(t, tt.bounded, res)
			},
		)
	}
}
