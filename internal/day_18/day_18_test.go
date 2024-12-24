package day_18

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "22" {
		t.Errorf("day 18 part 1 should solve with '22' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "?" {
		t.Errorf("day 18 part 2 should solve with '?' given input, not '%s'", result)
	}
}
