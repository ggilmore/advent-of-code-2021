package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const boardWidth = 5
const boardSize = boardWidth * boardWidth

func main() {
	input := os.Stdin

	if len(os.Args) > 1 {
		fileName := os.Args[1]
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("unable to open %q: %s", fileName, err)
		}

		input = f
	}
	defer input.Close()

	s := bufio.NewScanner(input)
	chosenNumbers, err := parseChosenNumbers(s)
	if err != nil {
		log.Fatalf("failed to parse chosen numbers: %s", err)
	}

	boards, err := parseBoards(s)
	if err != nil {
		log.Fatalf("failed to parse boards: %s", err)
	}

	for _, b := range boards {
		fmt.Println(b)
	}

	var scores []uint16
	var nextBoards []*board

	for _, n := range chosenNumbers {
		nextBoards = []*board{}

		for i := 0; i < len(boards); i++ {
			b := boards[i]

			b.mark(n)
			if b.bingo() {
				scores = append(scores, b.score(n))
				fmt.Printf("chosen number:%d \n", n)
				fmt.Println(b)
				continue
			}

			nextBoards = append(nextBoards, b)
		}
		boards = nextBoards
	}

	fmt.Println(scores[len(scores)-1])
}

type board struct {
	slots [boardSize]uint8
}

func (b board) String() string {
	var sb strings.Builder

	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardWidth; j++ {
			n := b.slots[i*boardWidth+j]
			sb.WriteString(fmt.Sprintf("%d ", int(n)))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (b *board) score(final uint8) uint16 {
	sum := uint16(0)
	for _, n := range b.slots {
		sum += uint16(n)
	}

	return sum * uint16(final)
}

func (b *board) mark(n uint8) {
	for i := 0; i < boardSize; i++ {
		if b.slots[i] == n {
			b.slots[i] = uint8(0)
		}
	}
}

func (b *board) bingo() bool {
	// horizontal
	for i := 0; i < boardWidth; i++ {
		allZeroes := true

		for j := 0; j < boardWidth; j++ {
			if b.slots[i*boardWidth+j] != 0 {
				allZeroes = false
				break
			}
		}

		if allZeroes {
			return true
		}
	}

	// vertical
	for i := 0; i < boardWidth; i++ {
		allZeroes := true
		for j := 0; j < boardWidth; j++ {
			if b.slots[i+boardWidth*j] != 0 {
				allZeroes = false
				break
			}
		}

		if allZeroes {
			return true
		}
	}

	return false
}

func parseBoards(s *bufio.Scanner) ([]*board, error) {
	var out []*board

	var slots [boardSize]uint8
	var idx = 0

	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		if line == "" {
			switch idx {
			case boardSize:
				out = append(out, &board{slots})
				idx = 0
				fallthrough
			case 0:
				continue

			default:
				return nil, fmt.Errorf("expected board with %d slots, got %d", boardSize, idx)
			}
		}

		rawNumbers := strings.Split(line, " ")
		for _, r := range rawNumbers {
			if r == "" {
				continue
			}

			v, err := strconv.ParseUint(r, 10, 8)
			if err != nil {
				return nil, fmt.Errorf("failed to parse %q into uint8: %w", r, err)
			}

			slots[idx] = uint8(v)
			idx++
		}
	}

	return out, nil
}

func parseChosenNumbers(s *bufio.Scanner) ([]uint8, error) {
	if !s.Scan() {
		if err := s.Err(); err != nil {
			return nil, fmt.Errorf("scanning: %w", err)
		}

		return nil, fmt.Errorf("unexpected end of file")
	}

	var out []uint8

	rawNumbers := strings.Split(s.Text(), ",")
	for _, r := range rawNumbers {
		v, err := strconv.ParseUint(r, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q into uint8: %w", r, err)
		}

		out = append(out, uint8(v))
	}

	return out, nil
}
