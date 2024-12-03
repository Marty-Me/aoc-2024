package day3

import (
	_ "embed"
	"testing"
)

//go:embed test/examples.txt
var part_1_input string

//go:embed test/example_2.txt
var part_2_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "161" {
		t.Errorf("day 3 part 1 should solve with '161' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part_2_input)

	if result != "48" {
		t.Errorf("day 3 part 2 should solve with '48' given input, not '%s'", result)
	}
}
