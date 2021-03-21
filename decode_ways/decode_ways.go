package decode_ways

import (
	"strconv"
)

var mem = map[string]int{}

func NumDecodings(s string) int {
	if res, ok := mem[s]; ok {
		return res
	}

	if len(s) == 0 {
		return 1
	}
	if len(s) == 1 && s != "0" {
		return 1
	}

	r := []rune(s)
	if string(r[0]) == "0" {
		mem[s] = 0
		return 0
	}

	decodeWays := NumDecodings(string(r[1:]))
	if n, _ := strconv.Atoi(string(r[0:2])); n <= 26 {
		decodeWays += NumDecodings(string(r[2:]))
	}

	mem[s] = decodeWays
	return decodeWays
}
