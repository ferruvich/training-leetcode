package canonical_path

import (
	"fmt"
	"strings"
)

func SimplifyPath(path string) string {
	canonicalPath := ""

	for _, el := range strings.Split(path, "/") {
		switch el {
		case "", ".":
			continue
		case "..":
			if len(canonicalPath) == 0 {
				continue
			}
			canonicalPath = string([]rune(canonicalPath)[:strings.LastIndex(canonicalPath, "/")])
		default:
			//prec = el
			canonicalPath = fmt.Sprintf("%s/%s", canonicalPath, el)
		}
	}

	if len(canonicalPath) == 0 {
		return "/"
	}
	return canonicalPath
}
