package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/seven"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := seven.CalcFuel(seven.ParseFile(data))
	fmt.Println("Part 1:", part1Answer)

	part2Answer := seven.CalcFuel2(seven.ParseFile(data))
	fmt.Println("Part 2:", part2Answer)
}
