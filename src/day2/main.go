package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	horizontal, depth, aim := 0, 0, 0

	s := bufio.NewScanner(input)
	for s.Scan() {
		rawCommand := s.Text()
		command := strings.Fields(rawCommand)

		if len(command) != 2 {
			log.Fatalf("expected command %q to have two components, got %d", rawCommand, len(command))
		}

		var distance int
		direction, rawDistance := command[0], command[1]
		distance, err := strconv.Atoi(rawDistance)
		if err != nil {
			log.Fatalf("could not convert distance %q to integer: %s", rawDistance, err)
		}

		switch direction {
		case "forward":
			horizontal += distance
			depth += distance * aim
		case "up":
			aim -= distance
		case "down":
			aim += distance
		default:
			log.Fatalf("unknown direction %q", direction)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("while reading input: %s", err)
	}

	fmt.Println(horizontal * depth)
}
