package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/twentyfive"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := twentyfive.IterateTillStopped(twentyfive.ParseFile(data))
	fmt.Println("Part 1:", part1Answer)
}
