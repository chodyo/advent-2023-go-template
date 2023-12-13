//nolint:goconst // Don't wanna make "file" str a const
package commands

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/chodyo/advent-2023-go-template/internal/input"
	"github.com/google/go-cmp/cmp"
)

func Test_Commands_TableTesting(t *testing.T) {
	testGetters := map[string]map[string]func() map[string]Test{
		`day 01`: {
			`part 1`: getDay01Part1Tests,
			`part 2`: getDay01Part2Tests,
		},
		`day 02`: {
			`part 1`: getDay02Part1Tests,
			`part 2`: getDay02Part2Tests,
		},
		`day 03`: {
			`part 1`: getDay03Part1Tests,
			`part 2`: getDay03Part2Tests,
		},
		`day 04`: {
			`part 1`: getDay04Part1Tests,
			`part 2`: getDay04Part2Tests,
		},
	}

	for day, getParts := range testGetters {
		for part, getTestCases := range getParts {
			for name, testCase := range getTestCases() {
				tc := testCase
				t.Run(fmt.Sprintf("%s %s %s", day, part, name), func(t *testing.T) {
					t.Parallel()

					actual := tc.command.Process(nil)
					diff := cmp.Diff(tc.expected, actual)
					if diff != "" {
						t.Log(diff)
						t.Fail()
					}
				})
			}
		}
	}
}

type Test struct {
	command  AdventCommand
	expected int
}

func getDay01Part1Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`01-1_input_sample`: adventAnswers[`day 01 part 1 sample 1`],
		`01-2_input_sample`: adventAnswers[`day 01 part 1 sample 2`],
		`01_input_full`:     adventAnswers[`day 01 part 1 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day01Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                false,
			},
			expected: out,
		}
	}
	return tests
}

func getDay01Part2Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`01-1_input_sample`: adventAnswers[`day 01 part 2 sample 1`],
		`01-2_input_sample`: adventAnswers[`day 01 part 2 sample 2`],
		`01_input_full`:     adventAnswers[`day 01 part 2 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day01Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                true,
			},
			expected: out,
		}
	}
	return tests
}

func getDay02Part1Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`02_input_sample`: adventAnswers[`day 02 part 1 sample`],
		`02_input_full`:   adventAnswers[`day 02 part 1 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day02Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                false,
			},
			expected: out,
		}
	}
	return tests
}

func getDay02Part2Tests() map[string]Test {
	tests := map[string]Test{}
	expectations := map[string]int{
		`02_input_sample`: adventAnswers[`day 02 part 2 sample`],
		`02_input_full`:   adventAnswers[`day 02 part 2 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day02Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                true,
			},
			expected: out,
		}
	}
	return tests
}

func getDay03Part1Tests() map[string]Test {
	tests := map[string]Test{}

	expectations := map[string]int{
		`03_input_sample`: adventAnswers[`day 03 part 1 sample`],
		`03_input_full`:   adventAnswers[`day 03 part 1 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day03Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                false,
			},
			expected: out,
		}
	}

	return tests
}

func getDay03Part2Tests() map[string]Test {
	tests := map[string]Test{}

	expectations := map[string]int{
		`03_input_sample`: adventAnswers[`day 03 part 2 sample`],
		`03_input_full`:   adventAnswers[`day 03 part 2 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day03Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                true,
			},
			expected: out,
		}
	}

	return tests
}

func getDay04Part1Tests() map[string]Test {
	tests := map[string]Test{}

	expectations := map[string]int{
		`04_input_sample`: adventAnswers[`day 04 part 1 sample`],
		`04_input_full`:   adventAnswers[`day 04 part 1 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day04Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                false,
			},
			expected: out,
		}
	}

	return tests
}

func getDay04Part2Tests() map[string]Test {
	tests := map[string]Test{}

	expectations := map[string]int{
		`04_input_sample`: adventAnswers[`day 04 part 2 sample`],
		`04_input_full`:   adventAnswers[`day 04 part 2 full`],
	}
	for in, out := range expectations {
		tests["file "+in] = Test{
			command: Day04Command{
				DefaultAdventCommand: makeDefaultAdventCommand(in),
				Part2:                true,
			},
			expected: out,
		}
	}

	return tests
}

func makeDefaultAdventCommand(filename string) DefaultAdventCommand {
	return DefaultAdventCommand{
		Option: input.Option{
			Filename: filepath.Join("input", filename) + ".txt",
		},
	}
}
