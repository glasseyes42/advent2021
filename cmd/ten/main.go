package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/ten"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := ten.ScoreIllegalLines(data)
	fmt.Println("Part 1:", part1Answer)

	part2Answer := ten.ScoreCompletions(data)
	fmt.Println("Part 2:", part2Answer)
}
