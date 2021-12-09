package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/nine"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := nine.RiskOfLowLevels(data)
	fmt.Println("Part 1:", part1Answer)

	part2Answer := nine.LargestBasins(data)
	fmt.Println("Part 2:", part2Answer)

}
