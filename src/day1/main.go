package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := os.Stdin

	if len(os.Args) > 1 {
		fileName := os.Args[1]
		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("unable to open %q: %w", fileName, err)
		}

		input = f
	}
	defer input.Close()

	twoBefore, previousDepth := 0, 0
	previousWindowSum := 0

	numIncreases := 0
	lineCount := 0

	s := bufio.NewScanner(input)
	for s.Scan() {
		rawDepth := s.Text()

		depth, err := strconv.Atoi(rawDepth)
		if err != nil {
			log.Fatalf("could not convert %q to integer: %w", rawDepth, err)
		}

		windowSum := twoBefore + previousDepth + depth

		if lineCount >= 3 && windowSum > previousWindowSum {
			// skip past the first three lines since we have no comparsion point
			numIncreases++
		}

		previousWindowSum = windowSum

		twoBefore = previousDepth
		previousDepth = depth

		lineCount++
	}

	if err := s.Err(); err != nil {
		log.Fatalf("while reading input: %w", err)
	}

	fmt.Println(numIncreases)
}
