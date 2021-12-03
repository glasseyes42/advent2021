package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/glasseyes42/advent2021/internal/one"
)

//go:embed data/1.1.txt
var part1 string

func main() {
	fmt.Println("Part 1:", one.CountIncreases(strings.Split(part1, "\n")))
	fmt.Println("Part 2:", one.CountIncreasesWindowed(strings.Split(part1, "\n")))
}
