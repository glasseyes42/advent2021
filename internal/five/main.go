package five

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ventLine struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func (vl *ventLine) GetPointsAsKeys() []string {
	points := []string{}
	xIncr := 1
	yIncr := 1
	if vl.EndX < vl.StartX {
		xIncr = -1
	}
	if vl.EndY < vl.StartY {
		yIncr = -1
	}

	currentX := vl.StartX
	currentY := vl.StartY
	completedX := false
	completedY := false
	for !(completedX && completedY) {
		points = append(points, fmt.Sprintf("%d,%d", currentX, currentY))
		if currentX != vl.EndX {
			currentX += xIncr
		} else {
			completedX = true
		}
		if currentY != vl.EndY {
			currentY += yIncr
		} else {
			completedY = true
		}

	}

	return points
}

func parseIntOrFail(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func ParseFile(file string) []ventLine {
	lines := strings.Split(strings.TrimSpace(file), "\n")

	lineRegex := regexp.MustCompile(`(\d+),(\d+)\s->\s(\d+),(\d+)`)
	vents := make([]ventLine, len(lines))
	for idx, line := range lines {
		matches := lineRegex.FindStringSubmatch(line)

		vents[idx] = ventLine{
			StartX: parseIntOrFail(matches[1]),
			StartY: parseIntOrFail(matches[2]),
			EndX:   parseIntOrFail(matches[3]),
			EndY:   parseIntOrFail(matches[4]),
		}
	}
	return vents
}

func FilterStraight(lines []ventLine) []ventLine {
	filtered := []ventLine{}

	for _, vl := range lines {
		if (vl.StartX == vl.EndX) || (vl.StartY == vl.EndY) {
			filtered = append(filtered, vl)
		}
	}

	return filtered
}

func CountPointsWithOverlap(min int, lines []ventLine) int {
	points := map[string]int{}

	for _, vl := range lines {
		linePoints := vl.GetPointsAsKeys()

		for _, p := range linePoints {
			points[p]++
		}
	}

	sum := 0
	for _, p := range points {
		if p >= min {
			sum++
		}
	}

	return sum
}
