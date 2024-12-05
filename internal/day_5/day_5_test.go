package day5

import (
	_ "embed"
	"testing"
)

//go:embed test/example.txt
var part1Input string

func TestPartOne(t *testing.T) {
	result := SolvePartOne(part1Input)

	if result != "143" {
		t.Errorf("day 5 part 1 should solve with '143' given input, not '%s'", result)
	}
}

func TestPartTwo(t *testing.T) {
	result := SolvePartTwo(part1Input)

	if result != "123" {
		t.Errorf("day 5 part 2 should solve with '123' given input, not '%s'", result)
	}
}

//go:embed test/complex-example.txt
var complexInput string

func TestComplexInput(t *testing.T) {
	result := SolvePartTwo(complexInput)

	if result != "61" {
		t.Errorf("Input should solve with '61' given input, not '%s'", result)
	}
}

//go:embed test/example_five.txt
var example5 string

func TestExample5Input(t *testing.T) {
	result := SolvePartTwo(example5)

	if result != "29" {
		t.Errorf("Input should solve with '29' given input, not '%s'", result)
	}
}

//go:embed test/example_six.txt
var example6 string

func TestExample6Input(t *testing.T) {
	result := SolvePartTwo(example6)

	if result != "47" {
		t.Errorf("Input should solve with '47' given input, not '%s'", result)
	}
}
