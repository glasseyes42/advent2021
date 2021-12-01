package one

import (
	"strconv"
)

func CountIncreases(lines []string) int {
	count := 0
	last, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	remaining := lines[1:]

	for len(remaining) > 0 {
		current, err := strconv.Atoi(remaining[0])
		if err != nil {
			panic(err)
		}

		if current > last {
			count++
		}

		last = current
		remaining = remaining[1:]
	}

	return count
}
