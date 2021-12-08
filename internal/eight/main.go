package eight

import (
	"strings"
)

func CountSimple(file string) int {
	lines := strings.Split(strings.TrimSpace(file), "\n")

	total := 0
	for _, line := range lines {
		hyphenSplit := strings.Split(strings.TrimSpace(line), "|")
		rhs := hyphenSplit[1]
		segments := strings.Split(strings.TrimSpace(rhs), " ")

		for _, segment := range segments {
			length := len(segment)

			if length == 4 || length == 3 || length == 7 || length == 2 {
				total++
			}
		}
	}

	return total
}
