package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/glasseyes42/advent2021/internal/three"
)

//go:embed data/3.1.txt
var part1 string

func main() {
	fmt.Println("Part 1:", three.CalculateConsumption(strings.Split(part1, "\n")))
	fmt.Println("Part 2:", three.CalculateSupport(strings.Split(part1, "\n")))
}
