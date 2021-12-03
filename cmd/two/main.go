package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/glasseyes42/advent2021/internal/two"
)

//go:embed data/2.1.txt
var part1 string

func main() {
	part1Answer := two.ProcessInstructions(strings.Split(part1, "\n"))
	fmt.Println("Part 1:", part1Answer.Depth*part1Answer.Horizonal)

	part2Answer := two.ProcessInstructionsWithAim(strings.Split(part1, "\n"))
	fmt.Println("Part 2:", part2Answer.Depth*part2Answer.Horizonal)
}
