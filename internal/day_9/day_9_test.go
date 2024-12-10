package day9

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "1928" {
		t.Errorf("day 9 part 1 should solve with '1928' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "2858" {
		t.Errorf("day 9 part 2 should solve with '?' given input, not '%s'", result)
	}
}

//go:embed test/example_2.txt
var part2Input string

func TestPartTwoExample(t *testing.T) {
	result := SolvePartTwo(part2Input)

	if result != "132" {
		t.Errorf("day 9 part 2 should solve with '?' given input, not '%s'", result)
	}
}
