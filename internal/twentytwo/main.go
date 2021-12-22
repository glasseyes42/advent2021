package twentytwo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Toggle string
	XMin   int
	XMax   int
	YMin   int
	YMax   int
	ZMin   int
	ZMax   int
}

func (i *Instruction) RangeWithLimits(atLeast, atMost int) []string {
	if i.XMin > atMost || i.YMin > atMost || i.ZMin > atMost {
		return []string{}
	}
	if i.XMax < atLeast || i.YMax < atLeast || i.ZMax < atLeast {
		return []string{}
	}

	cappedXMin := atLeast
	cappedYMin := atLeast
	cappedZMin := atLeast
	cappedXMax := atMost
	cappedYMax := atMost
	cappedZMax := atMost
	if i.XMin > atLeast {
		cappedXMin = i.XMin
	}
	if i.YMin > atLeast {
		cappedYMin = i.YMin
	}
	if i.ZMin > atLeast {
		cappedZMin = i.ZMin
	}
	if i.XMax < atMost {
		cappedXMax = i.XMax
	}
	if i.YMax < atMost {
		cappedYMax = i.YMax
	}
	if i.ZMax < atMost {
		cappedZMax = i.ZMax
	}

	list := []string{}

	for x := cappedXMin; x <= cappedXMax; x++ {
		for y := cappedYMin; y <= cappedYMax; y++ {
			for z := cappedZMin; z <= cappedZMax; z++ {
				list = append(list, fmt.Sprintf("%d,%d,%d", x, y, z))
			}
		}
	}

	return list
}

func parseIntOrFail(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func ParseFile(file string) []*Instruction {
	lines := strings.Split(strings.TrimSpace(file), "\n")
	instrs := make([]*Instruction, len(lines))

	lineRgx := regexp.MustCompile(`(on|off)\sx=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)
	for idx, line := range lines {
		matches := lineRgx.FindStringSubmatch(line)
		instr := Instruction{
			Toggle: matches[1],
		}

		if parseIntOrFail(matches[2]) <= parseIntOrFail(matches[3]) {
			instr.XMin = parseIntOrFail(matches[2])
			instr.XMax = parseIntOrFail(matches[3])
		} else {
			instr.XMin = parseIntOrFail(matches[3])
			instr.XMax = parseIntOrFail(matches[2])
		}

		if parseIntOrFail(matches[4]) <= parseIntOrFail(matches[5]) {
			instr.YMin = parseIntOrFail(matches[4])
			instr.YMax = parseIntOrFail(matches[5])
		} else {
			instr.YMin = parseIntOrFail(matches[5])
			instr.YMax = parseIntOrFail(matches[4])
		}

		if parseIntOrFail(matches[6]) <= parseIntOrFail(matches[7]) {
			instr.ZMin = parseIntOrFail(matches[6])
			instr.ZMax = parseIntOrFail(matches[7])
		} else {
			instr.ZMin = parseIntOrFail(matches[7])
			instr.ZMax = parseIntOrFail(matches[6])
		}

		instrs[idx] = &instr
	}

	return instrs
}

func Init(list []*Instruction) int {
	cubes := map[string]string{}

	for _, instr := range list {
		cubeList := instr.RangeWithLimits(-50, 50)

		for _, cube := range cubeList {
			cubes[cube] = instr.Toggle
		}
	}

	count := 0
	for _, cube := range cubes {
		if cube == "on" {
			count++
		}
	}

	return count
}
