package dayone

import (
	_ "embed"
	"testing"
)

//go:embed test/day_one_example.txt
var part_1_input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part_1_input)

	if result != "11" {
		t.Errorf("day one Part 1 should solve with '11' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part_1_input)

	if result != "31" {
		t.Errorf("day one Part 1 should solve with '31' given input, not '%s'", result)
	}
}
