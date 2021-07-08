package k_closest

import (
	"math"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return distanceFromOrigin(points[i][0], points[i][1]) <
			distanceFromOrigin(points[j][0], points[j][1])
	})

	return points[:k]
}

func distanceFromOrigin(x, y int) float64 {
	return math.Sqrt(
		float64(x*x + y*y),
	)
}
