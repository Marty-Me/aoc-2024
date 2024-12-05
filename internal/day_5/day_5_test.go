package day5

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part_1_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "143" {
		t.Errorf("day 5 part 1 should solve with '143' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part_1_input)

	if result != "?" {
		t.Errorf("day 5 part 2 should solve with '?' given input, not '%s'", result)
	}
}
