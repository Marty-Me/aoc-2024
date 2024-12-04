package day4

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part_1_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "18" {
		t.Errorf("day 4 part 1 should solve with '18' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part_1_input)

	if result != "9" {
		t.Errorf("day 4 part 1 should solve with '18' given input, not '%s'", result)
	}
}
