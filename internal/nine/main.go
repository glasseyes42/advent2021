package nine

import (
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

func parseFile(file string) [][]*location {
	split := strings.Split(strings.TrimSpace(file), "\n")
	grid := make([][]*location, len(split))
	for idx, line := range split {
		toConvert := strings.Split(strings.TrimSpace(line), "")
		nums := make([]*location, len(toConvert))
		for nIdx, num := range toConvert {
			nums[nIdx] = &location{
				Value: parseIntOrFail(num),
			}
		}
		grid[idx] = nums
	}

	for y, line := range grid {
		for x, loc := range line {
			if y-1 != -1 {
				loc.Above = grid[y-1][x]
			}
			if y+1 != len(grid) {
				loc.Below = grid[y+1][x]
			}
			if x-1 != -1 {
				loc.Left = grid[y][x-1]
			}
			if x+1 != len(grid[y]) {
				loc.Right = grid[y][x+1]
			}
		}
	}
	return grid
}

func RiskOfLowLevels(file string) int {
	grid := parseFile(file)

	sum := 0
	for _, line := range grid {
		for _, loc := range line {
			if loc.IsLowPoint() {
				sum += loc.Risk()
			}
		}
	}

	return sum
}
