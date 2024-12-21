package day_16

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "11048" {
		t.Errorf("day 16 part 1 should solve with '11048' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "?" {
		t.Errorf("day 16 part 2 should solve with '?' given input, not '%s'", result)
	}
}
