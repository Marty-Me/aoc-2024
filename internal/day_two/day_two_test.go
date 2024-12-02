package daytwo

import (
	_ "embed"
	"testing"
)

//go:embed test/day_two_example.txt
var part_1_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "2" {
		t.Errorf("day two Part 1 should solve with '2' given input, not '%s'", result)
	}
}
