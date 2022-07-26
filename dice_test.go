package dice_test

import (
	"fmt"
	"regexp"
	"testing"

	dice "github.com/atallison/dicer"
	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {

	var tests = []struct {
		equation    string
		count       int
		value       int
		keepHigh    int
		keepLow     int
		outputRegex *regexp.Regexp
	}{
		{"1d6", 1, 6, 0, 0, regexp.MustCompile("^\\d$")},
		{"2d4", 2, 4, 0, 0, regexp.MustCompile("^Rolled: \\d+ = \\d \\+ \\d$")},
		{"6d20", 6, 20, 0, 0, regexp.MustCompile("^Rolled: \\d+ = \\d{1,2}( \\+ \\d{1,2}){5}$")},
		{"3d20kh2", 3, 20, 2, 0, regexp.MustCompile("^Rolled: \\d+ = \\d{1,2}( \\+ \\d{1,2})$")},
		{"8d8kl3", 8, 8, 0, 3, regexp.MustCompile("^Rolled: \\d+ = \\d{1}( \\+ \\d{1}){2}$")},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Simple_roll_%s", test.equation)
		t.Run(testName, func(t *testing.T) {
			roll, err := dice.Roll(test.equation)
			assert.NoError(t, err)
			assert.Equal(t, test.count, len(roll.Rolls))
			assert.Equal(t, test.value, roll.Rolls[0].(dice.RollDie).Sides)
			assert.Equal(t, test.keepHigh, roll.KeepHighest)
			assert.Equal(t, test.keepLow, roll.KeepLowest)
			assert.Regexp(t, test.outputRegex, roll.ToString())
		})
	}

}
