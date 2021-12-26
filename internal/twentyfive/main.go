package twentyfive

import (
	"fmt"
	"strings"
)

type GridPoint struct {
	Right              *GridPoint
	Below              *GridPoint
	Contains           string
	MovedThisIteration bool
}

func (gp *GridPoint) TryMoveRight() bool {
	if !gp.MovedThisIteration && !gp.Right.MovedThisIteration && gp.Contains == ">" && gp.Right.Contains == "." {
		gp.Right.Contains = gp.Contains
		gp.Contains = "."
		gp.MovedThisIteration = true
		gp.Right.MovedThisIteration = true
		return true
	}
	return false
}

func (gp *GridPoint) TryMovedown() bool {
	if !gp.MovedThisIteration && !gp.Below.MovedThisIteration && gp.Contains == "v" && gp.Below.Contains == "." {
		gp.Below.Contains = gp.Contains
		gp.Contains = "."
		gp.MovedThisIteration = true
		gp.Below.MovedThisIteration = true
		return true
	}
	return false
}

func (gp *GridPoint) String() string {
	return gp.Contains
}

func printMap(m [][]*GridPoint) {
	for _, line := range m {
		fmt.Println(line)
	}
	fmt.Println()
}

func ParseFile(file string) [][]*GridPoint {
	lines := strings.Split(strings.TrimSpace(file), "\n")

	locations := make([][]*GridPoint, len(lines))

	for lineIdx, line := range lines {
		points := strings.Split(line, "")

		linePoints := make([]*GridPoint, len(points))
		for xIdx, p := range points {
			linePoints[xIdx] = &GridPoint{Contains: p}
		}

		locations[lineIdx] = linePoints
	}

	for lIdx, line := range locations {
		for xIdx, gp := range line {
			if xIdx == (len(line) - 1) {
				gp.Right = line[0]
			} else {
				gp.Right = line[xIdx+1]
			}

			if lIdx == (len(locations) - 1) {
				gp.Below = locations[0][xIdx]
			} else {
				gp.Below = locations[lIdx+1][xIdx]
			}
		}
	}

	return locations
}

func IterateTillStopped(seedMap [][]*GridPoint) int {
	iteration := 0
	thingsMoved := true

	for thingsMoved {
		iteration++
		thingsMoved = false

		for lIdx := len(seedMap) - 1; lIdx >= 0; lIdx-- {
			line := seedMap[lIdx]
			for xIdx := len(line) - 1; xIdx >= 0; xIdx-- {
				p := line[xIdx]
				didMove := p.TryMoveRight()
				thingsMoved = thingsMoved || didMove
			}
		}
		for _, line := range seedMap {
			for _, gp := range line {
				gp.MovedThisIteration = false
			}
		}

		for lIdx := len(seedMap) - 1; lIdx >= 0; lIdx-- {
			line := seedMap[lIdx]
			for xIdx := len(line) - 1; xIdx >= 0; xIdx-- {
				p := line[xIdx]
				didMove := p.TryMovedown()
				thingsMoved = thingsMoved || didMove
			}
		}

		for _, line := range seedMap {
			for _, gp := range line {
				gp.MovedThisIteration = false
			}
		}
	}

	return iteration
}
