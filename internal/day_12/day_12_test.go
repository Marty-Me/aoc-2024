package day12

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "1930" {
		t.Errorf("day 12 part 1 should solve with '1930' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "1206" {
		t.Errorf("day 12 part 2 should solve with '1206' given input, not '%s'", result)
	}
}

//go:embed test/example_custom_1.txt
var custom1Input string

func TestCustom1(t *testing.T) {
	result := SolvePartOne(custom1Input)

	if result != "80" {
		t.Errorf("day 12 custom 1 should solve with '80' given input, not '%s'", result)
	}
}

//go:embed test/example_custom_2.txt
var custom2Input string

func TestCustom2(t *testing.T) {
	result := SolvePartOne(custom2Input)

	if result != "120" {
		t.Errorf("day 12 custom 1 should solve with '120' given input, not '%s'", result)
	}
}
