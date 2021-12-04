package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const numberWidth = 12

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

	allNumbers := []uint16{}

	s := bufio.NewScanner(input)
	for s.Scan() {
		raw := s.Text()
		v, err := strconv.ParseUint(raw, 2, 16)
		if err != nil {
			log.Fatalf("unable to convert %q into uint16: %s", err)
		}

		value := uint16(v)

		fmt.Printf("%016b\n", value)
		allNumbers = append(allNumbers, value)
	}
	if err := s.Err(); err != nil {
		log.Fatalf("while reading input: %s", err)
	}

	o2 := oxygen(allNumbers)
	co2 := carbon(allNumbers)

	fmt.Println(o2 * co2)
}

func oxygen(numbers []uint16) uint64 {
	remaining := numbers

	for pos := numberWidth - 1; pos >= 0; pos-- {
		zeroes, ones := split(remaining, pos)

		mostCommon := ones
		if len(zeroes) > len(ones) {
			mostCommon = zeroes
		}

		if len(mostCommon) == 1 {
			return uint64(mostCommon[0])
		}

		remaining = mostCommon
	}

	panic("uh oh, we didn't find an oxygen count")
}

func carbon(numbers []uint16) uint64 {
	remaining := numbers

	for pos := numberWidth - 1; pos >= 0; pos-- {
		zeroes, ones := split(remaining, pos)

		leastCommon := zeroes
		if len(ones) < len(zeroes) {
			leastCommon = ones
		}

		if len(leastCommon) == 1 {
			return uint64(leastCommon[0])
		}

		remaining = leastCommon
	}

	panic("uh oh, we didn't find a co2 count")
}

func split(numbers []uint16, pos int) (zeroes, ones []uint16) {
	for i, n := range numbers {
		n &= uint16(1 << pos)
		n >>= pos

		switch n {
		case 1:
			ones = append(ones, numbers[i])
		case 0:
			zeroes = append(zeroes, numbers[i])
		}
	}

	return zeroes, ones
}
