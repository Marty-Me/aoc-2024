package day5

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "41" {
		t.Errorf("day 6 part 1 should solve with '41' given input, not '%s'", result)
	}
}

//go:embed test/example_right.txt
var exampleRight string

func TestRight(t *testing.T) {
	result := SolvePartOne(exampleRight)

	if result != "2" {
		t.Errorf("Should solve for 2 not '%s'", result)
	}
}

//go:embed test/example_left.txt
var exampleLeft string

func TestLeft(t *testing.T) {
	result := SolvePartOne(exampleLeft)

	if result != "3" {
		t.Errorf("Should solve for 3 not '%s'", result)
	}
}

//go:embed test/example_down.txt
var exampleDown string

func TestDown(t *testing.T) {
	result := SolvePartOne(exampleDown)
	if result != "3" {
		t.Errorf("Should solve for 3 not '%s'", result)
	}
}

//go:embed test/example_up.txt
var exampleUp string

func TestUp(t *testing.T) {
	result := SolvePartOne(exampleUp)
	if result != "5" {
		t.Errorf("Should solve for 5 not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "?" {
		t.Errorf("day 6 part 2 should solve with '?' given input, not '%s'", result)
	}
}
