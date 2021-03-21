package merge_intervals

import (
	"sort"
)

func MergeIntervals(intervals [][]int) [][]int {
	merged := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for _, el := range intervals {
		if len(merged) == 0 || el[0] > merged[len(merged)-1][1] {
			merged = append(merged, el)
			continue
		}

		if merged[len(merged)-1][1] < el[1] {
			merged[len(merged)-1][1] = el[1]
		}
	}

	return merged
}
