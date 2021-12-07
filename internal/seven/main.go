package seven

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseIntOrFail(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func ParseFile(file string) []int {
	split := strings.Split(strings.TrimSpace(file), ",")

	vals := make([]int, len(split))

	for idx, s := range split {
		vals[idx] = parseIntOrFail(s)
	}

	return vals
}

func findMedian(vals []int) int {
	sorted := make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)

	midIdx := len(sorted) / 2
	if midIdx%2 == 0 {
		return (sorted[midIdx-1] + sorted[midIdx]) / 2
	} else {
		return sorted[midIdx]
	}
}

func CalcFuel(crabs []int) int {
	median := findMedian(crabs)
	fuel := 0

	for _, crab := range crabs {
		fuel += int(math.Abs(float64(crab - median)))
	}
	return fuel
}

func calcFuelIncrForPosition(crabs []int, position int) int {
	fuel := 0
	for _, crab := range crabs {
		diff := int(math.Abs(float64(crab - position)))

		fuel += (diff * (diff + 1)) / 2
	}
	return fuel
}

func getMinMaxCrab(crabs []int) (int, int) {
	sorted := make([]int, len(crabs))
	copy(sorted, crabs)
	sort.Ints(sorted)

	return sorted[0], sorted[len(sorted)-1]
}

func CalcFuel2(crabs []int) int {
	min, max := getMinMaxCrab(crabs)
	fuel := calcFuelIncrForPosition(crabs, min)
	for position := min + 1; position <= max; position++ {
		result := calcFuelIncrForPosition(crabs, position)
		if result < fuel {
			fuel = result
		}
	}

	return fuel
}
