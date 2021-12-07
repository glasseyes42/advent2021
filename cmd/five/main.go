package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/five"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := five.CountPointsWithOverlap(2, five.FilterStraight(five.ParseFile(data)))
	fmt.Println("Part 1:", part1Answer)

	part2Answer := five.CountPointsWithOverlap(2, five.ParseFile(data))
	fmt.Println("Part 2:", part2Answer)

}
