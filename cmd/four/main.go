package main

import (
	_ "embed"
	"fmt"

	"github.com/glasseyes42/advent2021/internal/four"
)

//go:embed data/data.txt
var data string

func main() {
	part1Answer := four.Bingo(data)
	fmt.Println("Part 1:", part1Answer)

	part2Answer := four.BingoLastWinner(data)
	fmt.Println("Part 2:", part2Answer)

}
