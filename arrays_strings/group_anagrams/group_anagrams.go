package group_anagrams

import "sort"

// This is the solution I came up with after reading the explanation.
// Time complexity: O(nlogn) -> a sorting algorithm executed once for each entry
// Space complexity O(n) -> allocating space for the map and res
func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	var res [][]string

	for _, str := range strs {
		ordered := orderString(str)
		strings, ok := m[ordered]
		if !ok {
			strings = []string{}
		}
		m[ordered] = append(strings, str)
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}

func orderString(str string) string {
	// this is not space allocated, we're simply
	// treating THE SAME piece of memory in a different way
	r := []rune(str)

	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })

	return string(r)
}

// This is the first solution that came in my mind.
// Time complexity: O(n*m*s)
// Space complexity: O(m*s)
// where n=len(strs) and m=overall len(res) and s=average length of each string.
func GroupAnagramsUnefficient(strs []string) [][]string {
	var res [][]string

	for _, str := range strs {
		var found bool
		for i, el := range res {
			if len(str) != len(el[0]) {
				continue
			}
			if sameMap(createMap(str), createMap(el[0])) {
				res[i] = append(el, str)
				found = true
				break
			}
		}
		if !found {
			res = append(res, []string{str})
		}
	}

	return res
}

func createMap(str string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range str {
		if n, ok := m[c]; ok {
			m[c] = n + 1
			continue
		}
		m[c] = 1
	}

	return m
}

func sameMap(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}
