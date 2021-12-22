package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/twentytwo"
)

//go:embed data/input.txt
var data string

func main() {
	part1Answer := twentytwo.Init(twentytwo.ParseFile(data))
	fmt.Println("Part 1:", part1Answer)
}
