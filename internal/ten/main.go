package ten

import (
	"fmt"
	"strings"
)

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}
var closing = []string{")", "]", "}", ">"}

func isClose(s string) bool {
	for _, c := range closing {
		if c == s {
			return true
		}
	}
	return false
}

type Chunk struct {
	OpeningChar string
	Chunks      []*Chunk
}

func (c *Chunk) Parse(tokens []string) ([]string, error) {
	t, rest := tokens[0], tokens[1:]
	c.OpeningChar = t

	for len(rest) > 0 {
		if rest[0] == pairs[c.OpeningChar] {
			return rest[1:], nil
		}

		if isClose(rest[0]) {
			return rest, fmt.Errorf("expected %s, but found %s", pairs[c.OpeningChar], rest[0])
		}

		inner := Chunk{}
		c.Chunks = append(c.Chunks, &inner)
		var err error
		rest, err = inner.Parse(rest)
		if err != nil {
			return rest, err
		}
	}

	return rest, nil
}

func (c *Chunk) String(indent int) string {
	output := fmt.Sprintf("%s\n", c.OpeningChar)

	for _, inner := range c.Chunks {
		indentStr := "   "
		for i := 0; i < indent; i++ {
			indentStr = fmt.Sprintf("%s ", indentStr)
		}

		output = fmt.Sprintf("%s\n%s | => %s", output, indentStr, inner.String(indent+1))
	}
	return output
}

type Program struct {
	Chunks  []*Chunk
	Invalid string
}

func (p *Program) Parse(tokens []string) error {
	remaining := tokens
	for len(remaining) > 0 {
		c := Chunk{}
		p.Chunks = append(p.Chunks, &c)

		var err error
		remaining, err = c.Parse(remaining)
		if err != nil {
			p.Invalid = remaining[0]
			return err
		}
	}

	return nil
}

func (p *Program) String() string {
	output := ""
	for _, c := range p.Chunks {
		output = fmt.Sprintf("%s\n%% %s", output, c.String(0))
	}
	return output
}

func Parse(line string) (Program, error) {
	// fmt.Printf("Processing line: '%s'\n", line)
	tokens := strings.Split(line, "")
	p := Program{}
	err := p.Parse(tokens)
	// fmt.Println(p.String())
	return p, err
}

func ScoreIllegalLines(file string) int {
	lines := strings.Split(strings.TrimSpace(file), "\n")
	score := 0

	for _, line := range lines {
		p, err := Parse(line)
		if err != nil {
			score += scores[p.Invalid]
		}
	}
	return score
}
