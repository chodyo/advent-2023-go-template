package trebuchet

import (
	"log"

	"github.com/chodyo/advent-2023/internal/input"
)

type Command struct {
	Calibrate CalibrateCommand `command:"calibrate" description:"Parse numbers from text"`
}

type CalibrateCommand struct {
	input.Option

	DigitsOnly bool `long:"digits-only" description:"Day 1 Part 1; only parse digits from input"`
}

func (c CalibrateCommand) Execute(_ []string) error {
	log.Printf("sum=%d\n", c.Process())
	return nil
}

func (c CalibrateCommand) Process() int {
	input, err := input.LoadInput(c.Filename)
	if err != nil {
		return 0
	}

	if c.DigitsOnly {
		return calibrateDigits(input)
	}
	return calibrateDigitsAndWords(input)
}

// calibrateDigits is your implementation of advent of code 2023 day 1 part 1
func calibrateDigits(lines []string) int {
	var sum int

	// Your implementation here

	return sum
}

// calibrateDigits is your implementation of advent of code 2023 day 1 part 2
func calibrateDigitsAndWords(lines []string) int {
	var sum int

	// Your implementation here

	return sum
}
