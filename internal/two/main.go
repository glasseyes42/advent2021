package two

import (
	"regexp"
	"strconv"
)

type Position struct {
	Depth     int
	Horizonal int
}

type PositionWithAim struct {
	Depth     int
	Horizonal int
	Aim       int
}

type instruction struct {
	Desc   string
	Amount int
}

var instrMatcher = regexp.MustCompile(`(.+)\s([0-9]+)`)

func parseInstructions(lines []string) []instruction {
	instructions := []instruction{}

	for _, line := range lines {
		matches := instrMatcher.FindStringSubmatch(line)

		amount, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{
			Desc:   matches[1],
			Amount: amount,
		})
	}

	return instructions
}

func ProcessInstructions(lines []string) Position {
	instructions := parseInstructions(lines)

	location := Position{}
	for _, instr := range instructions {
		switch instr.Desc {
		case "forward":
			location.Horizonal += instr.Amount
		case "up":
			location.Depth -= instr.Amount
		case "down":
			location.Depth += instr.Amount
		}
	}

	return location
}

func ProcessInstructionsWithAim(lines []string) PositionWithAim {
	instructions := parseInstructions(lines)

	location := PositionWithAim{}
	for _, instr := range instructions {
		switch instr.Desc {
		case "forward":
			location.Horizonal += instr.Amount
			location.Depth += instr.Amount * location.Aim
		case "up":
			location.Aim -= instr.Amount
		case "down":
			location.Aim += instr.Amount
		}
	}

	return location
}
