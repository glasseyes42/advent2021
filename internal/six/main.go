package six

import (
	"strconv"
	"strings"
)

func ParseFile(file string) []int {
	vals := strings.Split(strings.TrimSpace(file), ",")

	fishCounters := make([]int, 9)
	for _, fish := range vals {
		val, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}

		fishCounters[val]++
	}

	return fishCounters
}

func ProcessDay(previousDay []int) []int {
	nextDay := make([]int, 9)

	nextDay[6] = previousDay[0]
	nextDay[8] = previousDay[0]

	for i := 1; i < 9; i++ {
		nextDay[i-1] += previousDay[i]
	}

	return nextDay
}

func ProcessDays(count int, seed []int) int {
	result := seed
	for day := 0; day < count; day++ {
		result = ProcessDay(result)
	}

	sum := 0
	for _, count := range result {
		sum += count
	}

	return sum
}
