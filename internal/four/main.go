package four

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseOutput struct {
	Numbers []int
	Boards  [][]string
}

func ParseFile(file string) (ParseOutput, error) {
	sections := strings.Split(strings.TrimSpace(file), "\n\n")

	numbersLine := strings.Split(sections[0], ",")
	numbers := make([]int, len(numbersLine))
	for idx, num := range numbersLine {
		val, err := strconv.Atoi(num)
		if err != nil {
			return ParseOutput{}, fmt.Errorf("could not convert to int. %s. %w", num, err)
		}
		numbers[idx] = val
	}

	sections = sections[1:]
	boards := make([][]string, len(sections))
	for idx, block := range sections {
		boards[idx] = strings.Split(block, "\n")
	}

	return ParseOutput{
		Numbers: numbers,
		Boards:  boards,
	}, nil
}

func Bingo(file string) int {
	parsed, err := ParseFile(file)
	if err != nil {
		panic(err)
	}

	boards := make([]*Board, len(parsed.Boards))
	for idx, data := range parsed.Boards {
		boards[idx] = (&Board{}).Init()
		boards[idx].Parse(data)
	}

	remaining := parsed.Numbers

	for len(remaining) > 4 {
		numbers := remaining[:5]
		remaining = remaining[5:]

		for _, num := range numbers {
			for _, board := range boards {
				board.Call(num)
				if board.CheckWon() {
					return board.Score(num)
				}
			}
		}
	}

	return 0
}

func BingoLastWinner(file string) int {
	type boardWithStatus struct {
		Board         *Board
		WinningNumber int
		Skip          bool
	}

	parsed, err := ParseFile(file)
	if err != nil {
		panic(err)
	}

	boards := make([]*boardWithStatus, len(parsed.Boards))
	for idx, data := range parsed.Boards {
		board := (&Board{}).Init()
		board.Parse(data)

		boards[idx] = &boardWithStatus{
			Board: board,
		}
	}

	winningBoards := []*boardWithStatus{}
	remaining := parsed.Numbers

	for len(remaining) > 4 {
		numbers := remaining[:5]
		remaining = remaining[5:]

		for _, num := range numbers {
			for _, board := range boards {
				if board.Skip {
					continue
				}

				board.Board.Call(num)
				if board.Board.CheckWon() {
					board.WinningNumber = num
					board.Skip = true
					winningBoards = append(winningBoards, board)
				}
			}
		}
	}

	return winningBoards[len(winningBoards)-1].Board.Score(winningBoards[len(winningBoards)-1].WinningNumber)
}
