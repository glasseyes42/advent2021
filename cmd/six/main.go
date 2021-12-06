package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/six"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := six.ProcessDays(80, six.ParseFile(data))
	fmt.Println("Part 1:", part1Answer)
}
