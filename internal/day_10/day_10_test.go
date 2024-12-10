package day10

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "36" {
		t.Errorf("day 10 part 1 should solve with '36' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "81" {
		t.Errorf("day 10 part 2 should solve with '81' given input, not '%s'", result)
	}
}
