package day_17

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "4,6,3,5,6,3,5,2,1,0" {
		t.Errorf("day 16 part 1 should solve with '4,6,3,5,6,3,5,2,1,0' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "?" {
		t.Errorf("day 16 part 2 should solve with '?' given input, not '%s'", result)
	}
}
