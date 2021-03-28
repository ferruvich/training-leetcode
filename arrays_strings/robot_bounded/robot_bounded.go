package robot_bounded

import (
	"strings"
)

const (
	forward = "G"
	left    = "L"
	right   = "R"
)

const (
	north = iota
	east
	south
	west
)

func IsRobotBounded(instructions string) bool {
	x, y := 0, 0
	direction := north
	commands := strings.Split(instructions, "")

	for i := 0; i < 4; i++ {
		for _, command := range commands {
			if command == forward {
				nextPoint := move(x, y, direction)
				x, y = nextPoint[0], nextPoint[1]
				continue
			}

			direction = changeDirection(direction, command)
		}
		if x == 0 && y == 0 {
			return true
		}
	}

	return false
}

func move(x, y, direction int) []int {
	switch direction {
	case north:
		return []int{x + 1, y}
	case east:
		return []int{x, y + 1}
	case south:
		return []int{x - 1, y}
	case west:
		return []int{x, y - 1}
	}

	return nil
}

func changeDirection(currentDirection int, command string) int {
	switch command {
	case left:
		if currentDirection == north {
			return west
		}
		return currentDirection - 1
	case right:
		if currentDirection == west {
			return north
		}
		return currentDirection + 1
	}

	return currentDirection
}
