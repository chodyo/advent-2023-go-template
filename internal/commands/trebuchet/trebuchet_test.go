package trebuchet

import (
	"testing"

	"github.com/chodyo/advent-2023/internal/input"
	"github.com/google/go-cmp/cmp"
)

func Test_Trebuchet_Calibration_DigitsOnly_Table(t *testing.T) {
	for name, testCase := range getTrebuchetCalibrationDigitsOnlyTableTests() {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			command := CalibrateCommand{
				Option: input.Option{
					Filename: tc.input,
				},
				DigitsOnly: true,
			}

			actual := command.Process()
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Log(diff)
				t.Fail()
			}
		})
	}
}

func Test_Trebuchet_Calibration_DigitsAndWords_Table(t *testing.T) {
	for name, testCase := range getTrebuchetCalibrationDigitsAndWordsTableTests() {
		tc := testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			command := CalibrateCommand{
				Option: input.Option{
					Filename: tc.input,
				},
				DigitsOnly: false,
			}

			actual := command.Process()
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Log(diff)
				t.Fail()
			}
		})
	}
}

type TrebuchetCalibrationTest struct {
	input    string
	expected int
}

func getTrebuchetCalibrationDigitsOnlyTableTests() map[string]TrebuchetCalibrationTest {
	return map[string]TrebuchetCalibrationTest{
		`sample day 1 part 1 test`: {
			input:    `calibrate_digitsOnly_input_sample.txt`,
			expected: 142,
		},
		`cody's day 1 part 1 test`: {
			input:    `calibrate_input_cody.txt`,
			expected: 56042,
		},
	}
}

func getTrebuchetCalibrationDigitsAndWordsTableTests() map[string]TrebuchetCalibrationTest {
	return map[string]TrebuchetCalibrationTest{
		`sample day 1 part 2 test`: {
			input:    `calibrate_digitsAndWords_input_sample.txt`,
			expected: 281,
		},
		`cody's day 1 part 2 test`: {
			input:    `calibrate_input_cody.txt`,
			expected: 55358,
		},
	}
}
