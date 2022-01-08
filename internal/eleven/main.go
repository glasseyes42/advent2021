package eleven

import (
	"fmt"
	"strconv"
	"strings"
)

type Grid struct {
	Octopi     [10][10]*Octopus
	octopiList [100]*Octopus
	Counter    *Counter
}

func (g *Grid) Iteration() {
	for _, o := range g.octopiList {
		o.AddEnergy()
	}

	for _, o := range g.octopiList {
		o.ResetFlash()
	}
	fmt.Println()
	fmt.Println(g)
}

func (g *Grid) String() string {
	output := ""

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			output = fmt.Sprintf("%s%s", output, g.Octopi[x][y])
		}
		output = fmt.Sprintf("%s\n", output)
	}

	return output
}

func (g *Grid) AddOctopus(o *Octopus, xIdx, yIdx int) {
	g.Octopi[xIdx][yIdx] = o
	g.octopiList[xIdx+(yIdx*10)] = o
}

func parseIntOrFail(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func BuildGrid(file string) Grid {
	lines := strings.Split(strings.TrimSpace(file), "\n")
	g := Grid{
		Counter: &Counter{},
	}

	for yIdx, line := range lines {
		octopi := strings.Split(line, "")

		for xIdx, oVal := range octopi {
			o := Octopus{}
			o.Init(&g, xIdx, yIdx, parseIntOrFail(oVal))
			g.AddOctopus(&o, xIdx, yIdx)
		}
	}

	return g
}

func Iterate(g Grid, iterations int) int {
	for i := 0; i < iterations; i++ {
		g.Iteration()
	}

	return g.Counter.Count()
}
