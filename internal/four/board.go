package four

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type GridPoint struct {
	Value  int
	Called bool
}

type Board struct {
	Lookup  map[int]*GridPoint
	Columns [5][5]*GridPoint
	Rows    [5][5]*GridPoint
}

func (b *Board) Init() *Board {
	b.Lookup = make(map[int]*GridPoint)
	return b
}

func (b *Board) Parse(lines []string) {
	if len(lines) > 5 {
		panic(fmt.Errorf("cannot parse a board with lines more than 5. Received %v", len(lines)))
	}
	for rowIdx, line := range lines {
		spaceRegex := regexp.MustCompile(`\s+`)
		vals := spaceRegex.Split(strings.TrimSpace(line), -1)
		if len(vals) > 5 {
			panic(fmt.Errorf("cannot parse a board with columns more than 5. Received %v", len(vals)))
		}
		for columnIdx, val := range vals {

			intVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}

			gp := GridPoint{
				Value: intVal,
			}

			b.Rows[rowIdx][columnIdx] = &gp
			b.Columns[columnIdx][rowIdx] = &gp
			b.Lookup[intVal] = &gp
		}
	}
}

func (b *Board) Call(val int) {
	gp, ok := b.Lookup[val]
	if ok {
		gp.Called = true
	}
}

func (b *Board) CheckWon() bool {
	won := false

	for _, row := range b.Rows {
		combined := true
		for _, gp := range row {
			combined = combined && gp.Called
		}

		if combined {
			won = true
			return won
		}
	}

	for _, col := range b.Columns {
		combined := true
		for _, gp := range col {
			combined = combined && gp.Called
		}

		if combined {
			won = true
			return won
		}
	}

	return won
}

func (b *Board) Score(lastCalled int) int {
	sum := 0
	for _, gp := range b.Lookup {
		if !gp.Called {
			sum += gp.Value
		}
	}
	return sum * lastCalled
}
