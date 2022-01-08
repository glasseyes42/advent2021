package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/eleven"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := eleven.Iterate(eleven.BuildGrid(data), 100)
	fmt.Println("Part 1:", part1Answer)

	part2Answer := eleven.IterateTillAllFlash(eleven.BuildGrid(data))
	fmt.Println("Part 1:", part2Answer)
}
