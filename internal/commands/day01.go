package commands

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Day01Command struct {
	DefaultAdventCommand

	Part2 bool `long:"part2" description:"Day 2 Part 2; parse digits and words"`
}

func (c Day01Command) Execute(args []string) error {
	log.Printf("%d\n", c.Process(args))
	return nil
}

func (c Day01Command) Process(_ []string) int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return -1
	}

	if c.Part2 {
		return day01part2(input)
	}
	return day01part1(input)
}

// day01part1 is your implementation of advent of code 2023 day 1 part 1
func day01part1(_ []string) (sum int) {
	// Your implementation here

	return sum
}

// day01part2 is your implementation of advent of code 2023 day 1 part 2
func day01part2(_ []string) (sum int) {
	// Your implementation here

	return sum
}
