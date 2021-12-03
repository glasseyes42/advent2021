package one

import (
	"strconv"
)

func toInt(lines []string) ([]int, error) {
	toReturn := make([]int, len(lines))

	for idx, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		toReturn[idx] = i
	}

	return toReturn, nil
}

func CountIncreases(linesRaw []string) int {
	lines, err := toInt(linesRaw)
	if err != nil {
		panic(err)
	}
	count := 0
	last := lines[0]
	remaining := lines[1:]

	for len(remaining) > 0 {
		current := remaining[0]
		if current > last {
			count++
		}

		last = current
		remaining = remaining[1:]
	}

	return count
}

func CountIncreasesWindowed(linesRaw []string) int {
	lines, err := toInt(linesRaw)
	if err != nil {
		panic(err)
	}

	sums := []int{}
	currentSum, nextSum := lines[0]+lines[1], lines[1]
	lines = lines[2:]

	for _, line := range lines {
		currentSum += line
		nextSum += line

		sums = append(sums, currentSum)
		currentSum = nextSum
		nextSum = line

	}

	count := 0
	last := sums[0]
	remaining := sums[1:]

	for len(remaining) > 0 {
		current := remaining[0]
		if current > last {
			count++
		}

		last = current
		remaining = remaining[1:]
	}
	return count
}
