package day8

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "14" {
		t.Errorf("day 8 part 1 should solve with '14' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "34" {
		t.Errorf("day 8 part 2 should solve with '34' given input, not '%s'", result)
	}
}

//go:embed test/simpler_example.txt
var simplerExample string

func TestSimplerExample(t *testing.T) {
	result := SolvePartTwo(simplerExample)

	if result != "6" {
		t.Errorf("should solve with '6' given input, not '%s'", result)
	}
}
